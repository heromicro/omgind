// Code generated by ent, DO NOT EDIT.

package orgposition

import (
	"time"
)

const (
	// Label holds the string label denoting the orgposition type in the database.
	Label = "org_position"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// EdgeOrgan holds the string denoting the organ edge name in mutations.
	EdgeOrgan = "organ"
	// EdgeStaffs holds the string denoting the staffs edge name in mutations.
	EdgeStaffs = "staffs"
	// Table holds the table name of the orgposition in the database.
	Table = "org_positions"
	// OrganTable is the table that holds the organ relation/edge.
	OrganTable = "org_positions"
	// OrganInverseTable is the table name for the OrgOrgan entity.
	// It exists in this package in order to avoid circular dependency with the "orgorgan" package.
	OrganInverseTable = "org_organs"
	// OrganColumn is the table column denoting the organ relation/edge.
	OrganColumn = "org_id"
	// StaffsTable is the table that holds the staffs relation/edge.
	StaffsTable = "org_staffs"
	// StaffsInverseTable is the table name for the OrgStaff entity.
	// It exists in this package in order to avoid circular dependency with the "orgstaff" package.
	StaffsInverseTable = "org_staffs"
	// StaffsColumn is the table column denoting the staffs relation/edge.
	StaffsColumn = "posi_id"
)

// Columns holds all SQL columns for orgposition fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsActive,
	FieldMemo,
	FieldName,
	FieldCode,
	FieldOrgID,
	FieldCreator,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDel holds the default value on creation for the "is_del" field.
	DefaultIsDel bool
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// OrgIDValidator is a validator for the "org_id" field. It is called by the builders before save.
	OrgIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
