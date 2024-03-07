package pbsql

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSelectField(t *testing.T) {
	testCases := []struct {
		name     string
		field    *field
		validate func(*queryBuilder)
	}{
		{
			name: "nullable, string",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "string",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(`groups`.`value`, '') as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "nullable, float32",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "float32",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(`groups`.`value`, 0.0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "nullable, float64",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "float64",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(`groups`.`value`, 0.0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "nullable, bool",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "bool",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(`groups`.`value`, 0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "nullable, other number",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "int",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(`groups`.`value`, 0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "non nullable",
			field: &field{
				isNullable: false,
				table:      "groups",
				name:       "value",
				typeStr:    "string",
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "`groups`.`value`,", strings.TrimSpace(qb.Fields.String()))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qb := &queryBuilder{
				Fields: strings.Builder{},
			}

			qb.writeSelectField(tc.field)
			tc.validate(qb)
		})
	}
}

func TestWriteSelectFunc(t *testing.T) {
	testCases := []struct {
		name     string
		field    *field
		validate func(*queryBuilder)
	}{
		{
			name: "string",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "string",
				selectFunc: &selectFuncData{
					name:    "concat",
					argName: "value",
				},
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(concat(`groups`.`value`), '') as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "float32",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "float32",
				selectFunc: &selectFuncData{
					name:    "concat",
					argName: "value",
				},
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(concat(`groups`.`value`), 0.0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "float64",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "float64",
				selectFunc: &selectFuncData{
					name:    "concat",
					argName: "value",
				},
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(concat(`groups`.`value`), 0.0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "bool",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "bool",
				selectFunc: &selectFuncData{
					name:    "concat",
					argName: "value",
				},
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(concat(`groups`.`value`), 0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
		{
			name: "other numbers",
			field: &field{
				isNullable: true,
				table:      "groups",
				name:       "value",
				typeStr:    "int",
				selectFunc: &selectFuncData{
					name:    "concat",
					argName: "value",
				},
			},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "ifnull(concat(`groups`.`value`), 0) as value,", strings.TrimSpace(qb.Fields.String()))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qb := &queryBuilder{
				Fields: strings.Builder{},
			}

			qb.writeSelectFunc(tc.field)
			tc.validate(qb)
		})
	}
}

func TestWritePredicate(t *testing.T) {
	testCases := []struct {
		name      string
		field     *field
		fieldMask []string
		format    string
		validate  func(*queryBuilder)
	}{
		{
			name: "and predicate, not default, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, not default, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, not default, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.6),
				isMultiValue: false,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` = :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.85),
				isMultiValue: false,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` = :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.65),
				isMultiValue: false,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` = :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.68),
				isMultiValue: false,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` = :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qb := &queryBuilder{
				Predicate: strings.Builder{},
			}

			qb.writePredicate(tc.field, tc.fieldMask, tc.format)
			tc.validate(qb)
		})
	}
}

func TestWriteNotPredicate(t *testing.T) {
	testCases := []struct {
		name      string
		field     *field
		fieldMask []string
		format    string
		validate  func(*queryBuilder)
	}{
		{
			name: "and predicate, not default, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` NOT IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, not default, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` NOT LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, not default, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.6),
				isMultiValue: false,
			},
			format: andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` != :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` NOT IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` NOT LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "and predicate, in mask, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.85),
				isMultiValue: false,
			},
			fieldMask: []string{"value"},
			format:    andPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "AND `groups`.`value` != :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` NOT IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` NOT LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, not default, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.65),
				isMultiValue: false,
			},
			format: orPredicate,
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` != :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, multi value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: true,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` NOT IN (test)", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, single value, string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "string",
				value:        reflect.ValueOf("test"),
				isMultiValue: false,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` NOT LIKE :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
		{
			name: "or predicate, in mask, single value, non string",
			field: &field{
				table:        "groups",
				name:         "value",
				typeStr:      "float64",
				value:        reflect.ValueOf(17.68),
				isMultiValue: false,
			},
			format:    orPredicate,
			fieldMask: []string{"value"},
			validate: func(qb *queryBuilder) {
				assert.Equal(t, "OR `groups`.`value` != :value", strings.TrimSpace(qb.Predicate.String()))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qb := &queryBuilder{
				Predicate: strings.Builder{},
			}

			qb.writeNotPredicate(tc.field, tc.fieldMask, tc.format)
			tc.validate(qb)
		})
	}
}
