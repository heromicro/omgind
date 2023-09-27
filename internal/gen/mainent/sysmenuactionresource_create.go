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
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuactionresource"
)

// SysMenuActionResourceCreate is the builder for creating a SysMenuActionResource entity.
type SysMenuActionResourceCreate struct {
	config
	mutation *SysMenuActionResourceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (smarc *SysMenuActionResourceCreate) SetIsDel(b bool) *SysMenuActionResourceCreate {
	smarc.mutation.SetIsDel(b)
	return smarc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableIsDel(b *bool) *SysMenuActionResourceCreate {
	if b != nil {
		smarc.SetIsDel(*b)
	}
	return smarc
}

// SetSort sets the "sort" field.
func (smarc *SysMenuActionResourceCreate) SetSort(i int32) *SysMenuActionResourceCreate {
	smarc.mutation.SetSort(i)
	return smarc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableSort(i *int32) *SysMenuActionResourceCreate {
	if i != nil {
		smarc.SetSort(*i)
	}
	return smarc
}

// SetMemo sets the "memo" field.
func (smarc *SysMenuActionResourceCreate) SetMemo(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetMemo(s)
	return smarc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableMemo(s *string) *SysMenuActionResourceCreate {
	if s != nil {
		smarc.SetMemo(*s)
	}
	return smarc
}

// SetCreatedAt sets the "created_at" field.
func (smarc *SysMenuActionResourceCreate) SetCreatedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetCreatedAt(t)
	return smarc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableCreatedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetCreatedAt(*t)
	}
	return smarc
}

// SetUpdatedAt sets the "updated_at" field.
func (smarc *SysMenuActionResourceCreate) SetUpdatedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetUpdatedAt(t)
	return smarc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetUpdatedAt(*t)
	}
	return smarc
}

// SetDeletedAt sets the "deleted_at" field.
func (smarc *SysMenuActionResourceCreate) SetDeletedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetDeletedAt(t)
	return smarc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableDeletedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetDeletedAt(*t)
	}
	return smarc
}

// SetIsActive sets the "is_active" field.
func (smarc *SysMenuActionResourceCreate) SetIsActive(b bool) *SysMenuActionResourceCreate {
	smarc.mutation.SetIsActive(b)
	return smarc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableIsActive(b *bool) *SysMenuActionResourceCreate {
	if b != nil {
		smarc.SetIsActive(*b)
	}
	return smarc
}

// SetMethod sets the "method" field.
func (smarc *SysMenuActionResourceCreate) SetMethod(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetMethod(s)
	return smarc
}

// SetPath sets the "path" field.
func (smarc *SysMenuActionResourceCreate) SetPath(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetPath(s)
	return smarc
}

// SetActionID sets the "action_id" field.
func (smarc *SysMenuActionResourceCreate) SetActionID(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetActionID(s)
	return smarc
}

// SetID sets the "id" field.
func (smarc *SysMenuActionResourceCreate) SetID(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetID(s)
	return smarc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableID(s *string) *SysMenuActionResourceCreate {
	if s != nil {
		smarc.SetID(*s)
	}
	return smarc
}

// Mutation returns the SysMenuActionResourceMutation object of the builder.
func (smarc *SysMenuActionResourceCreate) Mutation() *SysMenuActionResourceMutation {
	return smarc.mutation
}

// Save creates the SysMenuActionResource in the database.
func (smarc *SysMenuActionResourceCreate) Save(ctx context.Context) (*SysMenuActionResource, error) {
	smarc.defaults()
	return withHooks(ctx, smarc.sqlSave, smarc.mutation, smarc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smarc *SysMenuActionResourceCreate) SaveX(ctx context.Context) *SysMenuActionResource {
	v, err := smarc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smarc *SysMenuActionResourceCreate) Exec(ctx context.Context) error {
	_, err := smarc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smarc *SysMenuActionResourceCreate) ExecX(ctx context.Context) {
	if err := smarc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smarc *SysMenuActionResourceCreate) defaults() {
	if _, ok := smarc.mutation.IsDel(); !ok {
		v := sysmenuactionresource.DefaultIsDel
		smarc.mutation.SetIsDel(v)
	}
	if _, ok := smarc.mutation.Sort(); !ok {
		v := sysmenuactionresource.DefaultSort
		smarc.mutation.SetSort(v)
	}
	if _, ok := smarc.mutation.Memo(); !ok {
		v := sysmenuactionresource.DefaultMemo
		smarc.mutation.SetMemo(v)
	}
	if _, ok := smarc.mutation.CreatedAt(); !ok {
		v := sysmenuactionresource.DefaultCreatedAt()
		smarc.mutation.SetCreatedAt(v)
	}
	if _, ok := smarc.mutation.UpdatedAt(); !ok {
		v := sysmenuactionresource.DefaultUpdatedAt()
		smarc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smarc.mutation.IsActive(); !ok {
		v := sysmenuactionresource.DefaultIsActive
		smarc.mutation.SetIsActive(v)
	}
	if _, ok := smarc.mutation.ID(); !ok {
		v := sysmenuactionresource.DefaultID()
		smarc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smarc *SysMenuActionResourceCreate) check() error {
	if _, ok := smarc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`mainent: missing required field "SysMenuActionResource.is_del"`)}
	}
	if _, ok := smarc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`mainent: missing required field "SysMenuActionResource.sort"`)}
	}
	if v, ok := smarc.mutation.Memo(); ok {
		if err := sysmenuactionresource.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysMenuActionResource.memo": %w`, err)}
		}
	}
	if _, ok := smarc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`mainent: missing required field "SysMenuActionResource.is_active"`)}
	}
	if _, ok := smarc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`mainent: missing required field "SysMenuActionResource.method"`)}
	}
	if v, ok := smarc.mutation.Method(); ok {
		if err := sysmenuactionresource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`mainent: validator failed for field "SysMenuActionResource.method": %w`, err)}
		}
	}
	if _, ok := smarc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`mainent: missing required field "SysMenuActionResource.path"`)}
	}
	if v, ok := smarc.mutation.Path(); ok {
		if err := sysmenuactionresource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`mainent: validator failed for field "SysMenuActionResource.path": %w`, err)}
		}
	}
	if _, ok := smarc.mutation.ActionID(); !ok {
		return &ValidationError{Name: "action_id", err: errors.New(`mainent: missing required field "SysMenuActionResource.action_id"`)}
	}
	if v, ok := smarc.mutation.ActionID(); ok {
		if err := sysmenuactionresource.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf(`mainent: validator failed for field "SysMenuActionResource.action_id": %w`, err)}
		}
	}
	if v, ok := smarc.mutation.ID(); ok {
		if err := sysmenuactionresource.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`mainent: validator failed for field "SysMenuActionResource.id": %w`, err)}
		}
	}
	return nil
}

func (smarc *SysMenuActionResourceCreate) sqlSave(ctx context.Context) (*SysMenuActionResource, error) {
	if err := smarc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smarc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smarc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysMenuActionResource.ID type: %T", _spec.ID.Value)
		}
	}
	smarc.mutation.id = &_node.ID
	smarc.mutation.done = true
	return _node, nil
}

func (smarc *SysMenuActionResourceCreate) createSpec() (*SysMenuActionResource, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuActionResource{config: smarc.config}
		_spec = sqlgraph.NewCreateSpec(sysmenuactionresource.Table, sqlgraph.NewFieldSpec(sysmenuactionresource.FieldID, field.TypeString))
	)
	_spec.OnConflict = smarc.conflict
	if id, ok := smarc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smarc.mutation.IsDel(); ok {
		_spec.SetField(sysmenuactionresource.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := smarc.mutation.Sort(); ok {
		_spec.SetField(sysmenuactionresource.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := smarc.mutation.Memo(); ok {
		_spec.SetField(sysmenuactionresource.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := smarc.mutation.CreatedAt(); ok {
		_spec.SetField(sysmenuactionresource.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := smarc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysmenuactionresource.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := smarc.mutation.DeletedAt(); ok {
		_spec.SetField(sysmenuactionresource.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := smarc.mutation.IsActive(); ok {
		_spec.SetField(sysmenuactionresource.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := smarc.mutation.Method(); ok {
		_spec.SetField(sysmenuactionresource.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := smarc.mutation.Path(); ok {
		_spec.SetField(sysmenuactionresource.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := smarc.mutation.ActionID(); ok {
		_spec.SetField(sysmenuactionresource.FieldActionID, field.TypeString, value)
		_node.ActionID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuActionResource.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuActionResourceUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (smarc *SysMenuActionResourceCreate) OnConflict(opts ...sql.ConflictOption) *SysMenuActionResourceUpsertOne {
	smarc.conflict = opts
	return &SysMenuActionResourceUpsertOne{
		create: smarc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smarc *SysMenuActionResourceCreate) OnConflictColumns(columns ...string) *SysMenuActionResourceUpsertOne {
	smarc.conflict = append(smarc.conflict, sql.ConflictColumns(columns...))
	return &SysMenuActionResourceUpsertOne{
		create: smarc,
	}
}

type (
	// SysMenuActionResourceUpsertOne is the builder for "upsert"-ing
	//  one SysMenuActionResource node.
	SysMenuActionResourceUpsertOne struct {
		create *SysMenuActionResourceCreate
	}

	// SysMenuActionResourceUpsert is the "OnConflict" setter.
	SysMenuActionResourceUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionResourceUpsert) SetIsDel(v bool) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateIsDel() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldIsDel)
	return u
}

// SetSort sets the "sort" field.
func (u *SysMenuActionResourceUpsert) SetSort(v int32) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateSort() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionResourceUpsert) AddSort(v int32) *SysMenuActionResourceUpsert {
	u.Add(sysmenuactionresource.FieldSort, v)
	return u
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionResourceUpsert) SetMemo(v string) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateMemo() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionResourceUpsert) ClearMemo() *SysMenuActionResourceUpsert {
	u.SetNull(sysmenuactionresource.FieldMemo)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionResourceUpsert) SetUpdatedAt(v time.Time) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateUpdatedAt() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionResourceUpsert) ClearUpdatedAt() *SysMenuActionResourceUpsert {
	u.SetNull(sysmenuactionresource.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionResourceUpsert) SetDeletedAt(v time.Time) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateDeletedAt() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionResourceUpsert) ClearDeletedAt() *SysMenuActionResourceUpsert {
	u.SetNull(sysmenuactionresource.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionResourceUpsert) SetIsActive(v bool) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateIsActive() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldIsActive)
	return u
}

// SetMethod sets the "method" field.
func (u *SysMenuActionResourceUpsert) SetMethod(v string) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateMethod() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldMethod)
	return u
}

// SetPath sets the "path" field.
func (u *SysMenuActionResourceUpsert) SetPath(v string) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdatePath() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldPath)
	return u
}

// SetActionID sets the "action_id" field.
func (u *SysMenuActionResourceUpsert) SetActionID(v string) *SysMenuActionResourceUpsert {
	u.Set(sysmenuactionresource.FieldActionID, v)
	return u
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsert) UpdateActionID() *SysMenuActionResourceUpsert {
	u.SetExcluded(sysmenuactionresource.FieldActionID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuactionresource.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuActionResourceUpsertOne) UpdateNewValues() *SysMenuActionResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysmenuactionresource.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysmenuactionresource.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysMenuActionResourceUpsertOne) Ignore() *SysMenuActionResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuActionResourceUpsertOne) DoNothing() *SysMenuActionResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuActionResourceCreate.OnConflict
// documentation for more info.
func (u *SysMenuActionResourceUpsertOne) Update(set func(*SysMenuActionResourceUpsert)) *SysMenuActionResourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuActionResourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionResourceUpsertOne) SetIsDel(v bool) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateIsDel() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysMenuActionResourceUpsertOne) SetSort(v int32) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionResourceUpsertOne) AddSort(v int32) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateSort() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateSort()
	})
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionResourceUpsertOne) SetMemo(v string) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateMemo() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionResourceUpsertOne) ClearMemo() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionResourceUpsertOne) SetUpdatedAt(v time.Time) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateUpdatedAt() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionResourceUpsertOne) ClearUpdatedAt() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionResourceUpsertOne) SetDeletedAt(v time.Time) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateDeletedAt() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionResourceUpsertOne) ClearDeletedAt() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionResourceUpsertOne) SetIsActive(v bool) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateIsActive() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateIsActive()
	})
}

// SetMethod sets the "method" field.
func (u *SysMenuActionResourceUpsertOne) SetMethod(v string) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateMethod() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateMethod()
	})
}

// SetPath sets the "path" field.
func (u *SysMenuActionResourceUpsertOne) SetPath(v string) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdatePath() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdatePath()
	})
}

// SetActionID sets the "action_id" field.
func (u *SysMenuActionResourceUpsertOne) SetActionID(v string) *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetActionID(v)
	})
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertOne) UpdateActionID() *SysMenuActionResourceUpsertOne {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateActionID()
	})
}

// Exec executes the query.
func (u *SysMenuActionResourceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysMenuActionResourceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuActionResourceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysMenuActionResourceUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("mainent: SysMenuActionResourceUpsertOne.ID is not supported by MySQL driver. Use SysMenuActionResourceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysMenuActionResourceUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysMenuActionResourceCreateBulk is the builder for creating many SysMenuActionResource entities in bulk.
type SysMenuActionResourceCreateBulk struct {
	config
	err      error
	builders []*SysMenuActionResourceCreate
	conflict []sql.ConflictOption
}

// Save creates the SysMenuActionResource entities in the database.
func (smarcb *SysMenuActionResourceCreateBulk) Save(ctx context.Context) ([]*SysMenuActionResource, error) {
	if smarcb.err != nil {
		return nil, smarcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smarcb.builders))
	nodes := make([]*SysMenuActionResource, len(smarcb.builders))
	mutators := make([]Mutator, len(smarcb.builders))
	for i := range smarcb.builders {
		func(i int, root context.Context) {
			builder := smarcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuActionResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, smarcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = smarcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smarcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smarcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smarcb *SysMenuActionResourceCreateBulk) SaveX(ctx context.Context) []*SysMenuActionResource {
	v, err := smarcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smarcb *SysMenuActionResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := smarcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smarcb *SysMenuActionResourceCreateBulk) ExecX(ctx context.Context) {
	if err := smarcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuActionResource.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuActionResourceUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (smarcb *SysMenuActionResourceCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysMenuActionResourceUpsertBulk {
	smarcb.conflict = opts
	return &SysMenuActionResourceUpsertBulk{
		create: smarcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smarcb *SysMenuActionResourceCreateBulk) OnConflictColumns(columns ...string) *SysMenuActionResourceUpsertBulk {
	smarcb.conflict = append(smarcb.conflict, sql.ConflictColumns(columns...))
	return &SysMenuActionResourceUpsertBulk{
		create: smarcb,
	}
}

// SysMenuActionResourceUpsertBulk is the builder for "upsert"-ing
// a bulk of SysMenuActionResource nodes.
type SysMenuActionResourceUpsertBulk struct {
	create *SysMenuActionResourceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuactionresource.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuActionResourceUpsertBulk) UpdateNewValues() *SysMenuActionResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysmenuactionresource.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysmenuactionresource.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuActionResource.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysMenuActionResourceUpsertBulk) Ignore() *SysMenuActionResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuActionResourceUpsertBulk) DoNothing() *SysMenuActionResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuActionResourceCreateBulk.OnConflict
// documentation for more info.
func (u *SysMenuActionResourceUpsertBulk) Update(set func(*SysMenuActionResourceUpsert)) *SysMenuActionResourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuActionResourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionResourceUpsertBulk) SetIsDel(v bool) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateIsDel() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysMenuActionResourceUpsertBulk) SetSort(v int32) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionResourceUpsertBulk) AddSort(v int32) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateSort() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateSort()
	})
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionResourceUpsertBulk) SetMemo(v string) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateMemo() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionResourceUpsertBulk) ClearMemo() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionResourceUpsertBulk) SetUpdatedAt(v time.Time) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateUpdatedAt() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionResourceUpsertBulk) ClearUpdatedAt() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionResourceUpsertBulk) SetDeletedAt(v time.Time) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateDeletedAt() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionResourceUpsertBulk) ClearDeletedAt() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionResourceUpsertBulk) SetIsActive(v bool) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateIsActive() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateIsActive()
	})
}

// SetMethod sets the "method" field.
func (u *SysMenuActionResourceUpsertBulk) SetMethod(v string) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateMethod() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateMethod()
	})
}

// SetPath sets the "path" field.
func (u *SysMenuActionResourceUpsertBulk) SetPath(v string) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdatePath() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdatePath()
	})
}

// SetActionID sets the "action_id" field.
func (u *SysMenuActionResourceUpsertBulk) SetActionID(v string) *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.SetActionID(v)
	})
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *SysMenuActionResourceUpsertBulk) UpdateActionID() *SysMenuActionResourceUpsertBulk {
	return u.Update(func(s *SysMenuActionResourceUpsert) {
		s.UpdateActionID()
	})
}

// Exec executes the query.
func (u *SysMenuActionResourceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("mainent: OnConflict was set for builder %d. Set it on the SysMenuActionResourceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysMenuActionResourceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuActionResourceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
