// Code generated by ent, DO NOT EDIT.

package orgorgan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orgorgan type in the database.
	Label = "org_organ"
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
	// FieldSname holds the string denoting the sname field in the database.
	FieldSname = "sname"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldIdenNo holds the string denoting the iden_no field in the database.
	FieldIdenNo = "iden_no"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldHaddrID holds the string denoting the haddr_id field in the database.
	FieldHaddrID = "haddr_id"
	// EdgeHaddr holds the string denoting the haddr edge name in mutations.
	EdgeHaddr = "haddr"
	// EdgeDepts holds the string denoting the depts edge name in mutations.
	EdgeDepts = "depts"
	// EdgeStaffs holds the string denoting the staffs edge name in mutations.
	EdgeStaffs = "staffs"
	// EdgePositions holds the string denoting the positions edge name in mutations.
	EdgePositions = "positions"
	// Table holds the table name of the orgorgan in the database.
	Table = "org_organs"
	// HaddrTable is the table that holds the haddr relation/edge.
	HaddrTable = "org_organs"
	// HaddrInverseTable is the table name for the SysAddress entity.
	// It exists in this package in order to avoid circular dependency with the "sysaddress" package.
	HaddrInverseTable = "sys_addresses"
	// HaddrColumn is the table column denoting the haddr relation/edge.
	HaddrColumn = "haddr_id"
	// DeptsTable is the table that holds the depts relation/edge.
	DeptsTable = "org_depts"
	// DeptsInverseTable is the table name for the OrgDept entity.
	// It exists in this package in order to avoid circular dependency with the "orgdept" package.
	DeptsInverseTable = "org_depts"
	// DeptsColumn is the table column denoting the depts relation/edge.
	DeptsColumn = "org_id"
	// StaffsTable is the table that holds the staffs relation/edge.
	StaffsTable = "org_staffs"
	// StaffsInverseTable is the table name for the OrgStaff entity.
	// It exists in this package in order to avoid circular dependency with the "orgstaff" package.
	StaffsInverseTable = "org_staffs"
	// StaffsColumn is the table column denoting the staffs relation/edge.
	StaffsColumn = "org_id"
	// PositionsTable is the table that holds the positions relation/edge.
	PositionsTable = "org_positions"
	// PositionsInverseTable is the table name for the OrgPosition entity.
	// It exists in this package in order to avoid circular dependency with the "orgposition" package.
	PositionsInverseTable = "org_positions"
	// PositionsColumn is the table column denoting the positions relation/edge.
	PositionsColumn = "org_id"
)

// Columns holds all SQL columns for orgorgan fields.
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
	FieldSname,
	FieldCode,
	FieldIdenNo,
	FieldOwnerID,
	FieldHaddrID,
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
	// SnameValidator is a validator for the "sname" field. It is called by the builders before save.
	SnameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// IdenNoValidator is a validator for the "iden_no" field. It is called by the builders before save.
	IdenNoValidator func(string) error
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(string) error
	// HaddrIDValidator is a validator for the "haddr_id" field. It is called by the builders before save.
	HaddrIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the OrgOrgan queries.
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

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySname orders the results by the sname field.
func BySname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSname, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByIdenNo orders the results by the iden_no field.
func ByIdenNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdenNo, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByHaddrID orders the results by the haddr_id field.
func ByHaddrID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHaddrID, opts...).ToFunc()
}

// ByHaddrField orders the results by haddr field.
func ByHaddrField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHaddrStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeptsCount orders the results by depts count.
func ByDeptsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeptsStep(), opts...)
	}
}

// ByDepts orders the results by depts terms.
func ByDepts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeptsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStaffsCount orders the results by staffs count.
func ByStaffsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStaffsStep(), opts...)
	}
}

// ByStaffs orders the results by staffs terms.
func ByStaffs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStaffsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPositionsCount orders the results by positions count.
func ByPositionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPositionsStep(), opts...)
	}
}

// ByPositions orders the results by positions terms.
func ByPositions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPositionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHaddrStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HaddrInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, HaddrTable, HaddrColumn),
	)
}
func newDeptsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeptsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeptsTable, DeptsColumn),
	)
}
func newStaffsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StaffsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StaffsTable, StaffsColumn),
	)
}
func newPositionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PositionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PositionsTable, PositionsColumn),
	)
}
