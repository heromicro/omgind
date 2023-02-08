// Code generated by entc, DO NOT EDIT.

package sysrolemenu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuID), v))
	})
}

// ActionID applies equality check predicate on the "action_id" field. It's identical to ActionIDEQ.
func ActionID(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActionID), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoleID), v))
	})
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoleID), v...))
	})
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoleID), v...))
	})
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoleID), v))
	})
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoleID), v))
	})
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoleID), v))
	})
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoleID), v))
	})
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoleID), v))
	})
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoleID), v))
	})
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoleID), v))
	})
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoleID), v))
	})
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoleID), v))
	})
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuID), v))
	})
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMenuID), v))
	})
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMenuID), v...))
	})
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMenuID), v...))
	})
}

// MenuIDGT applies the GT predicate on the "menu_id" field.
func MenuIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMenuID), v))
	})
}

// MenuIDGTE applies the GTE predicate on the "menu_id" field.
func MenuIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMenuID), v))
	})
}

// MenuIDLT applies the LT predicate on the "menu_id" field.
func MenuIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMenuID), v))
	})
}

// MenuIDLTE applies the LTE predicate on the "menu_id" field.
func MenuIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMenuID), v))
	})
}

// MenuIDContains applies the Contains predicate on the "menu_id" field.
func MenuIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMenuID), v))
	})
}

// MenuIDHasPrefix applies the HasPrefix predicate on the "menu_id" field.
func MenuIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMenuID), v))
	})
}

// MenuIDHasSuffix applies the HasSuffix predicate on the "menu_id" field.
func MenuIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMenuID), v))
	})
}

// MenuIDEqualFold applies the EqualFold predicate on the "menu_id" field.
func MenuIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMenuID), v))
	})
}

// MenuIDContainsFold applies the ContainsFold predicate on the "menu_id" field.
func MenuIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMenuID), v))
	})
}

// ActionIDEQ applies the EQ predicate on the "action_id" field.
func ActionIDEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActionID), v))
	})
}

// ActionIDNEQ applies the NEQ predicate on the "action_id" field.
func ActionIDNEQ(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActionID), v))
	})
}

// ActionIDIn applies the In predicate on the "action_id" field.
func ActionIDIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActionID), v...))
	})
}

// ActionIDNotIn applies the NotIn predicate on the "action_id" field.
func ActionIDNotIn(vs ...string) predicate.SysRoleMenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActionID), v...))
	})
}

// ActionIDGT applies the GT predicate on the "action_id" field.
func ActionIDGT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActionID), v))
	})
}

// ActionIDGTE applies the GTE predicate on the "action_id" field.
func ActionIDGTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActionID), v))
	})
}

// ActionIDLT applies the LT predicate on the "action_id" field.
func ActionIDLT(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActionID), v))
	})
}

// ActionIDLTE applies the LTE predicate on the "action_id" field.
func ActionIDLTE(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActionID), v))
	})
}

// ActionIDContains applies the Contains predicate on the "action_id" field.
func ActionIDContains(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActionID), v))
	})
}

// ActionIDHasPrefix applies the HasPrefix predicate on the "action_id" field.
func ActionIDHasPrefix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActionID), v))
	})
}

// ActionIDHasSuffix applies the HasSuffix predicate on the "action_id" field.
func ActionIDHasSuffix(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActionID), v))
	})
}

// ActionIDIsNil applies the IsNil predicate on the "action_id" field.
func ActionIDIsNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldActionID)))
	})
}

// ActionIDNotNil applies the NotNil predicate on the "action_id" field.
func ActionIDNotNil() predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldActionID)))
	})
}

// ActionIDEqualFold applies the EqualFold predicate on the "action_id" field.
func ActionIDEqualFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActionID), v))
	})
}

// ActionIDContainsFold applies the ContainsFold predicate on the "action_id" field.
func ActionIDContainsFold(v string) predicate.SysRoleMenu {
	return predicate.SysRoleMenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActionID), v))
	})
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
