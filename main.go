package pbsql

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

// BuildCreateQuery accepts a target table name and a struct and attempts to build a valid SQL insert statement for use
// with sqlx.Named, ignoring any struct fields with default values. Fields must be tagged with `db:""` in order to be
// included in the result string.
func BuildCreateQuery(target string, source interface{}) (string, []interface{}, error) {
	t := reflect.ValueOf(source).Elem()
	var cols strings.Builder
	var vals strings.Builder
	cols.WriteString("INSERT INTO ")
	cols.WriteString(target)
	cols.WriteString(" (")
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
			cols.WriteString(tag)
			vals.WriteString(":")
			vals.WriteString(tag)
		}
	}
	cols.WriteString(") VALUES ")
	vals.WriteString(")")
	result := strings.ReplaceAll(cols.String()+vals.String(), "(, ", "(")
	return sqlx.Named(result, target)
}

// BuildDeleteQuery accepts a target table name and a struct and attempts to build a valid SQL delete statement for use
// with sqlx.Named by attempting to identify the struct field tagged as `primary_key:"y"`. This function returns a
// nullsafe query if nullable struct fields are properly tagged as `nullable:"y"`
func BuildDeleteQuery(target string, source interface{}) (string, []interface{}, error) {
	v := reflect.ValueOf(source).Elem()
	t := v.Type()
	var builder strings.Builder

	isActive, hasIsActive := t.FieldByName("IsActive")
	if hasIsActive {
		dbName := isActive.Tag.Get("db")
		builder.WriteString("UPDATE " + target + "SET ")
		builder.WriteString(dbName + " = :" + dbName)
	} else {
		builder.WriteString("DELETE FROM " + target + " WHERE ")
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		isPkey := typeField.Tag.Get("primary_key") != ""
		if isPkey {
			dbName := typeField.Tag.Get("db")
			builder.WriteString(dbName)
			builder.WriteString(" = :")
			builder.WriteString(dbName)
			break
		}
	}

	return sqlx.Named(builder.String(), target)
}

// BuildReadQuery accepts a target table name and a struct and attempts to build a valid SQL select statement for use
// with sqlx.Named, ignoring any struct fields with default values. Fields must be tagged with `db:""` in order to be
// included in the result string.
func BuildReadQuery(target string, source interface{}) (string, []interface{}, error) {
	nullHandler := "ifnull("
	if sqlDriver := os.Getenv("GRPC_SQL_DRIVER"); sqlDriver == "pgsql" {
		nullHandler = "coalesce("
	}

	t := reflect.ValueOf(source).Elem()

	var core strings.Builder
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

		if nullable != "" {
			fields.WriteString(nullHandler)
			fields.WriteString(dbName)
			fields.WriteString(", ")
			fields.WriteString(getDefault(typeName))
			fields.WriteString(") as ")
			fields.WriteString(dbName)
			fields.WriteString(", ")
		} else if dbName != "" {
			fields.WriteString(dbName)
			fields.WriteString(", ")
		}

		if valField.CanInterface() && notDefault(typeName, valField.Interface()) && dbName != "" {
			predicate.WriteString(" AND ")
			predicate.WriteString(dbName)
			if typeName == "string" {
				predicate.WriteString(" LIKE :")
				predicate.WriteString(dbName)
			} else {
				predicate.WriteString(" = :")
				predicate.WriteString(dbName)
			}
		}
	}

	core.WriteString(fields.String())
	core.WriteString("FROM ")
	core.WriteString(target)
	core.WriteString(predicate.String())
	result := strings.Replace(core.String(), ", FROM", " FROM", 1)
	return sqlx.Named(result, target)
}

// BuildUpdateQuery accepts a target table name `target`, a struct `source`, and a list of struct fields `fieldMask`
// and attempts to build a valid sql update statement for use with sqlx.Named, ignoring any struct fields not present
// in `fieldMask`. Struct fields must also be tagged with `db:""`, and the primary key should be tagged as
// `primary_key` otherwise this function will return an invalid query
func BuildUpdateQuery(target string, source interface{}, fieldMask map[string]int) (string, []interface{}, error) {
	v := reflect.ValueOf(source).Elem()
	t := v.Type()

	var builder strings.Builder
	builder.WriteString("UPDATE ")
	builder.WriteString(target)
	builder.WriteString(" SET ")

	var predicate strings.Builder
	for i := 0; i < v.NumField(); i++ {
		valField := v.Field(i)
		typeField := t.Field(i)
		dbName := typeField.Tag.Get("db")

		if valField.CanInterface() && dbName != "" {
			isPrimaryKey := typeField.Tag.Get("primary_key") != ""
			if isPrimaryKey {
				predicate.WriteString("WHERE ")
				predicate.WriteString(dbName)
				predicate.WriteString(" = :")
				predicate.WriteString(dbName)
			} else if _, ok := fieldMask[typeField.Name]; ok {
				builder.WriteString(dbName)
				builder.WriteString(" = :")
				builder.WriteString(dbName)
				builder.WriteString(", ")
			}
		}
	}

	result := strings.Replace(builder.String()+predicate.String(), ", WHERE", " WHERE", 1)
	return sqlx.Named(result, target)
}

// `notDefault` checks if a value is set to it's unitialized default, e.g. whether or not an `int32` value is `0`
// returns `true` if not default.
func notDefault(typeName string, fieldVal interface{}) bool {
	switch typeName {
	case "int32":
		return fieldVal.(int32) != 0
	case "float64":
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
