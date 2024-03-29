// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysrolemenu"
)

// SysRoleMenuUpdate is the builder for updating SysRoleMenu entities.
type SysRoleMenuUpdate struct {
	config
	hooks     []Hook
	mutation  *SysRoleMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysRoleMenuUpdate builder.
func (srmu *SysRoleMenuUpdate) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuUpdate {
	srmu.mutation.Where(ps...)
	return srmu
}

// SetIsDel sets the "is_del" field.
func (srmu *SysRoleMenuUpdate) SetIsDel(b bool) *SysRoleMenuUpdate {
	srmu.mutation.SetIsDel(b)
	return srmu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableIsDel(b *bool) *SysRoleMenuUpdate {
	if b != nil {
		srmu.SetIsDel(*b)
	}
	return srmu
}

// SetUpdatedAt sets the "updated_at" field.
func (srmu *SysRoleMenuUpdate) SetUpdatedAt(t time.Time) *SysRoleMenuUpdate {
	srmu.mutation.SetUpdatedAt(t)
	return srmu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (srmu *SysRoleMenuUpdate) ClearUpdatedAt() *SysRoleMenuUpdate {
	srmu.mutation.ClearUpdatedAt()
	return srmu
}

// SetDeletedAt sets the "deleted_at" field.
func (srmu *SysRoleMenuUpdate) SetDeletedAt(t time.Time) *SysRoleMenuUpdate {
	srmu.mutation.SetDeletedAt(t)
	return srmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableDeletedAt(t *time.Time) *SysRoleMenuUpdate {
	if t != nil {
		srmu.SetDeletedAt(*t)
	}
	return srmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (srmu *SysRoleMenuUpdate) ClearDeletedAt() *SysRoleMenuUpdate {
	srmu.mutation.ClearDeletedAt()
	return srmu
}

// SetRoleID sets the "role_id" field.
func (srmu *SysRoleMenuUpdate) SetRoleID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetRoleID(s)
	return srmu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableRoleID(s *string) *SysRoleMenuUpdate {
	if s != nil {
		srmu.SetRoleID(*s)
	}
	return srmu
}

// SetMenuID sets the "menu_id" field.
func (srmu *SysRoleMenuUpdate) SetMenuID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetMenuID(s)
	return srmu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableMenuID(s *string) *SysRoleMenuUpdate {
	if s != nil {
		srmu.SetMenuID(*s)
	}
	return srmu
}

// SetActionID sets the "action_id" field.
func (srmu *SysRoleMenuUpdate) SetActionID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetActionID(s)
	return srmu
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableActionID(s *string) *SysRoleMenuUpdate {
	if s != nil {
		srmu.SetActionID(*s)
	}
	return srmu
}

// ClearActionID clears the value of the "action_id" field.
func (srmu *SysRoleMenuUpdate) ClearActionID() *SysRoleMenuUpdate {
	srmu.mutation.ClearActionID()
	return srmu
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmu *SysRoleMenuUpdate) Mutation() *SysRoleMenuMutation {
	return srmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (srmu *SysRoleMenuUpdate) Save(ctx context.Context) (int, error) {
	srmu.defaults()
	return withHooks(ctx, srmu.sqlSave, srmu.mutation, srmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := srmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (srmu *SysRoleMenuUpdate) Exec(ctx context.Context) error {
	_, err := srmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) ExecX(ctx context.Context) {
	if err := srmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srmu *SysRoleMenuUpdate) defaults() {
	if _, ok := srmu.mutation.UpdatedAt(); !ok && !srmu.mutation.UpdatedAtCleared() {
		v := sysrolemenu.UpdateDefaultUpdatedAt()
		srmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srmu *SysRoleMenuUpdate) check() error {
	if v, ok := srmu.mutation.RoleID(); ok {
		if err := sysrolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.role_id": %w`, err)}
		}
	}
	if v, ok := srmu.mutation.MenuID(); ok {
		if err := sysrolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.menu_id": %w`, err)}
		}
	}
	if v, ok := srmu.mutation.ActionID(); ok {
		if err := sysrolemenu.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.action_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (srmu *SysRoleMenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysRoleMenuUpdate {
	srmu.modifiers = append(srmu.modifiers, modifiers...)
	return srmu
}

func (srmu *SysRoleMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := srmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysrolemenu.Table, sysrolemenu.Columns, sqlgraph.NewFieldSpec(sysrolemenu.FieldID, field.TypeString))
	if ps := srmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmu.mutation.IsDel(); ok {
		_spec.SetField(sysrolemenu.FieldIsDel, field.TypeBool, value)
	}
	if srmu.mutation.CreatedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := srmu.mutation.UpdatedAt(); ok {
		_spec.SetField(sysrolemenu.FieldUpdatedAt, field.TypeTime, value)
	}
	if srmu.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := srmu.mutation.DeletedAt(); ok {
		_spec.SetField(sysrolemenu.FieldDeletedAt, field.TypeTime, value)
	}
	if srmu.mutation.DeletedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := srmu.mutation.RoleID(); ok {
		_spec.SetField(sysrolemenu.FieldRoleID, field.TypeString, value)
	}
	if value, ok := srmu.mutation.MenuID(); ok {
		_spec.SetField(sysrolemenu.FieldMenuID, field.TypeString, value)
	}
	if value, ok := srmu.mutation.ActionID(); ok {
		_spec.SetField(sysrolemenu.FieldActionID, field.TypeString, value)
	}
	if srmu.mutation.ActionIDCleared() {
		_spec.ClearField(sysrolemenu.FieldActionID, field.TypeString)
	}
	_spec.AddModifiers(srmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, srmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	srmu.mutation.done = true
	return n, nil
}

// SysRoleMenuUpdateOne is the builder for updating a single SysRoleMenu entity.
type SysRoleMenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysRoleMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (srmuo *SysRoleMenuUpdateOne) SetIsDel(b bool) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetIsDel(b)
	return srmuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableIsDel(b *bool) *SysRoleMenuUpdateOne {
	if b != nil {
		srmuo.SetIsDel(*b)
	}
	return srmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (srmuo *SysRoleMenuUpdateOne) SetUpdatedAt(t time.Time) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetUpdatedAt(t)
	return srmuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (srmuo *SysRoleMenuUpdateOne) ClearUpdatedAt() *SysRoleMenuUpdateOne {
	srmuo.mutation.ClearUpdatedAt()
	return srmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (srmuo *SysRoleMenuUpdateOne) SetDeletedAt(t time.Time) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetDeletedAt(t)
	return srmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableDeletedAt(t *time.Time) *SysRoleMenuUpdateOne {
	if t != nil {
		srmuo.SetDeletedAt(*t)
	}
	return srmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (srmuo *SysRoleMenuUpdateOne) ClearDeletedAt() *SysRoleMenuUpdateOne {
	srmuo.mutation.ClearDeletedAt()
	return srmuo
}

// SetRoleID sets the "role_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetRoleID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetRoleID(s)
	return srmuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableRoleID(s *string) *SysRoleMenuUpdateOne {
	if s != nil {
		srmuo.SetRoleID(*s)
	}
	return srmuo
}

// SetMenuID sets the "menu_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetMenuID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetMenuID(s)
	return srmuo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableMenuID(s *string) *SysRoleMenuUpdateOne {
	if s != nil {
		srmuo.SetMenuID(*s)
	}
	return srmuo
}

// SetActionID sets the "action_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetActionID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetActionID(s)
	return srmuo
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableActionID(s *string) *SysRoleMenuUpdateOne {
	if s != nil {
		srmuo.SetActionID(*s)
	}
	return srmuo
}

// ClearActionID clears the value of the "action_id" field.
func (srmuo *SysRoleMenuUpdateOne) ClearActionID() *SysRoleMenuUpdateOne {
	srmuo.mutation.ClearActionID()
	return srmuo
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmuo *SysRoleMenuUpdateOne) Mutation() *SysRoleMenuMutation {
	return srmuo.mutation
}

// Where appends a list predicates to the SysRoleMenuUpdate builder.
func (srmuo *SysRoleMenuUpdateOne) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuUpdateOne {
	srmuo.mutation.Where(ps...)
	return srmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (srmuo *SysRoleMenuUpdateOne) Select(field string, fields ...string) *SysRoleMenuUpdateOne {
	srmuo.fields = append([]string{field}, fields...)
	return srmuo
}

// Save executes the query and returns the updated SysRoleMenu entity.
func (srmuo *SysRoleMenuUpdateOne) Save(ctx context.Context) (*SysRoleMenu, error) {
	srmuo.defaults()
	return withHooks(ctx, srmuo.sqlSave, srmuo.mutation, srmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) SaveX(ctx context.Context) *SysRoleMenu {
	node, err := srmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (srmuo *SysRoleMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := srmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) ExecX(ctx context.Context) {
	if err := srmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srmuo *SysRoleMenuUpdateOne) defaults() {
	if _, ok := srmuo.mutation.UpdatedAt(); !ok && !srmuo.mutation.UpdatedAtCleared() {
		v := sysrolemenu.UpdateDefaultUpdatedAt()
		srmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srmuo *SysRoleMenuUpdateOne) check() error {
	if v, ok := srmuo.mutation.RoleID(); ok {
		if err := sysrolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.role_id": %w`, err)}
		}
	}
	if v, ok := srmuo.mutation.MenuID(); ok {
		if err := sysrolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.menu_id": %w`, err)}
		}
	}
	if v, ok := srmuo.mutation.ActionID(); ok {
		if err := sysrolemenu.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf(`mainent: validator failed for field "SysRoleMenu.action_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (srmuo *SysRoleMenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysRoleMenuUpdateOne {
	srmuo.modifiers = append(srmuo.modifiers, modifiers...)
	return srmuo
}

func (srmuo *SysRoleMenuUpdateOne) sqlSave(ctx context.Context) (_node *SysRoleMenu, err error) {
	if err := srmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysrolemenu.Table, sysrolemenu.Columns, sqlgraph.NewFieldSpec(sysrolemenu.FieldID, field.TypeString))
	id, ok := srmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "SysRoleMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := srmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrolemenu.FieldID)
		for _, f := range fields {
			if !sysrolemenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != sysrolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := srmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmuo.mutation.IsDel(); ok {
		_spec.SetField(sysrolemenu.FieldIsDel, field.TypeBool, value)
	}
	if srmuo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := srmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysrolemenu.FieldUpdatedAt, field.TypeTime, value)
	}
	if srmuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := srmuo.mutation.DeletedAt(); ok {
		_spec.SetField(sysrolemenu.FieldDeletedAt, field.TypeTime, value)
	}
	if srmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysrolemenu.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := srmuo.mutation.RoleID(); ok {
		_spec.SetField(sysrolemenu.FieldRoleID, field.TypeString, value)
	}
	if value, ok := srmuo.mutation.MenuID(); ok {
		_spec.SetField(sysrolemenu.FieldMenuID, field.TypeString, value)
	}
	if value, ok := srmuo.mutation.ActionID(); ok {
		_spec.SetField(sysrolemenu.FieldActionID, field.TypeString, value)
	}
	if srmuo.mutation.ActionIDCleared() {
		_spec.ClearField(sysrolemenu.FieldActionID, field.TypeString)
	}
	_spec.AddModifiers(srmuo.modifiers...)
	_node = &SysRoleMenu{config: srmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, srmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	srmuo.mutation.done = true
	return _node, nil
}
