// Code generated by ent, DO NOT EDIT.

package sysmenu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sysmenu type in the database.
	Label = "sys_menu"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldRouter holds the string denoting the router field in the database.
	FieldRouter = "router"
	// FieldIsShow holds the string denoting the is_show field in the database.
	FieldIsShow = "is_show"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "pid"
	// FieldParentPath holds the string denoting the parent_path field in the database.
	FieldParentPath = "ppath"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldIsLeaf holds the string denoting the is_leaf field in the database.
	FieldIsLeaf = "is_leaf"
	// FieldOpenBlank holds the string denoting the open_blank field in the database.
	FieldOpenBlank = "open_blank"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the sysmenu in the database.
	Table = "sys_menus"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "pid"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "pid"
)

// Columns holds all SQL columns for sysmenu fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldMemo,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsActive,
	FieldName,
	FieldIcon,
	FieldRouter,
	FieldIsShow,
	FieldParentID,
	FieldParentPath,
	FieldLevel,
	FieldIsLeaf,
	FieldOpenBlank,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(string) error
	// RouterValidator is a validator for the "router" field. It is called by the builders before save.
	RouterValidator func(string) error
	// DefaultIsShow holds the default value on creation for the "is_show" field.
	DefaultIsShow bool
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	ParentPathValidator func(string) error
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int32) error
	// DefaultIsLeaf holds the default value on creation for the "is_leaf" field.
	DefaultIsLeaf bool
	// DefaultOpenBlank holds the default value on creation for the "open_blank" field.
	DefaultOpenBlank bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the SysMenu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsDel orders the results by the is_del field.
func ByIsDel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDel, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIcon orders the results by the icon field.
func ByIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIcon, opts...).ToFunc()
}

// ByRouter orders the results by the router field.
func ByRouter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRouter, opts...).ToFunc()
}

// ByIsShow orders the results by the is_show field.
func ByIsShow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsShow, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByParentPath orders the results by the parent_path field.
func ByParentPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentPath, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByIsLeaf orders the results by the is_leaf field.
func ByIsLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsLeaf, opts...).ToFunc()
}

// ByOpenBlank orders the results by the open_blank field.
func ByOpenBlank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOpenBlank, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}