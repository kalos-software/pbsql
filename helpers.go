package pbsql

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
  snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
  snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
  return strings.ToLower(snake)
}

func addOrder(t reflect.Value, core *strings.Builder) {
	orderBy := t.FieldByName("OrderBy")
	orderDir := t.FieldByName("OrderDir")
	
	if orderBy.CanAddr() && orderBy.String() != "" {
		orderStr := fmt.Sprintf(" order by %s", orderBy.String())
		if orderDir.CanAddr() && orderDir.String() != "" {
			orderStr = fmt.Sprintf("%s %s", orderStr, orderDir.String())
		} else {
			orderStr = fmt.Sprintf("%s asc", orderStr)
		}
		fmt.Fprint(core, orderStr)
	}
}

func handleForeignKeys(
	t reflect.Value,
	target string,
	typeField reflect.StructField,
	predicate *strings.Builder,
	joins *strings.Builder,
	) {
	foreignKey := typeField.Tag.Get("foreign_key")
	foreignTable := typeField.Tag.Get("foreign_table")
	localName := typeField.Tag.Get("local_name")

	if localName == "" {
		localName = foreignKey
	}

	related := reflect.Indirect(t)
	if related.CanAddr() && foreignKey != "" && foreignTable != "" {
		for j := 0; j < related.NumField(); j++ {
			relatedValField := related.Field(j)
			relatedTypeField := related.Type().Field(j)
			relatedTypeName := relatedValField.Type().Name()
			relatedDBName := relatedTypeField.Tag.Get("db")
			
			if relatedDBName != "" &&
				relatedValField.CanInterface() && 
				notDefault(relatedTypeName, relatedValField.Interface()) {
				fmt.Fprintf(predicate, " AND %s.%s", foreignTable, relatedDBName)
				if relatedTypeName == "string" {
					fmt.Fprintf(predicate, " LIKE '%s'", relatedValField)
				} else {
					fmt.Fprintf(predicate, " = %s", relatedValField)
				}
			}
		}
		fmt.Fprintf(
			joins,
			" LEFT JOIN %s on %s.%s = %s.%s",
			foreignTable,
			foreignTable,
			foreignKey, 
			target, 
			localName,
		)
	}
}

// addDateRange relies on the date_target proto tag to write a predicate based on
// supplied property and date strings
func oldddDateRange(target string, t reflect.Value, predicate *strings.Builder) {
	dateRange := t.FieldByName("DateRange")
	//dateTarget := t.FieldByName("DateTarget")
	// before proceeding, we need to be extra careful and make sure that dateRange is an addressable slice
	if  dateRange.CanAddr() && reflect.TypeOf(dateRange.Interface()).Kind() == reflect.Slice {
		for j := 0; j < dateRange.Len(); j = j + 1 {
			arr := dateRange.Index(j)
			if arr.CanAddr() && reflect.TypeOf(arr.Interface()).Kind() == reflect.Slice {
				dateTargetStr := arr.Index(0)
				dateStrArr := arr.Index(1)
				for i := 1; i < dateStrArr.Len(); i = i + 2 {
					fmt.Fprintf(
						predicate,
						" AND %s.%s %s '%v'",
						target,
						toSnakeCase(dateTargetStr.String()),
						dateStrArr.Index(i).String(),
						dateStrArr.Index(i + 1).String(),
					)
				}
			}
		}
	}
}
// a date range looks like this [string, [string, string, string?, string?]]
// our date range list is an array of [date range]

func addDateRange(target string, t reflect.Value, predicate *strings.Builder) {
	dateRange := t.FieldByName("DateRange")
	dateTarget := t.FieldByName("DateTarget").String()
	fmt.Println("normal", dateTarget)
	fmt.Println("camel", toSnakeCase(dateTarget))
	if dateTarget == "" {
		dateRangeTypeField, ok := t.Type().FieldByName("DateRange");
		if ok {
			dateTarget = dateRangeTypeField.Tag.Get("date_target");
		}
	}


	if dateRange.CanAddr() && reflect.TypeOf(dateRange.Interface()).Kind() == reflect.Slice {
		for i := 0; i < dateRange.Len(); i = i + 2 {
			fmt.Fprintf(
				predicate,
				" AND %s.%s %s '%v'",
				target,
				toSnakeCase(dateTarget),
				dateRange.Index(i),
				dateRange.Index(i + 1),
			)
		}
	}
}