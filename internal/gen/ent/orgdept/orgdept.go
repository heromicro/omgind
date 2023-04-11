// Code generated by ent, DO NOT EDIT.

package orgdept

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orgdept type in the database.
	Label = "org_dept"
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
	// FieldTreeID holds the string denoting the tree_id field in the database.
	FieldTreeID = "tree_id"
	// FieldTreeLevel holds the string denoting the tree_level field in the database.
	FieldTreeLevel = "tree_level"
	// FieldTreeLeft holds the string denoting the tree_left field in the database.
	FieldTreeLeft = "tree_left"
	// FieldTreeRight holds the string denoting the tree_right field in the database.
	FieldTreeRight = "tree_right"
	// FieldIsLeaf holds the string denoting the is_leaf field in the database.
	FieldIsLeaf = "is_leaf"
	// FieldTreePath holds the string denoting the tree_path field in the database.
	FieldTreePath = "t_path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldMergeName holds the string denoting the merge_name field in the database.
	FieldMergeName = "mname"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "pid"
	// FieldIsReal holds the string denoting the is_real field in the database.
	FieldIsReal = "is_rl"
	// FieldIsShow holds the string denoting the is_show field in the database.
	FieldIsShow = "is_sh"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeOrgan holds the string denoting the organ edge name in mutations.
	EdgeOrgan = "organ"
	// EdgeStaffs holds the string denoting the staffs edge name in mutations.
	EdgeStaffs = "staffs"
	// Table holds the table name of the orgdept in the database.
	Table = "org_depts"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "org_depts"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "pid"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "org_depts"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "pid"
	// OrganTable is the table that holds the organ relation/edge.
	OrganTable = "org_depts"
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
	StaffsColumn = "dept_id"
)

// Columns holds all SQL columns for orgdept fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsActive,
	FieldMemo,
	FieldTreeID,
	FieldTreeLevel,
	FieldTreeLeft,
	FieldTreeRight,
	FieldIsLeaf,
	FieldTreePath,
	FieldName,
	FieldCode,
	FieldMergeName,
	FieldOrgID,
	FieldParentID,
	FieldIsReal,
	FieldIsShow,
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
	// DefaultIsLeaf holds the default value on creation for the "is_leaf" field.
	DefaultIsLeaf bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// MergeNameValidator is a validator for the "merge_name" field. It is called by the builders before save.
	MergeNameValidator func(string) error
	// OrgIDValidator is a validator for the "org_id" field. It is called by the builders before save.
	OrgIDValidator func(string) error
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// DefaultIsReal holds the default value on creation for the "is_real" field.
	DefaultIsReal bool
	// DefaultIsShow holds the default value on creation for the "is_show" field.
	DefaultIsShow bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Order defines the ordering method for the OrgDept queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsDel orders the results by the is_del field.
func ByIsDel(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsDel, opts...).ToFunc()
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

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByTreeID orders the results by the tree_id field.
func ByTreeID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTreeID, opts...).ToFunc()
}

// ByTreeLevel orders the results by the tree_level field.
func ByTreeLevel(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTreeLevel, opts...).ToFunc()
}

// ByTreeLeft orders the results by the tree_left field.
func ByTreeLeft(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTreeLeft, opts...).ToFunc()
}

// ByTreeRight orders the results by the tree_right field.
func ByTreeRight(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTreeRight, opts...).ToFunc()
}

// ByIsLeaf orders the results by the is_leaf field.
func ByIsLeaf(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsLeaf, opts...).ToFunc()
}

// ByTreePath orders the results by the tree_path field.
func ByTreePath(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTreePath, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByMergeName orders the results by the merge_name field.
func ByMergeName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldMergeName, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByIsReal orders the results by the is_real field.
func ByIsReal(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsReal, opts...).ToFunc()
}

// ByIsShow orders the results by the is_show field.
func ByIsShow(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIsShow, opts...).ToFunc()
}

// ByCreator orders the results by the creator field.
func ByCreator(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreator, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrganField orders the results by organ field.
func ByOrganField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganStep(), sql.OrderByField(field, opts...))
	}
}

// ByStaffsCount orders the results by staffs count.
func ByStaffsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStaffsStep(), opts...)
	}
}

// ByStaffs orders the results by staffs terms.
func ByStaffs(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStaffsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newOrganStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganTable, OrganColumn),
	)
}
func newStaffsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StaffsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StaffsTable, StaffsColumn),
	)
}
