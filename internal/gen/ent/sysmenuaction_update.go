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
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuActionUpdate is the builder for updating SysMenuAction entities.
type SysMenuActionUpdate struct {
	config
	hooks     []Hook
	mutation  *SysMenuActionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysMenuActionUpdate builder.
func (smau *SysMenuActionUpdate) Where(ps ...predicate.SysMenuAction) *SysMenuActionUpdate {
	smau.mutation.Where(ps...)
	return smau
}

// SetIsDel sets the "is_del" field.
func (smau *SysMenuActionUpdate) SetIsDel(b bool) *SysMenuActionUpdate {
	smau.mutation.SetIsDel(b)
	return smau
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smau *SysMenuActionUpdate) SetNillableIsDel(b *bool) *SysMenuActionUpdate {
	if b != nil {
		smau.SetIsDel(*b)
	}
	return smau
}

// SetSort sets the "sort" field.
func (smau *SysMenuActionUpdate) SetSort(i int32) *SysMenuActionUpdate {
	smau.mutation.ResetSort()
	smau.mutation.SetSort(i)
	return smau
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smau *SysMenuActionUpdate) SetNillableSort(i *int32) *SysMenuActionUpdate {
	if i != nil {
		smau.SetSort(*i)
	}
	return smau
}

// AddSort adds i to the "sort" field.
func (smau *SysMenuActionUpdate) AddSort(i int32) *SysMenuActionUpdate {
	smau.mutation.AddSort(i)
	return smau
}

// SetIsActive sets the "is_active" field.
func (smau *SysMenuActionUpdate) SetIsActive(b bool) *SysMenuActionUpdate {
	smau.mutation.SetIsActive(b)
	return smau
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (smau *SysMenuActionUpdate) SetNillableIsActive(b *bool) *SysMenuActionUpdate {
	if b != nil {
		smau.SetIsActive(*b)
	}
	return smau
}

// SetMemo sets the "memo" field.
func (smau *SysMenuActionUpdate) SetMemo(s string) *SysMenuActionUpdate {
	smau.mutation.SetMemo(s)
	return smau
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smau *SysMenuActionUpdate) SetNillableMemo(s *string) *SysMenuActionUpdate {
	if s != nil {
		smau.SetMemo(*s)
	}
	return smau
}

// ClearMemo clears the value of the "memo" field.
func (smau *SysMenuActionUpdate) ClearMemo() *SysMenuActionUpdate {
	smau.mutation.ClearMemo()
	return smau
}

// SetUpdatedAt sets the "updated_at" field.
func (smau *SysMenuActionUpdate) SetUpdatedAt(t time.Time) *SysMenuActionUpdate {
	smau.mutation.SetUpdatedAt(t)
	return smau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (smau *SysMenuActionUpdate) ClearUpdatedAt() *SysMenuActionUpdate {
	smau.mutation.ClearUpdatedAt()
	return smau
}

// SetDeletedAt sets the "deleted_at" field.
func (smau *SysMenuActionUpdate) SetDeletedAt(t time.Time) *SysMenuActionUpdate {
	smau.mutation.SetDeletedAt(t)
	return smau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smau *SysMenuActionUpdate) SetNillableDeletedAt(t *time.Time) *SysMenuActionUpdate {
	if t != nil {
		smau.SetDeletedAt(*t)
	}
	return smau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smau *SysMenuActionUpdate) ClearDeletedAt() *SysMenuActionUpdate {
	smau.mutation.ClearDeletedAt()
	return smau
}

// SetMenuID sets the "menu_id" field.
func (smau *SysMenuActionUpdate) SetMenuID(s string) *SysMenuActionUpdate {
	smau.mutation.SetMenuID(s)
	return smau
}

// SetCode sets the "code" field.
func (smau *SysMenuActionUpdate) SetCode(s string) *SysMenuActionUpdate {
	smau.mutation.SetCode(s)
	return smau
}

// SetName sets the "name" field.
func (smau *SysMenuActionUpdate) SetName(s string) *SysMenuActionUpdate {
	smau.mutation.SetName(s)
	return smau
}

// Mutation returns the SysMenuActionMutation object of the builder.
func (smau *SysMenuActionUpdate) Mutation() *SysMenuActionMutation {
	return smau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smau *SysMenuActionUpdate) Save(ctx context.Context) (int, error) {
	smau.defaults()
	return withHooks[int, SysMenuActionMutation](ctx, smau.sqlSave, smau.mutation, smau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smau *SysMenuActionUpdate) SaveX(ctx context.Context) int {
	affected, err := smau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smau *SysMenuActionUpdate) Exec(ctx context.Context) error {
	_, err := smau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smau *SysMenuActionUpdate) ExecX(ctx context.Context) {
	if err := smau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smau *SysMenuActionUpdate) defaults() {
	if _, ok := smau.mutation.UpdatedAt(); !ok && !smau.mutation.UpdatedAtCleared() {
		v := sysmenuaction.UpdateDefaultUpdatedAt()
		smau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smau *SysMenuActionUpdate) check() error {
	if v, ok := smau.mutation.Memo(); ok {
		if err := sysmenuaction.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.memo": %w`, err)}
		}
	}
	if v, ok := smau.mutation.MenuID(); ok {
		if err := sysmenuaction.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.menu_id": %w`, err)}
		}
	}
	if v, ok := smau.mutation.Code(); ok {
		if err := sysmenuaction.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.code": %w`, err)}
		}
	}
	if v, ok := smau.mutation.Name(); ok {
		if err := sysmenuaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (smau *SysMenuActionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysMenuActionUpdate {
	smau.modifiers = append(smau.modifiers, modifiers...)
	return smau
}

func (smau *SysMenuActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := smau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysmenuaction.Table, sysmenuaction.Columns, sqlgraph.NewFieldSpec(sysmenuaction.FieldID, field.TypeString))
	if ps := smau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smau.mutation.IsDel(); ok {
		_spec.SetField(sysmenuaction.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := smau.mutation.Sort(); ok {
		_spec.SetField(sysmenuaction.FieldSort, field.TypeInt32, value)
	}
	if value, ok := smau.mutation.AddedSort(); ok {
		_spec.AddField(sysmenuaction.FieldSort, field.TypeInt32, value)
	}
	if value, ok := smau.mutation.IsActive(); ok {
		_spec.SetField(sysmenuaction.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := smau.mutation.Memo(); ok {
		_spec.SetField(sysmenuaction.FieldMemo, field.TypeString, value)
	}
	if smau.mutation.MemoCleared() {
		_spec.ClearField(sysmenuaction.FieldMemo, field.TypeString)
	}
	if smau.mutation.CreatedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := smau.mutation.UpdatedAt(); ok {
		_spec.SetField(sysmenuaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if smau.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := smau.mutation.DeletedAt(); ok {
		_spec.SetField(sysmenuaction.FieldDeletedAt, field.TypeTime, value)
	}
	if smau.mutation.DeletedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := smau.mutation.MenuID(); ok {
		_spec.SetField(sysmenuaction.FieldMenuID, field.TypeString, value)
	}
	if value, ok := smau.mutation.Code(); ok {
		_spec.SetField(sysmenuaction.FieldCode, field.TypeString, value)
	}
	if value, ok := smau.mutation.Name(); ok {
		_spec.SetField(sysmenuaction.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(smau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, smau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenuaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	smau.mutation.done = true
	return n, nil
}

// SysMenuActionUpdateOne is the builder for updating a single SysMenuAction entity.
type SysMenuActionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysMenuActionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (smauo *SysMenuActionUpdateOne) SetIsDel(b bool) *SysMenuActionUpdateOne {
	smauo.mutation.SetIsDel(b)
	return smauo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smauo *SysMenuActionUpdateOne) SetNillableIsDel(b *bool) *SysMenuActionUpdateOne {
	if b != nil {
		smauo.SetIsDel(*b)
	}
	return smauo
}

// SetSort sets the "sort" field.
func (smauo *SysMenuActionUpdateOne) SetSort(i int32) *SysMenuActionUpdateOne {
	smauo.mutation.ResetSort()
	smauo.mutation.SetSort(i)
	return smauo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smauo *SysMenuActionUpdateOne) SetNillableSort(i *int32) *SysMenuActionUpdateOne {
	if i != nil {
		smauo.SetSort(*i)
	}
	return smauo
}

// AddSort adds i to the "sort" field.
func (smauo *SysMenuActionUpdateOne) AddSort(i int32) *SysMenuActionUpdateOne {
	smauo.mutation.AddSort(i)
	return smauo
}

// SetIsActive sets the "is_active" field.
func (smauo *SysMenuActionUpdateOne) SetIsActive(b bool) *SysMenuActionUpdateOne {
	smauo.mutation.SetIsActive(b)
	return smauo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (smauo *SysMenuActionUpdateOne) SetNillableIsActive(b *bool) *SysMenuActionUpdateOne {
	if b != nil {
		smauo.SetIsActive(*b)
	}
	return smauo
}

// SetMemo sets the "memo" field.
func (smauo *SysMenuActionUpdateOne) SetMemo(s string) *SysMenuActionUpdateOne {
	smauo.mutation.SetMemo(s)
	return smauo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smauo *SysMenuActionUpdateOne) SetNillableMemo(s *string) *SysMenuActionUpdateOne {
	if s != nil {
		smauo.SetMemo(*s)
	}
	return smauo
}

// ClearMemo clears the value of the "memo" field.
func (smauo *SysMenuActionUpdateOne) ClearMemo() *SysMenuActionUpdateOne {
	smauo.mutation.ClearMemo()
	return smauo
}

// SetUpdatedAt sets the "updated_at" field.
func (smauo *SysMenuActionUpdateOne) SetUpdatedAt(t time.Time) *SysMenuActionUpdateOne {
	smauo.mutation.SetUpdatedAt(t)
	return smauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (smauo *SysMenuActionUpdateOne) ClearUpdatedAt() *SysMenuActionUpdateOne {
	smauo.mutation.ClearUpdatedAt()
	return smauo
}

// SetDeletedAt sets the "deleted_at" field.
func (smauo *SysMenuActionUpdateOne) SetDeletedAt(t time.Time) *SysMenuActionUpdateOne {
	smauo.mutation.SetDeletedAt(t)
	return smauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smauo *SysMenuActionUpdateOne) SetNillableDeletedAt(t *time.Time) *SysMenuActionUpdateOne {
	if t != nil {
		smauo.SetDeletedAt(*t)
	}
	return smauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smauo *SysMenuActionUpdateOne) ClearDeletedAt() *SysMenuActionUpdateOne {
	smauo.mutation.ClearDeletedAt()
	return smauo
}

// SetMenuID sets the "menu_id" field.
func (smauo *SysMenuActionUpdateOne) SetMenuID(s string) *SysMenuActionUpdateOne {
	smauo.mutation.SetMenuID(s)
	return smauo
}

// SetCode sets the "code" field.
func (smauo *SysMenuActionUpdateOne) SetCode(s string) *SysMenuActionUpdateOne {
	smauo.mutation.SetCode(s)
	return smauo
}

// SetName sets the "name" field.
func (smauo *SysMenuActionUpdateOne) SetName(s string) *SysMenuActionUpdateOne {
	smauo.mutation.SetName(s)
	return smauo
}

// Mutation returns the SysMenuActionMutation object of the builder.
func (smauo *SysMenuActionUpdateOne) Mutation() *SysMenuActionMutation {
	return smauo.mutation
}

// Where appends a list predicates to the SysMenuActionUpdate builder.
func (smauo *SysMenuActionUpdateOne) Where(ps ...predicate.SysMenuAction) *SysMenuActionUpdateOne {
	smauo.mutation.Where(ps...)
	return smauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smauo *SysMenuActionUpdateOne) Select(field string, fields ...string) *SysMenuActionUpdateOne {
	smauo.fields = append([]string{field}, fields...)
	return smauo
}

// Save executes the query and returns the updated SysMenuAction entity.
func (smauo *SysMenuActionUpdateOne) Save(ctx context.Context) (*SysMenuAction, error) {
	smauo.defaults()
	return withHooks[*SysMenuAction, SysMenuActionMutation](ctx, smauo.sqlSave, smauo.mutation, smauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smauo *SysMenuActionUpdateOne) SaveX(ctx context.Context) *SysMenuAction {
	node, err := smauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smauo *SysMenuActionUpdateOne) Exec(ctx context.Context) error {
	_, err := smauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smauo *SysMenuActionUpdateOne) ExecX(ctx context.Context) {
	if err := smauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smauo *SysMenuActionUpdateOne) defaults() {
	if _, ok := smauo.mutation.UpdatedAt(); !ok && !smauo.mutation.UpdatedAtCleared() {
		v := sysmenuaction.UpdateDefaultUpdatedAt()
		smauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smauo *SysMenuActionUpdateOne) check() error {
	if v, ok := smauo.mutation.Memo(); ok {
		if err := sysmenuaction.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.memo": %w`, err)}
		}
	}
	if v, ok := smauo.mutation.MenuID(); ok {
		if err := sysmenuaction.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.menu_id": %w`, err)}
		}
	}
	if v, ok := smauo.mutation.Code(); ok {
		if err := sysmenuaction.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.code": %w`, err)}
		}
	}
	if v, ok := smauo.mutation.Name(); ok {
		if err := sysmenuaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (smauo *SysMenuActionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysMenuActionUpdateOne {
	smauo.modifiers = append(smauo.modifiers, modifiers...)
	return smauo
}

func (smauo *SysMenuActionUpdateOne) sqlSave(ctx context.Context) (_node *SysMenuAction, err error) {
	if err := smauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysmenuaction.Table, sysmenuaction.Columns, sqlgraph.NewFieldSpec(sysmenuaction.FieldID, field.TypeString))
	id, ok := smauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysMenuAction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenuaction.FieldID)
		for _, f := range fields {
			if !sysmenuaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysmenuaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smauo.mutation.IsDel(); ok {
		_spec.SetField(sysmenuaction.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := smauo.mutation.Sort(); ok {
		_spec.SetField(sysmenuaction.FieldSort, field.TypeInt32, value)
	}
	if value, ok := smauo.mutation.AddedSort(); ok {
		_spec.AddField(sysmenuaction.FieldSort, field.TypeInt32, value)
	}
	if value, ok := smauo.mutation.IsActive(); ok {
		_spec.SetField(sysmenuaction.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := smauo.mutation.Memo(); ok {
		_spec.SetField(sysmenuaction.FieldMemo, field.TypeString, value)
	}
	if smauo.mutation.MemoCleared() {
		_spec.ClearField(sysmenuaction.FieldMemo, field.TypeString)
	}
	if smauo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := smauo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysmenuaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if smauo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := smauo.mutation.DeletedAt(); ok {
		_spec.SetField(sysmenuaction.FieldDeletedAt, field.TypeTime, value)
	}
	if smauo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysmenuaction.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := smauo.mutation.MenuID(); ok {
		_spec.SetField(sysmenuaction.FieldMenuID, field.TypeString, value)
	}
	if value, ok := smauo.mutation.Code(); ok {
		_spec.SetField(sysmenuaction.FieldCode, field.TypeString, value)
	}
	if value, ok := smauo.mutation.Name(); ok {
		_spec.SetField(sysmenuaction.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(smauo.modifiers...)
	_node = &SysMenuAction{config: smauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenuaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	smauo.mutation.done = true
	return _node, nil
}
