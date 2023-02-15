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
	"github.com/heromicro/omgind/internal/gen/ent/sysuser"
)

// SysUserUpdate is the builder for updating SysUser entities.
type SysUserUpdate struct {
	config
	hooks    []Hook
	mutation *SysUserMutation
}

// Where appends a list predicates to the SysUserUpdate builder.
func (suu *SysUserUpdate) Where(ps ...predicate.SysUser) *SysUserUpdate {
	suu.mutation.Where(ps...)
	return suu
}

// SetIsDel sets the "is_del" field.
func (suu *SysUserUpdate) SetIsDel(b bool) *SysUserUpdate {
	suu.mutation.SetIsDel(b)
	return suu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableIsDel(b *bool) *SysUserUpdate {
	if b != nil {
		suu.SetIsDel(*b)
	}
	return suu
}

// SetSort sets the "sort" field.
func (suu *SysUserUpdate) SetSort(i int32) *SysUserUpdate {
	suu.mutation.ResetSort()
	suu.mutation.SetSort(i)
	return suu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableSort(i *int32) *SysUserUpdate {
	if i != nil {
		suu.SetSort(*i)
	}
	return suu
}

// AddSort adds i to the "sort" field.
func (suu *SysUserUpdate) AddSort(i int32) *SysUserUpdate {
	suu.mutation.AddSort(i)
	return suu
}

// SetUpdatedAt sets the "updated_at" field.
func (suu *SysUserUpdate) SetUpdatedAt(t time.Time) *SysUserUpdate {
	suu.mutation.SetUpdatedAt(t)
	return suu
}

// SetDeletedAt sets the "deleted_at" field.
func (suu *SysUserUpdate) SetDeletedAt(t time.Time) *SysUserUpdate {
	suu.mutation.SetDeletedAt(t)
	return suu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableDeletedAt(t *time.Time) *SysUserUpdate {
	if t != nil {
		suu.SetDeletedAt(*t)
	}
	return suu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suu *SysUserUpdate) ClearDeletedAt() *SysUserUpdate {
	suu.mutation.ClearDeletedAt()
	return suu
}

// SetIsActive sets the "is_active" field.
func (suu *SysUserUpdate) SetIsActive(b bool) *SysUserUpdate {
	suu.mutation.SetIsActive(b)
	return suu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableIsActive(b *bool) *SysUserUpdate {
	if b != nil {
		suu.SetIsActive(*b)
	}
	return suu
}

// SetRealName sets the "real_name" field.
func (suu *SysUserUpdate) SetRealName(s string) *SysUserUpdate {
	suu.mutation.SetRealName(s)
	return suu
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableRealName(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetRealName(*s)
	}
	return suu
}

// ClearRealName clears the value of the "real_name" field.
func (suu *SysUserUpdate) ClearRealName() *SysUserUpdate {
	suu.mutation.ClearRealName()
	return suu
}

// SetFirstName sets the "first_name" field.
func (suu *SysUserUpdate) SetFirstName(s string) *SysUserUpdate {
	suu.mutation.SetFirstName(s)
	return suu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableFirstName(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetFirstName(*s)
	}
	return suu
}

// ClearFirstName clears the value of the "first_name" field.
func (suu *SysUserUpdate) ClearFirstName() *SysUserUpdate {
	suu.mutation.ClearFirstName()
	return suu
}

// SetLastName sets the "last_name" field.
func (suu *SysUserUpdate) SetLastName(s string) *SysUserUpdate {
	suu.mutation.SetLastName(s)
	return suu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableLastName(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetLastName(*s)
	}
	return suu
}

// ClearLastName clears the value of the "last_name" field.
func (suu *SysUserUpdate) ClearLastName() *SysUserUpdate {
	suu.mutation.ClearLastName()
	return suu
}

// SetPassword sets the "Password" field.
func (suu *SysUserUpdate) SetPassword(s string) *SysUserUpdate {
	suu.mutation.SetPassword(s)
	return suu
}

// SetEmail sets the "Email" field.
func (suu *SysUserUpdate) SetEmail(s string) *SysUserUpdate {
	suu.mutation.SetEmail(s)
	return suu
}

// SetPhone sets the "Phone" field.
func (suu *SysUserUpdate) SetPhone(s string) *SysUserUpdate {
	suu.mutation.SetPhone(s)
	return suu
}

// SetSalt sets the "salt" field.
func (suu *SysUserUpdate) SetSalt(s string) *SysUserUpdate {
	suu.mutation.SetSalt(s)
	return suu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableSalt(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetSalt(*s)
	}
	return suu
}

// Mutation returns the SysUserMutation object of the builder.
func (suu *SysUserUpdate) Mutation() *SysUserMutation {
	return suu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suu *SysUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	suu.defaults()
	if len(suu.hooks) == 0 {
		if err = suu.check(); err != nil {
			return 0, err
		}
		affected, err = suu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suu.check(); err != nil {
				return 0, err
			}
			suu.mutation = mutation
			affected, err = suu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(suu.hooks) - 1; i >= 0; i-- {
			if suu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (suu *SysUserUpdate) SaveX(ctx context.Context) int {
	affected, err := suu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suu *SysUserUpdate) Exec(ctx context.Context) error {
	_, err := suu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suu *SysUserUpdate) ExecX(ctx context.Context) {
	if err := suu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suu *SysUserUpdate) defaults() {
	if _, ok := suu.mutation.UpdatedAt(); !ok {
		v := sysuser.UpdateDefaultUpdatedAt()
		suu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suu *SysUserUpdate) check() error {
	if v, ok := suu.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if v, ok := suu.mutation.FirstName(); ok {
		if err := sysuser.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.first_name": %w`, err)}
		}
	}
	if v, ok := suu.mutation.LastName(); ok {
		if err := sysuser.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.last_name": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "Password", err: fmt.Errorf(`ent: validator failed for field "SysUser.Password": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Email(); ok {
		if err := sysuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "SysUser.Email": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Phone(); ok {
		if err := sysuser.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "Phone", err: fmt.Errorf(`ent: validator failed for field "SysUser.Phone": %w`, err)}
		}
	}
	return nil
}

func (suu *SysUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysuser.Table,
			Columns: sysuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysuser.FieldID,
			},
		},
	}
	if ps := suu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysuser.FieldIsDel,
		})
	}
	if value, ok := suu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSort,
		})
	}
	if value, ok := suu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSort,
		})
	}
	if value, ok := suu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldUpdatedAt,
		})
	}
	if value, ok := suu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldDeletedAt,
		})
	}
	if suu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysuser.FieldDeletedAt,
		})
	}
	if value, ok := suu.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysuser.FieldIsActive,
		})
	}
	if value, ok := suu.mutation.RealName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldRealName,
		})
	}
	if suu.mutation.RealNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldRealName,
		})
	}
	if value, ok := suu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldFirstName,
		})
	}
	if suu.mutation.FirstNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldFirstName,
		})
	}
	if value, ok := suu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldLastName,
		})
	}
	if suu.mutation.LastNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldLastName,
		})
	}
	if value, ok := suu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPassword,
		})
	}
	if value, ok := suu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldEmail,
		})
	}
	if value, ok := suu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPhone,
		})
	}
	if value, ok := suu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldSalt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, suu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SysUserUpdateOne is the builder for updating a single SysUser entity.
type SysUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysUserMutation
}

// SetIsDel sets the "is_del" field.
func (suuo *SysUserUpdateOne) SetIsDel(b bool) *SysUserUpdateOne {
	suuo.mutation.SetIsDel(b)
	return suuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableIsDel(b *bool) *SysUserUpdateOne {
	if b != nil {
		suuo.SetIsDel(*b)
	}
	return suuo
}

// SetSort sets the "sort" field.
func (suuo *SysUserUpdateOne) SetSort(i int32) *SysUserUpdateOne {
	suuo.mutation.ResetSort()
	suuo.mutation.SetSort(i)
	return suuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableSort(i *int32) *SysUserUpdateOne {
	if i != nil {
		suuo.SetSort(*i)
	}
	return suuo
}

// AddSort adds i to the "sort" field.
func (suuo *SysUserUpdateOne) AddSort(i int32) *SysUserUpdateOne {
	suuo.mutation.AddSort(i)
	return suuo
}

// SetUpdatedAt sets the "updated_at" field.
func (suuo *SysUserUpdateOne) SetUpdatedAt(t time.Time) *SysUserUpdateOne {
	suuo.mutation.SetUpdatedAt(t)
	return suuo
}

// SetDeletedAt sets the "deleted_at" field.
func (suuo *SysUserUpdateOne) SetDeletedAt(t time.Time) *SysUserUpdateOne {
	suuo.mutation.SetDeletedAt(t)
	return suuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableDeletedAt(t *time.Time) *SysUserUpdateOne {
	if t != nil {
		suuo.SetDeletedAt(*t)
	}
	return suuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suuo *SysUserUpdateOne) ClearDeletedAt() *SysUserUpdateOne {
	suuo.mutation.ClearDeletedAt()
	return suuo
}

// SetIsActive sets the "is_active" field.
func (suuo *SysUserUpdateOne) SetIsActive(b bool) *SysUserUpdateOne {
	suuo.mutation.SetIsActive(b)
	return suuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableIsActive(b *bool) *SysUserUpdateOne {
	if b != nil {
		suuo.SetIsActive(*b)
	}
	return suuo
}

// SetRealName sets the "real_name" field.
func (suuo *SysUserUpdateOne) SetRealName(s string) *SysUserUpdateOne {
	suuo.mutation.SetRealName(s)
	return suuo
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableRealName(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetRealName(*s)
	}
	return suuo
}

// ClearRealName clears the value of the "real_name" field.
func (suuo *SysUserUpdateOne) ClearRealName() *SysUserUpdateOne {
	suuo.mutation.ClearRealName()
	return suuo
}

// SetFirstName sets the "first_name" field.
func (suuo *SysUserUpdateOne) SetFirstName(s string) *SysUserUpdateOne {
	suuo.mutation.SetFirstName(s)
	return suuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableFirstName(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetFirstName(*s)
	}
	return suuo
}

// ClearFirstName clears the value of the "first_name" field.
func (suuo *SysUserUpdateOne) ClearFirstName() *SysUserUpdateOne {
	suuo.mutation.ClearFirstName()
	return suuo
}

// SetLastName sets the "last_name" field.
func (suuo *SysUserUpdateOne) SetLastName(s string) *SysUserUpdateOne {
	suuo.mutation.SetLastName(s)
	return suuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableLastName(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetLastName(*s)
	}
	return suuo
}

// ClearLastName clears the value of the "last_name" field.
func (suuo *SysUserUpdateOne) ClearLastName() *SysUserUpdateOne {
	suuo.mutation.ClearLastName()
	return suuo
}

// SetPassword sets the "Password" field.
func (suuo *SysUserUpdateOne) SetPassword(s string) *SysUserUpdateOne {
	suuo.mutation.SetPassword(s)
	return suuo
}

// SetEmail sets the "Email" field.
func (suuo *SysUserUpdateOne) SetEmail(s string) *SysUserUpdateOne {
	suuo.mutation.SetEmail(s)
	return suuo
}

// SetPhone sets the "Phone" field.
func (suuo *SysUserUpdateOne) SetPhone(s string) *SysUserUpdateOne {
	suuo.mutation.SetPhone(s)
	return suuo
}

// SetSalt sets the "salt" field.
func (suuo *SysUserUpdateOne) SetSalt(s string) *SysUserUpdateOne {
	suuo.mutation.SetSalt(s)
	return suuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableSalt(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetSalt(*s)
	}
	return suuo
}

// Mutation returns the SysUserMutation object of the builder.
func (suuo *SysUserUpdateOne) Mutation() *SysUserMutation {
	return suuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suuo *SysUserUpdateOne) Select(field string, fields ...string) *SysUserUpdateOne {
	suuo.fields = append([]string{field}, fields...)
	return suuo
}

// Save executes the query and returns the updated SysUser entity.
func (suuo *SysUserUpdateOne) Save(ctx context.Context) (*SysUser, error) {
	var (
		err  error
		node *SysUser
	)
	suuo.defaults()
	if len(suuo.hooks) == 0 {
		if err = suuo.check(); err != nil {
			return nil, err
		}
		node, err = suuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suuo.check(); err != nil {
				return nil, err
			}
			suuo.mutation = mutation
			node, err = suuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suuo.hooks) - 1; i >= 0; i-- {
			if suuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suuo *SysUserUpdateOne) SaveX(ctx context.Context) *SysUser {
	node, err := suuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suuo *SysUserUpdateOne) Exec(ctx context.Context) error {
	_, err := suuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suuo *SysUserUpdateOne) ExecX(ctx context.Context) {
	if err := suuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suuo *SysUserUpdateOne) defaults() {
	if _, ok := suuo.mutation.UpdatedAt(); !ok {
		v := sysuser.UpdateDefaultUpdatedAt()
		suuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suuo *SysUserUpdateOne) check() error {
	if v, ok := suuo.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.FirstName(); ok {
		if err := sysuser.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.first_name": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.LastName(); ok {
		if err := sysuser.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.last_name": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "Password", err: fmt.Errorf(`ent: validator failed for field "SysUser.Password": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Email(); ok {
		if err := sysuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`ent: validator failed for field "SysUser.Email": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Phone(); ok {
		if err := sysuser.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "Phone", err: fmt.Errorf(`ent: validator failed for field "SysUser.Phone": %w`, err)}
		}
	}
	return nil
}

func (suuo *SysUserUpdateOne) sqlSave(ctx context.Context) (_node *SysUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysuser.Table,
			Columns: sysuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysuser.FieldID,
			},
		},
	}
	id, ok := suuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuser.FieldID)
		for _, f := range fields {
			if !sysuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suuo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysuser.FieldIsDel,
		})
	}
	if value, ok := suuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSort,
		})
	}
	if value, ok := suuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSort,
		})
	}
	if value, ok := suuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldUpdatedAt,
		})
	}
	if value, ok := suuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldDeletedAt,
		})
	}
	if suuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysuser.FieldDeletedAt,
		})
	}
	if value, ok := suuo.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysuser.FieldIsActive,
		})
	}
	if value, ok := suuo.mutation.RealName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldRealName,
		})
	}
	if suuo.mutation.RealNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldRealName,
		})
	}
	if value, ok := suuo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldFirstName,
		})
	}
	if suuo.mutation.FirstNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldFirstName,
		})
	}
	if value, ok := suuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldLastName,
		})
	}
	if suuo.mutation.LastNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysuser.FieldLastName,
		})
	}
	if value, ok := suuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPassword,
		})
	}
	if value, ok := suuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldEmail,
		})
	}
	if value, ok := suuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPhone,
		})
	}
	if value, ok := suuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldSalt,
		})
	}
	_node = &SysUser{config: suuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
