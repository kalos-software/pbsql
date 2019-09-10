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
	createQry, args, err := BuildCreateQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildCreateQuery failed", err)
	}
	fmt.Println(createQry)
	fmt.Printf("%v", args)

	readQry, args, err := BuildReadQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildReadQuery failed", err)
	}
	fmt.Println(readQry)
	fmt.Printf("%v", args)

	fieldMask := make(map[string]int, 2)
	fieldMask["GeoLat"] = 0
	fieldMask["GeoLng"] = 1
	updateQry, args, err := BuildUpdateQuery("test_table", &target, fieldMask)
	if err != nil {
		t.Fatal("BuildUpdateQuery failed", err)
	}
	fmt.Println(updateQry)
	fmt.Printf("%v", args)

	deleteQry, args, err := BuildDeleteQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildDeleteQuery failed", err)
	}
	fmt.Println(deleteQry)
	fmt.Printf("%v", args)
}
