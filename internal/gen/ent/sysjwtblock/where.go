// Code generated by ent, DO NOT EDIT.

package sysjwtblock

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// Jwt applies equality check predicate on the "jwt" field. It's identical to JwtEQ.
func Jwt(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJwt), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemo), v))
	})
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMemo), v...))
	})
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMemo), v...))
	})
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemo), v))
	})
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemo), v))
	})
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemo), v))
	})
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemo), v))
	})
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemo), v))
	})
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemo), v))
	})
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemo), v))
	})
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemo), v))
	})
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsActive), v))
	})
}

// JwtEQ applies the EQ predicate on the "jwt" field.
func JwtEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJwt), v))
	})
}

// JwtNEQ applies the NEQ predicate on the "jwt" field.
func JwtNEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJwt), v))
	})
}

// JwtIn applies the In predicate on the "jwt" field.
func JwtIn(vs ...string) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJwt), v...))
	})
}

// JwtNotIn applies the NotIn predicate on the "jwt" field.
func JwtNotIn(vs ...string) predicate.SysJwtBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJwt), v...))
	})
}

// JwtGT applies the GT predicate on the "jwt" field.
func JwtGT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJwt), v))
	})
}

// JwtGTE applies the GTE predicate on the "jwt" field.
func JwtGTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJwt), v))
	})
}

// JwtLT applies the LT predicate on the "jwt" field.
func JwtLT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJwt), v))
	})
}

// JwtLTE applies the LTE predicate on the "jwt" field.
func JwtLTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJwt), v))
	})
}

// JwtContains applies the Contains predicate on the "jwt" field.
func JwtContains(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldJwt), v))
	})
}

// JwtHasPrefix applies the HasPrefix predicate on the "jwt" field.
func JwtHasPrefix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldJwt), v))
	})
}

// JwtHasSuffix applies the HasSuffix predicate on the "jwt" field.
func JwtHasSuffix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldJwt), v))
	})
}

// JwtEqualFold applies the EqualFold predicate on the "jwt" field.
func JwtEqualFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldJwt), v))
	})
}

// JwtContainsFold applies the ContainsFold predicate on the "jwt" field.
func JwtContainsFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldJwt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysJwtBlock) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysJwtBlock) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
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
func Not(p predicate.SysJwtBlock) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(func(s *sql.Selector) {
		p(s.Not())
	})
}
