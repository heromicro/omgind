// Code generated by ent, DO NOT EDIT.

package syslogging

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// TraceID applies equality check predicate on the "trace_id" field. It's identical to TraceIDEQ.
func TraceID(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTraceID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// ErrorStack applies equality check predicate on the "error_stack" field. It's identical to ErrorStackEQ.
func ErrorStack(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorStack), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemo), v))
	})
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
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
func MemoNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
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
func MemoGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemo), v))
	})
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemo), v))
	})
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemo), v))
	})
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemo), v))
	})
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemo), v))
	})
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemo), v))
	})
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemo), v))
	})
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMemo)))
	})
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMemo)))
	})
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemo), v))
	})
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemo), v))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// LevelContains applies the Contains predicate on the "level" field.
func LevelContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLevel), v))
	})
}

// LevelHasPrefix applies the HasPrefix predicate on the "level" field.
func LevelHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLevel), v))
	})
}

// LevelHasSuffix applies the HasSuffix predicate on the "level" field.
func LevelHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLevel), v))
	})
}

// LevelEqualFold applies the EqualFold predicate on the "level" field.
func LevelEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLevel), v))
	})
}

// LevelContainsFold applies the ContainsFold predicate on the "level" field.
func LevelContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLevel), v))
	})
}

// TraceIDEQ applies the EQ predicate on the "trace_id" field.
func TraceIDEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTraceID), v))
	})
}

// TraceIDNEQ applies the NEQ predicate on the "trace_id" field.
func TraceIDNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTraceID), v))
	})
}

// TraceIDIn applies the In predicate on the "trace_id" field.
func TraceIDIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTraceID), v...))
	})
}

// TraceIDNotIn applies the NotIn predicate on the "trace_id" field.
func TraceIDNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTraceID), v...))
	})
}

// TraceIDGT applies the GT predicate on the "trace_id" field.
func TraceIDGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTraceID), v))
	})
}

// TraceIDGTE applies the GTE predicate on the "trace_id" field.
func TraceIDGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTraceID), v))
	})
}

// TraceIDLT applies the LT predicate on the "trace_id" field.
func TraceIDLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTraceID), v))
	})
}

// TraceIDLTE applies the LTE predicate on the "trace_id" field.
func TraceIDLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTraceID), v))
	})
}

// TraceIDContains applies the Contains predicate on the "trace_id" field.
func TraceIDContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTraceID), v))
	})
}

// TraceIDHasPrefix applies the HasPrefix predicate on the "trace_id" field.
func TraceIDHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTraceID), v))
	})
}

// TraceIDHasSuffix applies the HasSuffix predicate on the "trace_id" field.
func TraceIDHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTraceID), v))
	})
}

// TraceIDEqualFold applies the EqualFold predicate on the "trace_id" field.
func TraceIDEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTraceID), v))
	})
}

// TraceIDContainsFold applies the ContainsFold predicate on the "trace_id" field.
func TraceIDContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTraceID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTag), v))
	})
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTag), v...))
	})
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTag), v...))
	})
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTag), v))
	})
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTag), v))
	})
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTag), v))
	})
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTag), v))
	})
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTag), v))
	})
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTag), v))
	})
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTag), v))
	})
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTag), v))
	})
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTag), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVersion), v))
	})
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVersion), v))
	})
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVersion), v))
	})
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVersion), v))
	})
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVersion), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldData), v))
	})
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldData), v...))
	})
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldData), v...))
	})
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldData), v))
	})
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldData), v))
	})
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldData), v))
	})
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldData), v))
	})
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldData), v))
	})
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldData), v))
	})
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldData), v))
	})
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldData)))
	})
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldData)))
	})
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldData), v))
	})
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldData), v))
	})
}

// ErrorStackEQ applies the EQ predicate on the "error_stack" field.
func ErrorStackEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorStack), v))
	})
}

// ErrorStackNEQ applies the NEQ predicate on the "error_stack" field.
func ErrorStackNEQ(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldErrorStack), v))
	})
}

// ErrorStackIn applies the In predicate on the "error_stack" field.
func ErrorStackIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldErrorStack), v...))
	})
}

// ErrorStackNotIn applies the NotIn predicate on the "error_stack" field.
func ErrorStackNotIn(vs ...string) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldErrorStack), v...))
	})
}

// ErrorStackGT applies the GT predicate on the "error_stack" field.
func ErrorStackGT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldErrorStack), v))
	})
}

// ErrorStackGTE applies the GTE predicate on the "error_stack" field.
func ErrorStackGTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldErrorStack), v))
	})
}

// ErrorStackLT applies the LT predicate on the "error_stack" field.
func ErrorStackLT(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldErrorStack), v))
	})
}

// ErrorStackLTE applies the LTE predicate on the "error_stack" field.
func ErrorStackLTE(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldErrorStack), v))
	})
}

// ErrorStackContains applies the Contains predicate on the "error_stack" field.
func ErrorStackContains(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldErrorStack), v))
	})
}

// ErrorStackHasPrefix applies the HasPrefix predicate on the "error_stack" field.
func ErrorStackHasPrefix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldErrorStack), v))
	})
}

// ErrorStackHasSuffix applies the HasSuffix predicate on the "error_stack" field.
func ErrorStackHasSuffix(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldErrorStack), v))
	})
}

// ErrorStackEqualFold applies the EqualFold predicate on the "error_stack" field.
func ErrorStackEqualFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldErrorStack), v))
	})
}

// ErrorStackContainsFold applies the ContainsFold predicate on the "error_stack" field.
func ErrorStackContainsFold(v string) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldErrorStack), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.SysLogging {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysLogging(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysLogging) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysLogging) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
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
func Not(p predicate.SysLogging) predicate.SysLogging {
	return predicate.SysLogging(func(s *sql.Selector) {
		p(s.Not())
	})
}
