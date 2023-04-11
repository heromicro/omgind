// Code generated by ent, DO NOT EDIT.

package sysdict

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldNameCn holds the string denoting the name_cn field in the database.
	FieldNameCn = "name_cn"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// FieldTipe holds the string denoting the tipe field in the database.
	FieldTipe = "tipe"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// Table holds the table name of the sysdict in the database.
	Table = "sys_dicts"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "sys_dict_items"
	// ItemsInverseTable is the table name for the SysDictItem entity.
	// It exists in this package in order to avoid circular dependency with the "sysdictitem" package.
	ItemsInverseTable = "sys_dict_items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "dict_id"
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
	FieldIsActive,
	FieldNameCn,
	FieldNameEn,
	FieldTipe,
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// NameCnValidator is a validator for the "name_cn" field. It is called by the builders before save.
	NameCnValidator func(string) error
	// NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	NameEnValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Tipe defines the type for the "tipe" enum field.
type Tipe string

// TipeInt is the default value of the Tipe enum.
const DefaultTipe = TipeInt

// Tipe values.
const (
	TipeInt    Tipe = "int"
	TipeString Tipe = "string"
)

func (t Tipe) String() string {
	return string(t)
}

// TipeValidator is a validator for the "tipe" field enum values. It is called by the builders before save.
func TipeValidator(t Tipe) error {
	switch t {
	case TipeInt, TipeString:
		return nil
	default:
		return fmt.Errorf("sysdict: invalid enum value for tipe field: %q", t)
	}
}

// Order defines the ordering method for the SysDict queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsDel orders the results by the is_del field.
func ByIsDel(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsDel, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByNameCn orders the results by the name_cn field.
func ByNameCn(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldNameCn, opts...).ToFunc()
}

// ByNameEn orders the results by the name_en field.
func ByNameEn(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldNameEn, opts...).ToFunc()
}

// ByTipe orders the results by the tipe field.
func ByTipe(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTipe, opts...).ToFunc()
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}
