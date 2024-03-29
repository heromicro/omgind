// Code generated by ent, DO NOT EDIT.

package sysmenu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsDel, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldMemo, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldSort, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsActive, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIcon, v))
}

// Router applies equality check predicate on the "router" field. It's identical to RouterEQ.
func Router(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldRouter, v))
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsShow, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentID, v))
}

// ParentPath applies equality check predicate on the "parent_path" field. It's identical to ParentPathEQ.
func ParentPath(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentPath, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldLevel, v))
}

// IsLeaf applies equality check predicate on the "is_leaf" field. It's identical to IsLeafEQ.
func IsLeaf(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsLeaf, v))
}

// OpenBlank applies equality check predicate on the "open_blank" field. It's identical to OpenBlankEQ.
func OpenBlank(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldOpenBlank, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIsDel, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldMemo, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIsActive, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldName, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldIcon, v))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldIcon, v))
}

// RouterEQ applies the EQ predicate on the "router" field.
func RouterEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldRouter, v))
}

// RouterNEQ applies the NEQ predicate on the "router" field.
func RouterNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldRouter, v))
}

// RouterIn applies the In predicate on the "router" field.
func RouterIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldRouter, vs...))
}

// RouterNotIn applies the NotIn predicate on the "router" field.
func RouterNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldRouter, vs...))
}

// RouterGT applies the GT predicate on the "router" field.
func RouterGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldRouter, v))
}

// RouterGTE applies the GTE predicate on the "router" field.
func RouterGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldRouter, v))
}

// RouterLT applies the LT predicate on the "router" field.
func RouterLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldRouter, v))
}

// RouterLTE applies the LTE predicate on the "router" field.
func RouterLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldRouter, v))
}

// RouterContains applies the Contains predicate on the "router" field.
func RouterContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldRouter, v))
}

// RouterHasPrefix applies the HasPrefix predicate on the "router" field.
func RouterHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldRouter, v))
}

// RouterHasSuffix applies the HasSuffix predicate on the "router" field.
func RouterHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldRouter, v))
}

// RouterEqualFold applies the EqualFold predicate on the "router" field.
func RouterEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldRouter, v))
}

// RouterContainsFold applies the ContainsFold predicate on the "router" field.
func RouterContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldRouter, v))
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsShow, v))
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIsShow, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldParentID, v))
}

// ParentIDContains applies the Contains predicate on the "parent_id" field.
func ParentIDContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldParentID, v))
}

// ParentIDHasPrefix applies the HasPrefix predicate on the "parent_id" field.
func ParentIDHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldParentID, v))
}

// ParentIDHasSuffix applies the HasSuffix predicate on the "parent_id" field.
func ParentIDHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldParentID, v))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldParentID))
}

// ParentIDEqualFold applies the EqualFold predicate on the "parent_id" field.
func ParentIDEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldParentID, v))
}

// ParentIDContainsFold applies the ContainsFold predicate on the "parent_id" field.
func ParentIDContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldParentID, v))
}

// ParentPathEQ applies the EQ predicate on the "parent_path" field.
func ParentPathEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldParentPath, v))
}

// ParentPathNEQ applies the NEQ predicate on the "parent_path" field.
func ParentPathNEQ(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldParentPath, v))
}

// ParentPathIn applies the In predicate on the "parent_path" field.
func ParentPathIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldParentPath, vs...))
}

// ParentPathNotIn applies the NotIn predicate on the "parent_path" field.
func ParentPathNotIn(vs ...string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldParentPath, vs...))
}

// ParentPathGT applies the GT predicate on the "parent_path" field.
func ParentPathGT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldParentPath, v))
}

// ParentPathGTE applies the GTE predicate on the "parent_path" field.
func ParentPathGTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldParentPath, v))
}

// ParentPathLT applies the LT predicate on the "parent_path" field.
func ParentPathLT(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldParentPath, v))
}

// ParentPathLTE applies the LTE predicate on the "parent_path" field.
func ParentPathLTE(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldParentPath, v))
}

// ParentPathContains applies the Contains predicate on the "parent_path" field.
func ParentPathContains(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContains(FieldParentPath, v))
}

// ParentPathHasPrefix applies the HasPrefix predicate on the "parent_path" field.
func ParentPathHasPrefix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasPrefix(FieldParentPath, v))
}

// ParentPathHasSuffix applies the HasSuffix predicate on the "parent_path" field.
func ParentPathHasSuffix(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldHasSuffix(FieldParentPath, v))
}

// ParentPathIsNil applies the IsNil predicate on the "parent_path" field.
func ParentPathIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldParentPath))
}

// ParentPathNotNil applies the NotNil predicate on the "parent_path" field.
func ParentPathNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldParentPath))
}

// ParentPathEqualFold applies the EqualFold predicate on the "parent_path" field.
func ParentPathEqualFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEqualFold(FieldParentPath, v))
}

// ParentPathContainsFold applies the ContainsFold predicate on the "parent_path" field.
func ParentPathContainsFold(v string) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldContainsFold(FieldParentPath, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int32) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldLTE(FieldLevel, v))
}

// IsLeafEQ applies the EQ predicate on the "is_leaf" field.
func IsLeafEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldIsLeaf, v))
}

// IsLeafNEQ applies the NEQ predicate on the "is_leaf" field.
func IsLeafNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldIsLeaf, v))
}

// IsLeafIsNil applies the IsNil predicate on the "is_leaf" field.
func IsLeafIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldIsLeaf))
}

// IsLeafNotNil applies the NotNil predicate on the "is_leaf" field.
func IsLeafNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldIsLeaf))
}

// OpenBlankEQ applies the EQ predicate on the "open_blank" field.
func OpenBlankEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldEQ(FieldOpenBlank, v))
}

// OpenBlankNEQ applies the NEQ predicate on the "open_blank" field.
func OpenBlankNEQ(v bool) predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNEQ(FieldOpenBlank, v))
}

// OpenBlankIsNil applies the IsNil predicate on the "open_blank" field.
func OpenBlankIsNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldIsNull(FieldOpenBlank))
}

// OpenBlankNotNil applies the NotNil predicate on the "open_blank" field.
func OpenBlankNotNil() predicate.SysMenu {
	return predicate.SysMenu(sql.FieldNotNull(FieldOpenBlank))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysMenu) predicate.SysMenu {
	return predicate.SysMenu(sql.NotPredicates(p))
}
