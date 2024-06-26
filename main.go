package pbsql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

// BuildCountQuery_OLD is deprecated is a convenience wrapper for getting the result count of a query already generated by pbsql
// value based, does not affect the initially supplied query string
func BuildCountQuery_OLD(selectQry string) string {
	return "SELECT COUNT(*) as count FROM (" + selectQry + ") as count"
}

// BuildCreateQuery accepts a target table name and a protobuf message and attempts to build a valid SQL insert statement for use
// with sqlx.Named, ignoring any struct fields with default values. Fields must be tagged with `db:""` in order to be
// included in the result string.
func BuildCreateQuery(target string, source interface{}) (string, []interface{}, error) {
	t := reflect.ValueOf(source).Elem()
	var qb queryBuilder
	fmt.Fprintf(&qb.Columns, "INSERT INTO %s (", target)
	qb.Values.WriteString("(")

	for i := 0; i < t.NumField(); i++ {
		field := parseReflection(t, i, target)
		if field.value.CanInterface() {
			if notDefault(field.typeStr, field.value.Interface()) && field.name != "" {
				if i != 0 {
					qb.Columns.WriteString(", ")
					qb.Values.WriteString(", ")
				}
				fmt.Fprintf(&qb.Columns, "%s.%s", target, field.name)
				fmt.Fprintf(&qb.Values, ":%s", field.name)
			}
		}
	}
	qb.Values.WriteString(")")
	fmt.Fprintf(&qb.Columns, ") VALUES %s", qb.Values.String())
	result := strings.ReplaceAll(qb.Columns.String(), "(, ", "(")
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
	reflectedValue := reflect.ValueOf(source).Elem()
	var builder strings.Builder

	if _, hasIsActive := reflectedValue.Type().FieldByName("IsActive"); hasIsActive {
		fmt.Fprintf(&builder, "UPDATE %s SET %s.is_active = 0 WHERE ", target, target)
	} else {
		fmt.Fprintf(&builder, "DELETE FROM %s WHERE ", target)
	}

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, target)
		if field.isPrimaryKey {
			fmt.Fprintf(&builder, "%s.%s = :%s", target, field.name, field.name)
			break
		}
	}

	return sqlx.Named(builder.String(), source)
}

// BuildSearchQuery builds a search query
func BuildSearchQuery(target string, source interface{}, searchPhrase string) (string, []interface{}, error) {
	var qb queryBuilder
	qb.Core.WriteString("SELECT ")
	qb.Predicate.WriteString(" WHERE true")
	reflectedValue := reflect.ValueOf(source).Elem()
	fieldMask := make([]string, 0)
	fields := make([]*field, 0)
	n := reflectedValue.NumField()

	for i := 0; i < n; i++ {
		field := parseReflection(reflectedValue, i, target)
		if field.selectFunc.ok {
			field.shouldIgnore = true
		}
		fields = append(fields, field)
		if field.name != "" && !field.shouldIgnore {
			if field.typeStr == "string" && field.value.String() == "" {
				fieldMask = append(fieldMask, field.self.Name)
			} else if field.value.CanAddr() {
				qb.writePredicate(field, fieldMask, andPredicate)
			}
		} else if field.selectFunc.ok {
			qb.writeSelectFunc(field)
		}
	}

	qb.Predicate.WriteString(" AND (")
	for i := 0; i < n; i++ {
		field := fields[i]
		if field.name != "" && !field.shouldIgnore {
			qb.writeSelectField(field)
			if field.value.CanAddr() {
				if field.typeStr == "string" && field.value.String() == "" {
					qb.writePredicate(field, fieldMask, orPredicate)
				}
			}
		}
		if field.hasForeignKey {
			qb.handleForeignKey(field)
		}
	}
	qb.Predicate.WriteString(")")
	/* here we choose to use the args returned from BuildReadQuery*/
	qry, falseArgs, err := sqlx.Named(qb.getReadResult(target, &reflectedValue), source)
	_, altArgs, _ := BuildReadQuery(target, source)
	searchArgs := getSearchArgs(len(falseArgs)-len(altArgs), searchPhrase)
	return qry, append(altArgs, searchArgs...), err
}

// BuildCountQuery is a convenience wrapper for getting the result count of a query already generated by pbsql
// value based, does not affect the initially supplied query string
func BuildCountQuery(target string, source interface{}, fieldMask ...string) (string, []interface{}, error) {
	reflectedValue := reflect.ValueOf(source).Elem()
	var qb queryBuilder
	qb.Core.WriteString("SELECT COUNT(*) ")
	qb.Predicate.WriteString(" WHERE TRUE")
	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, target)
		if field.value.CanInterface() {
			if field.name != "" && field.value.CanAddr() {
				qb.writePredicate(field, fieldMask, andPredicate)
			}
			if field.hasForeignKey {
				qb.handleForeignKey(field)
			}
		}
	}
	result := qb.getReadResult(target, &reflectedValue)
	return sqlx.Named(result, source)
}

// BuildReadQuery accepts a target table name and a protobuf message and attempts to build a valid SQL select statement,
// ignoring any struct fields with default values when writing predicates. Fields must be tagged with `db:""` in order to be
// included in the result string.
//
// Returns a SQL statement as a string, a slice of args to interpolate, and an error
func BuildReadQuery(target string, source interface{}, fieldMask ...string) (string, []interface{}, error) {
	reflectedValue := reflect.ValueOf(source).Elem()
	var qb queryBuilder
	qb.Core.WriteString("SELECT ")
	qb.Predicate.WriteString(" WHERE true")

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, target)
		if field.name != "" {
			if !field.shouldIgnore && !field.selectFunc.ok {
				qb.writeSelectField(field)
				if field.value.CanAddr() {
					qb.writePredicate(field, fieldMask, andPredicate)
				}
			} else if field.selectFunc.ok {
				qb.writeSelectFunc(field)
			} else if field.isMultiValue && field.value.CanAddr() {
				qb.writePredicate(field, fieldMask, andPredicate)
			}
		}
		if field.hasForeignKey {
			qb.handleForeignKey(field)
		}
	}
	qb.handleDateRange(target, &reflectedValue)
	result := qb.getReadResult(target, &reflectedValue)
	return sqlx.Named(result, source)
}

// BuildReadQueryWithNotList accepts a target table name and a protobuf message and attempts to build a valid SQL select statement,
// ignoring any struct fields with default values when writing predicates. Fields must be tagged with `db:""` in order to be
// included in the result string.
//
// Returns a SQL statement as a string, a slice of args to interpolate, and an error
func BuildReadQueryWithNotList(target string, source interface{}, notList []string, fieldMask ...string) (string, []interface{}, error) {
	reflectedValue := reflect.ValueOf(source).Elem()
	var qb queryBuilder
	qb.Core.WriteString("SELECT ")
	qb.Predicate.WriteString(" WHERE true")

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, target)
		if field.name != "" {
			if !field.shouldIgnore && !field.selectFunc.ok {
				qb.writeSelectField(field)
				if field.value.CanAddr() {
					if findInMask(notList, field.self.Name) {
						qb.writeNotPredicate(field, notList, andPredicate)
					} else {
						qb.writePredicate(field, fieldMask, andPredicate)
					}
				}
			} else if field.selectFunc.ok {
				qb.writeSelectFunc(field)
			} else if field.isMultiValue && field.value.CanAddr() {
				if findInMask(notList, field.self.Name) {
					qb.writeNotPredicate(field, notList, andPredicate)
				} else {
					qb.writePredicate(field, fieldMask, andPredicate)
				}
			}
		}
		if field.hasForeignKey {
			qb.handleForeignKey(field)
		}
	}
	qb.handleDateRange(target, &reflectedValue)
	result := qb.getReadResult(target, &reflectedValue)
	return sqlx.Named(result, source)
}

type Query struct {
	Target    string
	Source    interface{}
	NotList   []string
	FieldMask []string
	Collate   bool
}

func (q *Query) BuildRead() (string, []interface{}, error) {
	reflectedValue := reflect.ValueOf(q.Source).Elem()
	var qb queryBuilder
	qb.Core.WriteString("SELECT ")
	qb.Predicate.WriteString(" WHERE true")

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, q.Target)
		if field.name != "" {
			if !field.shouldIgnore && !field.selectFunc.ok {
				qb.writeSelectField(field)
				if field.value.CanAddr() {
					if findInMask(q.NotList, field.self.Name) {
						qb.writeNotPredicate(field, q.NotList, andPredicate)
					} else {
						qb.writePredicate(field, q.FieldMask, andPredicate)
					}
				}
			} else if field.selectFunc.ok {
				qb.writeSelectFunc(field)
			} else if field.isMultiValue && field.value.CanAddr() {
				if findInMask(q.NotList, field.self.Name) {
					qb.writeNotPredicate(field, q.NotList, andPredicate)
				} else {
					qb.writePredicate(field, q.FieldMask, andPredicate)
				}
			}
		}
		if field.hasForeignKey {
			qb.handleForeignKey(field)
		}
	}
	qb.handleDateRange(q.Target, &reflectedValue)
	result := qb.getReadResult(q.Target, &reflectedValue)
	return sqlx.Named(result, q.Source)
}

// BuildUpdateQuery accepts a target table name `target`, a struct `source`, and a list of struct fields `fieldMask`
// and attempts to build a valid sql update statement for use with sqlx.Named, ignoring any struct fields not present
// in `fieldMask`. Struct fields must also be tagged with `db:""`, and the primary key should be tagged as
// `primary_key` otherwise this function will return an invalid query
func BuildUpdateQuery(target string, source interface{}, fieldMask []string) (string, []interface{}, error) {
	reflectedValue := reflect.ValueOf(source).Elem()
	var qb queryBuilder
	fmt.Fprintf(&qb.Core, "UPDATE %s SET ", target)

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, target)

		if field.value.CanInterface() && field.name != "" {
			if field.isPrimaryKey {
				fmt.Fprintf(&qb.Predicate, "WHERE %s.%s = :%s", target, field.name, field.name)
			} else if findInMask(fieldMask, field.self.Name) && !field.shouldIgnore && field.value.CanInterface() {
				fmt.Fprintf(&qb.Core, "%s.%s = :%s, ", target, field.name, field.name)
			}
		}
	}

	return sqlx.Named(qb.getUpdateResult(), source)
}

// BuildRelatedReadQuery can be used to quickly build queries for many to one relationships
// This method is still experimental
func BuildRelatedReadQuery(source interface{}, foreignKey string, foreignValue interface{}) string {
	var qb queryBuilder
	reflectedValue := reflect.ValueOf(source).Elem()

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := parseReflection(reflectedValue, i, "")
		foreignKeyTag := field.self.Tag.Get("foreign_key")
		foreignTable := field.self.Tag.Get("foreign_table")
		localName := field.self.Tag.Get("local_name")

		if foreignKeyTag == foreignKey && foreignTable != "" && localName != "" {
			related := reflect.Indirect(field.value)
			fmt.Fprintf(&qb.Core, "SELECT ")
			if related.CanAddr() {
				for j := 0; j < related.NumField(); j++ {
					f := parseReflection(related, j, foreignTable)
					if f.name != "" && f.value.CanInterface() {
						qb.writeSelectField(f)
					}
				}
				fmt.Fprintf(
					&qb.Core,
					"%sFROM %s where %s.%s = %v",
					qb.Fields.String(),
					foreignTable,
					foreignTable,
					foreignKey,
					foreignValue,
				)
			}
		}
	}

	return strings.Replace(qb.Core.String(), ", FROM", " FROM", 1)
}
