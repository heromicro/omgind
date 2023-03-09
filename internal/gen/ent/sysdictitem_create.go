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
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemCreate is the builder for creating a SysDictItem entity.
type SysDictItemCreate struct {
	config
	mutation *SysDictItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (sdic *SysDictItemCreate) SetIsDel(b bool) *SysDictItemCreate {
	sdic.mutation.SetIsDel(b)
	return sdic
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableIsDel(b *bool) *SysDictItemCreate {
	if b != nil {
		sdic.SetIsDel(*b)
	}
	return sdic
}

// SetMemo sets the "memo" field.
func (sdic *SysDictItemCreate) SetMemo(s string) *SysDictItemCreate {
	sdic.mutation.SetMemo(s)
	return sdic
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableMemo(s *string) *SysDictItemCreate {
	if s != nil {
		sdic.SetMemo(*s)
	}
	return sdic
}

// SetSort sets the "sort" field.
func (sdic *SysDictItemCreate) SetSort(i int32) *SysDictItemCreate {
	sdic.mutation.SetSort(i)
	return sdic
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableSort(i *int32) *SysDictItemCreate {
	if i != nil {
		sdic.SetSort(*i)
	}
	return sdic
}

// SetCreatedAt sets the "created_at" field.
func (sdic *SysDictItemCreate) SetCreatedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetCreatedAt(t)
	return sdic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableCreatedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetCreatedAt(*t)
	}
	return sdic
}

// SetUpdatedAt sets the "updated_at" field.
func (sdic *SysDictItemCreate) SetUpdatedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetUpdatedAt(t)
	return sdic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableUpdatedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetUpdatedAt(*t)
	}
	return sdic
}

// SetDeletedAt sets the "deleted_at" field.
func (sdic *SysDictItemCreate) SetDeletedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetDeletedAt(t)
	return sdic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableDeletedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetDeletedAt(*t)
	}
	return sdic
}

// SetIsActive sets the "is_active" field.
func (sdic *SysDictItemCreate) SetIsActive(b bool) *SysDictItemCreate {
	sdic.mutation.SetIsActive(b)
	return sdic
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableIsActive(b *bool) *SysDictItemCreate {
	if b != nil {
		sdic.SetIsActive(*b)
	}
	return sdic
}

// SetLabel sets the "label" field.
func (sdic *SysDictItemCreate) SetLabel(s string) *SysDictItemCreate {
	sdic.mutation.SetLabel(s)
	return sdic
}

// SetValue sets the "value" field.
func (sdic *SysDictItemCreate) SetValue(i int) *SysDictItemCreate {
	sdic.mutation.SetValue(i)
	return sdic
}

// SetDictID sets the "dict_id" field.
func (sdic *SysDictItemCreate) SetDictID(s string) *SysDictItemCreate {
	sdic.mutation.SetDictID(s)
	return sdic
}

// SetID sets the "id" field.
func (sdic *SysDictItemCreate) SetID(s string) *SysDictItemCreate {
	sdic.mutation.SetID(s)
	return sdic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableID(s *string) *SysDictItemCreate {
	if s != nil {
		sdic.SetID(*s)
	}
	return sdic
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdic *SysDictItemCreate) Mutation() *SysDictItemMutation {
	return sdic.mutation
}

// Save creates the SysDictItem in the database.
func (sdic *SysDictItemCreate) Save(ctx context.Context) (*SysDictItem, error) {
	sdic.defaults()
	return withHooks[*SysDictItem, SysDictItemMutation](ctx, sdic.sqlSave, sdic.mutation, sdic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdic *SysDictItemCreate) SaveX(ctx context.Context) *SysDictItem {
	v, err := sdic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdic *SysDictItemCreate) Exec(ctx context.Context) error {
	_, err := sdic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdic *SysDictItemCreate) ExecX(ctx context.Context) {
	if err := sdic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdic *SysDictItemCreate) defaults() {
	if _, ok := sdic.mutation.IsDel(); !ok {
		v := sysdictitem.DefaultIsDel
		sdic.mutation.SetIsDel(v)
	}
	if _, ok := sdic.mutation.Memo(); !ok {
		v := sysdictitem.DefaultMemo
		sdic.mutation.SetMemo(v)
	}
	if _, ok := sdic.mutation.Sort(); !ok {
		v := sysdictitem.DefaultSort
		sdic.mutation.SetSort(v)
	}
	if _, ok := sdic.mutation.CreatedAt(); !ok {
		v := sysdictitem.DefaultCreatedAt()
		sdic.mutation.SetCreatedAt(v)
	}
	if _, ok := sdic.mutation.UpdatedAt(); !ok {
		v := sysdictitem.DefaultUpdatedAt()
		sdic.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdic.mutation.IsActive(); !ok {
		v := sysdictitem.DefaultIsActive
		sdic.mutation.SetIsActive(v)
	}
	if _, ok := sdic.mutation.ID(); !ok {
		v := sysdictitem.DefaultID()
		sdic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdic *SysDictItemCreate) check() error {
	if _, ok := sdic.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysDictItem.is_del"`)}
	}
	if v, ok := sdic.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysDictItem.memo": %w`, err)}
		}
	}
	if _, ok := sdic.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "SysDictItem.sort"`)}
	}
	if _, ok := sdic.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "SysDictItem.is_active"`)}
	}
	if _, ok := sdic.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "SysDictItem.label"`)}
	}
	if v, ok := sdic.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "SysDictItem.label": %w`, err)}
		}
	}
	if _, ok := sdic.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SysDictItem.value"`)}
	}
	if _, ok := sdic.mutation.DictID(); !ok {
		return &ValidationError{Name: "dict_id", err: errors.New(`ent: missing required field "SysDictItem.dict_id"`)}
	}
	if v, ok := sdic.mutation.DictID(); ok {
		if err := sysdictitem.DictIDValidator(v); err != nil {
			return &ValidationError{Name: "dict_id", err: fmt.Errorf(`ent: validator failed for field "SysDictItem.dict_id": %w`, err)}
		}
	}
	if v, ok := sdic.mutation.ID(); ok {
		if err := sysdictitem.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysDictItem.id": %w`, err)}
		}
	}
	return nil
}

func (sdic *SysDictItemCreate) sqlSave(ctx context.Context) (*SysDictItem, error) {
	if err := sdic.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdic.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysDictItem.ID type: %T", _spec.ID.Value)
		}
	}
	sdic.mutation.id = &_node.ID
	sdic.mutation.done = true
	return _node, nil
}

func (sdic *SysDictItemCreate) createSpec() (*SysDictItem, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDictItem{config: sdic.config}
		_spec = sqlgraph.NewCreateSpec(sysdictitem.Table, sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString))
	)
	_spec.Schema = sdic.schemaConfig.SysDictItem
	_spec.OnConflict = sdic.conflict
	if id, ok := sdic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdic.mutation.IsDel(); ok {
		_spec.SetField(sysdictitem.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := sdic.mutation.Memo(); ok {
		_spec.SetField(sysdictitem.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := sdic.mutation.Sort(); ok {
		_spec.SetField(sysdictitem.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := sdic.mutation.CreatedAt(); ok {
		_spec.SetField(sysdictitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := sdic.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdictitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sdic.mutation.DeletedAt(); ok {
		_spec.SetField(sysdictitem.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := sdic.mutation.IsActive(); ok {
		_spec.SetField(sysdictitem.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := sdic.mutation.Label(); ok {
		_spec.SetField(sysdictitem.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := sdic.mutation.Value(); ok {
		_spec.SetField(sysdictitem.FieldValue, field.TypeInt, value)
		_node.Value = value
	}
	if value, ok := sdic.mutation.DictID(); ok {
		_spec.SetField(sysdictitem.FieldDictID, field.TypeString, value)
		_node.DictID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDictItem.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictItemUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sdic *SysDictItemCreate) OnConflict(opts ...sql.ConflictOption) *SysDictItemUpsertOne {
	sdic.conflict = opts
	return &SysDictItemUpsertOne{
		create: sdic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdic *SysDictItemCreate) OnConflictColumns(columns ...string) *SysDictItemUpsertOne {
	sdic.conflict = append(sdic.conflict, sql.ConflictColumns(columns...))
	return &SysDictItemUpsertOne{
		create: sdic,
	}
}

type (
	// SysDictItemUpsertOne is the builder for "upsert"-ing
	//  one SysDictItem node.
	SysDictItemUpsertOne struct {
		create *SysDictItemCreate
	}

	// SysDictItemUpsert is the "OnConflict" setter.
	SysDictItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysDictItemUpsert) SetIsDel(v bool) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateIsDel() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldIsDel)
	return u
}

// SetMemo sets the "memo" field.
func (u *SysDictItemUpsert) SetMemo(v string) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateMemo() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictItemUpsert) ClearMemo() *SysDictItemUpsert {
	u.SetNull(sysdictitem.FieldMemo)
	return u
}

// SetSort sets the "sort" field.
func (u *SysDictItemUpsert) SetSort(v int32) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateSort() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysDictItemUpsert) AddSort(v int32) *SysDictItemUpsert {
	u.Add(sysdictitem.FieldSort, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictItemUpsert) SetUpdatedAt(v time.Time) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateUpdatedAt() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictItemUpsert) ClearUpdatedAt() *SysDictItemUpsert {
	u.SetNull(sysdictitem.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictItemUpsert) SetDeletedAt(v time.Time) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateDeletedAt() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictItemUpsert) ClearDeletedAt() *SysDictItemUpsert {
	u.SetNull(sysdictitem.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysDictItemUpsert) SetIsActive(v bool) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateIsActive() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldIsActive)
	return u
}

// SetLabel sets the "label" field.
func (u *SysDictItemUpsert) SetLabel(v string) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldLabel, v)
	return u
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateLabel() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldLabel)
	return u
}

// SetValue sets the "value" field.
func (u *SysDictItemUpsert) SetValue(v int) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateValue() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldValue)
	return u
}

// AddValue adds v to the "value" field.
func (u *SysDictItemUpsert) AddValue(v int) *SysDictItemUpsert {
	u.Add(sysdictitem.FieldValue, v)
	return u
}

// SetDictID sets the "dict_id" field.
func (u *SysDictItemUpsert) SetDictID(v string) *SysDictItemUpsert {
	u.Set(sysdictitem.FieldDictID, v)
	return u
}

// UpdateDictID sets the "dict_id" field to the value that was provided on create.
func (u *SysDictItemUpsert) UpdateDictID() *SysDictItemUpsert {
	u.SetExcluded(sysdictitem.FieldDictID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdictitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictItemUpsertOne) UpdateNewValues() *SysDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysdictitem.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysdictitem.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysDictItemUpsertOne) Ignore() *SysDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictItemUpsertOne) DoNothing() *SysDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictItemCreate.OnConflict
// documentation for more info.
func (u *SysDictItemUpsertOne) Update(set func(*SysDictItemUpsert)) *SysDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysDictItemUpsertOne) SetIsDel(v bool) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateIsDel() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysDictItemUpsertOne) SetMemo(v string) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateMemo() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictItemUpsertOne) ClearMemo() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysDictItemUpsertOne) SetSort(v int32) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysDictItemUpsertOne) AddSort(v int32) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateSort() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictItemUpsertOne) SetUpdatedAt(v time.Time) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateUpdatedAt() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictItemUpsertOne) ClearUpdatedAt() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictItemUpsertOne) SetDeletedAt(v time.Time) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateDeletedAt() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictItemUpsertOne) ClearDeletedAt() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysDictItemUpsertOne) SetIsActive(v bool) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateIsActive() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateIsActive()
	})
}

// SetLabel sets the "label" field.
func (u *SysDictItemUpsertOne) SetLabel(v string) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetLabel(v)
	})
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateLabel() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateLabel()
	})
}

// SetValue sets the "value" field.
func (u *SysDictItemUpsertOne) SetValue(v int) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SysDictItemUpsertOne) AddValue(v int) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateValue() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateValue()
	})
}

// SetDictID sets the "dict_id" field.
func (u *SysDictItemUpsertOne) SetDictID(v string) *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetDictID(v)
	})
}

// UpdateDictID sets the "dict_id" field to the value that was provided on create.
func (u *SysDictItemUpsertOne) UpdateDictID() *SysDictItemUpsertOne {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateDictID()
	})
}

// Exec executes the query.
func (u *SysDictItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysDictItemUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SysDictItemUpsertOne.ID is not supported by MySQL driver. Use SysDictItemUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysDictItemUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysDictItemCreateBulk is the builder for creating many SysDictItem entities in bulk.
type SysDictItemCreateBulk struct {
	config
	builders []*SysDictItemCreate
	conflict []sql.ConflictOption
}

// Save creates the SysDictItem entities in the database.
func (sdicb *SysDictItemCreateBulk) Save(ctx context.Context) ([]*SysDictItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdicb.builders))
	nodes := make([]*SysDictItem, len(sdicb.builders))
	mutators := make([]Mutator, len(sdicb.builders))
	for i := range sdicb.builders {
		func(i int, root context.Context) {
			builder := sdicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictItemMutation)
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
					_, err = mutators[i+1].Mutate(root, sdicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sdicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdicb *SysDictItemCreateBulk) SaveX(ctx context.Context) []*SysDictItem {
	v, err := sdicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdicb *SysDictItemCreateBulk) Exec(ctx context.Context) error {
	_, err := sdicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdicb *SysDictItemCreateBulk) ExecX(ctx context.Context) {
	if err := sdicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDictItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictItemUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sdicb *SysDictItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysDictItemUpsertBulk {
	sdicb.conflict = opts
	return &SysDictItemUpsertBulk{
		create: sdicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdicb *SysDictItemCreateBulk) OnConflictColumns(columns ...string) *SysDictItemUpsertBulk {
	sdicb.conflict = append(sdicb.conflict, sql.ConflictColumns(columns...))
	return &SysDictItemUpsertBulk{
		create: sdicb,
	}
}

// SysDictItemUpsertBulk is the builder for "upsert"-ing
// a bulk of SysDictItem nodes.
type SysDictItemUpsertBulk struct {
	create *SysDictItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdictitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictItemUpsertBulk) UpdateNewValues() *SysDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysdictitem.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysdictitem.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDictItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysDictItemUpsertBulk) Ignore() *SysDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictItemUpsertBulk) DoNothing() *SysDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictItemCreateBulk.OnConflict
// documentation for more info.
func (u *SysDictItemUpsertBulk) Update(set func(*SysDictItemUpsert)) *SysDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysDictItemUpsertBulk) SetIsDel(v bool) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateIsDel() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysDictItemUpsertBulk) SetMemo(v string) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateMemo() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictItemUpsertBulk) ClearMemo() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysDictItemUpsertBulk) SetSort(v int32) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysDictItemUpsertBulk) AddSort(v int32) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateSort() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictItemUpsertBulk) SetUpdatedAt(v time.Time) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateUpdatedAt() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictItemUpsertBulk) ClearUpdatedAt() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictItemUpsertBulk) SetDeletedAt(v time.Time) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateDeletedAt() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictItemUpsertBulk) ClearDeletedAt() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysDictItemUpsertBulk) SetIsActive(v bool) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateIsActive() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateIsActive()
	})
}

// SetLabel sets the "label" field.
func (u *SysDictItemUpsertBulk) SetLabel(v string) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetLabel(v)
	})
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateLabel() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateLabel()
	})
}

// SetValue sets the "value" field.
func (u *SysDictItemUpsertBulk) SetValue(v int) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SysDictItemUpsertBulk) AddValue(v int) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateValue() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateValue()
	})
}

// SetDictID sets the "dict_id" field.
func (u *SysDictItemUpsertBulk) SetDictID(v string) *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.SetDictID(v)
	})
}

// UpdateDictID sets the "dict_id" field to the value that was provided on create.
func (u *SysDictItemUpsertBulk) UpdateDictID() *SysDictItemUpsertBulk {
	return u.Update(func(s *SysDictItemUpsert) {
		s.UpdateDictID()
	})
}

// Exec executes the query.
func (u *SysDictItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysDictItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
