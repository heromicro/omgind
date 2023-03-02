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
	"github.com/heromicro/omgind/internal/gen/ent/sysuser"
)

// SysUserCreate is the builder for creating a SysUser entity.
type SysUserCreate struct {
	config
	mutation *SysUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (suc *SysUserCreate) SetIsDel(b bool) *SysUserCreate {
	suc.mutation.SetIsDel(b)
	return suc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableIsDel(b *bool) *SysUserCreate {
	if b != nil {
		suc.SetIsDel(*b)
	}
	return suc
}

// SetSort sets the "sort" field.
func (suc *SysUserCreate) SetSort(i int32) *SysUserCreate {
	suc.mutation.SetSort(i)
	return suc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSort(i *int32) *SysUserCreate {
	if i != nil {
		suc.SetSort(*i)
	}
	return suc
}

// SetCreatedAt sets the "created_at" field.
func (suc *SysUserCreate) SetCreatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetCreatedAt(t)
	return suc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetCreatedAt(*t)
	}
	return suc
}

// SetUpdatedAt sets the "updated_at" field.
func (suc *SysUserCreate) SetUpdatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetUpdatedAt(t)
	return suc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetUpdatedAt(*t)
	}
	return suc
}

// SetDeletedAt sets the "deleted_at" field.
func (suc *SysUserCreate) SetDeletedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetDeletedAt(t)
	return suc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeletedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetDeletedAt(*t)
	}
	return suc
}

// SetIsActive sets the "is_active" field.
func (suc *SysUserCreate) SetIsActive(b bool) *SysUserCreate {
	suc.mutation.SetIsActive(b)
	return suc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableIsActive(b *bool) *SysUserCreate {
	if b != nil {
		suc.SetIsActive(*b)
	}
	return suc
}

// SetUserName sets the "user_name" field.
func (suc *SysUserCreate) SetUserName(s string) *SysUserCreate {
	suc.mutation.SetUserName(s)
	return suc
}

// SetRealName sets the "real_name" field.
func (suc *SysUserCreate) SetRealName(s string) *SysUserCreate {
	suc.mutation.SetRealName(s)
	return suc
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRealName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetRealName(*s)
	}
	return suc
}

// SetFirstName sets the "first_name" field.
func (suc *SysUserCreate) SetFirstName(s string) *SysUserCreate {
	suc.mutation.SetFirstName(s)
	return suc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableFirstName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetFirstName(*s)
	}
	return suc
}

// SetLastName sets the "last_name" field.
func (suc *SysUserCreate) SetLastName(s string) *SysUserCreate {
	suc.mutation.SetLastName(s)
	return suc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableLastName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetLastName(*s)
	}
	return suc
}

// SetPassword sets the "password" field.
func (suc *SysUserCreate) SetPassword(s string) *SysUserCreate {
	suc.mutation.SetPassword(s)
	return suc
}

// SetEmail sets the "email" field.
func (suc *SysUserCreate) SetEmail(s string) *SysUserCreate {
	suc.mutation.SetEmail(s)
	return suc
}

// SetMobile sets the "mobile" field.
func (suc *SysUserCreate) SetMobile(s string) *SysUserCreate {
	suc.mutation.SetMobile(s)
	return suc
}

// SetSalt sets the "salt" field.
func (suc *SysUserCreate) SetSalt(s string) *SysUserCreate {
	suc.mutation.SetSalt(s)
	return suc
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSalt(s *string) *SysUserCreate {
	if s != nil {
		suc.SetSalt(*s)
	}
	return suc
}

// SetID sets the "id" field.
func (suc *SysUserCreate) SetID(s string) *SysUserCreate {
	suc.mutation.SetID(s)
	return suc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableID(s *string) *SysUserCreate {
	if s != nil {
		suc.SetID(*s)
	}
	return suc
}

// Mutation returns the SysUserMutation object of the builder.
func (suc *SysUserCreate) Mutation() *SysUserMutation {
	return suc.mutation
}

// Save creates the SysUser in the database.
func (suc *SysUserCreate) Save(ctx context.Context) (*SysUser, error) {
	suc.defaults()
	return withHooks[*SysUser, SysUserMutation](ctx, suc.sqlSave, suc.mutation, suc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SysUserCreate) SaveX(ctx context.Context) *SysUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *SysUserCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *SysUserCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suc *SysUserCreate) defaults() {
	if _, ok := suc.mutation.IsDel(); !ok {
		v := sysuser.DefaultIsDel
		suc.mutation.SetIsDel(v)
	}
	if _, ok := suc.mutation.Sort(); !ok {
		v := sysuser.DefaultSort
		suc.mutation.SetSort(v)
	}
	if _, ok := suc.mutation.CreatedAt(); !ok {
		v := sysuser.DefaultCreatedAt()
		suc.mutation.SetCreatedAt(v)
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		v := sysuser.DefaultUpdatedAt()
		suc.mutation.SetUpdatedAt(v)
	}
	if _, ok := suc.mutation.IsActive(); !ok {
		v := sysuser.DefaultIsActive
		suc.mutation.SetIsActive(v)
	}
	if _, ok := suc.mutation.Salt(); !ok {
		v := sysuser.DefaultSalt()
		suc.mutation.SetSalt(v)
	}
	if _, ok := suc.mutation.ID(); !ok {
		v := sysuser.DefaultID()
		suc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SysUserCreate) check() error {
	if _, ok := suc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysUser.is_del"`)}
	}
	if _, ok := suc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "SysUser.sort"`)}
	}
	if _, ok := suc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysUser.created_at"`)}
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysUser.updated_at"`)}
	}
	if _, ok := suc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "SysUser.is_active"`)}
	}
	if _, ok := suc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`ent: missing required field "SysUser.user_name"`)}
	}
	if v, ok := suc.mutation.UserName(); ok {
		if err := sysuser.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.user_name": %w`, err)}
		}
	}
	if v, ok := suc.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if v, ok := suc.mutation.FirstName(); ok {
		if err := sysuser.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.first_name": %w`, err)}
		}
	}
	if v, ok := suc.mutation.LastName(); ok {
		if err := sysuser.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.last_name": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "SysUser.password"`)}
	}
	if v, ok := suc.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "SysUser.password": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "SysUser.email"`)}
	}
	if v, ok := suc.mutation.Email(); ok {
		if err := sysuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "SysUser.email": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "SysUser.mobile"`)}
	}
	if v, ok := suc.mutation.Mobile(); ok {
		if err := sysuser.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "SysUser.mobile": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "SysUser.salt"`)}
	}
	if v, ok := suc.mutation.ID(); ok {
		if err := sysuser.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysUser.id": %w`, err)}
		}
	}
	return nil
}

func (suc *SysUserCreate) sqlSave(ctx context.Context) (*SysUser, error) {
	if err := suc.check(); err != nil {
		return nil, err
	}
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysUser.ID type: %T", _spec.ID.Value)
		}
	}
	suc.mutation.id = &_node.ID
	suc.mutation.done = true
	return _node, nil
}

func (suc *SysUserCreate) createSpec() (*SysUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUser{config: suc.config}
		_spec = sqlgraph.NewCreateSpec(sysuser.Table, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString))
	)
	_spec.OnConflict = suc.conflict
	if id, ok := suc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suc.mutation.IsDel(); ok {
		_spec.SetField(sysuser.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := suc.mutation.Sort(); ok {
		_spec.SetField(sysuser.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := suc.mutation.CreatedAt(); ok {
		_spec.SetField(sysuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := suc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := suc.mutation.DeletedAt(); ok {
		_spec.SetField(sysuser.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := suc.mutation.IsActive(); ok {
		_spec.SetField(sysuser.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := suc.mutation.UserName(); ok {
		_spec.SetField(sysuser.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := suc.mutation.RealName(); ok {
		_spec.SetField(sysuser.FieldRealName, field.TypeString, value)
		_node.RealName = &value
	}
	if value, ok := suc.mutation.FirstName(); ok {
		_spec.SetField(sysuser.FieldFirstName, field.TypeString, value)
		_node.FirstName = &value
	}
	if value, ok := suc.mutation.LastName(); ok {
		_spec.SetField(sysuser.FieldLastName, field.TypeString, value)
		_node.LastName = &value
	}
	if value, ok := suc.mutation.Password(); ok {
		_spec.SetField(sysuser.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := suc.mutation.Email(); ok {
		_spec.SetField(sysuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := suc.mutation.Mobile(); ok {
		_spec.SetField(sysuser.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := suc.mutation.Salt(); ok {
		_spec.SetField(sysuser.FieldSalt, field.TypeString, value)
		_node.Salt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysUser.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysUserUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (suc *SysUserCreate) OnConflict(opts ...sql.ConflictOption) *SysUserUpsertOne {
	suc.conflict = opts
	return &SysUserUpsertOne{
		create: suc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (suc *SysUserCreate) OnConflictColumns(columns ...string) *SysUserUpsertOne {
	suc.conflict = append(suc.conflict, sql.ConflictColumns(columns...))
	return &SysUserUpsertOne{
		create: suc,
	}
}

type (
	// SysUserUpsertOne is the builder for "upsert"-ing
	//  one SysUser node.
	SysUserUpsertOne struct {
		create *SysUserCreate
	}

	// SysUserUpsert is the "OnConflict" setter.
	SysUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysUserUpsert) SetIsDel(v bool) *SysUserUpsert {
	u.Set(sysuser.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateIsDel() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldIsDel)
	return u
}

// SetSort sets the "sort" field.
func (u *SysUserUpsert) SetSort(v int32) *SysUserUpsert {
	u.Set(sysuser.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateSort() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysUserUpsert) AddSort(v int32) *SysUserUpsert {
	u.Add(sysuser.FieldSort, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserUpsert) SetUpdatedAt(v time.Time) *SysUserUpsert {
	u.Set(sysuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateUpdatedAt() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserUpsert) SetDeletedAt(v time.Time) *SysUserUpsert {
	u.Set(sysuser.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateDeletedAt() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserUpsert) ClearDeletedAt() *SysUserUpsert {
	u.SetNull(sysuser.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysUserUpsert) SetIsActive(v bool) *SysUserUpsert {
	u.Set(sysuser.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateIsActive() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldIsActive)
	return u
}

// SetRealName sets the "real_name" field.
func (u *SysUserUpsert) SetRealName(v string) *SysUserUpsert {
	u.Set(sysuser.FieldRealName, v)
	return u
}

// UpdateRealName sets the "real_name" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateRealName() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldRealName)
	return u
}

// ClearRealName clears the value of the "real_name" field.
func (u *SysUserUpsert) ClearRealName() *SysUserUpsert {
	u.SetNull(sysuser.FieldRealName)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *SysUserUpsert) SetFirstName(v string) *SysUserUpsert {
	u.Set(sysuser.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateFirstName() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldFirstName)
	return u
}

// ClearFirstName clears the value of the "first_name" field.
func (u *SysUserUpsert) ClearFirstName() *SysUserUpsert {
	u.SetNull(sysuser.FieldFirstName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *SysUserUpsert) SetLastName(v string) *SysUserUpsert {
	u.Set(sysuser.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateLastName() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldLastName)
	return u
}

// ClearLastName clears the value of the "last_name" field.
func (u *SysUserUpsert) ClearLastName() *SysUserUpsert {
	u.SetNull(sysuser.FieldLastName)
	return u
}

// SetPassword sets the "password" field.
func (u *SysUserUpsert) SetPassword(v string) *SysUserUpsert {
	u.Set(sysuser.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *SysUserUpsert) UpdatePassword() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldPassword)
	return u
}

// SetEmail sets the "email" field.
func (u *SysUserUpsert) SetEmail(v string) *SysUserUpsert {
	u.Set(sysuser.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateEmail() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldEmail)
	return u
}

// SetMobile sets the "mobile" field.
func (u *SysUserUpsert) SetMobile(v string) *SysUserUpsert {
	u.Set(sysuser.FieldMobile, v)
	return u
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateMobile() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldMobile)
	return u
}

// SetSalt sets the "salt" field.
func (u *SysUserUpsert) SetSalt(v string) *SysUserUpsert {
	u.Set(sysuser.FieldSalt, v)
	return u
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *SysUserUpsert) UpdateSalt() *SysUserUpsert {
	u.SetExcluded(sysuser.FieldSalt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysUserUpsertOne) UpdateNewValues() *SysUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysuser.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysuser.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UserName(); exists {
			s.SetIgnore(sysuser.FieldUserName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysUserUpsertOne) Ignore() *SysUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysUserUpsertOne) DoNothing() *SysUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysUserCreate.OnConflict
// documentation for more info.
func (u *SysUserUpsertOne) Update(set func(*SysUserUpsert)) *SysUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysUserUpsertOne) SetIsDel(v bool) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateIsDel() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysUserUpsertOne) SetSort(v int32) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysUserUpsertOne) AddSort(v int32) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateSort() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserUpsertOne) SetUpdatedAt(v time.Time) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateUpdatedAt() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserUpsertOne) SetDeletedAt(v time.Time) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateDeletedAt() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserUpsertOne) ClearDeletedAt() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysUserUpsertOne) SetIsActive(v bool) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateIsActive() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateIsActive()
	})
}

// SetRealName sets the "real_name" field.
func (u *SysUserUpsertOne) SetRealName(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetRealName(v)
	})
}

// UpdateRealName sets the "real_name" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateRealName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateRealName()
	})
}

// ClearRealName clears the value of the "real_name" field.
func (u *SysUserUpsertOne) ClearRealName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearRealName()
	})
}

// SetFirstName sets the "first_name" field.
func (u *SysUserUpsertOne) SetFirstName(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateFirstName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateFirstName()
	})
}

// ClearFirstName clears the value of the "first_name" field.
func (u *SysUserUpsertOne) ClearFirstName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *SysUserUpsertOne) SetLastName(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateLastName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateLastName()
	})
}

// ClearLastName clears the value of the "last_name" field.
func (u *SysUserUpsertOne) ClearLastName() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearLastName()
	})
}

// SetPassword sets the "password" field.
func (u *SysUserUpsertOne) SetPassword(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdatePassword() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *SysUserUpsertOne) SetEmail(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateEmail() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateEmail()
	})
}

// SetMobile sets the "mobile" field.
func (u *SysUserUpsertOne) SetMobile(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateMobile() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateMobile()
	})
}

// SetSalt sets the "salt" field.
func (u *SysUserUpsertOne) SetSalt(v string) *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.SetSalt(v)
	})
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *SysUserUpsertOne) UpdateSalt() *SysUserUpsertOne {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateSalt()
	})
}

// Exec executes the query.
func (u *SysUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysUserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SysUserUpsertOne.ID is not supported by MySQL driver. Use SysUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysUserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysUserCreateBulk is the builder for creating many SysUser entities in bulk.
type SysUserCreateBulk struct {
	config
	builders []*SysUserCreate
	conflict []sql.ConflictOption
}

// Save creates the SysUser entities in the database.
func (sucb *SysUserCreateBulk) Save(ctx context.Context) ([]*SysUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SysUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SysUserCreateBulk) SaveX(ctx context.Context) []*SysUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *SysUserCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *SysUserCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysUserUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sucb *SysUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysUserUpsertBulk {
	sucb.conflict = opts
	return &SysUserUpsertBulk{
		create: sucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sucb *SysUserCreateBulk) OnConflictColumns(columns ...string) *SysUserUpsertBulk {
	sucb.conflict = append(sucb.conflict, sql.ConflictColumns(columns...))
	return &SysUserUpsertBulk{
		create: sucb,
	}
}

// SysUserUpsertBulk is the builder for "upsert"-ing
// a bulk of SysUser nodes.
type SysUserUpsertBulk struct {
	create *SysUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysUserUpsertBulk) UpdateNewValues() *SysUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysuser.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysuser.FieldCreatedAt)
			}
			if _, exists := b.mutation.UserName(); exists {
				s.SetIgnore(sysuser.FieldUserName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysUserUpsertBulk) Ignore() *SysUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysUserUpsertBulk) DoNothing() *SysUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysUserCreateBulk.OnConflict
// documentation for more info.
func (u *SysUserUpsertBulk) Update(set func(*SysUserUpsert)) *SysUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysUserUpsertBulk) SetIsDel(v bool) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateIsDel() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *SysUserUpsertBulk) SetSort(v int32) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysUserUpsertBulk) AddSort(v int32) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateSort() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysUserUpsertBulk) SetUpdatedAt(v time.Time) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateUpdatedAt() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysUserUpsertBulk) SetDeletedAt(v time.Time) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateDeletedAt() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysUserUpsertBulk) ClearDeletedAt() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysUserUpsertBulk) SetIsActive(v bool) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateIsActive() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateIsActive()
	})
}

// SetRealName sets the "real_name" field.
func (u *SysUserUpsertBulk) SetRealName(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetRealName(v)
	})
}

// UpdateRealName sets the "real_name" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateRealName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateRealName()
	})
}

// ClearRealName clears the value of the "real_name" field.
func (u *SysUserUpsertBulk) ClearRealName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearRealName()
	})
}

// SetFirstName sets the "first_name" field.
func (u *SysUserUpsertBulk) SetFirstName(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateFirstName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateFirstName()
	})
}

// ClearFirstName clears the value of the "first_name" field.
func (u *SysUserUpsertBulk) ClearFirstName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *SysUserUpsertBulk) SetLastName(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateLastName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateLastName()
	})
}

// ClearLastName clears the value of the "last_name" field.
func (u *SysUserUpsertBulk) ClearLastName() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.ClearLastName()
	})
}

// SetPassword sets the "password" field.
func (u *SysUserUpsertBulk) SetPassword(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdatePassword() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *SysUserUpsertBulk) SetEmail(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateEmail() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateEmail()
	})
}

// SetMobile sets the "mobile" field.
func (u *SysUserUpsertBulk) SetMobile(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateMobile() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateMobile()
	})
}

// SetSalt sets the "salt" field.
func (u *SysUserUpsertBulk) SetSalt(v string) *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.SetSalt(v)
	})
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *SysUserUpsertBulk) UpdateSalt() *SysUserUpsertBulk {
	return u.Update(func(s *SysUserUpsert) {
		s.UpdateSalt()
	})
}

// Exec executes the query.
func (u *SysUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
