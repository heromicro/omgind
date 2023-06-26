// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/syslogging"
)

// SysLoggingUpdate is the builder for updating SysLogging entities.
type SysLoggingUpdate struct {
	config
	hooks     []Hook
	mutation  *SysLoggingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysLoggingUpdate builder.
func (slu *SysLoggingUpdate) Where(ps ...predicate.SysLogging) *SysLoggingUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetUpdatedAt sets the "updated_at" field.
func (slu *SysLoggingUpdate) SetUpdatedAt(t time.Time) *SysLoggingUpdate {
	slu.mutation.SetUpdatedAt(t)
	return slu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (slu *SysLoggingUpdate) ClearUpdatedAt() *SysLoggingUpdate {
	slu.mutation.ClearUpdatedAt()
	return slu
}

// SetDeletedAt sets the "deleted_at" field.
func (slu *SysLoggingUpdate) SetDeletedAt(t time.Time) *SysLoggingUpdate {
	slu.mutation.SetDeletedAt(t)
	return slu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableDeletedAt(t *time.Time) *SysLoggingUpdate {
	if t != nil {
		slu.SetDeletedAt(*t)
	}
	return slu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (slu *SysLoggingUpdate) ClearDeletedAt() *SysLoggingUpdate {
	slu.mutation.ClearDeletedAt()
	return slu
}

// SetIsDel sets the "is_del" field.
func (slu *SysLoggingUpdate) SetIsDel(b bool) *SysLoggingUpdate {
	slu.mutation.SetIsDel(b)
	return slu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableIsDel(b *bool) *SysLoggingUpdate {
	if b != nil {
		slu.SetIsDel(*b)
	}
	return slu
}

// SetMemo sets the "memo" field.
func (slu *SysLoggingUpdate) SetMemo(s string) *SysLoggingUpdate {
	slu.mutation.SetMemo(s)
	return slu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableMemo(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetMemo(*s)
	}
	return slu
}

// ClearMemo clears the value of the "memo" field.
func (slu *SysLoggingUpdate) ClearMemo() *SysLoggingUpdate {
	slu.mutation.ClearMemo()
	return slu
}

// SetLevel sets the "level" field.
func (slu *SysLoggingUpdate) SetLevel(s string) *SysLoggingUpdate {
	slu.mutation.SetLevel(s)
	return slu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableLevel(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetLevel(*s)
	}
	return slu
}

// ClearLevel clears the value of the "level" field.
func (slu *SysLoggingUpdate) ClearLevel() *SysLoggingUpdate {
	slu.mutation.ClearLevel()
	return slu
}

// SetTraceID sets the "trace_id" field.
func (slu *SysLoggingUpdate) SetTraceID(s string) *SysLoggingUpdate {
	slu.mutation.SetTraceID(s)
	return slu
}

// SetNillableTraceID sets the "trace_id" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableTraceID(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetTraceID(*s)
	}
	return slu
}

// ClearTraceID clears the value of the "trace_id" field.
func (slu *SysLoggingUpdate) ClearTraceID() *SysLoggingUpdate {
	slu.mutation.ClearTraceID()
	return slu
}

// SetUserID sets the "user_id" field.
func (slu *SysLoggingUpdate) SetUserID(s string) *SysLoggingUpdate {
	slu.mutation.SetUserID(s)
	return slu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableUserID(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetUserID(*s)
	}
	return slu
}

// ClearUserID clears the value of the "user_id" field.
func (slu *SysLoggingUpdate) ClearUserID() *SysLoggingUpdate {
	slu.mutation.ClearUserID()
	return slu
}

// SetTag sets the "tag" field.
func (slu *SysLoggingUpdate) SetTag(s string) *SysLoggingUpdate {
	slu.mutation.SetTag(s)
	return slu
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableTag(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetTag(*s)
	}
	return slu
}

// ClearTag clears the value of the "tag" field.
func (slu *SysLoggingUpdate) ClearTag() *SysLoggingUpdate {
	slu.mutation.ClearTag()
	return slu
}

// SetVersion sets the "version" field.
func (slu *SysLoggingUpdate) SetVersion(s string) *SysLoggingUpdate {
	slu.mutation.SetVersion(s)
	return slu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableVersion(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetVersion(*s)
	}
	return slu
}

// ClearVersion clears the value of the "version" field.
func (slu *SysLoggingUpdate) ClearVersion() *SysLoggingUpdate {
	slu.mutation.ClearVersion()
	return slu
}

// SetMessage sets the "message" field.
func (slu *SysLoggingUpdate) SetMessage(s string) *SysLoggingUpdate {
	slu.mutation.SetMessage(s)
	return slu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableMessage(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetMessage(*s)
	}
	return slu
}

// ClearMessage clears the value of the "message" field.
func (slu *SysLoggingUpdate) ClearMessage() *SysLoggingUpdate {
	slu.mutation.ClearMessage()
	return slu
}

// SetErrorStack sets the "error_stack" field.
func (slu *SysLoggingUpdate) SetErrorStack(s string) *SysLoggingUpdate {
	slu.mutation.SetErrorStack(s)
	return slu
}

// SetNillableErrorStack sets the "error_stack" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableErrorStack(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetErrorStack(*s)
	}
	return slu
}

// ClearErrorStack clears the value of the "error_stack" field.
func (slu *SysLoggingUpdate) ClearErrorStack() *SysLoggingUpdate {
	slu.mutation.ClearErrorStack()
	return slu
}

// Mutation returns the SysLoggingMutation object of the builder.
func (slu *SysLoggingUpdate) Mutation() *SysLoggingMutation {
	return slu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *SysLoggingUpdate) Save(ctx context.Context) (int, error) {
	if err := slu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, slu.sqlSave, slu.mutation, slu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (slu *SysLoggingUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *SysLoggingUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *SysLoggingUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slu *SysLoggingUpdate) defaults() error {
	if _, ok := slu.mutation.UpdatedAt(); !ok && !slu.mutation.UpdatedAtCleared() {
		if syslogging.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized syslogging.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := syslogging.UpdateDefaultUpdatedAt()
		slu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (slu *SysLoggingUpdate) check() error {
	if v, ok := slu.mutation.Memo(); ok {
		if err := syslogging.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysLogging.memo": %w`, err)}
		}
	}
	if v, ok := slu.mutation.Level(); ok {
		if err := syslogging.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "SysLogging.level": %w`, err)}
		}
	}
	if v, ok := slu.mutation.TraceID(); ok {
		if err := syslogging.TraceIDValidator(v); err != nil {
			return &ValidationError{Name: "trace_id", err: fmt.Errorf(`ent: validator failed for field "SysLogging.trace_id": %w`, err)}
		}
	}
	if v, ok := slu.mutation.UserID(); ok {
		if err := syslogging.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "SysLogging.user_id": %w`, err)}
		}
	}
	if v, ok := slu.mutation.Tag(); ok {
		if err := syslogging.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "SysLogging.tag": %w`, err)}
		}
	}
	if v, ok := slu.mutation.Version(); ok {
		if err := syslogging.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "SysLogging.version": %w`, err)}
		}
	}
	if v, ok := slu.mutation.Message(); ok {
		if err := syslogging.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "SysLogging.message": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (slu *SysLoggingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysLoggingUpdate {
	slu.modifiers = append(slu.modifiers, modifiers...)
	return slu
}

func (slu *SysLoggingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := slu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(syslogging.Table, syslogging.Columns, sqlgraph.NewFieldSpec(syslogging.FieldID, field.TypeString))
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if slu.mutation.CreatedAtCleared() {
		_spec.ClearField(syslogging.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := slu.mutation.UpdatedAt(); ok {
		_spec.SetField(syslogging.FieldUpdatedAt, field.TypeTime, value)
	}
	if slu.mutation.UpdatedAtCleared() {
		_spec.ClearField(syslogging.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := slu.mutation.DeletedAt(); ok {
		_spec.SetField(syslogging.FieldDeletedAt, field.TypeTime, value)
	}
	if slu.mutation.DeletedAtCleared() {
		_spec.ClearField(syslogging.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := slu.mutation.IsDel(); ok {
		_spec.SetField(syslogging.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := slu.mutation.Memo(); ok {
		_spec.SetField(syslogging.FieldMemo, field.TypeString, value)
	}
	if slu.mutation.MemoCleared() {
		_spec.ClearField(syslogging.FieldMemo, field.TypeString)
	}
	if value, ok := slu.mutation.Level(); ok {
		_spec.SetField(syslogging.FieldLevel, field.TypeString, value)
	}
	if slu.mutation.LevelCleared() {
		_spec.ClearField(syslogging.FieldLevel, field.TypeString)
	}
	if value, ok := slu.mutation.TraceID(); ok {
		_spec.SetField(syslogging.FieldTraceID, field.TypeString, value)
	}
	if slu.mutation.TraceIDCleared() {
		_spec.ClearField(syslogging.FieldTraceID, field.TypeString)
	}
	if value, ok := slu.mutation.UserID(); ok {
		_spec.SetField(syslogging.FieldUserID, field.TypeString, value)
	}
	if slu.mutation.UserIDCleared() {
		_spec.ClearField(syslogging.FieldUserID, field.TypeString)
	}
	if value, ok := slu.mutation.Tag(); ok {
		_spec.SetField(syslogging.FieldTag, field.TypeString, value)
	}
	if slu.mutation.TagCleared() {
		_spec.ClearField(syslogging.FieldTag, field.TypeString)
	}
	if value, ok := slu.mutation.Version(); ok {
		_spec.SetField(syslogging.FieldVersion, field.TypeString, value)
	}
	if slu.mutation.VersionCleared() {
		_spec.ClearField(syslogging.FieldVersion, field.TypeString)
	}
	if value, ok := slu.mutation.Message(); ok {
		_spec.SetField(syslogging.FieldMessage, field.TypeString, value)
	}
	if slu.mutation.MessageCleared() {
		_spec.ClearField(syslogging.FieldMessage, field.TypeString)
	}
	if slu.mutation.DataCleared() {
		_spec.ClearField(syslogging.FieldData, field.TypeString)
	}
	if value, ok := slu.mutation.ErrorStack(); ok {
		_spec.SetField(syslogging.FieldErrorStack, field.TypeString, value)
	}
	if slu.mutation.ErrorStackCleared() {
		_spec.ClearField(syslogging.FieldErrorStack, field.TypeString)
	}
	_spec.AddModifiers(slu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syslogging.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	slu.mutation.done = true
	return n, nil
}

// SysLoggingUpdateOne is the builder for updating a single SysLogging entity.
type SysLoggingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysLoggingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (sluo *SysLoggingUpdateOne) SetUpdatedAt(t time.Time) *SysLoggingUpdateOne {
	sluo.mutation.SetUpdatedAt(t)
	return sluo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sluo *SysLoggingUpdateOne) ClearUpdatedAt() *SysLoggingUpdateOne {
	sluo.mutation.ClearUpdatedAt()
	return sluo
}

// SetDeletedAt sets the "deleted_at" field.
func (sluo *SysLoggingUpdateOne) SetDeletedAt(t time.Time) *SysLoggingUpdateOne {
	sluo.mutation.SetDeletedAt(t)
	return sluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableDeletedAt(t *time.Time) *SysLoggingUpdateOne {
	if t != nil {
		sluo.SetDeletedAt(*t)
	}
	return sluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sluo *SysLoggingUpdateOne) ClearDeletedAt() *SysLoggingUpdateOne {
	sluo.mutation.ClearDeletedAt()
	return sluo
}

// SetIsDel sets the "is_del" field.
func (sluo *SysLoggingUpdateOne) SetIsDel(b bool) *SysLoggingUpdateOne {
	sluo.mutation.SetIsDel(b)
	return sluo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableIsDel(b *bool) *SysLoggingUpdateOne {
	if b != nil {
		sluo.SetIsDel(*b)
	}
	return sluo
}

// SetMemo sets the "memo" field.
func (sluo *SysLoggingUpdateOne) SetMemo(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetMemo(s)
	return sluo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableMemo(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetMemo(*s)
	}
	return sluo
}

// ClearMemo clears the value of the "memo" field.
func (sluo *SysLoggingUpdateOne) ClearMemo() *SysLoggingUpdateOne {
	sluo.mutation.ClearMemo()
	return sluo
}

// SetLevel sets the "level" field.
func (sluo *SysLoggingUpdateOne) SetLevel(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetLevel(s)
	return sluo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableLevel(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetLevel(*s)
	}
	return sluo
}

// ClearLevel clears the value of the "level" field.
func (sluo *SysLoggingUpdateOne) ClearLevel() *SysLoggingUpdateOne {
	sluo.mutation.ClearLevel()
	return sluo
}

// SetTraceID sets the "trace_id" field.
func (sluo *SysLoggingUpdateOne) SetTraceID(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetTraceID(s)
	return sluo
}

// SetNillableTraceID sets the "trace_id" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableTraceID(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetTraceID(*s)
	}
	return sluo
}

// ClearTraceID clears the value of the "trace_id" field.
func (sluo *SysLoggingUpdateOne) ClearTraceID() *SysLoggingUpdateOne {
	sluo.mutation.ClearTraceID()
	return sluo
}

// SetUserID sets the "user_id" field.
func (sluo *SysLoggingUpdateOne) SetUserID(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetUserID(s)
	return sluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableUserID(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetUserID(*s)
	}
	return sluo
}

// ClearUserID clears the value of the "user_id" field.
func (sluo *SysLoggingUpdateOne) ClearUserID() *SysLoggingUpdateOne {
	sluo.mutation.ClearUserID()
	return sluo
}

// SetTag sets the "tag" field.
func (sluo *SysLoggingUpdateOne) SetTag(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetTag(s)
	return sluo
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableTag(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetTag(*s)
	}
	return sluo
}

// ClearTag clears the value of the "tag" field.
func (sluo *SysLoggingUpdateOne) ClearTag() *SysLoggingUpdateOne {
	sluo.mutation.ClearTag()
	return sluo
}

// SetVersion sets the "version" field.
func (sluo *SysLoggingUpdateOne) SetVersion(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetVersion(s)
	return sluo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableVersion(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetVersion(*s)
	}
	return sluo
}

// ClearVersion clears the value of the "version" field.
func (sluo *SysLoggingUpdateOne) ClearVersion() *SysLoggingUpdateOne {
	sluo.mutation.ClearVersion()
	return sluo
}

// SetMessage sets the "message" field.
func (sluo *SysLoggingUpdateOne) SetMessage(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetMessage(s)
	return sluo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableMessage(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetMessage(*s)
	}
	return sluo
}

// ClearMessage clears the value of the "message" field.
func (sluo *SysLoggingUpdateOne) ClearMessage() *SysLoggingUpdateOne {
	sluo.mutation.ClearMessage()
	return sluo
}

// SetErrorStack sets the "error_stack" field.
func (sluo *SysLoggingUpdateOne) SetErrorStack(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetErrorStack(s)
	return sluo
}

// SetNillableErrorStack sets the "error_stack" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableErrorStack(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetErrorStack(*s)
	}
	return sluo
}

// ClearErrorStack clears the value of the "error_stack" field.
func (sluo *SysLoggingUpdateOne) ClearErrorStack() *SysLoggingUpdateOne {
	sluo.mutation.ClearErrorStack()
	return sluo
}

// Mutation returns the SysLoggingMutation object of the builder.
func (sluo *SysLoggingUpdateOne) Mutation() *SysLoggingMutation {
	return sluo.mutation
}

// Where appends a list predicates to the SysLoggingUpdate builder.
func (sluo *SysLoggingUpdateOne) Where(ps ...predicate.SysLogging) *SysLoggingUpdateOne {
	sluo.mutation.Where(ps...)
	return sluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *SysLoggingUpdateOne) Select(field string, fields ...string) *SysLoggingUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated SysLogging entity.
func (sluo *SysLoggingUpdateOne) Save(ctx context.Context) (*SysLogging, error) {
	if err := sluo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sluo.sqlSave, sluo.mutation, sluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *SysLoggingUpdateOne) SaveX(ctx context.Context) *SysLogging {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *SysLoggingUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *SysLoggingUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sluo *SysLoggingUpdateOne) defaults() error {
	if _, ok := sluo.mutation.UpdatedAt(); !ok && !sluo.mutation.UpdatedAtCleared() {
		if syslogging.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized syslogging.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := syslogging.UpdateDefaultUpdatedAt()
		sluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sluo *SysLoggingUpdateOne) check() error {
	if v, ok := sluo.mutation.Memo(); ok {
		if err := syslogging.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysLogging.memo": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.Level(); ok {
		if err := syslogging.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "SysLogging.level": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.TraceID(); ok {
		if err := syslogging.TraceIDValidator(v); err != nil {
			return &ValidationError{Name: "trace_id", err: fmt.Errorf(`ent: validator failed for field "SysLogging.trace_id": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.UserID(); ok {
		if err := syslogging.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "SysLogging.user_id": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.Tag(); ok {
		if err := syslogging.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "SysLogging.tag": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.Version(); ok {
		if err := syslogging.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "SysLogging.version": %w`, err)}
		}
	}
	if v, ok := sluo.mutation.Message(); ok {
		if err := syslogging.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "SysLogging.message": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sluo *SysLoggingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysLoggingUpdateOne {
	sluo.modifiers = append(sluo.modifiers, modifiers...)
	return sluo
}

func (sluo *SysLoggingUpdateOne) sqlSave(ctx context.Context) (_node *SysLogging, err error) {
	if err := sluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(syslogging.Table, syslogging.Columns, sqlgraph.NewFieldSpec(syslogging.FieldID, field.TypeString))
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysLogging.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syslogging.FieldID)
		for _, f := range fields {
			if !syslogging.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != syslogging.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if sluo.mutation.CreatedAtCleared() {
		_spec.ClearField(syslogging.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sluo.mutation.UpdatedAt(); ok {
		_spec.SetField(syslogging.FieldUpdatedAt, field.TypeTime, value)
	}
	if sluo.mutation.UpdatedAtCleared() {
		_spec.ClearField(syslogging.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sluo.mutation.DeletedAt(); ok {
		_spec.SetField(syslogging.FieldDeletedAt, field.TypeTime, value)
	}
	if sluo.mutation.DeletedAtCleared() {
		_spec.ClearField(syslogging.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sluo.mutation.IsDel(); ok {
		_spec.SetField(syslogging.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sluo.mutation.Memo(); ok {
		_spec.SetField(syslogging.FieldMemo, field.TypeString, value)
	}
	if sluo.mutation.MemoCleared() {
		_spec.ClearField(syslogging.FieldMemo, field.TypeString)
	}
	if value, ok := sluo.mutation.Level(); ok {
		_spec.SetField(syslogging.FieldLevel, field.TypeString, value)
	}
	if sluo.mutation.LevelCleared() {
		_spec.ClearField(syslogging.FieldLevel, field.TypeString)
	}
	if value, ok := sluo.mutation.TraceID(); ok {
		_spec.SetField(syslogging.FieldTraceID, field.TypeString, value)
	}
	if sluo.mutation.TraceIDCleared() {
		_spec.ClearField(syslogging.FieldTraceID, field.TypeString)
	}
	if value, ok := sluo.mutation.UserID(); ok {
		_spec.SetField(syslogging.FieldUserID, field.TypeString, value)
	}
	if sluo.mutation.UserIDCleared() {
		_spec.ClearField(syslogging.FieldUserID, field.TypeString)
	}
	if value, ok := sluo.mutation.Tag(); ok {
		_spec.SetField(syslogging.FieldTag, field.TypeString, value)
	}
	if sluo.mutation.TagCleared() {
		_spec.ClearField(syslogging.FieldTag, field.TypeString)
	}
	if value, ok := sluo.mutation.Version(); ok {
		_spec.SetField(syslogging.FieldVersion, field.TypeString, value)
	}
	if sluo.mutation.VersionCleared() {
		_spec.ClearField(syslogging.FieldVersion, field.TypeString)
	}
	if value, ok := sluo.mutation.Message(); ok {
		_spec.SetField(syslogging.FieldMessage, field.TypeString, value)
	}
	if sluo.mutation.MessageCleared() {
		_spec.ClearField(syslogging.FieldMessage, field.TypeString)
	}
	if sluo.mutation.DataCleared() {
		_spec.ClearField(syslogging.FieldData, field.TypeString)
	}
	if value, ok := sluo.mutation.ErrorStack(); ok {
		_spec.SetField(syslogging.FieldErrorStack, field.TypeString, value)
	}
	if sluo.mutation.ErrorStackCleared() {
		_spec.ClearField(syslogging.FieldErrorStack, field.TypeString)
	}
	_spec.AddModifiers(sluo.modifiers...)
	_node = &SysLogging{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syslogging.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sluo.mutation.done = true
	return _node, nil
}
