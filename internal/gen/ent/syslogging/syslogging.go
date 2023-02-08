// Code generated by entc, DO NOT EDIT.

package syslogging

import (
	"time"
)

const (
	// Label holds the string label denoting the syslogging type in the database.
	Label = "sys_logging"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldTraceID holds the string denoting the trace_id field in the database.
	FieldTraceID = "trace_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldErrorStack holds the string denoting the error_stack field in the database.
	FieldErrorStack = "error_stack"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// Table holds the table name of the syslogging in the database.
	Table = "sys_logging"
)

// Columns holds all SQL columns for syslogging fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldMemo,
	FieldLevel,
	FieldTraceID,
	FieldUserID,
	FieldTag,
	FieldVersion,
	FieldMessage,
	FieldData,
	FieldErrorStack,
	FieldCreatedAt,
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
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(string) error
	// TraceIDValidator is a validator for the "trace_id" field. It is called by the builders before save.
	TraceIDValidator func(string) error
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
	// VersionValidator is a validator for the "version" field. It is called by the builders before save.
	VersionValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
