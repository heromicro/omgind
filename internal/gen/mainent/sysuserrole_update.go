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
	"github.com/heromicro/omgind/internal/gen/mainent/sysuserrole"
)

// SysUserRoleUpdate is the builder for updating SysUserRole entities.
type SysUserRoleUpdate struct {
	config
	hooks     []Hook
	mutation  *SysUserRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysUserRoleUpdate builder.
func (suru *SysUserRoleUpdate) Where(ps ...predicate.SysUserRole) *SysUserRoleUpdate {
	suru.mutation.Where(ps...)
	return suru
}

// SetIsDel sets the "is_del" field.
func (suru *SysUserRoleUpdate) SetIsDel(b bool) *SysUserRoleUpdate {
	suru.mutation.SetIsDel(b)
	return suru
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (suru *SysUserRoleUpdate) SetNillableIsDel(b *bool) *SysUserRoleUpdate {
	if b != nil {
		suru.SetIsDel(*b)
	}
	return suru
}

// SetUpdatedAt sets the "updated_at" field.
func (suru *SysUserRoleUpdate) SetUpdatedAt(t time.Time) *SysUserRoleUpdate {
	suru.mutation.SetUpdatedAt(t)
	return suru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suru *SysUserRoleUpdate) ClearUpdatedAt() *SysUserRoleUpdate {
	suru.mutation.ClearUpdatedAt()
	return suru
}

// SetDeletedAt sets the "deleted_at" field.
func (suru *SysUserRoleUpdate) SetDeletedAt(t time.Time) *SysUserRoleUpdate {
	suru.mutation.SetDeletedAt(t)
	return suru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suru *SysUserRoleUpdate) SetNillableDeletedAt(t *time.Time) *SysUserRoleUpdate {
	if t != nil {
		suru.SetDeletedAt(*t)
	}
	return suru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suru *SysUserRoleUpdate) ClearDeletedAt() *SysUserRoleUpdate {
	suru.mutation.ClearDeletedAt()
	return suru
}

// SetUserID sets the "user_id" field.
func (suru *SysUserRoleUpdate) SetUserID(s string) *SysUserRoleUpdate {
	suru.mutation.SetUserID(s)
	return suru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suru *SysUserRoleUpdate) SetNillableUserID(s *string) *SysUserRoleUpdate {
	if s != nil {
		suru.SetUserID(*s)
	}
	return suru
}

// SetRoleID sets the "role_id" field.
func (suru *SysUserRoleUpdate) SetRoleID(s string) *SysUserRoleUpdate {
	suru.mutation.SetRoleID(s)
	return suru
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (suru *SysUserRoleUpdate) SetNillableRoleID(s *string) *SysUserRoleUpdate {
	if s != nil {
		suru.SetRoleID(*s)
	}
	return suru
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (suru *SysUserRoleUpdate) Mutation() *SysUserRoleMutation {
	return suru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suru *SysUserRoleUpdate) Save(ctx context.Context) (int, error) {
	suru.defaults()
	return withHooks(ctx, suru.sqlSave, suru.mutation, suru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suru *SysUserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := suru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suru *SysUserRoleUpdate) Exec(ctx context.Context) error {
	_, err := suru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suru *SysUserRoleUpdate) ExecX(ctx context.Context) {
	if err := suru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suru *SysUserRoleUpdate) defaults() {
	if _, ok := suru.mutation.UpdatedAt(); !ok && !suru.mutation.UpdatedAtCleared() {
		v := sysuserrole.UpdateDefaultUpdatedAt()
		suru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suru *SysUserRoleUpdate) check() error {
	if v, ok := suru.mutation.UserID(); ok {
		if err := sysuserrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`mainent: validator failed for field "SysUserRole.user_id": %w`, err)}
		}
	}
	if v, ok := suru.mutation.RoleID(); ok {
		if err := sysuserrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`mainent: validator failed for field "SysUserRole.role_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suru *SysUserRoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysUserRoleUpdate {
	suru.modifiers = append(suru.modifiers, modifiers...)
	return suru
}

func (suru *SysUserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := suru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuserrole.Table, sysuserrole.Columns, sqlgraph.NewFieldSpec(sysuserrole.FieldID, field.TypeString))
	if ps := suru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suru.mutation.IsDel(); ok {
		_spec.SetField(sysuserrole.FieldIsDel, field.TypeBool, value)
	}
	if suru.mutation.CreatedAtCleared() {
		_spec.ClearField(sysuserrole.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suru.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuserrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if suru.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysuserrole.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := suru.mutation.DeletedAt(); ok {
		_spec.SetField(sysuserrole.FieldDeletedAt, field.TypeTime, value)
	}
	if suru.mutation.DeletedAtCleared() {
		_spec.ClearField(sysuserrole.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := suru.mutation.UserID(); ok {
		_spec.SetField(sysuserrole.FieldUserID, field.TypeString, value)
	}
	if value, ok := suru.mutation.RoleID(); ok {
		_spec.SetField(sysuserrole.FieldRoleID, field.TypeString, value)
	}
	_spec.AddModifiers(suru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, suru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	suru.mutation.done = true
	return n, nil
}

// SysUserRoleUpdateOne is the builder for updating a single SysUserRole entity.
type SysUserRoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysUserRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (suruo *SysUserRoleUpdateOne) SetIsDel(b bool) *SysUserRoleUpdateOne {
	suruo.mutation.SetIsDel(b)
	return suruo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (suruo *SysUserRoleUpdateOne) SetNillableIsDel(b *bool) *SysUserRoleUpdateOne {
	if b != nil {
		suruo.SetIsDel(*b)
	}
	return suruo
}

// SetUpdatedAt sets the "updated_at" field.
func (suruo *SysUserRoleUpdateOne) SetUpdatedAt(t time.Time) *SysUserRoleUpdateOne {
	suruo.mutation.SetUpdatedAt(t)
	return suruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suruo *SysUserRoleUpdateOne) ClearUpdatedAt() *SysUserRoleUpdateOne {
	suruo.mutation.ClearUpdatedAt()
	return suruo
}

// SetDeletedAt sets the "deleted_at" field.
func (suruo *SysUserRoleUpdateOne) SetDeletedAt(t time.Time) *SysUserRoleUpdateOne {
	suruo.mutation.SetDeletedAt(t)
	return suruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suruo *SysUserRoleUpdateOne) SetNillableDeletedAt(t *time.Time) *SysUserRoleUpdateOne {
	if t != nil {
		suruo.SetDeletedAt(*t)
	}
	return suruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suruo *SysUserRoleUpdateOne) ClearDeletedAt() *SysUserRoleUpdateOne {
	suruo.mutation.ClearDeletedAt()
	return suruo
}

// SetUserID sets the "user_id" field.
func (suruo *SysUserRoleUpdateOne) SetUserID(s string) *SysUserRoleUpdateOne {
	suruo.mutation.SetUserID(s)
	return suruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suruo *SysUserRoleUpdateOne) SetNillableUserID(s *string) *SysUserRoleUpdateOne {
	if s != nil {
		suruo.SetUserID(*s)
	}
	return suruo
}

// SetRoleID sets the "role_id" field.
func (suruo *SysUserRoleUpdateOne) SetRoleID(s string) *SysUserRoleUpdateOne {
	suruo.mutation.SetRoleID(s)
	return suruo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (suruo *SysUserRoleUpdateOne) SetNillableRoleID(s *string) *SysUserRoleUpdateOne {
	if s != nil {
		suruo.SetRoleID(*s)
	}
	return suruo
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (suruo *SysUserRoleUpdateOne) Mutation() *SysUserRoleMutation {
	return suruo.mutation
}

// Where appends a list predicates to the SysUserRoleUpdate builder.
func (suruo *SysUserRoleUpdateOne) Where(ps ...predicate.SysUserRole) *SysUserRoleUpdateOne {
	suruo.mutation.Where(ps...)
	return suruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suruo *SysUserRoleUpdateOne) Select(field string, fields ...string) *SysUserRoleUpdateOne {
	suruo.fields = append([]string{field}, fields...)
	return suruo
}

// Save executes the query and returns the updated SysUserRole entity.
func (suruo *SysUserRoleUpdateOne) Save(ctx context.Context) (*SysUserRole, error) {
	suruo.defaults()
	return withHooks(ctx, suruo.sqlSave, suruo.mutation, suruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suruo *SysUserRoleUpdateOne) SaveX(ctx context.Context) *SysUserRole {
	node, err := suruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suruo *SysUserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := suruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suruo *SysUserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := suruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suruo *SysUserRoleUpdateOne) defaults() {
	if _, ok := suruo.mutation.UpdatedAt(); !ok && !suruo.mutation.UpdatedAtCleared() {
		v := sysuserrole.UpdateDefaultUpdatedAt()
		suruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suruo *SysUserRoleUpdateOne) check() error {
	if v, ok := suruo.mutation.UserID(); ok {
		if err := sysuserrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`mainent: validator failed for field "SysUserRole.user_id": %w`, err)}
		}
	}
	if v, ok := suruo.mutation.RoleID(); ok {
		if err := sysuserrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`mainent: validator failed for field "SysUserRole.role_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suruo *SysUserRoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysUserRoleUpdateOne {
	suruo.modifiers = append(suruo.modifiers, modifiers...)
	return suruo
}

func (suruo *SysUserRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysUserRole, err error) {
	if err := suruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuserrole.Table, sysuserrole.Columns, sqlgraph.NewFieldSpec(sysuserrole.FieldID, field.TypeString))
	id, ok := suruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "SysUserRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuserrole.FieldID)
		for _, f := range fields {
			if !sysuserrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != sysuserrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suruo.mutation.IsDel(); ok {
		_spec.SetField(sysuserrole.FieldIsDel, field.TypeBool, value)
	}
	if suruo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysuserrole.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suruo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuserrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if suruo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysuserrole.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := suruo.mutation.DeletedAt(); ok {
		_spec.SetField(sysuserrole.FieldDeletedAt, field.TypeTime, value)
	}
	if suruo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysuserrole.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := suruo.mutation.UserID(); ok {
		_spec.SetField(sysuserrole.FieldUserID, field.TypeString, value)
	}
	if value, ok := suruo.mutation.RoleID(); ok {
		_spec.SetField(sysuserrole.FieldRoleID, field.TypeString, value)
	}
	_spec.AddModifiers(suruo.modifiers...)
	_node = &SysUserRole{config: suruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suruo.mutation.done = true
	return _node, nil
}
