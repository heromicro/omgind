// Code generated by ent, DO NOT EDIT.

package orgorgan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldIsDel, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldSort, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldIsActive, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldMemo, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldName, v))
}

// Sname applies equality check predicate on the "sname" field. It's identical to SnameEQ.
func Sname(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldSname, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCode, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldOwnerID, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCreator, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldIsDel, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldIsActive, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldMemo, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldName, v))
}

// SnameEQ applies the EQ predicate on the "sname" field.
func SnameEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldSname, v))
}

// SnameNEQ applies the NEQ predicate on the "sname" field.
func SnameNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldSname, v))
}

// SnameIn applies the In predicate on the "sname" field.
func SnameIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldSname, vs...))
}

// SnameNotIn applies the NotIn predicate on the "sname" field.
func SnameNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldSname, vs...))
}

// SnameGT applies the GT predicate on the "sname" field.
func SnameGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldSname, v))
}

// SnameGTE applies the GTE predicate on the "sname" field.
func SnameGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldSname, v))
}

// SnameLT applies the LT predicate on the "sname" field.
func SnameLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldSname, v))
}

// SnameLTE applies the LTE predicate on the "sname" field.
func SnameLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldSname, v))
}

// SnameContains applies the Contains predicate on the "sname" field.
func SnameContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldSname, v))
}

// SnameHasPrefix applies the HasPrefix predicate on the "sname" field.
func SnameHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldSname, v))
}

// SnameHasSuffix applies the HasSuffix predicate on the "sname" field.
func SnameHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldSname, v))
}

// SnameIsNil applies the IsNil predicate on the "sname" field.
func SnameIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldSname))
}

// SnameNotNil applies the NotNil predicate on the "sname" field.
func SnameNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldSname))
}

// SnameEqualFold applies the EqualFold predicate on the "sname" field.
func SnameEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldSname, v))
}

// SnameContainsFold applies the ContainsFold predicate on the "sname" field.
func SnameContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldSname, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldCode, v))
}

// CodeIsNil applies the IsNil predicate on the "code" field.
func CodeIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldCode))
}

// CodeNotNil applies the NotNil predicate on the "code" field.
func CodeNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldCode))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldCode, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldOwnerID, v))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldOwnerID, v))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldOwnerID, v))
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldOwnerID))
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldOwnerID))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldOwnerID, v))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldOwnerID, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldLTE(FieldCreator, v))
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContains(FieldCreator, v))
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasPrefix(FieldCreator, v))
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldHasSuffix(FieldCreator, v))
}

// CreatorIsNil applies the IsNil predicate on the "creator" field.
func CreatorIsNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldIsNull(FieldCreator))
}

// CreatorNotNil applies the NotNil predicate on the "creator" field.
func CreatorNotNil() predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldNotNull(FieldCreator))
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldEqualFold(FieldCreator, v))
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.OrgOrgan {
	return predicate.OrgOrgan(sql.FieldContainsFold(FieldCreator, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgOrgan) predicate.OrgOrgan {
	return predicate.OrgOrgan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgOrgan) predicate.OrgOrgan {
	return predicate.OrgOrgan(func(s *sql.Selector) {
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
func Not(p predicate.OrgOrgan) predicate.OrgOrgan {
	return predicate.OrgOrgan(func(s *sql.Selector) {
		p(s.Not())
	})
}
