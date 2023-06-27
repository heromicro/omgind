// Code generated by ent, DO NOT EDIT.

package sysmenuactionresource

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sysmenuactionresource type in the database.
	Label = "sys_menu_action_resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldActionID holds the string denoting the action_id field in the database.
	FieldActionID = "action_id"
	// Table holds the table name of the sysmenuactionresource in the database.
	Table = "sys_menu_action_resources"
)

// Columns holds all SQL columns for sysmenuactionresource fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldSort,
	FieldMemo,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsActive,
	FieldMethod,
	FieldPath,
	FieldActionID,
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
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// MethodValidator is a validator for the "method" field. It is called by the builders before save.
	MethodValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// ActionIDValidator is a validator for the "action_id" field. It is called by the builders before save.
	ActionIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the SysMenuActionResource queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsDel orders the results by the is_del field.
func ByIsDel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDel, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByMethod orders the results by the method field.
func ByMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMethod, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByActionID orders the results by the action_id field.
func ByActionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionID, opts...).ToFunc()
}