// Code generated by ent, DO NOT EDIT.

package sysrolemenu

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sysrolemenu type in the database.
	Label = "sys_role_menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldMenuID holds the string denoting the menu_id field in the database.
	FieldMenuID = "menu_id"
	// FieldActionID holds the string denoting the action_id field in the database.
	FieldActionID = "action_id"
	// Table holds the table name of the sysrolemenu in the database.
	Table = "sys_role_menus"
)

// Columns holds all SQL columns for sysrolemenu fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldRoleID,
	FieldMenuID,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	RoleIDValidator func(string) error
	// MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	MenuIDValidator func(string) error
	// ActionIDValidator is a validator for the "action_id" field. It is called by the builders before save.
	ActionIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the SysRoleMenu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsDel orders the results by the is_del field.
func ByIsDel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDel, opts...).ToFunc()
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

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByMenuID orders the results by the menu_id field.
func ByMenuID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuID, opts...).ToFunc()
}

// ByActionID orders the results by the action_id field.
func ByActionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionID, opts...).ToFunc()
}