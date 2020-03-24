package pbsql

import (
	"fmt"
	"reflect"
	"strings"
)

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

func addDateRange(target string, t reflect.Value, predicate *strings.Builder) {
	var dateTarget string
	dateRange := t.FieldByName("DateRange");
	dateRangeTypeField, ok := t.Type().FieldByName("DateRange");
	if ok {
		dateTarget = dateRangeTypeField.Tag.Get("date_target");
	}

	if dateTarget != "" && dateRange.CanAddr() && reflect.TypeOf(dateRange.Interface()).Kind() == reflect.Slice {
		for i := 0; i < dateRange.Len(); i = i + 2 {
			fmt.Fprintf(
				predicate,
				" AND %s.%s %s '%v'",
				target,
				dateTarget,
				dateRange.Index(i),
				dateRange.Index(i + 1),
			)
		}
	}
}