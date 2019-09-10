package pbsql

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	ID     int32   `db:"id" primary_key:"y"`
	Name   string  `db:"name" nullable:"y"`
	Date   string  `db:"date" nullable:"y"`
	GeoLat float64 `db:"geolocation_lat" nullable:"y"`
	GeoLng float64 `db:"geolocation_lng" nullable:"y"`
}

func TestBuilders(t *testing.T) {
	target := TestStruct{ID: 1, Name: "Hello!", Date: "2019-01-01", GeoLat: 123.456, GeoLng: 654.321}
	createQry := BuildCreateQuery("test_table", &target)
	fmt.Println(createQry)
	fmt.Println()

	readQry := BuildReadQuery("test_table", &target)
	fmt.Println(readQry)
	fmt.Println()

	fieldMask := make(map[string]int, 2)
	fieldMask["GeoLat"] = 0
	fieldMask["GeoLng"] = 1
	updateQry := BuildUpdateQuery("test_table", &target, fieldMask)
	fmt.Println(updateQry)
	fmt.Println()

	deleteQry := BuildDeleteQuery("test_table", &target)
	fmt.Println(deleteQry)
	fmt.Println()
}
