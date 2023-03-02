// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRoleCreate is the builder for creating a SysUserRole entity.
type SysUserRoleCreate struct {
	config
	mutation *SysUserRoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (surc *SysUserRoleCreate) SetIsDel(b bool) *SysUserRoleCreate {
	surc.mutation.SetIsDel(b)
	return surc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableIsDel(b *bool) *SysUserRoleCreate {
	if b != nil {
		surc.SetIsDel(*b)
	}
	return surc
}

// SetCreatedAt sets the "created_at" field.
func (surc *SysUserRoleCreate) SetCreatedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetCreatedAt(t)
	return surc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableCreatedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetCreatedAt(*t)
	}
	return surc
}

// SetUpdatedAt sets the "updated_at" field.
func (surc *SysUserRoleCreate) SetUpdatedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetUpdatedAt(t)
	return surc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableUpdatedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetUpdatedAt(*t)
	}
	return surc
}

// SetDeletedAt sets the "deleted_at" field.
func (surc *SysUserRoleCreate) SetDeletedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetDeletedAt(t)
	return surc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableDeletedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetDeletedAt(*t)
	}
	return surc
}

// SetUserID sets the "user_id" field.
func (surc *SysUserRoleCreate) SetUserID(s string) *SysUserRoleCreate {
	surc.mutation.SetUserID(s)
	return surc
}

// SetRoleID sets the "role_id" field.
func (surc *SysUserRoleCreate) SetRoleID(s string) *SysUserRoleCreate {
	surc.mutation.SetRoleID(s)
	return surc
}

// SetID sets the "id" field.
func (surc *SysUserRoleCreate) SetID(s string) *SysUserRoleCreate {
	surc.mutation.SetID(s)
	return surc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableID(s *string) *SysUserRoleCreate {
	if s != nil {
		surc.SetID(*s)
	}
	return surc
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (surc *SysUserRoleCreate) Mutation() *SysUserRoleMutation {
	return surc.mutation
}

// Save creates the SysUserRole in the database.
func (surc *SysUserRoleCreate) Save(ctx context.Context) (*SysUserRole, error) {
	surc.defaults()
	return withHooks[*SysUserRole, SysUserRoleMutation](ctx, surc.sqlSave, surc.mutation, surc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (surc *SysUserRoleCreate) SaveX(ctx context.Context) *SysUserRole {
	v, err := surc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (surc *SysUserRoleCreate) Exec(ctx context.Context) error {
	_, err := surc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (surc *SysUserRoleCreate) ExecX(ctx context.Context) {
	if err := surc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (surc *SysUserRoleCreate) defaults() {
	if _, ok := surc.mutation.IsDel(); !ok {
		v := sysuserrole.DefaultIsDel
		surc.mutation.SetIsDel(v)
	}
	if _, ok := surc.mutation.CreatedAt(); !ok {
		v := sysuserrole.DefaultCreatedAt()
		surc.mutation.SetCreatedAt(v)
	}
	if _, ok := surc.mutation.UpdatedAt(); !ok {
		v := sysuserrole.DefaultUpdatedAt()
		surc.mutation.SetUpdatedAt(v)
	}
	if _, ok := surc.mutation.ID(); !ok {
		v := sysuserrole.DefaultID()
		surc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (surc *SysUserRoleCreate) check() error {
	if _, ok := surc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysUserRole.is_del"`)}
	}
	if _, ok := surc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysUserRole.created_at"`)}
	}
	if _, ok := surc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysUserRole.updated_at"`)}
	}
	if _, ok := surc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "SysUserRole.user_id"`)}
	}
	if v, ok := surc.mutation.UserID(); ok {
		if err := sysuserrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "SysUserRole.user_id": %w`, err)}
		}
	}
	if _, ok := surc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "SysUserRole.role_id"`)}
	}
	if v, ok := surc.mutation.RoleID(); ok {
		if err := sysuserrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "SysUserRole.role_id": %w`, err)}
		}
	}
	if v, ok := surc.mutation.ID(); ok {
		if err := sysuserrole.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysUserRole.id": %w`, err)}
		}
	}
	return nil
}

func (surc *SysUserRoleCreate) sqlSave(ctx context.Context) (*SysUserRole, error) {
	if err := surc.check(); err != nil {
		return nil, err
	}
	_node, _spec := surc.createSpec()
	if err := sqlgraph.CreateNode(ctx, surc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysUserRole.ID type: %T", _spec.ID.Value)
		}
	}
	surc.mutation.id = &_node.ID
	surc.mutation.done = true
	return _node, nil
}

func (surc *SysUserRoleCreate) createSpec() (*SysUserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUserRole{config: surc.config}
		_spec = sqlgraph.NewCreateSpec(sysuserrole.Table, sqlgraph.NewFieldSpec(sysuserrole.FieldID, field.TypeString))
	)
	_spec.OnConflict = surc.conflict
	if id, ok := surc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := surc.mutation.IsDel(); ok {
		_spec.SetField(sysuserrole.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := surc.mutation.CreatedAt(); ok {
		_spec.SetField(sysuserrole.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := surc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuserrole.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := surc.mutation.DeletedAt(); ok {
		_spec.SetField(sysuserrole.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := surc.mutation.UserID(); ok {
		_spec.SetField(sysuserrole.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := surc.mutation.RoleID(); ok {
		_spec.SetField(sysuserrole.FieldRoleID, field.TypeString, value)
		_node.RoleID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysUserRole.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysUserRoleUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (surc *SysUserRoleCreate) OnConflict(opts ...sql.ConflictOption) *SysUserRoleUpsertOne {
	surc.conflict = opts
	return &SysUserRoleUpsertOne{
		create: surc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (surc *SysUserRoleCreate) OnConflictColumns(columns ...string) *SysUserRoleUpsertOne {
	surc.conflict = append(surc.conflict, sql.ConflictColumns(columns...))
	return &SysUserRoleUpsertOne{
		create: surc,
	}
}

type (
	// SysUserRoleUpsertOne is the builder for "upsert"-ing
	//  one SysUserRole node.
	SysUserRoleUpsertOne struct {
		create *SysUserRoleCreate
	}

	// SysUserRoleUpsert is the "OnConflict" setter.
	SysUserRoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysUserRoleUpsert) SetIsDel(v bool) *SysUserRoleUpsert {
	u.Set(sysuserrole.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserRoleUpsert) UpdateIsDel() *SysUserRoleUpsert {
	u.SetExcluded(sysuserrole.FieldIsDel)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserRoleUpsert) SetUpdatedAt(v time.Time) *SysUserRoleUpsert {
	u.Set(sysuserrole.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserRoleUpsert) UpdateUpdatedAt() *SysUserRoleUpsert {
	u.SetExcluded(sysuserrole.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserRoleUpsert) SetDeletedAt(v time.Time) *SysUserRoleUpsert {
	u.Set(sysuserrole.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserRoleUpsert) UpdateDeletedAt() *SysUserRoleUpsert {
	u.SetExcluded(sysuserrole.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserRoleUpsert) ClearDeletedAt() *SysUserRoleUpsert {
	u.SetNull(sysuserrole.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *SysUserRoleUpsert) SetUserID(v string) *SysUserRoleUpsert {
	u.Set(sysuserrole.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SysUserRoleUpsert) UpdateUserID() *SysUserRoleUpsert {
	u.SetExcluded(sysuserrole.FieldUserID)
	return u
}

// SetRoleID sets the "role_id" field.
func (u *SysUserRoleUpsert) SetRoleID(v string) *SysUserRoleUpsert {
	u.Set(sysuserrole.FieldRoleID, v)
	return u
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *SysUserRoleUpsert) UpdateRoleID() *SysUserRoleUpsert {
	u.SetExcluded(sysuserrole.FieldRoleID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysuserrole.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysUserRoleUpsertOne) UpdateNewValues() *SysUserRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysuserrole.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysuserrole.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysUserRoleUpsertOne) Ignore() *SysUserRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysUserRoleUpsertOne) DoNothing() *SysUserRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysUserRoleCreate.OnConflict
// documentation for more info.
func (u *SysUserRoleUpsertOne) Update(set func(*SysUserRoleUpsert)) *SysUserRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysUserRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysUserRoleUpsertOne) SetIsDel(v bool) *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserRoleUpsertOne) UpdateIsDel() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateIsDel()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserRoleUpsertOne) SetUpdatedAt(v time.Time) *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserRoleUpsertOne) UpdateUpdatedAt() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserRoleUpsertOne) SetDeletedAt(v time.Time) *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserRoleUpsertOne) UpdateDeletedAt() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserRoleUpsertOne) ClearDeletedAt() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *SysUserRoleUpsertOne) SetUserID(v string) *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SysUserRoleUpsertOne) UpdateUserID() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateUserID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *SysUserRoleUpsertOne) SetRoleID(v string) *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *SysUserRoleUpsertOne) UpdateRoleID() *SysUserRoleUpsertOne {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateRoleID()
	})
}

// Exec executes the query.
func (u *SysUserRoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysUserRoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysUserRoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysUserRoleUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SysUserRoleUpsertOne.ID is not supported by MySQL driver. Use SysUserRoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysUserRoleUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysUserRoleCreateBulk is the builder for creating many SysUserRole entities in bulk.
type SysUserRoleCreateBulk struct {
	config
	builders []*SysUserRoleCreate
	conflict []sql.ConflictOption
}

// Save creates the SysUserRole entities in the database.
func (surcb *SysUserRoleCreateBulk) Save(ctx context.Context) ([]*SysUserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(surcb.builders))
	nodes := make([]*SysUserRole, len(surcb.builders))
	mutators := make([]Mutator, len(surcb.builders))
	for i := range surcb.builders {
		func(i int, root context.Context) {
			builder := surcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserRoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, surcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = surcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, surcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, surcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (surcb *SysUserRoleCreateBulk) SaveX(ctx context.Context) []*SysUserRole {
	v, err := surcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (surcb *SysUserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := surcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (surcb *SysUserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := surcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysUserRole.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysUserRoleUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (surcb *SysUserRoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysUserRoleUpsertBulk {
	surcb.conflict = opts
	return &SysUserRoleUpsertBulk{
		create: surcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (surcb *SysUserRoleCreateBulk) OnConflictColumns(columns ...string) *SysUserRoleUpsertBulk {
	surcb.conflict = append(surcb.conflict, sql.ConflictColumns(columns...))
	return &SysUserRoleUpsertBulk{
		create: surcb,
	}
}

// SysUserRoleUpsertBulk is the builder for "upsert"-ing
// a bulk of SysUserRole nodes.
type SysUserRoleUpsertBulk struct {
	create *SysUserRoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysuserrole.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysUserRoleUpsertBulk) UpdateNewValues() *SysUserRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysuserrole.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysuserrole.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysUserRole.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysUserRoleUpsertBulk) Ignore() *SysUserRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysUserRoleUpsertBulk) DoNothing() *SysUserRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysUserRoleCreateBulk.OnConflict
// documentation for more info.
func (u *SysUserRoleUpsertBulk) Update(set func(*SysUserRoleUpsert)) *SysUserRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysUserRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysUserRoleUpsertBulk) SetIsDel(v bool) *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserRoleUpsertBulk) UpdateIsDel() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateIsDel()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserRoleUpsertBulk) SetUpdatedAt(v time.Time) *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserRoleUpsertBulk) UpdateUpdatedAt() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserRoleUpsertBulk) SetDeletedAt(v time.Time) *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserRoleUpsertBulk) UpdateDeletedAt() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserRoleUpsertBulk) ClearDeletedAt() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *SysUserRoleUpsertBulk) SetUserID(v string) *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SysUserRoleUpsertBulk) UpdateUserID() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateUserID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *SysUserRoleUpsertBulk) SetRoleID(v string) *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *SysUserRoleUpsertBulk) UpdateRoleID() *SysUserRoleUpsertBulk {
	return u.Update(func(s *SysUserRoleUpsert) {
		s.UpdateRoleID()
	})
}

// Exec executes the query.
func (u *SysUserRoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysUserRoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysUserRoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysUserRoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
