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
	"github.com/heromicro/omgind/internal/gen/mainent/sysannex"
)

// SysAnnexCreate is the builder for creating a SysAnnex entity.
type SysAnnexCreate struct {
	config
	mutation *SysAnnexMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetMemo sets the "memo" field.
func (sac *SysAnnexCreate) SetMemo(s string) *SysAnnexCreate {
	sac.mutation.SetMemo(s)
	return sac
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableMemo(s *string) *SysAnnexCreate {
	if s != nil {
		sac.SetMemo(*s)
	}
	return sac
}

// SetSort sets the "sort" field.
func (sac *SysAnnexCreate) SetSort(i int32) *SysAnnexCreate {
	sac.mutation.SetSort(i)
	return sac
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableSort(i *int32) *SysAnnexCreate {
	if i != nil {
		sac.SetSort(*i)
	}
	return sac
}

// SetCreatedAt sets the "created_at" field.
func (sac *SysAnnexCreate) SetCreatedAt(t time.Time) *SysAnnexCreate {
	sac.mutation.SetCreatedAt(t)
	return sac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableCreatedAt(t *time.Time) *SysAnnexCreate {
	if t != nil {
		sac.SetCreatedAt(*t)
	}
	return sac
}

// SetUpdatedAt sets the "updated_at" field.
func (sac *SysAnnexCreate) SetUpdatedAt(t time.Time) *SysAnnexCreate {
	sac.mutation.SetUpdatedAt(t)
	return sac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableUpdatedAt(t *time.Time) *SysAnnexCreate {
	if t != nil {
		sac.SetUpdatedAt(*t)
	}
	return sac
}

// SetDeletedAt sets the "deleted_at" field.
func (sac *SysAnnexCreate) SetDeletedAt(t time.Time) *SysAnnexCreate {
	sac.mutation.SetDeletedAt(t)
	return sac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableDeletedAt(t *time.Time) *SysAnnexCreate {
	if t != nil {
		sac.SetDeletedAt(*t)
	}
	return sac
}

// SetIsActive sets the "is_active" field.
func (sac *SysAnnexCreate) SetIsActive(b bool) *SysAnnexCreate {
	sac.mutation.SetIsActive(b)
	return sac
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableIsActive(b *bool) *SysAnnexCreate {
	if b != nil {
		sac.SetIsActive(*b)
	}
	return sac
}

// SetIsDel sets the "is_del" field.
func (sac *SysAnnexCreate) SetIsDel(b bool) *SysAnnexCreate {
	sac.mutation.SetIsDel(b)
	return sac
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableIsDel(b *bool) *SysAnnexCreate {
	if b != nil {
		sac.SetIsDel(*b)
	}
	return sac
}

// SetName sets the "name" field.
func (sac *SysAnnexCreate) SetName(s string) *SysAnnexCreate {
	sac.mutation.SetName(s)
	return sac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableName(s *string) *SysAnnexCreate {
	if s != nil {
		sac.SetName(*s)
	}
	return sac
}

// SetFilePath sets the "file_path" field.
func (sac *SysAnnexCreate) SetFilePath(s string) *SysAnnexCreate {
	sac.mutation.SetFilePath(s)
	return sac
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableFilePath(s *string) *SysAnnexCreate {
	if s != nil {
		sac.SetFilePath(*s)
	}
	return sac
}

// SetID sets the "id" field.
func (sac *SysAnnexCreate) SetID(s string) *SysAnnexCreate {
	sac.mutation.SetID(s)
	return sac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sac *SysAnnexCreate) SetNillableID(s *string) *SysAnnexCreate {
	if s != nil {
		sac.SetID(*s)
	}
	return sac
}

// Mutation returns the SysAnnexMutation object of the builder.
func (sac *SysAnnexCreate) Mutation() *SysAnnexMutation {
	return sac.mutation
}

// Save creates the SysAnnex in the database.
func (sac *SysAnnexCreate) Save(ctx context.Context) (*SysAnnex, error) {
	sac.defaults()
	return withHooks(ctx, sac.sqlSave, sac.mutation, sac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sac *SysAnnexCreate) SaveX(ctx context.Context) *SysAnnex {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *SysAnnexCreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *SysAnnexCreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sac *SysAnnexCreate) defaults() {
	if _, ok := sac.mutation.Memo(); !ok {
		v := sysannex.DefaultMemo
		sac.mutation.SetMemo(v)
	}
	if _, ok := sac.mutation.Sort(); !ok {
		v := sysannex.DefaultSort
		sac.mutation.SetSort(v)
	}
	if _, ok := sac.mutation.CreatedAt(); !ok {
		v := sysannex.DefaultCreatedAt()
		sac.mutation.SetCreatedAt(v)
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		v := sysannex.DefaultUpdatedAt()
		sac.mutation.SetUpdatedAt(v)
	}
	if _, ok := sac.mutation.IsActive(); !ok {
		v := sysannex.DefaultIsActive
		sac.mutation.SetIsActive(v)
	}
	if _, ok := sac.mutation.IsDel(); !ok {
		v := sysannex.DefaultIsDel
		sac.mutation.SetIsDel(v)
	}
	if _, ok := sac.mutation.ID(); !ok {
		v := sysannex.DefaultID()
		sac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *SysAnnexCreate) check() error {
	if v, ok := sac.mutation.Memo(); ok {
		if err := sysannex.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysAnnex.memo": %w`, err)}
		}
	}
	if _, ok := sac.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`mainent: missing required field "SysAnnex.sort"`)}
	}
	if _, ok := sac.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`mainent: missing required field "SysAnnex.is_active"`)}
	}
	if _, ok := sac.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`mainent: missing required field "SysAnnex.is_del"`)}
	}
	if v, ok := sac.mutation.Name(); ok {
		if err := sysannex.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`mainent: validator failed for field "SysAnnex.name": %w`, err)}
		}
	}
	if v, ok := sac.mutation.ID(); ok {
		if err := sysannex.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`mainent: validator failed for field "SysAnnex.id": %w`, err)}
		}
	}
	return nil
}

func (sac *SysAnnexCreate) sqlSave(ctx context.Context) (*SysAnnex, error) {
	if err := sac.check(); err != nil {
		return nil, err
	}
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysAnnex.ID type: %T", _spec.ID.Value)
		}
	}
	sac.mutation.id = &_node.ID
	sac.mutation.done = true
	return _node, nil
}

func (sac *SysAnnexCreate) createSpec() (*SysAnnex, *sqlgraph.CreateSpec) {
	var (
		_node = &SysAnnex{config: sac.config}
		_spec = sqlgraph.NewCreateSpec(sysannex.Table, sqlgraph.NewFieldSpec(sysannex.FieldID, field.TypeString))
	)
	_spec.OnConflict = sac.conflict
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.Memo(); ok {
		_spec.SetField(sysannex.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := sac.mutation.Sort(); ok {
		_spec.SetField(sysannex.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := sac.mutation.CreatedAt(); ok {
		_spec.SetField(sysannex.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := sac.mutation.UpdatedAt(); ok {
		_spec.SetField(sysannex.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sac.mutation.DeletedAt(); ok {
		_spec.SetField(sysannex.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := sac.mutation.IsActive(); ok {
		_spec.SetField(sysannex.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := sac.mutation.IsDel(); ok {
		_spec.SetField(sysannex.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := sac.mutation.Name(); ok {
		_spec.SetField(sysannex.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := sac.mutation.FilePath(); ok {
		_spec.SetField(sysannex.FieldFilePath, field.TypeString, value)
		_node.FilePath = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysAnnex.Create().
//		SetMemo(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysAnnexUpsert) {
//			SetMemo(v+v).
//		}).
//		Exec(ctx)
func (sac *SysAnnexCreate) OnConflict(opts ...sql.ConflictOption) *SysAnnexUpsertOne {
	sac.conflict = opts
	return &SysAnnexUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sac *SysAnnexCreate) OnConflictColumns(columns ...string) *SysAnnexUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &SysAnnexUpsertOne{
		create: sac,
	}
}

type (
	// SysAnnexUpsertOne is the builder for "upsert"-ing
	//  one SysAnnex node.
	SysAnnexUpsertOne struct {
		create *SysAnnexCreate
	}

	// SysAnnexUpsert is the "OnConflict" setter.
	SysAnnexUpsert struct {
		*sql.UpdateSet
	}
)

// SetMemo sets the "memo" field.
func (u *SysAnnexUpsert) SetMemo(v string) *SysAnnexUpsert {
	u.Set(sysannex.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateMemo() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysAnnexUpsert) ClearMemo() *SysAnnexUpsert {
	u.SetNull(sysannex.FieldMemo)
	return u
}

// SetSort sets the "sort" field.
func (u *SysAnnexUpsert) SetSort(v int32) *SysAnnexUpsert {
	u.Set(sysannex.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateSort() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysAnnexUpsert) AddSort(v int32) *SysAnnexUpsert {
	u.Add(sysannex.FieldSort, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysAnnexUpsert) SetUpdatedAt(v time.Time) *SysAnnexUpsert {
	u.Set(sysannex.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateUpdatedAt() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysAnnexUpsert) ClearUpdatedAt() *SysAnnexUpsert {
	u.SetNull(sysannex.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysAnnexUpsert) SetDeletedAt(v time.Time) *SysAnnexUpsert {
	u.Set(sysannex.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateDeletedAt() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysAnnexUpsert) ClearDeletedAt() *SysAnnexUpsert {
	u.SetNull(sysannex.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysAnnexUpsert) SetIsActive(v bool) *SysAnnexUpsert {
	u.Set(sysannex.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateIsActive() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldIsActive)
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysAnnexUpsert) SetIsDel(v bool) *SysAnnexUpsert {
	u.Set(sysannex.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateIsDel() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldIsDel)
	return u
}

// SetName sets the "name" field.
func (u *SysAnnexUpsert) SetName(v string) *SysAnnexUpsert {
	u.Set(sysannex.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateName() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *SysAnnexUpsert) ClearName() *SysAnnexUpsert {
	u.SetNull(sysannex.FieldName)
	return u
}

// SetFilePath sets the "file_path" field.
func (u *SysAnnexUpsert) SetFilePath(v string) *SysAnnexUpsert {
	u.Set(sysannex.FieldFilePath, v)
	return u
}

// UpdateFilePath sets the "file_path" field to the value that was provided on create.
func (u *SysAnnexUpsert) UpdateFilePath() *SysAnnexUpsert {
	u.SetExcluded(sysannex.FieldFilePath)
	return u
}

// ClearFilePath clears the value of the "file_path" field.
func (u *SysAnnexUpsert) ClearFilePath() *SysAnnexUpsert {
	u.SetNull(sysannex.FieldFilePath)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysannex.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysAnnexUpsertOne) UpdateNewValues() *SysAnnexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysannex.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysannex.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysAnnexUpsertOne) Ignore() *SysAnnexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysAnnexUpsertOne) DoNothing() *SysAnnexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysAnnexCreate.OnConflict
// documentation for more info.
func (u *SysAnnexUpsertOne) Update(set func(*SysAnnexUpsert)) *SysAnnexUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysAnnexUpsert{UpdateSet: update})
	}))
	return u
}

// SetMemo sets the "memo" field.
func (u *SysAnnexUpsertOne) SetMemo(v string) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateMemo() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysAnnexUpsertOne) ClearMemo() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysAnnexUpsertOne) SetSort(v int32) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysAnnexUpsertOne) AddSort(v int32) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateSort() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysAnnexUpsertOne) SetUpdatedAt(v time.Time) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateUpdatedAt() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysAnnexUpsertOne) ClearUpdatedAt() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysAnnexUpsertOne) SetDeletedAt(v time.Time) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateDeletedAt() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysAnnexUpsertOne) ClearDeletedAt() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysAnnexUpsertOne) SetIsActive(v bool) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateIsActive() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateIsActive()
	})
}

// SetIsDel sets the "is_del" field.
func (u *SysAnnexUpsertOne) SetIsDel(v bool) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateIsDel() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateIsDel()
	})
}

// SetName sets the "name" field.
func (u *SysAnnexUpsertOne) SetName(v string) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateName() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *SysAnnexUpsertOne) ClearName() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearName()
	})
}

// SetFilePath sets the "file_path" field.
func (u *SysAnnexUpsertOne) SetFilePath(v string) *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetFilePath(v)
	})
}

// UpdateFilePath sets the "file_path" field to the value that was provided on create.
func (u *SysAnnexUpsertOne) UpdateFilePath() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateFilePath()
	})
}

// ClearFilePath clears the value of the "file_path" field.
func (u *SysAnnexUpsertOne) ClearFilePath() *SysAnnexUpsertOne {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearFilePath()
	})
}

// Exec executes the query.
func (u *SysAnnexUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysAnnexCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysAnnexUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysAnnexUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("mainent: SysAnnexUpsertOne.ID is not supported by MySQL driver. Use SysAnnexUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysAnnexUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysAnnexCreateBulk is the builder for creating many SysAnnex entities in bulk.
type SysAnnexCreateBulk struct {
	config
	err      error
	builders []*SysAnnexCreate
	conflict []sql.ConflictOption
}

// Save creates the SysAnnex entities in the database.
func (sacb *SysAnnexCreateBulk) Save(ctx context.Context) ([]*SysAnnex, error) {
	if sacb.err != nil {
		return nil, sacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SysAnnex, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysAnnexMutation)
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
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *SysAnnexCreateBulk) SaveX(ctx context.Context) []*SysAnnex {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *SysAnnexCreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *SysAnnexCreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysAnnex.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysAnnexUpsert) {
//			SetMemo(v+v).
//		}).
//		Exec(ctx)
func (sacb *SysAnnexCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysAnnexUpsertBulk {
	sacb.conflict = opts
	return &SysAnnexUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sacb *SysAnnexCreateBulk) OnConflictColumns(columns ...string) *SysAnnexUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &SysAnnexUpsertBulk{
		create: sacb,
	}
}

// SysAnnexUpsertBulk is the builder for "upsert"-ing
// a bulk of SysAnnex nodes.
type SysAnnexUpsertBulk struct {
	create *SysAnnexCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysannex.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysAnnexUpsertBulk) UpdateNewValues() *SysAnnexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysannex.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysannex.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysAnnex.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysAnnexUpsertBulk) Ignore() *SysAnnexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysAnnexUpsertBulk) DoNothing() *SysAnnexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysAnnexCreateBulk.OnConflict
// documentation for more info.
func (u *SysAnnexUpsertBulk) Update(set func(*SysAnnexUpsert)) *SysAnnexUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysAnnexUpsert{UpdateSet: update})
	}))
	return u
}

// SetMemo sets the "memo" field.
func (u *SysAnnexUpsertBulk) SetMemo(v string) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateMemo() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysAnnexUpsertBulk) ClearMemo() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysAnnexUpsertBulk) SetSort(v int32) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysAnnexUpsertBulk) AddSort(v int32) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateSort() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysAnnexUpsertBulk) SetUpdatedAt(v time.Time) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateUpdatedAt() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysAnnexUpsertBulk) ClearUpdatedAt() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysAnnexUpsertBulk) SetDeletedAt(v time.Time) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateDeletedAt() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysAnnexUpsertBulk) ClearDeletedAt() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysAnnexUpsertBulk) SetIsActive(v bool) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateIsActive() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateIsActive()
	})
}

// SetIsDel sets the "is_del" field.
func (u *SysAnnexUpsertBulk) SetIsDel(v bool) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateIsDel() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateIsDel()
	})
}

// SetName sets the "name" field.
func (u *SysAnnexUpsertBulk) SetName(v string) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateName() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *SysAnnexUpsertBulk) ClearName() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearName()
	})
}

// SetFilePath sets the "file_path" field.
func (u *SysAnnexUpsertBulk) SetFilePath(v string) *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.SetFilePath(v)
	})
}

// UpdateFilePath sets the "file_path" field to the value that was provided on create.
func (u *SysAnnexUpsertBulk) UpdateFilePath() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.UpdateFilePath()
	})
}

// ClearFilePath clears the value of the "file_path" field.
func (u *SysAnnexUpsertBulk) ClearFilePath() *SysAnnexUpsertBulk {
	return u.Update(func(s *SysAnnexUpsert) {
		s.ClearFilePath()
	})
}

// Exec executes the query.
func (u *SysAnnexUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("mainent: OnConflict was set for builder %d. Set it on the SysAnnexCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("mainent: missing options for SysAnnexCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysAnnexUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
