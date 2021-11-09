package pbsql

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const nullSelectField = "ifnull(%s.%s, %s) as %s, "
const selectField = "%s.%s, "
const selectFuncField = "ifnull(%s(%s.%s), %s) as %s, "
const andPredicate = " AND %s.%s"
const orPredicate = " OR %s.%s"
const strComparison = " LIKE :%s"
const notStrComparison = " NOT LIKE :%s"
const valComparison = " = :%s"
const notValComparison = " != :%s"
const queryCore = "%sFROM %s%s%s"
const isoDateFormat = "2006-01-02 15:04:05"
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
	hasSelectFunc bool
	selectFuncName string
	dateTarget string
	dateRange []string
	selectFunc *selectFuncData
	isMultiValue bool
	name string
}

type selectFuncData struct {
	ok bool
	name string
	argName string
}


func parseReflection(val reflect.Value, i int, target string) *field {
	self := val.Type().Field(i)
	value := val.Field(i)
	name := self.Tag.Get("db")
	if name == "" {
		name = self.Tag.Get("name")
	}
	foreignKey := self.Tag.Get("foreign_key")

	selectFuncName := self.Tag.Get("select_func")
	selectFunc := &selectFuncData{
		ok: selectFuncName != "",
		name: selectFuncName,
		argName: self.Tag.Get("func_arg_name"),
	}


	return &field{
		value: value,
		table: target,
		self: self,
		typeStr: value.Type().Name(),
		isNullable: self.Tag.Get("nullable") == "y",
		isPrimaryKey: self.Tag.Get("primary_key") != "",
		shouldIgnore: self.Tag.Get("ignore") != "",
		hasForeignKey: foreignKey != "",
		isMultiValue: self.Tag.Get("multi_value") != "",
		selectFunc: selectFunc,
		name: name,
	}
}

/** Field Tags
* __________________
* Standard Group    |
* db                | corresponding database property name 
* nullable          | y \ n if the field could be a null value
* primary_key       | y \ n if the field is the primary key of a table
* ignore            | y \ n if the field should be ignored (edge case)
* date_target       | default date field to use for date range searches
* __________________|
* Foreign Key Group |
* foreign_key       | corresponding database property name on the foreign entity table
* local_name        | name of key on local table (if "", uses foreign_key)
* foreign_table     | name of foreign table
* __________________|
* Function Group    |
* select_func       | name of SQL UDF used to select data
* func_arg_name     | name of the field which is used as an argument to the function
* __________________|
**/

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
		fmt.Fprintf(&qb.Fields, nullSelectField, f.table, f.name, getDefault(f.typeStr, f.name), f.name)
	} else {
		fmt.Fprintf(&qb.Fields, selectField, f.table, f.name)
	}
}

func (qb *queryBuilder) writeSelectFunc(f *field) {
	fmt.Fprintf(&qb.Fields, selectFuncField, f.selectFunc.name, f.table, f.selectFunc.argName, getDefault(f.typeStr, f.name), f.name)
}

func (qb *queryBuilder) writePredicate(f *field, fieldMask []string, predicateStr string) {
	if notDefault(f.typeStr, f.value.Interface()) || findInMask(fieldMask, f.self.Name) {
		fmt.Fprintf(&qb.Predicate, predicateStr, f.table, f.name)
		if f.isMultiValue && !f.value.IsZero() {
			fmt.Fprintf(&qb.Predicate, " IN (%s)", f.value)
		} else {
		if f.typeStr == "string" {
			fmt.Fprintf(&qb.Predicate, strComparison, f.name)
		} else {
			fmt.Fprintf(&qb.Predicate,  valComparison, f.name)
		}
	}
	}
}

func (qb *queryBuilder) writeNotPredicate(f *field, fieldMask []string, predicateStr string) {
	if notDefault(f.typeStr, f.value.Interface()) || findInMask(fieldMask, f.self.Name) {
		fmt.Fprintf(&qb.Predicate, predicateStr, f.table, f.name)
		if f.isMultiValue {
			fmt.Fprintf(&qb.Predicate, " NOT IN (%s)", f.value)
		} else {
		if f.typeStr == "string" {
			fmt.Fprintf(&qb.Predicate, notStrComparison, f.name)
		} else {
			fmt.Fprintf(&qb.Predicate,  notValComparison, f.name)
		}
	}
	}
}

/*
func (qb *queryBuilder) writeOrPredicate(f *field, fieldMask []string) {
	qb.writePredicate(f, fieldMask, orPredicate)
}

func (qb *queryBuilder) writeAndPredicate(f *field, fieldMask []string) {
	qb.writePredicate(f, fieldMask, andPredicate)
}*/

func (qb *queryBuilder) getReadResult(table string, v *reflect.Value) string {
	fmt.Fprintf(&qb.Core, queryCore, qb.Fields.String(), table, qb.Joins.String(), qb.Predicate.String())
	qb.handleGroupBy(v)
	qb.handleOrder(v)
	return strings.Replace(strings.Replace(qb.Core.String(), ", FROM", " FROM", 1), "( OR", "(", 1)
}

func (qb *queryBuilder) getUpdateResult() string {
	qb.Core.WriteString(qb.Predicate.String())
	return strings.Replace(qb.Core.String(), ", WHERE", " WHERE", 1)	
}

func (qb *queryBuilder) handleGroupBy(v *reflect.Value) {
	groupBy := v.FieldByName("GroupBy")
	if groupBy.CanAddr() && groupBy.String() != "" {
		groupStr := fmt.Sprintf(" group by %s", groupBy.String())
		fmt.Fprintf(&qb.Core, groupStr)
	}
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
	} else {
		for i := 0; i < v.NumField(); i++ {
			field := parseReflection(*v, i, "")
			if field.hasForeignKey {
				foreignKey := field.self.Tag.Get("foreign_key")
				foreignTable := field.self.Tag.Get("foreign_table")
				localName := field.self.Tag.Get("local_name")
			
				if localName == "" {
					localName = foreignKey
				}
			
				related := reflect.Indirect(field.value)
				if related.CanAddr() && foreignKey != "" && foreignTable != "" {
					orderBy := related.FieldByName("OrderBy")
					orderDir := related.FieldByName("OrderDir")
					if orderBy.CanAddr() && orderBy.String() != "" && orderBy.CanInterface() {
						orderStr := fmt.Sprintf(" order by %s.%s", foreignTable, orderBy.String())
						if orderDir.CanAddr() && orderDir.String() != "" {
							orderStr = fmt.Sprintf("%s %s", orderStr, orderDir.String())
						} else {
							orderStr = fmt.Sprintf("%s asc", orderStr)
						}
						fmt.Fprint(&qb.Core, orderStr)
					}
					return
				}
			}
		}
	}
}

func isEmptySlice(v reflect.Value) bool {
	return v.IsValid() && v.Kind() == reflect.Slice && v.Len() == 0
}

func (qb *queryBuilder) handleDateRange(target string, t *reflect.Value) {
	var dateTarget string
	dateRange := t.FieldByName("DateRange")
	dateTargetField := t.FieldByName("DateTarget")
	canIntefaceDateTarget := dateTargetField.IsValid() && dateTargetField.CanInterface()
	if dateRange.IsValid() {
		if canIntefaceDateTarget && dateRange.CanInterface() && dateRange.Kind() == reflect.Slice {
			if isEmptySlice(dateTargetField) || dateTargetField.Kind() == reflect.String {
				dateTarget = dateTargetField.String()
				if dateTarget == "" || dateTargetField.Kind() == reflect.Slice {
					dateRangeTypeField, ok := t.Type().FieldByName("DateRange");
					if ok {
						dateTarget = dateRangeTypeField.Tag.Get("date_target");
					}
				}

				if dateTarget != "" {
					for i := 0; i < dateRange.Len(); i = i + 2 {
						fmt.Fprintf(
							&qb.Predicate,
							" AND %s.%s %s %v",
							target,
							dateTarget,
							dateRange.Index(i),
							processDateRangeField(dateRange.Index(i + 1)),
						)
					}
				}
			} else {
				dateTargetSlice :=  dateTargetField.Interface().([]string)
				if len(dateTargetSlice) != 0 {
					for i := 0; i < dateRange.Len(); i = i + 2 {
						j := 0
						if len(dateTargetSlice) == 2 && i != 0 {
							j = 1
						}
						fmt.Fprintf(
							&qb.Predicate,
							" AND %s.%s %s '%v'",
							target,
							dateTargetSlice[j],
							dateRange.Index(i),
							dateRange.Index(i + 1),
						)
					}
				}
			}
		}
	}
}

func processDateRangeField(val reflect.Value) string {
	var dateStr string
	if strings.Contains(val.String(), "(") {
		dateStr = val.String()
	} else {
		dateStr = "'" + val.String() + "'"
	}
	testStr := strings.ToLower(dateStr)
	if strings.Contains(testStr, "select") || strings.Contains(testStr, "update") || strings.Contains(testStr, "create") || strings.Contains(testStr, "insert") {
		panic("someone is sql injecting us")
	}
	return dateStr
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
	if  related.IsValid() {
		qb.handleDateRange(foreignTable, &related)
	}
}

// `notDefault` checks if a value is set to it's unitialized default, e.g. whether or not an `int32` value is `0`
// returns `true` if not default.
func notDefault(typeName string, fieldVal interface{}) bool {
	switch typeName {
	case "uint":
		return fieldVal.(uint) != 0
	case "int":
		return fieldVal.(int) != 0
	case "uint8":
		return fieldVal.(uint8) != 0
	case "uint16":
		return fieldVal.(uint16) != 0
	case "uint32":
		return fieldVal.(uint32) != 0
	case "uint64":
		return fieldVal.(uint64) != 0
	case "int8":
		return fieldVal.(int8) != 0
	case "int16": 
		return fieldVal.(int16) != 0
	case "int32":
		return fieldVal.(int32) != 0
	case	"int64":
		return fieldVal.(int64) != 0
	case "byte":
		return fieldVal.(byte) != 0
	case "rune":
		return fieldVal.(rune) != 0 
	case "uintptr":
		return fieldVal.(uintptr) != 0
	case "float32":
		return fieldVal.(float32) != 0.0
	case "float64":
		return fieldVal.(float64) != 0.0
	case "complex64":
		return fieldVal.(complex64) != (0+0i)
	case "complex128":
		return fieldVal.(complex128) != (0+0i)
	case "string":
		return fieldVal.(string) != ""
	case "bool":
		return fieldVal.(bool)
	default:
		return false
	}
}

// `getDefault` returns the unitialized value of a type for sql ifnull statements
func getDefault(typeName string, fieldName string) string {
	switch typeName {
	case "byte", "rune", "uint", "int", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "0"
	case "string":
		lowerName := strings.ToLower(fieldName)
		if strings.Contains(lowerName, "date") || strings.Contains(lowerName, "timestamp") {
			return "'0001-01-01 00::00::00'"
		}
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

func getSearchArgs(n int, searchPhrase string) []interface{} {
	res := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		res = append(res, searchPhrase)
	}
	return res
}