package pbsql

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

// BuildCountQuery is a convenience wrapper for getting the result count of a query already generated by pbsql
// value based, does not affect the initially supplied query string
func BuildCountQuery(selectQry string) string {
	return "SELECT COUNT(*) as count FROM (" + selectQry + ") as count"
}

// BuildCreateQuery accepts a target table name and a protobuf message and attempts to build a valid SQL insert statement for use
// with sqlx.Named, ignoring any struct fields with default values. Fields must be tagged with `db:""` in order to be
// included in the result string.
func BuildCreateQuery(target string, source interface{}) (string, []interface{}, error) {
	t := reflect.ValueOf(source).Elem()
	var cols strings.Builder
	var vals strings.Builder
	fmt.Fprintf(&cols, "INSERT INTO %s (", target)
	vals.WriteString("(")

	for i := 0; i < t.NumField(); i++ {
		valField := t.Field(i)
		typeField := t.Type().Field(i)
		typeName := valField.Type().Name()
		isPrimaryKey := typeField.Tag.Get("primary_key") != ""
		tag := typeField.Tag.Get("db")

		if notDefault(typeName, valField.Interface()) && tag != "" && !isPrimaryKey {
			if i != 0 {
				cols.WriteString(", ")
				vals.WriteString(", ")
			}
			fmt.Fprintf(&cols, "%s.%s", target, tag)
			fmt.Fprintf(&vals, ":%s", tag)
		}
	}
	vals.WriteString(")")
	fmt.Fprintf(&cols, ") VALUES %s", vals.String())
	result := strings.ReplaceAll(cols.String(), "(, ", "(")
	return sqlx.Named(result, source)
}

// BuildDeleteQuery accepts a target table name and a protobuf message and attempts to build a valid SQL
// delete statement by utilizing struct tags to denote information such as database field names and
// whether something is a primary key. If successful, returns a SQL statement in the form of a string,
// a slice of args to interpolate, and a nil error.
//
// This function returns a nullsafe query if nullable struct fields are properly tagged as `nullable:"y"`.
//
// If an IsActive field is detected (is_active), this func returns an update statement that sets is_active to 0,
// otherwise it returns a delete statement
func BuildDeleteQuery(target string, source interface{}) (string, []interface{}, error) {
	v := reflect.ValueOf(source).Elem()
	t := v.Type()
	var builder strings.Builder

	isActive, hasIsActive := t.FieldByName("IsActive")
	if hasIsActive {
		dbName := isActive.Tag.Get("db")
		fmt.Fprintf(&builder, "UPDATE %s SET %s.%s = :%s WHERE ", target, target, dbName, dbName)
	} else {
		fmt.Fprintf(&builder, "DELETE FROM %s WHERE ", target)
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		isPkey := typeField.Tag.Get("primary_key") != ""
		if isPkey {
			dbName := typeField.Tag.Get("db")
			fmt.Fprintf(&builder, "%s.%s = :%s", target, dbName, dbName)
			break
		}
	}

	return sqlx.Named(builder.String(), source)
}

// BuildReadQuery accepts a target table name and a protobuf message and attempts to build a valid SQL select statement,
// ignoring any struct fields with default values when writing predicates. Fields must be tagged with `db:""` in order to be
// included in the result string.
//
// Returns a SQL statement as a string, a slice of args to interpolate, and an error
func BuildReadQuery(target string, source interface{}) (string, []interface{}, error) {
	nullHandler := "ifnull("
	if sqlDriver := os.Getenv("GRPC_SQL_DRIVER"); sqlDriver == "pgsql" {
		nullHandler = "coalesce("
	}

	t := reflect.ValueOf(source).Elem()

	var core strings.Builder
	var joins strings.Builder
	var fields strings.Builder
	var predicate strings.Builder
	core.WriteString("SELECT ")
	predicate.WriteString(" WHERE true")

	for i := 0; i < t.NumField(); i++ {
		valField := t.Field(i)
		typeField := t.Type().Field(i)
		typeName := valField.Type().Name()
		dbName := typeField.Tag.Get("db")
		nullable := typeField.Tag.Get("nullable")
		foreignKey := typeField.Tag.Get("foreign_key")
		foreignTable := typeField.Tag.Get("foreign_table")
		localName := typeField.Tag.Get("local_name")

		if dbName != "" {
			if nullable != "" {
				fmt.Fprintf(&fields, "%s%s.%s, %s) as %s, ", nullHandler, target, dbName, getDefault(typeName), dbName)
			} else {
				fmt.Fprintf(&fields, "%s.%s, ", target, dbName)
			}

			if valField.CanInterface() && notDefault(typeName, valField.Interface()) {
				fmt.Fprintf(&predicate, " AND %s.%s", target, dbName)
				if typeName == "string" {
					fmt.Fprintf(&predicate, " LIKE :%s", dbName)
				} else {
					fmt.Fprintf(&predicate, " = :%s", dbName)
				}
			}
		}

		if foreignKey != "" && foreignTable != "" && localName != "" {
			related := reflect.Indirect(valField)
			if related.CanAddr() {
				for j := 0; j < related.NumField(); j++ {
					relatedValField := related.Field(j)
					relatedTypeField := related.Type().Field(j)
					relatedTypeName := relatedValField.Type().Name()
					relatedDBName := relatedTypeField.Tag.Get("db")
					
					if relatedDBName != "" && relatedValField.CanInterface() && notDefault(relatedTypeName, relatedValField.Interface()) {
						fmt.Fprintf(&predicate, " AND %s.%s", foreignTable, relatedDBName)
						if relatedTypeName == "string" {
							fmt.Fprintf(&predicate, " LIKE '%s'", relatedValField)
						} else {
							fmt.Fprintf(&predicate, " = %s", relatedValField)
						}
					}
				}
				fmt.Fprintf(&joins, " LEFT JOIN %s on %s.%s = %s.%s", foreignTable, foreignTable, foreignKey, target, localName)
			}
		}
	}

	fmt.Fprintf(&core, "%sFROM %s%s%s", fields.String(), target, joins.String(), predicate.String())

	orderBy := t.FieldByName("OrderBy")
	orderDir := t.FieldByName("OrderDir")
	
	if orderBy.CanAddr() && orderBy.String() != "" {
		orderStr := fmt.Sprintf(" order by %s", orderBy.String())
		if orderDir.CanAddr() && orderDir.String() != "" {
			orderStr = fmt.Sprintf("%s %s", orderStr, orderDir.String())
		} else {
			orderStr = fmt.Sprintf("%s asc", orderStr)
		}
		fmt.Fprint(&core, orderStr)
	}

	result := strings.Replace(core.String(), ", FROM", " FROM", 1)
	return sqlx.Named(result, source)
}

// BuildUpdateQuery accepts a target table name `target`, a struct `source`, and a list of struct fields `fieldMask`
// and attempts to build a valid sql update statement for use with sqlx.Named, ignoring any struct fields not present
// in `fieldMask`. Struct fields must also be tagged with `db:""`, and the primary key should be tagged as
// `primary_key` otherwise this function will return an invalid query
func BuildUpdateQuery(target string, source interface{}, fieldMask []string) (string, []interface{}, error) {
	v := reflect.ValueOf(source).Elem()
	t := v.Type()

	var builder strings.Builder
	fmt.Fprintf(&builder, "UPDATE %s SET ", target)

	var predicate strings.Builder
	for i := 0; i < v.NumField(); i++ {
		valField := v.Field(i)
		typeField := t.Field(i)
		dbName := typeField.Tag.Get("db")

		if valField.CanInterface() && dbName != "" {
			isPrimaryKey := typeField.Tag.Get("primary_key") != ""
			if isPrimaryKey {
				fmt.Fprintf(&predicate, "WHERE %s.%s = :%s", target, dbName, dbName)
			} else if findInMask(fieldMask, typeField.Name) {
				fmt.Fprintf(&builder, "%s.%s = :%s, ", target, dbName, dbName)
			}
		}
	}

	builder.WriteString(predicate.String())
	result := strings.Replace(builder.String(), ", WHERE", " WHERE", 1)
	return sqlx.Named(result, source)
}

// `notDefault` checks if a value is set to it's unitialized default, e.g. whether or not an `int32` value is `0`
// returns `true` if not default.
func notDefault(typeName string, fieldVal interface{}) bool {
	switch typeName {
	case "uint", "int", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return fieldVal.(int32) != 0
	case "float32", "float64":
		return fieldVal.(float64) != 0
	case "string":
		return fieldVal.(string) != ""
	default:
		return fieldVal != nil
	}
}

// `getDefault` returns the unitialized value of a type for sql ifnull statements
func getDefault(typeName string) string {
	switch typeName {
	case "byte", "rune", "uint", "int", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "0"
	case "string":
		return "''"
	default:
		panic(fmt.Errorf("couldn't determine default value for provided type %s", typeName))
	}
}

func findInMask(fieldMask []string, field string) bool {
	for _, v := range fieldMask {
		if v == field {
			return true
		}
	}
	return false
}
