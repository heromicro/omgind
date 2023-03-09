// Code generated by ent, DO NOT EDIT.

package sysjwtblock

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldIsDel, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldMemo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldIsActive, v))
}

// Jwt applies equality check predicate on the "jwt" field. It's identical to JwtEQ.
func Jwt(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldJwt, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldIsDel, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldContainsFold(FieldMemo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldIsActive, v))
}

// JwtEQ applies the EQ predicate on the "jwt" field.
func JwtEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEQ(FieldJwt, v))
}

// JwtNEQ applies the NEQ predicate on the "jwt" field.
func JwtNEQ(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNEQ(FieldJwt, v))
}

// JwtIn applies the In predicate on the "jwt" field.
func JwtIn(vs ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldIn(FieldJwt, vs...))
}

// JwtNotIn applies the NotIn predicate on the "jwt" field.
func JwtNotIn(vs ...string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldNotIn(FieldJwt, vs...))
}

// JwtGT applies the GT predicate on the "jwt" field.
func JwtGT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGT(FieldJwt, v))
}

// JwtGTE applies the GTE predicate on the "jwt" field.
func JwtGTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldGTE(FieldJwt, v))
}

// JwtLT applies the LT predicate on the "jwt" field.
func JwtLT(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLT(FieldJwt, v))
}

// JwtLTE applies the LTE predicate on the "jwt" field.
func JwtLTE(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldLTE(FieldJwt, v))
}

// JwtContains applies the Contains predicate on the "jwt" field.
func JwtContains(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldContains(FieldJwt, v))
}

// JwtHasPrefix applies the HasPrefix predicate on the "jwt" field.
func JwtHasPrefix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldHasPrefix(FieldJwt, v))
}

// JwtHasSuffix applies the HasSuffix predicate on the "jwt" field.
func JwtHasSuffix(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldHasSuffix(FieldJwt, v))
}

// JwtEqualFold applies the EqualFold predicate on the "jwt" field.
func JwtEqualFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldEqualFold(FieldJwt, v))
}

// JwtContainsFold applies the ContainsFold predicate on the "jwt" field.
func JwtContainsFold(v string) predicate.SysJwtBlock {
	return predicate.SysJwtBlock(sql.FieldContainsFold(FieldJwt, v))
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
