// Code generated by ent, DO NOT EDIT.

package sysrolemenu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldIsDel, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// ActionID applies equality check predicate on the "action_id" field. It's identical to ActionIDEQ.
func ActionID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldActionID, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldIsDel, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotNull(FieldDeletedAt))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContainsFold(FieldRoleID, v))
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldMenuID, v))
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldMenuID, vs...))
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldMenuID, vs...))
}

// MenuIDGT applies the GT predicate on the "menu_id" field.
func MenuIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldMenuID, v))
}

// MenuIDGTE applies the GTE predicate on the "menu_id" field.
func MenuIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldMenuID, v))
}

// MenuIDLT applies the LT predicate on the "menu_id" field.
func MenuIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldMenuID, v))
}

// MenuIDLTE applies the LTE predicate on the "menu_id" field.
func MenuIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldMenuID, v))
}

// MenuIDContains applies the Contains predicate on the "menu_id" field.
func MenuIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContains(FieldMenuID, v))
}

// MenuIDHasPrefix applies the HasPrefix predicate on the "menu_id" field.
func MenuIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasPrefix(FieldMenuID, v))
}

// MenuIDHasSuffix applies the HasSuffix predicate on the "menu_id" field.
func MenuIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasSuffix(FieldMenuID, v))
}

// MenuIDEqualFold applies the EqualFold predicate on the "menu_id" field.
func MenuIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEqualFold(FieldMenuID, v))
}

// MenuIDContainsFold applies the ContainsFold predicate on the "menu_id" field.
func MenuIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContainsFold(FieldMenuID, v))
}

// ActionIDEQ applies the EQ predicate on the "action_id" field.
func ActionIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEQ(FieldActionID, v))
}

// ActionIDNEQ applies the NEQ predicate on the "action_id" field.
func ActionIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNEQ(FieldActionID, v))
}

// ActionIDIn applies the In predicate on the "action_id" field.
func ActionIDIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIn(FieldActionID, vs...))
}

// ActionIDNotIn applies the NotIn predicate on the "action_id" field.
func ActionIDNotIn(vs ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotIn(FieldActionID, vs...))
}

// ActionIDGT applies the GT predicate on the "action_id" field.
func ActionIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGT(FieldActionID, v))
}

// ActionIDGTE applies the GTE predicate on the "action_id" field.
func ActionIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldGTE(FieldActionID, v))
}

// ActionIDLT applies the LT predicate on the "action_id" field.
func ActionIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLT(FieldActionID, v))
}

// ActionIDLTE applies the LTE predicate on the "action_id" field.
func ActionIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldLTE(FieldActionID, v))
}

// ActionIDContains applies the Contains predicate on the "action_id" field.
func ActionIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContains(FieldActionID, v))
}

// ActionIDHasPrefix applies the HasPrefix predicate on the "action_id" field.
func ActionIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasPrefix(FieldActionID, v))
}

// ActionIDHasSuffix applies the HasSuffix predicate on the "action_id" field.
func ActionIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldHasSuffix(FieldActionID, v))
}

// ActionIDIsNil applies the IsNil predicate on the "action_id" field.
func ActionIDIsNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldIsNull(FieldActionID))
}

// ActionIDNotNil applies the NotNil predicate on the "action_id" field.
func ActionIDNotNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldNotNull(FieldActionID))
}

// ActionIDEqualFold applies the EqualFold predicate on the "action_id" field.
func ActionIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldEqualFold(FieldActionID, v))
}

// ActionIDContainsFold applies the ContainsFold predicate on the "action_id" field.
func ActionIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(sql.FieldContainsFold(FieldActionID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysRoleMenu) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysRoleMenu) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
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
func Not(p predicate.SysRoleMenu) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		p(s.Not())
	})
}
