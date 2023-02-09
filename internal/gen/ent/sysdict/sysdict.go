// Code generated by ent, DO NOT EDIT.

package sysdict

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdict type in the database.
	Label = "sys_dict"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldNameCn holds the string denoting the name_cn field in the database.
	FieldNameCn = "name_cn"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// Table holds the table name of the sysdict in the database.
	Table = "sys_dicts"
)

// Columns holds all SQL columns for sysdict fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldMemo,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldStatus,
	FieldNameCn,
	FieldNameEn,
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
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int16
	// NameCnValidator is a validator for the "name_cn" field. It is called by the builders before save.
	NameCnValidator func(string) error
	// NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	NameEnValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
