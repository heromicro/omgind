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
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuActionCreate is the builder for creating a SysMenuAction entity.
type SysMenuActionCreate struct {
	config
	mutation *SysMenuActionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (smac *SysMenuActionCreate) SetIsDel(b bool) *SysMenuActionCreate {
	smac.mutation.SetIsDel(b)
	return smac
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableIsDel(b *bool) *SysMenuActionCreate {
	if b != nil {
		smac.SetIsDel(*b)
	}
	return smac
}

// SetSort sets the "sort" field.
func (smac *SysMenuActionCreate) SetSort(i int32) *SysMenuActionCreate {
	smac.mutation.SetSort(i)
	return smac
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableSort(i *int32) *SysMenuActionCreate {
	if i != nil {
		smac.SetSort(*i)
	}
	return smac
}

// SetIsActive sets the "is_active" field.
func (smac *SysMenuActionCreate) SetIsActive(b bool) *SysMenuActionCreate {
	smac.mutation.SetIsActive(b)
	return smac
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableIsActive(b *bool) *SysMenuActionCreate {
	if b != nil {
		smac.SetIsActive(*b)
	}
	return smac
}

// SetMemo sets the "memo" field.
func (smac *SysMenuActionCreate) SetMemo(s string) *SysMenuActionCreate {
	smac.mutation.SetMemo(s)
	return smac
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableMemo(s *string) *SysMenuActionCreate {
	if s != nil {
		smac.SetMemo(*s)
	}
	return smac
}

// SetCreatedAt sets the "created_at" field.
func (smac *SysMenuActionCreate) SetCreatedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetCreatedAt(t)
	return smac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableCreatedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetCreatedAt(*t)
	}
	return smac
}

// SetUpdatedAt sets the "updated_at" field.
func (smac *SysMenuActionCreate) SetUpdatedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetUpdatedAt(t)
	return smac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetUpdatedAt(*t)
	}
	return smac
}

// SetDeletedAt sets the "deleted_at" field.
func (smac *SysMenuActionCreate) SetDeletedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetDeletedAt(t)
	return smac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableDeletedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetDeletedAt(*t)
	}
	return smac
}

// SetMenuID sets the "menu_id" field.
func (smac *SysMenuActionCreate) SetMenuID(s string) *SysMenuActionCreate {
	smac.mutation.SetMenuID(s)
	return smac
}

// SetCode sets the "code" field.
func (smac *SysMenuActionCreate) SetCode(s string) *SysMenuActionCreate {
	smac.mutation.SetCode(s)
	return smac
}

// SetName sets the "name" field.
func (smac *SysMenuActionCreate) SetName(s string) *SysMenuActionCreate {
	smac.mutation.SetName(s)
	return smac
}

// SetID sets the "id" field.
func (smac *SysMenuActionCreate) SetID(s string) *SysMenuActionCreate {
	smac.mutation.SetID(s)
	return smac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableID(s *string) *SysMenuActionCreate {
	if s != nil {
		smac.SetID(*s)
	}
	return smac
}

// Mutation returns the SysMenuActionMutation object of the builder.
func (smac *SysMenuActionCreate) Mutation() *SysMenuActionMutation {
	return smac.mutation
}

// Save creates the SysMenuAction in the database.
func (smac *SysMenuActionCreate) Save(ctx context.Context) (*SysMenuAction, error) {
	smac.defaults()
	return withHooks[*SysMenuAction, SysMenuActionMutation](ctx, smac.sqlSave, smac.mutation, smac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smac *SysMenuActionCreate) SaveX(ctx context.Context) *SysMenuAction {
	v, err := smac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smac *SysMenuActionCreate) Exec(ctx context.Context) error {
	_, err := smac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smac *SysMenuActionCreate) ExecX(ctx context.Context) {
	if err := smac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smac *SysMenuActionCreate) defaults() {
	if _, ok := smac.mutation.IsDel(); !ok {
		v := sysmenuaction.DefaultIsDel
		smac.mutation.SetIsDel(v)
	}
	if _, ok := smac.mutation.Sort(); !ok {
		v := sysmenuaction.DefaultSort
		smac.mutation.SetSort(v)
	}
	if _, ok := smac.mutation.IsActive(); !ok {
		v := sysmenuaction.DefaultIsActive
		smac.mutation.SetIsActive(v)
	}
	if _, ok := smac.mutation.Memo(); !ok {
		v := sysmenuaction.DefaultMemo
		smac.mutation.SetMemo(v)
	}
	if _, ok := smac.mutation.CreatedAt(); !ok {
		v := sysmenuaction.DefaultCreatedAt()
		smac.mutation.SetCreatedAt(v)
	}
	if _, ok := smac.mutation.UpdatedAt(); !ok {
		v := sysmenuaction.DefaultUpdatedAt()
		smac.mutation.SetUpdatedAt(v)
	}
	if _, ok := smac.mutation.ID(); !ok {
		v := sysmenuaction.DefaultID()
		smac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smac *SysMenuActionCreate) check() error {
	if _, ok := smac.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysMenuAction.is_del"`)}
	}
	if _, ok := smac.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "SysMenuAction.sort"`)}
	}
	if _, ok := smac.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "SysMenuAction.is_active"`)}
	}
	if v, ok := smac.mutation.Memo(); ok {
		if err := sysmenuaction.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.memo": %w`, err)}
		}
	}
	if _, ok := smac.mutation.MenuID(); !ok {
		return &ValidationError{Name: "menu_id", err: errors.New(`ent: missing required field "SysMenuAction.menu_id"`)}
	}
	if v, ok := smac.mutation.MenuID(); ok {
		if err := sysmenuaction.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.menu_id": %w`, err)}
		}
	}
	if _, ok := smac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "SysMenuAction.code"`)}
	}
	if v, ok := smac.mutation.Code(); ok {
		if err := sysmenuaction.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.code": %w`, err)}
		}
	}
	if _, ok := smac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SysMenuAction.name"`)}
	}
	if v, ok := smac.mutation.Name(); ok {
		if err := sysmenuaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.name": %w`, err)}
		}
	}
	if v, ok := smac.mutation.ID(); ok {
		if err := sysmenuaction.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysMenuAction.id": %w`, err)}
		}
	}
	return nil
}

func (smac *SysMenuActionCreate) sqlSave(ctx context.Context) (*SysMenuAction, error) {
	if err := smac.check(); err != nil {
		return nil, err
	}
	_node, _spec := smac.createSpec()
	if err := sqlgraph.CreateNode(ctx, smac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysMenuAction.ID type: %T", _spec.ID.Value)
		}
	}
	smac.mutation.id = &_node.ID
	smac.mutation.done = true
	return _node, nil
}

func (smac *SysMenuActionCreate) createSpec() (*SysMenuAction, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuAction{config: smac.config}
		_spec = sqlgraph.NewCreateSpec(sysmenuaction.Table, sqlgraph.NewFieldSpec(sysmenuaction.FieldID, field.TypeString))
	)
	_spec.OnConflict = smac.conflict
	if id, ok := smac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smac.mutation.IsDel(); ok {
		_spec.SetField(sysmenuaction.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := smac.mutation.Sort(); ok {
		_spec.SetField(sysmenuaction.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := smac.mutation.IsActive(); ok {
		_spec.SetField(sysmenuaction.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := smac.mutation.Memo(); ok {
		_spec.SetField(sysmenuaction.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := smac.mutation.CreatedAt(); ok {
		_spec.SetField(sysmenuaction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := smac.mutation.UpdatedAt(); ok {
		_spec.SetField(sysmenuaction.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := smac.mutation.DeletedAt(); ok {
		_spec.SetField(sysmenuaction.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := smac.mutation.MenuID(); ok {
		_spec.SetField(sysmenuaction.FieldMenuID, field.TypeString, value)
		_node.MenuID = value
	}
	if value, ok := smac.mutation.Code(); ok {
		_spec.SetField(sysmenuaction.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := smac.mutation.Name(); ok {
		_spec.SetField(sysmenuaction.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuAction.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuActionUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (smac *SysMenuActionCreate) OnConflict(opts ...sql.ConflictOption) *SysMenuActionUpsertOne {
	smac.conflict = opts
	return &SysMenuActionUpsertOne{
		create: smac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smac *SysMenuActionCreate) OnConflictColumns(columns ...string) *SysMenuActionUpsertOne {
	smac.conflict = append(smac.conflict, sql.ConflictColumns(columns...))
	return &SysMenuActionUpsertOne{
		create: smac,
	}
}

type (
	// SysMenuActionUpsertOne is the builder for "upsert"-ing
	//  one SysMenuAction node.
	SysMenuActionUpsertOne struct {
		create *SysMenuActionCreate
	}

	// SysMenuActionUpsert is the "OnConflict" setter.
	SysMenuActionUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionUpsert) SetIsDel(v bool) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateIsDel() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldIsDel)
	return u
}

// SetSort sets the "sort" field.
func (u *SysMenuActionUpsert) SetSort(v int32) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateSort() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionUpsert) AddSort(v int32) *SysMenuActionUpsert {
	u.Add(sysmenuaction.FieldSort, v)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionUpsert) SetIsActive(v bool) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateIsActive() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldIsActive)
	return u
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionUpsert) SetMemo(v string) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateMemo() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionUpsert) ClearMemo() *SysMenuActionUpsert {
	u.SetNull(sysmenuaction.FieldMemo)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionUpsert) SetUpdatedAt(v time.Time) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateUpdatedAt() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionUpsert) ClearUpdatedAt() *SysMenuActionUpsert {
	u.SetNull(sysmenuaction.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionUpsert) SetDeletedAt(v time.Time) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateDeletedAt() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionUpsert) ClearDeletedAt() *SysMenuActionUpsert {
	u.SetNull(sysmenuaction.FieldDeletedAt)
	return u
}

// SetMenuID sets the "menu_id" field.
func (u *SysMenuActionUpsert) SetMenuID(v string) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldMenuID, v)
	return u
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateMenuID() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldMenuID)
	return u
}

// SetCode sets the "code" field.
func (u *SysMenuActionUpsert) SetCode(v string) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateCode() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldCode)
	return u
}

// SetName sets the "name" field.
func (u *SysMenuActionUpsert) SetName(v string) *SysMenuActionUpsert {
	u.Set(sysmenuaction.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysMenuActionUpsert) UpdateName() *SysMenuActionUpsert {
	u.SetExcluded(sysmenuaction.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuActionUpsertOne) UpdateNewValues() *SysMenuActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysmenuaction.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysmenuaction.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysMenuActionUpsertOne) Ignore() *SysMenuActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuActionUpsertOne) DoNothing() *SysMenuActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuActionCreate.OnConflict
// documentation for more info.
func (u *SysMenuActionUpsertOne) Update(set func(*SysMenuActionUpsert)) *SysMenuActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuActionUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionUpsertOne) SetIsDel(v bool) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateIsDel() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysMenuActionUpsertOne) SetSort(v int32) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionUpsertOne) AddSort(v int32) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateSort() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateSort()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionUpsertOne) SetIsActive(v bool) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateIsActive() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateIsActive()
	})
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionUpsertOne) SetMemo(v string) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateMemo() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionUpsertOne) ClearMemo() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionUpsertOne) SetUpdatedAt(v time.Time) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateUpdatedAt() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionUpsertOne) ClearUpdatedAt() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionUpsertOne) SetDeletedAt(v time.Time) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateDeletedAt() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionUpsertOne) ClearDeletedAt() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetMenuID sets the "menu_id" field.
func (u *SysMenuActionUpsertOne) SetMenuID(v string) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetMenuID(v)
	})
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateMenuID() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateMenuID()
	})
}

// SetCode sets the "code" field.
func (u *SysMenuActionUpsertOne) SetCode(v string) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateCode() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateCode()
	})
}

// SetName sets the "name" field.
func (u *SysMenuActionUpsertOne) SetName(v string) *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysMenuActionUpsertOne) UpdateName() *SysMenuActionUpsertOne {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *SysMenuActionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysMenuActionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuActionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysMenuActionUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SysMenuActionUpsertOne.ID is not supported by MySQL driver. Use SysMenuActionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysMenuActionUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysMenuActionCreateBulk is the builder for creating many SysMenuAction entities in bulk.
type SysMenuActionCreateBulk struct {
	config
	builders []*SysMenuActionCreate
	conflict []sql.ConflictOption
}

// Save creates the SysMenuAction entities in the database.
func (smacb *SysMenuActionCreateBulk) Save(ctx context.Context) ([]*SysMenuAction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smacb.builders))
	nodes := make([]*SysMenuAction, len(smacb.builders))
	mutators := make([]Mutator, len(smacb.builders))
	for i := range smacb.builders {
		func(i int, root context.Context) {
			builder := smacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuActionMutation)
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
					_, err = mutators[i+1].Mutate(root, smacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = smacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smacb *SysMenuActionCreateBulk) SaveX(ctx context.Context) []*SysMenuAction {
	v, err := smacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smacb *SysMenuActionCreateBulk) Exec(ctx context.Context) error {
	_, err := smacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smacb *SysMenuActionCreateBulk) ExecX(ctx context.Context) {
	if err := smacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysMenuAction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysMenuActionUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (smacb *SysMenuActionCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysMenuActionUpsertBulk {
	smacb.conflict = opts
	return &SysMenuActionUpsertBulk{
		create: smacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (smacb *SysMenuActionCreateBulk) OnConflictColumns(columns ...string) *SysMenuActionUpsertBulk {
	smacb.conflict = append(smacb.conflict, sql.ConflictColumns(columns...))
	return &SysMenuActionUpsertBulk{
		create: smacb,
	}
}

// SysMenuActionUpsertBulk is the builder for "upsert"-ing
// a bulk of SysMenuAction nodes.
type SysMenuActionUpsertBulk struct {
	create *SysMenuActionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysmenuaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysMenuActionUpsertBulk) UpdateNewValues() *SysMenuActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysmenuaction.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysmenuaction.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysMenuAction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysMenuActionUpsertBulk) Ignore() *SysMenuActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysMenuActionUpsertBulk) DoNothing() *SysMenuActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysMenuActionCreateBulk.OnConflict
// documentation for more info.
func (u *SysMenuActionUpsertBulk) Update(set func(*SysMenuActionUpsert)) *SysMenuActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysMenuActionUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysMenuActionUpsertBulk) SetIsDel(v bool) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateIsDel() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysMenuActionUpsertBulk) SetSort(v int32) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysMenuActionUpsertBulk) AddSort(v int32) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateSort() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateSort()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysMenuActionUpsertBulk) SetIsActive(v bool) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateIsActive() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateIsActive()
	})
}

// SetMemo sets the "memo" field.
func (u *SysMenuActionUpsertBulk) SetMemo(v string) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateMemo() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysMenuActionUpsertBulk) ClearMemo() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearMemo()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysMenuActionUpsertBulk) SetUpdatedAt(v time.Time) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateUpdatedAt() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysMenuActionUpsertBulk) ClearUpdatedAt() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysMenuActionUpsertBulk) SetDeletedAt(v time.Time) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateDeletedAt() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysMenuActionUpsertBulk) ClearDeletedAt() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetMenuID sets the "menu_id" field.
func (u *SysMenuActionUpsertBulk) SetMenuID(v string) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetMenuID(v)
	})
}

// UpdateMenuID sets the "menu_id" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateMenuID() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateMenuID()
	})
}

// SetCode sets the "code" field.
func (u *SysMenuActionUpsertBulk) SetCode(v string) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateCode() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateCode()
	})
}

// SetName sets the "name" field.
func (u *SysMenuActionUpsertBulk) SetName(v string) *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SysMenuActionUpsertBulk) UpdateName() *SysMenuActionUpsertBulk {
	return u.Update(func(s *SysMenuActionUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *SysMenuActionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysMenuActionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysMenuActionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysMenuActionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
