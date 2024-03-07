package pbsql

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	ID         int32   `db:"id" primary_key:"y"`
	Name       string  `db:"name" nullable:"y"`
	Date       string  `db:"date" nullable:"y"`
	GeoLat     float64 `db:"geolocation_lat" nullable:"y"`
	GeoLng     float64 `db:"geolocation_lng" nullable:"y"`
	IsActive   int32   `db:"is_active"`
	PropertyID int32   `db:"property_id" foreign_key:"property_id" foreign_table:"properties"`
	NotEquals  []string
	OrderBy    string
	OrderDir   string
}

var target TestStruct

var expectedCreateQry = "INSERT INTO test_table (test_table.date, test_table.geolocation_lat, test_table.geolocation_lng) VALUES (?, ?, ?)"

var expectedReadQry = "SELECT test_table.id, ifnull(test_table.name, '') as name, ifnull(test_table.date, '') as date, ifnull(test_table.geolocation_lat, 0.0) as geolocation_lat, ifnull(test_table.geolocation_lng, 0.0) as geolocation_lng, test_table.is_active FROM test_table WHERE true AND test_table.id = ? AND test_table.date LIKE ? AND test_table.geolocation_lat = ? AND test_table.geolocation_lng = ? order by id asc"

var expectedUpdateQry = "UPDATE test_table SET test_table.date = ?, test_table.geolocation_lat = ?, test_table.geolocation_lng = ? WHERE test_table.id = ?"

var expectedDeleteQry = "UPDATE test_table SET test_table.is_active = 0 WHERE test_table.id = ?"

var testEvent Event

var testUser User

var testTxn Transaction

var testTask Task

var testTSL TimesheetLine

func TestMain(m *testing.M) {
	target.ID = 0
	target.Date = ""
	target.GeoLat = 123.456
	target.GeoLng = 654.321
	target.IsActive = 0
	target.OrderBy = ""
	os.Exit(m.Run())
}

func TestBuildCreate(t *testing.T) {
	task := &Task{
		ExternalCode:     "ext_code",
		ReferenceNumber:  "ref_number",
		CreatorUserId:    12,
		BriefDescription: "brief_description",
		Details:          "details",
		Notes:            "notes",
	}

	query, args, err := BuildCreateQuery("task", task)
	if err != nil {
		t.Fatal("BuildCreateQuery failed", err.Error())
	}

	assert.Equal(t, "INSERT INTO task (task.external_code, task.reference_number, task.creator_user_id, task.brief_description, task.details, task.notes) VALUES (?, ?, ?, ?, ?, ?)", query)
	require.Equal(t, 6, len(args))
	assert.Equal(t, "ext_code", args[0])
	assert.Equal(t, "ref_number", args[1])
	assert.Equal(t, int32(12), args[2])
	assert.Equal(t, "brief_description", args[3])
	assert.Equal(t, "details", args[4])
	assert.Equal(t, "notes", args[5])
}

func TestBuildCount(t *testing.T) {
	task := &Task{
		ExternalCode:     "ext_code",
		ReferenceNumber:  "ref_number",
		CreatorUserId:    12,
		BriefDescription: "brief_description",
		Details:          "details",
		Notes:            "notes",
		IsActive:         1,
		ExternalId:       101253,
	}
	query, args, err := BuildCountQuery("task", task)
	if err != nil {
		t.Fatal("BuildCountQuery failed", err.Error())
	}

	assert.Equal(t, "SELECT COUNT(*) FROM task WHERE TRUE AND `task`.`external_id` = ? AND `task`.`external_code` LIKE ? AND `task`.`reference_number` LIKE ? AND `task`.`creator_user_id` = ? AND `task`.`brief_description` LIKE ? AND `task`.`details` LIKE ? AND `task`.`notes` LIKE ? AND `task`.`isActive` = ?", query)
	require.Equal(t, 8, len(args))
	assert.Equal(t, int32(101253), args[0])
	assert.Equal(t, "ext_code", args[1])
	assert.Equal(t, "ref_number", args[2])
	assert.Equal(t, int32(12), args[3])
	assert.Equal(t, "brief_description", args[4])
	assert.Equal(t, "details", args[5])
	assert.Equal(t, "notes", args[6])
	assert.Equal(t, int32(1), args[7])
}

func TestBuildSearchQuery(t *testing.T) {
	task := &Task{
		ExternalCode:     "ext_code",
		ReferenceNumber:  "ref_number",
		CreatorUserId:    12,
		BriefDescription: "brief_description",
		Details:          "details",
		Notes:            "notes",
	}
	query, args, err := BuildSearchQuery("task", task, "")
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, "SELECT ifnull(name_of_user(`task`.`external_id`), '') as owner_name, `task`.`task_id`, `task`.`external_id`, ifnull(`task`.`external_code`, '') as external_code, ifnull(`task`.`reference_number`, '') as reference_number, `task`.`creator_user_id`, `task`.`time_created`, ifnull(`task`.`time_due`, '') as time_due, ifnull(`task`.`brief_description`, '') as brief_description, ifnull(`task`.`details`, '') as details, ifnull(`task`.`notes`, '') as notes, `task`.`status_id`, `task`.`priority_id`, ifnull(`task`.`reference_url`, '') as reference_url, ifnull(`task`.`isActive`, 0) as isActive, `task`.`billable`, ifnull(`task`.`task_billable_type`, '') as task_billable_type, ifnull(`task`.`flat_rate`, 0.0) as flat_rate, ifnull(`task`.`hourly_start`, '') as hourly_start, ifnull(`task`.`hourly_end`, '') as hourly_end, ifnull(`task`.`address`, '') as address, ifnull(`task`.`order_num`, '') as order_num, `task`.`spiff_amount`, ifnull(`task`.`spiff_jobNumber`, '') as spiff_jobNumber, ifnull(`task`.`spiff_type_id`, 0) as spiff_type_id, ifnull(`task`.`spiff_address`, '') as spiff_address, ifnull(`task`.`toolpurchase_date`, '0001-01-01 00:00:00') as toolpurchase_date, `task`.`toolpurchase_cost`, ifnull(`task`.`toolpurchase_file`, '') as toolpurchase_file, ifnull(`task`.`admin_action_id`, 0) as admin_action_id, ifnull(`task`.`date_performed`, '0001-01-01 00:00:00') as date_performed, ifnull(`task`.`spiff_tool_id`, '') as spiff_tool_id, ifnull(`task`.`spiff_tool_closeout_date`, '0001-01-01 00:00:00') as spiff_tool_closeout_date FROM task WHERE true AND `task`.`external_code` LIKE ? AND `task`.`reference_number` LIKE ? AND `task`.`creator_user_id` = ? AND `task`.`brief_description` LIKE ? AND `task`.`details` LIKE ? AND `task`.`notes` LIKE ? AND ( `task`.`time_created` LIKE ? OR `task`.`time_due` LIKE ? OR `task`.`reference_url` LIKE ? OR `task`.`task_billable_type` LIKE ? OR `task`.`hourly_start` LIKE ? OR `task`.`hourly_end` LIKE ? OR `task`.`address` LIKE ? OR `task`.`order_num` LIKE ? OR `task`.`spiff_jobNumber` LIKE ? OR `task`.`spiff_address` LIKE ? OR `task`.`toolpurchase_date` LIKE ? OR `task`.`toolpurchase_file` LIKE ? OR `task`.`date_performed` LIKE ? OR `task`.`spiff_tool_id` LIKE ? OR `task`.`spiff_tool_closeout_date` LIKE ?)", query)
	require.Equal(t, 21, len(args))
	assert.Equal(t, "ext_code", args[0])
	assert.Equal(t, "ref_number", args[1])
	assert.Equal(t, int32(12), args[2])
	assert.Equal(t, "brief_description", args[3])
	assert.Equal(t, "details", args[4])
	assert.Equal(t, "notes", args[5])
}

func TestBuildRelatedReadQuery(t *testing.T) {
	user := &User{
		Id:               8418,
		ServicesRendered: &ServicesRendered{},
	}
	// _ := BuildRelatedReadQuery(&testEvent, Relationship{ ForeignKey: "property_id", ForeignValue: testEvent.PropertyId})
	query := BuildRelatedReadQuery(user, "technician_user_id", user.Id)
	assert.Equal(t, "SELECT `services_rendered`.`sr_id`, ifnull(`services_rendered`.`event_id`, 0) as event_id, ifnull(`services_rendered`.`technician_user_id`, 0) as technician_user_id, `services_rendered`.`sr_name`, ifnull(`services_rendered`.`sr_materialsUsed`, '') as sr_materialsUsed, ifnull(`services_rendered`.`sr_serviceRendered`, '') as sr_serviceRendered, ifnull(`services_rendered`.`sr_techNotes`, '') as sr_techNotes, `services_rendered`.`sr_status`, `services_rendered`.`sr_datetime`, ifnull(`services_rendered`.`time_started`, '') as time_started, ifnull(`services_rendered`.`time_finished`, '') as time_finished, ifnull(`services_rendered`.`isactive`, 0) as isactive, ifnull(`services_rendered`.`hide_from_timesheet`, 0) as hide_from_timesheet, ifnull(`services_rendered`.`signature_id`, 0) as signature_id, ifnull(`services_rendered`.`signatureData`, '') as signatureData FROM services_rendered where services_rendered.technician_user_id = 8418", query)
}

func TestBuildUpdate(t *testing.T) {
	m := &TestStruct{
		ID:         12,
		Name:       "test_name",
		Date:       "2022-01-11",
		GeoLat:     1000,
		GeoLng:     7000,
		IsActive:   1,
		PropertyID: 78,
		NotEquals:  []string{"date"},
		OrderBy:    "id",
		OrderDir:   "desc",
	}
	fieldMask := []string{"is_active", "geolocation_lat", "geolocation_lng"}
	query, args, err := BuildUpdateQuery("test_table", m, fieldMask)
	if err != nil {
		t.Fatal("BuildUpdateQuery failed", err)
	}

	assert.Equal(t, "UPDATE test_table SET test_table.is_active = ? WHERE test_table.id = ?", query)
	require.Equal(t, 2, len(args))
	assert.Equal(t, int32(1), args[0])
	assert.Equal(t, int32(12), args[1])
}

func TestBuildDelete(t *testing.T) {
	m := &TestStruct{
		ID: 3,
	}
	query, args, err := BuildDeleteQuery("test_table", m)
	if err != nil {
		t.Fatal("BuildDeleteQuery failed", err)
	}

	assert.Equal(t, "UPDATE test_table SET test_table.is_active = 0 WHERE test_table.id = ?", query)
	require.Equal(t, 1, len(args))
	assert.Equal(t, int32(3), args[0])
}

func TestBuildReadQuery(t *testing.T) {
	task := &Task{
		ExternalCode:     "ext_code",
		ReferenceNumber:  "ref_number",
		CreatorUserId:    12,
		BriefDescription: "brief_description",
		Details:          "details",
		Notes:            "notes",
	}
	query, args, err := BuildReadQuery("tasks", task)
	require.NoError(t, err)

	assert.Equal(t, "SELECT `tasks`.`task_id`, `tasks`.`external_id`, ifnull(`tasks`.`external_code`, '') as external_code, ifnull(`tasks`.`reference_number`, '') as reference_number, `tasks`.`creator_user_id`, `tasks`.`time_created`, ifnull(`tasks`.`time_due`, '') as time_due, ifnull(`tasks`.`brief_description`, '') as brief_description, ifnull(`tasks`.`details`, '') as details, ifnull(`tasks`.`notes`, '') as notes, `tasks`.`status_id`, `tasks`.`priority_id`, ifnull(`tasks`.`reference_url`, '') as reference_url, ifnull(`tasks`.`isActive`, 0) as isActive, `tasks`.`billable`, ifnull(`tasks`.`task_billable_type`, '') as task_billable_type, ifnull(`tasks`.`flat_rate`, 0.0) as flat_rate, ifnull(`tasks`.`hourly_start`, '') as hourly_start, ifnull(`tasks`.`hourly_end`, '') as hourly_end, ifnull(`tasks`.`address`, '') as address, ifnull(`tasks`.`order_num`, '') as order_num, `tasks`.`spiff_amount`, ifnull(`tasks`.`spiff_jobNumber`, '') as spiff_jobNumber, ifnull(`tasks`.`spiff_type_id`, 0) as spiff_type_id, ifnull(`tasks`.`spiff_address`, '') as spiff_address, ifnull(`tasks`.`toolpurchase_date`, '0001-01-01 00:00:00') as toolpurchase_date, `tasks`.`toolpurchase_cost`, ifnull(`tasks`.`toolpurchase_file`, '') as toolpurchase_file, ifnull(`tasks`.`admin_action_id`, 0) as admin_action_id, ifnull(`tasks`.`date_performed`, '0001-01-01 00:00:00') as date_performed, ifnull(`tasks`.`spiff_tool_id`, '') as spiff_tool_id, ifnull(name_of_user(`tasks`.`external_id`), '') as owner_name, ifnull(`tasks`.`spiff_tool_closeout_date`, '0001-01-01 00:00:00') as spiff_tool_closeout_date FROM tasks WHERE true AND `tasks`.`external_code` LIKE ? AND `tasks`.`reference_number` LIKE ? AND `tasks`.`creator_user_id` = ? AND `tasks`.`brief_description` LIKE ? AND `tasks`.`details` LIKE ? AND `tasks`.`notes` LIKE ?", query)
	require.Equal(t, 6, len(args))
	assert.Equal(t, "ext_code", args[0])
	assert.Equal(t, "ref_number", args[1])
	assert.Equal(t, int32(12), args[2])
	assert.Equal(t, "brief_description", args[3])
	assert.Equal(t, "details", args[4])
	assert.Equal(t, "notes", args[5])
}
