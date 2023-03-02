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
	"github.com/heromicro/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlockUpdate is the builder for updating SysJwtBlock entities.
type SysJwtBlockUpdate struct {
	config
	hooks    []Hook
	mutation *SysJwtBlockMutation
}

// Where appends a list predicates to the SysJwtBlockUpdate builder.
func (sjbu *SysJwtBlockUpdate) Where(ps ...predicate.SysJwtBlock) *SysJwtBlockUpdate {
	sjbu.mutation.Where(ps...)
	return sjbu
}

// SetIsDel sets the "is_del" field.
func (sjbu *SysJwtBlockUpdate) SetIsDel(b bool) *SysJwtBlockUpdate {
	sjbu.mutation.SetIsDel(b)
	return sjbu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sjbu *SysJwtBlockUpdate) SetNillableIsDel(b *bool) *SysJwtBlockUpdate {
	if b != nil {
		sjbu.SetIsDel(*b)
	}
	return sjbu
}

// SetMemo sets the "memo" field.
func (sjbu *SysJwtBlockUpdate) SetMemo(s string) *SysJwtBlockUpdate {
	sjbu.mutation.SetMemo(s)
	return sjbu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sjbu *SysJwtBlockUpdate) SetNillableMemo(s *string) *SysJwtBlockUpdate {
	if s != nil {
		sjbu.SetMemo(*s)
	}
	return sjbu
}

// ClearMemo clears the value of the "memo" field.
func (sjbu *SysJwtBlockUpdate) ClearMemo() *SysJwtBlockUpdate {
	sjbu.mutation.ClearMemo()
	return sjbu
}

// SetUpdatedAt sets the "updated_at" field.
func (sjbu *SysJwtBlockUpdate) SetUpdatedAt(t time.Time) *SysJwtBlockUpdate {
	sjbu.mutation.SetUpdatedAt(t)
	return sjbu
}

// SetDeletedAt sets the "deleted_at" field.
func (sjbu *SysJwtBlockUpdate) SetDeletedAt(t time.Time) *SysJwtBlockUpdate {
	sjbu.mutation.SetDeletedAt(t)
	return sjbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sjbu *SysJwtBlockUpdate) SetNillableDeletedAt(t *time.Time) *SysJwtBlockUpdate {
	if t != nil {
		sjbu.SetDeletedAt(*t)
	}
	return sjbu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sjbu *SysJwtBlockUpdate) ClearDeletedAt() *SysJwtBlockUpdate {
	sjbu.mutation.ClearDeletedAt()
	return sjbu
}

// SetIsActive sets the "is_active" field.
func (sjbu *SysJwtBlockUpdate) SetIsActive(b bool) *SysJwtBlockUpdate {
	sjbu.mutation.SetIsActive(b)
	return sjbu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sjbu *SysJwtBlockUpdate) SetNillableIsActive(b *bool) *SysJwtBlockUpdate {
	if b != nil {
		sjbu.SetIsActive(*b)
	}
	return sjbu
}

// SetJwt sets the "jwt" field.
func (sjbu *SysJwtBlockUpdate) SetJwt(s string) *SysJwtBlockUpdate {
	sjbu.mutation.SetJwt(s)
	return sjbu
}

// Mutation returns the SysJwtBlockMutation object of the builder.
func (sjbu *SysJwtBlockUpdate) Mutation() *SysJwtBlockMutation {
	return sjbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sjbu *SysJwtBlockUpdate) Save(ctx context.Context) (int, error) {
	sjbu.defaults()
	return withHooks[int, SysJwtBlockMutation](ctx, sjbu.sqlSave, sjbu.mutation, sjbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sjbu *SysJwtBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := sjbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sjbu *SysJwtBlockUpdate) Exec(ctx context.Context) error {
	_, err := sjbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbu *SysJwtBlockUpdate) ExecX(ctx context.Context) {
	if err := sjbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sjbu *SysJwtBlockUpdate) defaults() {
	if _, ok := sjbu.mutation.UpdatedAt(); !ok {
		v := sysjwtblock.UpdateDefaultUpdatedAt()
		sjbu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sjbu *SysJwtBlockUpdate) check() error {
	if v, ok := sjbu.mutation.Memo(); ok {
		if err := sysjwtblock.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysJwtBlock.memo": %w`, err)}
		}
	}
	if v, ok := sjbu.mutation.Jwt(); ok {
		if err := sysjwtblock.JwtValidator(v); err != nil {
			return &ValidationError{Name: "jwt", err: fmt.Errorf(`ent: validator failed for field "SysJwtBlock.jwt": %w`, err)}
		}
	}
	return nil
}

func (sjbu *SysJwtBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sjbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysjwtblock.Table, sysjwtblock.Columns, sqlgraph.NewFieldSpec(sysjwtblock.FieldID, field.TypeString))
	if ps := sjbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sjbu.mutation.IsDel(); ok {
		_spec.SetField(sysjwtblock.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sjbu.mutation.Memo(); ok {
		_spec.SetField(sysjwtblock.FieldMemo, field.TypeString, value)
	}
	if sjbu.mutation.MemoCleared() {
		_spec.ClearField(sysjwtblock.FieldMemo, field.TypeString)
	}
	if value, ok := sjbu.mutation.UpdatedAt(); ok {
		_spec.SetField(sysjwtblock.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sjbu.mutation.DeletedAt(); ok {
		_spec.SetField(sysjwtblock.FieldDeletedAt, field.TypeTime, value)
	}
	if sjbu.mutation.DeletedAtCleared() {
		_spec.ClearField(sysjwtblock.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sjbu.mutation.IsActive(); ok {
		_spec.SetField(sysjwtblock.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sjbu.mutation.Jwt(); ok {
		_spec.SetField(sysjwtblock.FieldJwt, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sjbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysjwtblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sjbu.mutation.done = true
	return n, nil
}

// SysJwtBlockUpdateOne is the builder for updating a single SysJwtBlock entity.
type SysJwtBlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysJwtBlockMutation
}

// SetIsDel sets the "is_del" field.
func (sjbuo *SysJwtBlockUpdateOne) SetIsDel(b bool) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetIsDel(b)
	return sjbuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sjbuo *SysJwtBlockUpdateOne) SetNillableIsDel(b *bool) *SysJwtBlockUpdateOne {
	if b != nil {
		sjbuo.SetIsDel(*b)
	}
	return sjbuo
}

// SetMemo sets the "memo" field.
func (sjbuo *SysJwtBlockUpdateOne) SetMemo(s string) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetMemo(s)
	return sjbuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sjbuo *SysJwtBlockUpdateOne) SetNillableMemo(s *string) *SysJwtBlockUpdateOne {
	if s != nil {
		sjbuo.SetMemo(*s)
	}
	return sjbuo
}

// ClearMemo clears the value of the "memo" field.
func (sjbuo *SysJwtBlockUpdateOne) ClearMemo() *SysJwtBlockUpdateOne {
	sjbuo.mutation.ClearMemo()
	return sjbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sjbuo *SysJwtBlockUpdateOne) SetUpdatedAt(t time.Time) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetUpdatedAt(t)
	return sjbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (sjbuo *SysJwtBlockUpdateOne) SetDeletedAt(t time.Time) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetDeletedAt(t)
	return sjbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sjbuo *SysJwtBlockUpdateOne) SetNillableDeletedAt(t *time.Time) *SysJwtBlockUpdateOne {
	if t != nil {
		sjbuo.SetDeletedAt(*t)
	}
	return sjbuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sjbuo *SysJwtBlockUpdateOne) ClearDeletedAt() *SysJwtBlockUpdateOne {
	sjbuo.mutation.ClearDeletedAt()
	return sjbuo
}

// SetIsActive sets the "is_active" field.
func (sjbuo *SysJwtBlockUpdateOne) SetIsActive(b bool) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetIsActive(b)
	return sjbuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sjbuo *SysJwtBlockUpdateOne) SetNillableIsActive(b *bool) *SysJwtBlockUpdateOne {
	if b != nil {
		sjbuo.SetIsActive(*b)
	}
	return sjbuo
}

// SetJwt sets the "jwt" field.
func (sjbuo *SysJwtBlockUpdateOne) SetJwt(s string) *SysJwtBlockUpdateOne {
	sjbuo.mutation.SetJwt(s)
	return sjbuo
}

// Mutation returns the SysJwtBlockMutation object of the builder.
func (sjbuo *SysJwtBlockUpdateOne) Mutation() *SysJwtBlockMutation {
	return sjbuo.mutation
}

// Where appends a list predicates to the SysJwtBlockUpdate builder.
func (sjbuo *SysJwtBlockUpdateOne) Where(ps ...predicate.SysJwtBlock) *SysJwtBlockUpdateOne {
	sjbuo.mutation.Where(ps...)
	return sjbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sjbuo *SysJwtBlockUpdateOne) Select(field string, fields ...string) *SysJwtBlockUpdateOne {
	sjbuo.fields = append([]string{field}, fields...)
	return sjbuo
}

// Save executes the query and returns the updated SysJwtBlock entity.
func (sjbuo *SysJwtBlockUpdateOne) Save(ctx context.Context) (*SysJwtBlock, error) {
	sjbuo.defaults()
	return withHooks[*SysJwtBlock, SysJwtBlockMutation](ctx, sjbuo.sqlSave, sjbuo.mutation, sjbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sjbuo *SysJwtBlockUpdateOne) SaveX(ctx context.Context) *SysJwtBlock {
	node, err := sjbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sjbuo *SysJwtBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := sjbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbuo *SysJwtBlockUpdateOne) ExecX(ctx context.Context) {
	if err := sjbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sjbuo *SysJwtBlockUpdateOne) defaults() {
	if _, ok := sjbuo.mutation.UpdatedAt(); !ok {
		v := sysjwtblock.UpdateDefaultUpdatedAt()
		sjbuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sjbuo *SysJwtBlockUpdateOne) check() error {
	if v, ok := sjbuo.mutation.Memo(); ok {
		if err := sysjwtblock.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysJwtBlock.memo": %w`, err)}
		}
	}
	if v, ok := sjbuo.mutation.Jwt(); ok {
		if err := sysjwtblock.JwtValidator(v); err != nil {
			return &ValidationError{Name: "jwt", err: fmt.Errorf(`ent: validator failed for field "SysJwtBlock.jwt": %w`, err)}
		}
	}
	return nil
}

func (sjbuo *SysJwtBlockUpdateOne) sqlSave(ctx context.Context) (_node *SysJwtBlock, err error) {
	if err := sjbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysjwtblock.Table, sysjwtblock.Columns, sqlgraph.NewFieldSpec(sysjwtblock.FieldID, field.TypeString))
	id, ok := sjbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysJwtBlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sjbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysjwtblock.FieldID)
		for _, f := range fields {
			if !sysjwtblock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysjwtblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sjbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sjbuo.mutation.IsDel(); ok {
		_spec.SetField(sysjwtblock.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sjbuo.mutation.Memo(); ok {
		_spec.SetField(sysjwtblock.FieldMemo, field.TypeString, value)
	}
	if sjbuo.mutation.MemoCleared() {
		_spec.ClearField(sysjwtblock.FieldMemo, field.TypeString)
	}
	if value, ok := sjbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysjwtblock.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sjbuo.mutation.DeletedAt(); ok {
		_spec.SetField(sysjwtblock.FieldDeletedAt, field.TypeTime, value)
	}
	if sjbuo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysjwtblock.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sjbuo.mutation.IsActive(); ok {
		_spec.SetField(sysjwtblock.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sjbuo.mutation.Jwt(); ok {
		_spec.SetField(sysjwtblock.FieldJwt, field.TypeString, value)
	}
	_node = &SysJwtBlock{config: sjbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sjbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysjwtblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sjbuo.mutation.done = true
	return _node, nil
}
