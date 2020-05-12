package pbsql

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const nullSelectField = "ifnull(%s.%s, %s) as %s, "
const selectField = "%s.%s, "
const andPredicate = " AND %s.%s"
const orPredicate = " OR %s.%s"
const strComparison = " LIKE :%s"
const valComparison = " = :%s"
const queryCore = "%sFROM %s%s%s"
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

type field struct {
	value reflect.Value
	self reflect.StructField
	typeStr string
	table string
	isNullable bool
	shouldIgnore bool
	isPrimaryKey bool
	hasForeignKey bool
	name string
}

func parseReflection(val reflect.Value, i int, target string) *field {
	self := val.Type().Field(i)
	value := val.Field(i)
	return &field{
		value: value,
		table: target,
		self: self,
		typeStr: value.Type().Name(),
		isNullable: self.Tag.Get("nullable") == "y",
		isPrimaryKey: self.Tag.Get("primary_key") != "",
		shouldIgnore: self.Tag.Get("ignore") != "",
		hasForeignKey: self.Tag.Get("foreign_key") != "",
		name: self.Tag.Get("db"),
	}
}

type queryBuilder struct {
	Core strings.Builder
	Joins strings.Builder
	Fields strings.Builder
	Predicate strings.Builder
	Columns strings.Builder
	Values strings.Builder
}

func (qb *queryBuilder) writeSelectField(f *field) {
	if f.isNullable {
		fmt.Fprintf(&qb.Fields, nullSelectField, f.table, f.name, getDefault(f.typeStr), f.name)
	} else {
		fmt.Fprintf(&qb.Fields, selectField, f.table, f.name)
	}
}

func (qb *queryBuilder) writePredicate(f *field, fieldMask []string, predicateStr string) {
	if notDefault(f.typeStr, f.value.Interface()) || findInMask(fieldMask, f.self.Name) {
		fmt.Fprintf(&qb.Predicate, predicateStr, f.table, f.name)
		if f.typeStr == "string" {
			fmt.Fprintf(&qb.Predicate, strComparison, f.name)
		} else {
			fmt.Fprintf(&qb.Predicate,  valComparison, f.name)
		}
	}
}

func (qb *queryBuilder) writeOrPredicate(f *field, fieldMask []string) {
	qb.writePredicate(f, fieldMask, orPredicate)
}

func (qb *queryBuilder) writeAndPredicate(f *field, fieldMask []string) {
	qb.writePredicate(f, fieldMask, andPredicate)
}

func (qb *queryBuilder) getReadResult(table string, v *reflect.Value) string {
	qb.handleDateRange(table, v)
	fmt.Fprintf(&qb.Core, queryCore, qb.Fields.String(), table, qb.Joins.String(), qb.Predicate.String())
	qb.handleOrder(v)
	return strings.Replace(qb.Core.String(), ", FROM", " FROM", 1)
}

func (qb *queryBuilder) getUpdateResult() string {
	qb.Core.WriteString(qb.Predicate.String())
	return strings.Replace(qb.Core.String(), ", WHERE", " WHERE", 1)	
}

func (qb *queryBuilder) handleOrder(v *reflect.Value) {
	orderBy := v.FieldByName("OrderBy")
	orderDir := v.FieldByName("OrderDir")
	
	if orderBy.CanAddr() && orderBy.String() != "" {
		orderStr := fmt.Sprintf(" order by %s", orderBy.String())
		if orderDir.CanAddr() && orderDir.String() != "" {
			orderStr = fmt.Sprintf("%s %s", orderStr, orderDir.String())
		} else {
			orderStr = fmt.Sprintf("%s asc", orderStr)
		}
		fmt.Fprint(&qb.Core, orderStr)
	}
}

func (qb *queryBuilder) handleDateRange(target string, t *reflect.Value) {
	var dateTarget string
	dateRange := t.FieldByName("DateRange")
	dateTargetField := t.FieldByName("DateTarget")
	if dateTargetField.CanAddr() {
		dateTarget = dateTargetField.String()
		if dateTarget == "" {
			dateRangeTypeField, ok := t.Type().FieldByName("DateRange");
			if ok {
				dateTarget = dateRangeTypeField.Tag.Get("date_target");
			}
		}
	}

	if dateTarget != "" && dateRange.CanAddr() && reflect.TypeOf(dateRange.Interface()).Kind() == reflect.Slice {
		for i := 0; i < dateRange.Len(); i = i + 2 {
			fmt.Fprintf(
				&qb.Predicate,
				" AND %s.%s %s '%v'",
				target,
				toSnakeCase(dateTarget),
				dateRange.Index(i),
				dateRange.Index(i + 1),
			)
		}
	}
}

func (qb *queryBuilder) handleForeignKey(f *field) {
	foreignKey := f.self.Tag.Get("foreign_key")
	foreignTable := f.self.Tag.Get("foreign_table")
	localName := f.self.Tag.Get("local_name")

	if localName == "" {
		localName = foreignKey
	}

	related := reflect.Indirect(f.value)
	if related.CanAddr() && foreignKey != "" && foreignTable != "" {
		for j := 0; j < related.NumField(); j++ {
			field := parseReflection(related, j, foreignTable)
			
			if field.name != "" && field.value.CanInterface() && notDefault(field.typeStr, field.value.Interface()) {
				fmt.Fprintf(&qb.Predicate, " AND %s.%s", field.table, field.name)
				if field.typeStr == "string" {
					fmt.Fprintf(&qb.Predicate, " LIKE '%s'", field.value)
				} else {
					fmt.Fprintf(&qb.Predicate, " = %v", field.value)
				}
			}
		}
		fmt.Fprintf(
			&qb.Joins,
			" LEFT JOIN %s on %s.%s = %s.%s",
			foreignTable,
			foreignTable,
			foreignKey, 
			f.table, 
			localName,
		)
	}
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
	case "bool":
		return fieldVal.(bool)
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

func toSnakeCase(str string) string {
  snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
  snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
  return strings.ToLower(snake)
}