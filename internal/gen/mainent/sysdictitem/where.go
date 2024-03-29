// Code generated by ent, DO NOT EDIT.

package sysdictitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContainsFold(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldIsDel, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldMemo, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldSort, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldIsActive, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldValue, v))
}

// DictID applies equality check predicate on the "dict_id" field. It's identical to DictIDEQ.
func DictID(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldDictID, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldIsDel, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContainsFold(FieldMemo, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldIsActive, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContainsFold(FieldLabel, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldValue, v))
}

// DictIDEQ applies the EQ predicate on the "dict_id" field.
func DictIDEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEQ(FieldDictID, v))
}

// DictIDNEQ applies the NEQ predicate on the "dict_id" field.
func DictIDNEQ(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNEQ(FieldDictID, v))
}

// DictIDIn applies the In predicate on the "dict_id" field.
func DictIDIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIn(FieldDictID, vs...))
}

// DictIDNotIn applies the NotIn predicate on the "dict_id" field.
func DictIDNotIn(vs ...string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotIn(FieldDictID, vs...))
}

// DictIDGT applies the GT predicate on the "dict_id" field.
func DictIDGT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGT(FieldDictID, v))
}

// DictIDGTE applies the GTE predicate on the "dict_id" field.
func DictIDGTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldGTE(FieldDictID, v))
}

// DictIDLT applies the LT predicate on the "dict_id" field.
func DictIDLT(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLT(FieldDictID, v))
}

// DictIDLTE applies the LTE predicate on the "dict_id" field.
func DictIDLTE(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldLTE(FieldDictID, v))
}

// DictIDContains applies the Contains predicate on the "dict_id" field.
func DictIDContains(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContains(FieldDictID, v))
}

// DictIDHasPrefix applies the HasPrefix predicate on the "dict_id" field.
func DictIDHasPrefix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasPrefix(FieldDictID, v))
}

// DictIDHasSuffix applies the HasSuffix predicate on the "dict_id" field.
func DictIDHasSuffix(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldHasSuffix(FieldDictID, v))
}

// DictIDIsNil applies the IsNil predicate on the "dict_id" field.
func DictIDIsNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldIsNull(FieldDictID))
}

// DictIDNotNil applies the NotNil predicate on the "dict_id" field.
func DictIDNotNil() predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldNotNull(FieldDictID))
}

// DictIDEqualFold applies the EqualFold predicate on the "dict_id" field.
func DictIDEqualFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldEqualFold(FieldDictID, v))
}

// DictIDContainsFold applies the ContainsFold predicate on the "dict_id" field.
func DictIDContainsFold(v string) predicate.SysDictItem {
	return predicate.SysDictItem(sql.FieldContainsFold(FieldDictID, v))
}

// HasDict applies the HasEdge predicate on the "dict" edge.
func HasDict() predicate.SysDictItem {
	return predicate.SysDictItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DictTable, DictColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDictWith applies the HasEdge predicate on the "dict" edge with a given conditions (other predicates).
func HasDictWith(preds ...predicate.SysDict) predicate.SysDictItem {
	return predicate.SysDictItem(func(s *sql.Selector) {
		step := newDictStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysDictItem) predicate.SysDictItem {
	return predicate.SysDictItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysDictItem) predicate.SysDictItem {
	return predicate.SysDictItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysDictItem) predicate.SysDictItem {
	return predicate.SysDictItem(sql.NotPredicates(p))
}
