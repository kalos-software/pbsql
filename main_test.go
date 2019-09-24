package pbsql

import (
	"os"
	"testing"
)

type TestStruct struct {
	ID       int32   `db:"id" primary_key:"y"`
	Name     string  `db:"name" nullable:"y"`
	Date     string  `db:"date" nullable:"y"`
	GeoLat   float64 `db:"geolocation_lat" nullable:"y"`
	GeoLng   float64 `db:"geolocation_lng" nullable:"y"`
	IsActive int32   `db:"is_active"`
}

var target TestStruct

func TestMain(m *testing.M) {
	target.ID = 1
	target.Date = "2019-01-01"
	target.GeoLat = 123.456
	target.GeoLng = 654.321
	os.Exit(m.Run())
}

func TestBuildCreate(t *testing.T) {
	qry, _, err := BuildCreateQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildCreateQuery failed", err)
	}
	println(qry)
}

func TestBuildRead(t *testing.T) {
	qry, _, err := BuildReadQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildReadQuery failed", err)
	}
	println(qry)
}

func TestBuildUpdate(t *testing.T) {
	fieldMask := make(map[string]int32, 2)
	fieldMask["GeoLat"] = 0
	fieldMask["GeoLng"] = 1
	qry, _, err := BuildUpdateQuery("test_table", &target, fieldMask)
	if err != nil {
		t.Fatal("BuildUpdateQuery failed", err)
	}
	println(qry)
}

func TestBuildDelete(t *testing.T) {
	qry, _, err := BuildDeleteQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildDeleteQuery failed", err)
	}
	println(qry)
}
