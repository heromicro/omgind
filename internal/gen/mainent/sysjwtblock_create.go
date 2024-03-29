// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/sysjwtblock"
)

// SysJwtBlockCreate is the builder for creating a SysJwtBlock entity.
type SysJwtBlockCreate struct {
	config
	mutation *SysJwtBlockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (sjbc *SysJwtBlockCreate) SetIsDel(b bool) *SysJwtBlockCreate {
	sjbc.mutation.SetIsDel(b)
	return sjbc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableIsDel(b *bool) *SysJwtBlockCreate {
	if b != nil {
		sjbc.SetIsDel(*b)
	}
	return sjbc
}

// SetMemo sets the "memo" field.
func (sjbc *SysJwtBlockCreate) SetMemo(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetMemo(s)
	return sjbc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableMemo(s *string) *SysJwtBlockCreate {
	if s != nil {
		sjbc.SetMemo(*s)
	}
	return sjbc
}

// SetCreatedAt sets the "created_at" field.
func (sjbc *SysJwtBlockCreate) SetCreatedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetCreatedAt(t)
	return sjbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableCreatedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetCreatedAt(*t)
	}
	return sjbc
}

// SetUpdatedAt sets the "updated_at" field.
func (sjbc *SysJwtBlockCreate) SetUpdatedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetUpdatedAt(t)
	return sjbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableUpdatedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetUpdatedAt(*t)
	}
	return sjbc
}

// SetDeletedAt sets the "deleted_at" field.
func (sjbc *SysJwtBlockCreate) SetDeletedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetDeletedAt(t)
	return sjbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableDeletedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetDeletedAt(*t)
	}
	return sjbc
}

// SetIsActive sets the "is_active" field.
func (sjbc *SysJwtBlockCreate) SetIsActive(b bool) *SysJwtBlockCreate {
	sjbc.mutation.SetIsActive(b)
	return sjbc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableIsActive(b *bool) *SysJwtBlockCreate {
	if b != nil {
		sjbc.SetIsActive(*b)
	}
	return sjbc
}

// SetJwt sets the "jwt" field.
func (sjbc *SysJwtBlockCreate) SetJwt(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetJwt(s)
	return sjbc
}

// SetID sets the "id" field.
func (sjbc *SysJwtBlockCreate) SetID(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetID(s)
	return sjbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableID(s *string) *SysJwtBlockCreate {
	if s != nil {
		sjbc.SetID(*s)
	}
	return sjbc
}

// Mutation returns the SysJwtBlockMutation object of the builder.
func (sjbc *SysJwtBlockCreate) Mutation() *SysJwtBlockMutation {
	return sjbc.mutation
}

// Save creates the SysJwtBlock in the database.
func (sjbc *SysJwtBlockCreate) Save(ctx context.Context) (*SysJwtBlock, error) {
	sjbc.defaults()
	return withHooks(ctx, sjbc.sqlSave, sjbc.mutation, sjbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sjbc *SysJwtBlockCreate) SaveX(ctx context.Context) *SysJwtBlock {
	v, err := sjbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sjbc *SysJwtBlockCreate) Exec(ctx context.Context) error {
	_, err := sjbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbc *SysJwtBlockCreate) ExecX(ctx context.Context) {
	if err := sjbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sjbc *SysJwtBlockCreate) defaults() {
	if _, ok := sjbc.mutation.IsDel(); !ok {
		v := sysjwtblock.DefaultIsDel
		sjbc.mutation.SetIsDel(v)
	}
	if _, ok := sjbc.mutation.Memo(); !ok {
		v := sysjwtblock.DefaultMemo
		sjbc.mutation.SetMemo(v)
	}
	if _, ok := sjbc.mutation.CreatedAt(); !ok {
		v := sysjwtblock.DefaultCreatedAt()
		sjbc.mutation.SetCreatedAt(v)
	}
	if _, ok := sjbc.mutation.UpdatedAt(); !ok {
		v := sysjwtblock.DefaultUpdatedAt()
		sjbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sjbc.mutation.IsActive(); !ok {
		v := sysjwtblock.DefaultIsActive
		sjbc.mutation.SetIsActive(v)
	}
	if _, ok := sjbc.mutation.ID(); !ok {
		v := sysjwtblock.DefaultID()
		sjbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sjbc *SysJwtBlockCreate) check() error {
	if _, ok := sjbc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`mainent: missing required field "SysJwtBlock.is_del"`)}
	}
	if v, ok := sjbc.mutation.Memo(); ok {
		if err := sysjwtblock.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysJwtBlock.memo": %w`, err)}
		}
	}
	if _, ok := sjbc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`mainent: missing required field "SysJwtBlock.is_active"`)}
	}
	if _, ok := sjbc.mutation.Jwt(); !ok {
		return &ValidationError{Name: "jwt", err: errors.New(`mainent: missing required field "SysJwtBlock.jwt"`)}
	}
	if v, ok := sjbc.mutation.Jwt(); ok {
		if err := sysjwtblock.JwtValidator(v); err != nil {
			return &ValidationError{Name: "jwt", err: fmt.Errorf(`mainent: validator failed for field "SysJwtBlock.jwt": %w`, err)}
		}
	}
	if v, ok := sjbc.mutation.ID(); ok {
		if err := sysjwtblock.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`mainent: validator failed for field "SysJwtBlock.id": %w`, err)}
		}
	}
	return nil
}

func (sjbc *SysJwtBlockCreate) sqlSave(ctx context.Context) (*SysJwtBlock, error) {
	if err := sjbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sjbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sjbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysJwtBlock.ID type: %T", _spec.ID.Value)
		}
	}
	sjbc.mutation.id = &_node.ID
	sjbc.mutation.done = true
	return _node, nil
}

func (sjbc *SysJwtBlockCreate) createSpec() (*SysJwtBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &SysJwtBlock{config: sjbc.config}
		_spec = sqlgraph.NewCreateSpec(sysjwtblock.Table, sqlgraph.NewFieldSpec(sysjwtblock.FieldID, field.TypeString))
	)
	_spec.OnConflict = sjbc.conflict
	if id, ok := sjbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sjbc.mutation.IsDel(); ok {
		_spec.SetField(sysjwtblock.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := sjbc.mutation.Memo(); ok {
		_spec.SetField(sysjwtblock.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := sjbc.mutation.CreatedAt(); ok {
		_spec.SetField(sysjwtblock.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := sjbc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysjwtblock.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sjbc.mutation.DeletedAt(); ok {
		_spec.SetField(sysjwtblock.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := sjbc.mutation.IsActive(); ok {
		_spec.SetField(sysjwtblock.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := sjbc.mutation.Jwt(); ok {
		_spec.SetField(sysjwtblock.FieldJwt, field.TypeString, value)
		_node.Jwt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysJwtBlock.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysJwtBlockUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sjbc *SysJwtBlockCreate) OnConflict(opts ...sql.ConflictOption) *SysJwtBlockUpsertOne {
	sjbc.conflict = opts
	return &SysJwtBlockUpsertOne{
		create: sjbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sjbc *SysJwtBlockCreate) OnConflictColumns(columns ...string) *SysJwtBlockUpsertOne {
	sjbc.conflict = append(sjbc.conflict, sql.ConflictColumns(columns...))
	return &SysJwtBlockUpsertOne{
		create: sjbc,
	}
}

type (
	// SysJwtBlockUpsertOne is the builder for "upsert"-ing
	//  one SysJwtBlock node.
	SysJwtBlockUpsertOne struct {
		create *SysJwtBlockCreate
	}

	// SysJwtBlockUpsert is the "OnConflict" setter.
	SysJwtBlockUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysJwtBlockUpsert) SetIsDel(v bool) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateIsDel() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldIsDel)
	return u
}

// SetMemo sets the "memo" field.
func (u *SysJwtBlockUpsert) SetMemo(v string) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateMemo() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysJwtBlockUpsert) ClearMemo() *SysJwtBlockUpsert {
	u.SetNull(sysjwtblock.FieldMemo)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysJwtBlockUpsert) SetUpdatedAt(v time.Time) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateUpdatedAt() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysJwtBlockUpsert) ClearUpdatedAt() *SysJwtBlockUpsert {
	u.SetNull(sysjwtblock.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysJwtBlockUpsert) SetDeletedAt(v time.Time) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateDeletedAt() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysJwtBlockUpsert) ClearDeletedAt() *SysJwtBlockUpsert {
	u.SetNull(sysjwtblock.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysJwtBlockUpsert) SetIsActive(v bool) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateIsActive() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldIsActive)
	return u
}

// SetJwt sets the "jwt" field.
func (u *SysJwtBlockUpsert) SetJwt(v string) *SysJwtBlockUpsert {
	u.Set(sysjwtblock.FieldJwt, v)
	return u
}

// UpdateJwt sets the "jwt" field to the value that was provided on create.
func (u *SysJwtBlockUpsert) UpdateJwt() *SysJwtBlockUpsert {
	u.SetExcluded(sysjwtblock.FieldJwt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysjwtblock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysJwtBlockUpsertOne) UpdateNewValues() *SysJwtBlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysjwtblock.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysjwtblock.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysJwtBlockUpsertOne) Ignore() *SysJwtBlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysJwtBlockUpsertOne) DoNothing() *SysJwtBlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysJwtBlockCreate.OnConflict
// documentation for more info.
func (u *SysJwtBlockUpsertOne) Update(set func(*SysJwtBlockUpsert)) *SysJwtBlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysJwtBlockUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysJwtBlockUpsertOne) SetIsDel(v bool) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateIsDel() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysJwtBlockUpsertOne) SetMemo(v string) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateMemo() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysJwtBlockUpsertOne) ClearMemo() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysJwtBlockUpsertOne) SetUpdatedAt(v time.Time) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateUpdatedAt() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysJwtBlockUpsertOne) ClearUpdatedAt() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysJwtBlockUpsertOne) SetDeletedAt(v time.Time) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateDeletedAt() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysJwtBlockUpsertOne) ClearDeletedAt() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysJwtBlockUpsertOne) SetIsActive(v bool) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateIsActive() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateIsActive()
	})
}

// SetJwt sets the "jwt" field.
func (u *SysJwtBlockUpsertOne) SetJwt(v string) *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetJwt(v)
	})
}

// UpdateJwt sets the "jwt" field to the value that was provided on create.
func (u *SysJwtBlockUpsertOne) UpdateJwt() *SysJwtBlockUpsertOne {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateJwt()
	})
}

// Exec executes the query.
func (u *SysJwtBlockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysJwtBlockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysJwtBlockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysJwtBlockUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("mainent: SysJwtBlockUpsertOne.ID is not supported by MySQL driver. Use SysJwtBlockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysJwtBlockUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysJwtBlockCreateBulk is the builder for creating many SysJwtBlock entities in bulk.
type SysJwtBlockCreateBulk struct {
	config
	err      error
	builders []*SysJwtBlockCreate
	conflict []sql.ConflictOption
}

// Save creates the SysJwtBlock entities in the database.
func (sjbcb *SysJwtBlockCreateBulk) Save(ctx context.Context) ([]*SysJwtBlock, error) {
	if sjbcb.err != nil {
		return nil, sjbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sjbcb.builders))
	nodes := make([]*SysJwtBlock, len(sjbcb.builders))
	mutators := make([]Mutator, len(sjbcb.builders))
	for i := range sjbcb.builders {
		func(i int, root context.Context) {
			builder := sjbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysJwtBlockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sjbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sjbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sjbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sjbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sjbcb *SysJwtBlockCreateBulk) SaveX(ctx context.Context) []*SysJwtBlock {
	v, err := sjbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sjbcb *SysJwtBlockCreateBulk) Exec(ctx context.Context) error {
	_, err := sjbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbcb *SysJwtBlockCreateBulk) ExecX(ctx context.Context) {
	if err := sjbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysJwtBlock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysJwtBlockUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sjbcb *SysJwtBlockCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysJwtBlockUpsertBulk {
	sjbcb.conflict = opts
	return &SysJwtBlockUpsertBulk{
		create: sjbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sjbcb *SysJwtBlockCreateBulk) OnConflictColumns(columns ...string) *SysJwtBlockUpsertBulk {
	sjbcb.conflict = append(sjbcb.conflict, sql.ConflictColumns(columns...))
	return &SysJwtBlockUpsertBulk{
		create: sjbcb,
	}
}

// SysJwtBlockUpsertBulk is the builder for "upsert"-ing
// a bulk of SysJwtBlock nodes.
type SysJwtBlockUpsertBulk struct {
	create *SysJwtBlockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysjwtblock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysJwtBlockUpsertBulk) UpdateNewValues() *SysJwtBlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysjwtblock.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysjwtblock.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysJwtBlock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysJwtBlockUpsertBulk) Ignore() *SysJwtBlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysJwtBlockUpsertBulk) DoNothing() *SysJwtBlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysJwtBlockCreateBulk.OnConflict
// documentation for more info.
func (u *SysJwtBlockUpsertBulk) Update(set func(*SysJwtBlockUpsert)) *SysJwtBlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysJwtBlockUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysJwtBlockUpsertBulk) SetIsDel(v bool) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateIsDel() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysJwtBlockUpsertBulk) SetMemo(v string) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateMemo() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysJwtBlockUpsertBulk) ClearMemo() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysJwtBlockUpsertBulk) SetUpdatedAt(v time.Time) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateUpdatedAt() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysJwtBlockUpsertBulk) ClearUpdatedAt() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysJwtBlockUpsertBulk) SetDeletedAt(v time.Time) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateDeletedAt() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysJwtBlockUpsertBulk) ClearDeletedAt() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysJwtBlockUpsertBulk) SetIsActive(v bool) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateIsActive() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateIsActive()
	})
}

// SetJwt sets the "jwt" field.
func (u *SysJwtBlockUpsertBulk) SetJwt(v string) *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.SetJwt(v)
	})
}

// UpdateJwt sets the "jwt" field to the value that was provided on create.
func (u *SysJwtBlockUpsertBulk) UpdateJwt() *SysJwtBlockUpsertBulk {
	return u.Update(func(s *SysJwtBlockUpsert) {
		s.UpdateJwt()
	})
}

// Exec executes the query.
func (u *SysJwtBlockUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("mainent: OnConflict was set for builder %d. Set it on the SysJwtBlockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysJwtBlockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysJwtBlockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
