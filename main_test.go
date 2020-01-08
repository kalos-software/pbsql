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
	OrderBy string
	OrderDir string
}

var target TestStruct

var expectedCreateQry = "INSERT INTO test_table (test_table.date, test_table.geolocation_lat, test_table.geolocation_lng) VALUES (?, ?, ?)"

var expectedReadQry = "SELECT test_table.id, ifnull(test_table.name, '') as name, ifnull(test_table.date, '') as date, ifnull(test_table.geolocation_lat, 0.0) as geolocation_lat, ifnull(test_table.geolocation_lng, 0.0) as geolocation_lng, test_table.is_active FROM test_table WHERE true AND test_table.id = ? AND test_table.date LIKE ? AND test_table.geolocation_lat = ? AND test_table.geolocation_lng = ? order by id desc"

var expectedUpdateQry = "UPDATE test_table SET test_table.geolocation_lat = ?, test_table.geolocation_lng = ? WHERE test_table.id = ?"

var expectedDeleteQry = "UPDATE test_table SET test_table.is_active = ? WHERE test_table.id = ?"

func TestMain(m *testing.M) {
	target.ID = 1
	target.Date = "2019-01-01"
	target.GeoLat = 123.456
	target.GeoLng = 654.321
	target.OrderBy = "id"
	target.OrderDir = "desc"
	os.Exit(m.Run())
}

func TestBuildCreate(t *testing.T) {
	qry, _, err := BuildCreateQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildCreateQuery failed", err)
	}

	if qry != expectedCreateQry {
		t.Log("Got: ", qry)
		t.Fatal("Expected:", expectedCreateQry)
	}
}

func TestBuildRead(t *testing.T) {
	qry, _, err := BuildReadQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildReadQuery failed", err)
	}

	if qry != expectedReadQry {
		t.Log("Got:", qry)
		t.Fatal("Expected:", expectedReadQry)
	}
}

func TestBuildUpdate(t *testing.T) {
	fieldMask := []string{"GeoLat", "GeoLng"}
	qry, _, err := BuildUpdateQuery("test_table", &target, fieldMask)
	if err != nil {
		t.Fatal("BuildUpdateQuery failed", err)
	}

	if qry != expectedUpdateQry {
		t.Log("Got:", qry)
		t.Fatal("Expected:", expectedUpdateQry)
	}

}

func TestBuildDelete(t *testing.T) {
	qry, _, err := BuildDeleteQuery("test_table", &target)
	if err != nil {
		t.Fatal("BuildDeleteQuery failed", err)
	}

	if qry != expectedDeleteQry {
		t.Log("Got:", qry)
		t.Fatal("Expected:", expectedDeleteQry)
	}
}
