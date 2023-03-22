// Code generated by ent, DO NOT EDIT.

package sysdict

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContainsFold(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldIsDel, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldMemo, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldSort, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldIsActive, v))
}

// NameCn applies equality check predicate on the "name_cn" field. It's identical to NameCnEQ.
func NameCn(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldNameCn, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldNameEn, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldIsDel, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContainsFold(FieldMemo, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysDict {
	return predicate.SysDict(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldIsActive, v))
}

// NameCnEQ applies the EQ predicate on the "name_cn" field.
func NameCnEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldNameCn, v))
}

// NameCnNEQ applies the NEQ predicate on the "name_cn" field.
func NameCnNEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldNameCn, v))
}

// NameCnIn applies the In predicate on the "name_cn" field.
func NameCnIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldNameCn, vs...))
}

// NameCnNotIn applies the NotIn predicate on the "name_cn" field.
func NameCnNotIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldNameCn, vs...))
}

// NameCnGT applies the GT predicate on the "name_cn" field.
func NameCnGT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldNameCn, v))
}

// NameCnGTE applies the GTE predicate on the "name_cn" field.
func NameCnGTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldNameCn, v))
}

// NameCnLT applies the LT predicate on the "name_cn" field.
func NameCnLT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldNameCn, v))
}

// NameCnLTE applies the LTE predicate on the "name_cn" field.
func NameCnLTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldNameCn, v))
}

// NameCnContains applies the Contains predicate on the "name_cn" field.
func NameCnContains(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContains(FieldNameCn, v))
}

// NameCnHasPrefix applies the HasPrefix predicate on the "name_cn" field.
func NameCnHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasPrefix(FieldNameCn, v))
}

// NameCnHasSuffix applies the HasSuffix predicate on the "name_cn" field.
func NameCnHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasSuffix(FieldNameCn, v))
}

// NameCnEqualFold applies the EqualFold predicate on the "name_cn" field.
func NameCnEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEqualFold(FieldNameCn, v))
}

// NameCnContainsFold applies the ContainsFold predicate on the "name_cn" field.
func NameCnContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContainsFold(FieldNameCn, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(sql.FieldContainsFold(FieldNameEn, v))
}

// TipeEQ applies the EQ predicate on the "tipe" field.
func TipeEQ(v Tipe) predicate.SysDict {
	return predicate.SysDict(sql.FieldEQ(FieldTipe, v))
}

// TipeNEQ applies the NEQ predicate on the "tipe" field.
func TipeNEQ(v Tipe) predicate.SysDict {
	return predicate.SysDict(sql.FieldNEQ(FieldTipe, v))
}

// TipeIn applies the In predicate on the "tipe" field.
func TipeIn(vs ...Tipe) predicate.SysDict {
	return predicate.SysDict(sql.FieldIn(FieldTipe, vs...))
}

// TipeNotIn applies the NotIn predicate on the "tipe" field.
func TipeNotIn(vs ...Tipe) predicate.SysDict {
	return predicate.SysDict(sql.FieldNotIn(FieldTipe, vs...))
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.SysDictItem
		step.Edge.Schema = schemaConfig.SysDictItem
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.SysDictItem) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.SysDictItem
		step.Edge.Schema = schemaConfig.SysDictItem
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysDict) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysDict) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysDict) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		p(s.Not())
	})
}
