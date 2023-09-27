// Code generated by ent, DO NOT EDIT.

package sysmenuactionresource

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContainsFold(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldIsDel, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldSort, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldMemo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldIsActive, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldMethod, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldPath, v))
}

// ActionID applies equality check predicate on the "action_id" field. It's identical to ActionIDEQ.
func ActionID(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldActionID, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldIsDel, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldSort, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContainsFold(FieldMemo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldIsActive, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContainsFold(FieldMethod, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContainsFold(FieldPath, v))
}

// ActionIDEQ applies the EQ predicate on the "action_id" field.
func ActionIDEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEQ(FieldActionID, v))
}

// ActionIDNEQ applies the NEQ predicate on the "action_id" field.
func ActionIDNEQ(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNEQ(FieldActionID, v))
}

// ActionIDIn applies the In predicate on the "action_id" field.
func ActionIDIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldIn(FieldActionID, vs...))
}

// ActionIDNotIn applies the NotIn predicate on the "action_id" field.
func ActionIDNotIn(vs ...string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldNotIn(FieldActionID, vs...))
}

// ActionIDGT applies the GT predicate on the "action_id" field.
func ActionIDGT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGT(FieldActionID, v))
}

// ActionIDGTE applies the GTE predicate on the "action_id" field.
func ActionIDGTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldGTE(FieldActionID, v))
}

// ActionIDLT applies the LT predicate on the "action_id" field.
func ActionIDLT(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLT(FieldActionID, v))
}

// ActionIDLTE applies the LTE predicate on the "action_id" field.
func ActionIDLTE(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldLTE(FieldActionID, v))
}

// ActionIDContains applies the Contains predicate on the "action_id" field.
func ActionIDContains(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContains(FieldActionID, v))
}

// ActionIDHasPrefix applies the HasPrefix predicate on the "action_id" field.
func ActionIDHasPrefix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasPrefix(FieldActionID, v))
}

// ActionIDHasSuffix applies the HasSuffix predicate on the "action_id" field.
func ActionIDHasSuffix(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldHasSuffix(FieldActionID, v))
}

// ActionIDEqualFold applies the EqualFold predicate on the "action_id" field.
func ActionIDEqualFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldEqualFold(FieldActionID, v))
}

// ActionIDContainsFold applies the ContainsFold predicate on the "action_id" field.
func ActionIDContainsFold(v string) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.FieldContainsFold(FieldActionID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysMenuActionResource) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysMenuActionResource) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysMenuActionResource) predicate.SysMenuActionResource {
	return predicate.SysMenuActionResource(sql.NotPredicates(p))
}
