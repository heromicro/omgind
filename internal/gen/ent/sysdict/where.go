// Code generated by ent, DO NOT EDIT.

package sysdict

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// NameCn applies equality check predicate on the "name_cn" field. It's identical to NameCnEQ.
func NameCn(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameCn), v))
	})
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameEn), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemo), v))
	})
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func MemoNotIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func MemoGT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemo), v))
	})
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemo), v))
	})
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemo), v))
	})
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemo), v))
	})
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemo), v))
	})
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemo), v))
	})
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemo), v))
	})
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemo), v))
	})
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemo), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsActive), v))
	})
}

// NameCnEQ applies the EQ predicate on the "name_cn" field.
func NameCnEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameCn), v))
	})
}

// NameCnNEQ applies the NEQ predicate on the "name_cn" field.
func NameCnNEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameCn), v))
	})
}

// NameCnIn applies the In predicate on the "name_cn" field.
func NameCnIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNameCn), v...))
	})
}

// NameCnNotIn applies the NotIn predicate on the "name_cn" field.
func NameCnNotIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNameCn), v...))
	})
}

// NameCnGT applies the GT predicate on the "name_cn" field.
func NameCnGT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameCn), v))
	})
}

// NameCnGTE applies the GTE predicate on the "name_cn" field.
func NameCnGTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameCn), v))
	})
}

// NameCnLT applies the LT predicate on the "name_cn" field.
func NameCnLT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameCn), v))
	})
}

// NameCnLTE applies the LTE predicate on the "name_cn" field.
func NameCnLTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameCn), v))
	})
}

// NameCnContains applies the Contains predicate on the "name_cn" field.
func NameCnContains(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameCn), v))
	})
}

// NameCnHasPrefix applies the HasPrefix predicate on the "name_cn" field.
func NameCnHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameCn), v))
	})
}

// NameCnHasSuffix applies the HasSuffix predicate on the "name_cn" field.
func NameCnHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameCn), v))
	})
}

// NameCnEqualFold applies the EqualFold predicate on the "name_cn" field.
func NameCnEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameCn), v))
	})
}

// NameCnContainsFold applies the ContainsFold predicate on the "name_cn" field.
func NameCnContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameCn), v))
	})
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameEn), v))
	})
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameEn), v))
	})
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNameEn), v...))
	})
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.SysDict {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysDict(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNameEn), v...))
	})
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameEn), v))
	})
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameEn), v))
	})
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameEn), v))
	})
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameEn), v))
	})
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameEn), v))
	})
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameEn), v))
	})
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameEn), v))
	})
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameEn), v))
	})
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.SysDict {
	return predicate.SysDict(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameEn), v))
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
