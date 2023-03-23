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
	"github.com/heromicro/omgind/internal/gen/ent/orgstaff"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictCreate is the builder for creating a SysDict entity.
type SysDictCreate struct {
	config
	mutation *SysDictMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (sdc *SysDictCreate) SetIsDel(b bool) *SysDictCreate {
	sdc.mutation.SetIsDel(b)
	return sdc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableIsDel(b *bool) *SysDictCreate {
	if b != nil {
		sdc.SetIsDel(*b)
	}
	return sdc
}

// SetMemo sets the "memo" field.
func (sdc *SysDictCreate) SetMemo(s string) *SysDictCreate {
	sdc.mutation.SetMemo(s)
	return sdc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableMemo(s *string) *SysDictCreate {
	if s != nil {
		sdc.SetMemo(*s)
	}
	return sdc
}

// SetSort sets the "sort" field.
func (sdc *SysDictCreate) SetSort(i int32) *SysDictCreate {
	sdc.mutation.SetSort(i)
	return sdc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableSort(i *int32) *SysDictCreate {
	if i != nil {
		sdc.SetSort(*i)
	}
	return sdc
}

// SetCreatedAt sets the "created_at" field.
func (sdc *SysDictCreate) SetCreatedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableCreatedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetCreatedAt(*t)
	}
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *SysDictCreate) SetUpdatedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableUpdatedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetDeletedAt sets the "deleted_at" field.
func (sdc *SysDictCreate) SetDeletedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetDeletedAt(t)
	return sdc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableDeletedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetDeletedAt(*t)
	}
	return sdc
}

// SetIsActive sets the "is_active" field.
func (sdc *SysDictCreate) SetIsActive(b bool) *SysDictCreate {
	sdc.mutation.SetIsActive(b)
	return sdc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableIsActive(b *bool) *SysDictCreate {
	if b != nil {
		sdc.SetIsActive(*b)
	}
	return sdc
}

// SetNameCn sets the "name_cn" field.
func (sdc *SysDictCreate) SetNameCn(s string) *SysDictCreate {
	sdc.mutation.SetNameCn(s)
	return sdc
}

// SetNameEn sets the "name_en" field.
func (sdc *SysDictCreate) SetNameEn(s string) *SysDictCreate {
	sdc.mutation.SetNameEn(s)
	return sdc
}

// SetTipe sets the "tipe" field.
func (sdc *SysDictCreate) SetTipe(s sysdict.Tipe) *SysDictCreate {
	sdc.mutation.SetTipe(s)
	return sdc
}

// SetNillableTipe sets the "tipe" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableTipe(s *sysdict.Tipe) *SysDictCreate {
	if s != nil {
		sdc.SetTipe(*s)
	}
	return sdc
}

// SetID sets the "id" field.
func (sdc *SysDictCreate) SetID(s string) *SysDictCreate {
	sdc.mutation.SetID(s)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableID(s *string) *SysDictCreate {
	if s != nil {
		sdc.SetID(*s)
	}
	return sdc
}

// AddItemIDs adds the "items" edge to the SysDictItem entity by IDs.
func (sdc *SysDictCreate) AddItemIDs(ids ...string) *SysDictCreate {
	sdc.mutation.AddItemIDs(ids...)
	return sdc
}

// AddItems adds the "items" edges to the SysDictItem entity.
func (sdc *SysDictCreate) AddItems(s ...*SysDictItem) *SysDictCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdc.AddItemIDs(ids...)
}

// SetStaffGenderID sets the "staff_gender" edge to the OrgStaff entity by ID.
func (sdc *SysDictCreate) SetStaffGenderID(id string) *SysDictCreate {
	sdc.mutation.SetStaffGenderID(id)
	return sdc
}

// SetNillableStaffGenderID sets the "staff_gender" edge to the OrgStaff entity by ID if the given value is not nil.
func (sdc *SysDictCreate) SetNillableStaffGenderID(id *string) *SysDictCreate {
	if id != nil {
		sdc = sdc.SetStaffGenderID(*id)
	}
	return sdc
}

// SetStaffGender sets the "staff_gender" edge to the OrgStaff entity.
func (sdc *SysDictCreate) SetStaffGender(o *OrgStaff) *SysDictCreate {
	return sdc.SetStaffGenderID(o.ID)
}

// SetStaffEmpystID sets the "staff_empyst" edge to the OrgStaff entity by ID.
func (sdc *SysDictCreate) SetStaffEmpystID(id string) *SysDictCreate {
	sdc.mutation.SetStaffEmpystID(id)
	return sdc
}

// SetNillableStaffEmpystID sets the "staff_empyst" edge to the OrgStaff entity by ID if the given value is not nil.
func (sdc *SysDictCreate) SetNillableStaffEmpystID(id *string) *SysDictCreate {
	if id != nil {
		sdc = sdc.SetStaffEmpystID(*id)
	}
	return sdc
}

// SetStaffEmpyst sets the "staff_empyst" edge to the OrgStaff entity.
func (sdc *SysDictCreate) SetStaffEmpyst(o *OrgStaff) *SysDictCreate {
	return sdc.SetStaffEmpystID(o.ID)
}

// Mutation returns the SysDictMutation object of the builder.
func (sdc *SysDictCreate) Mutation() *SysDictMutation {
	return sdc.mutation
}

// Save creates the SysDict in the database.
func (sdc *SysDictCreate) Save(ctx context.Context) (*SysDict, error) {
	sdc.defaults()
	return withHooks[*SysDict, SysDictMutation](ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SysDictCreate) SaveX(ctx context.Context) *SysDict {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SysDictCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SysDictCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *SysDictCreate) defaults() {
	if _, ok := sdc.mutation.IsDel(); !ok {
		v := sysdict.DefaultIsDel
		sdc.mutation.SetIsDel(v)
	}
	if _, ok := sdc.mutation.Memo(); !ok {
		v := sysdict.DefaultMemo
		sdc.mutation.SetMemo(v)
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		v := sysdict.DefaultSort
		sdc.mutation.SetSort(v)
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		v := sysdict.DefaultCreatedAt()
		sdc.mutation.SetCreatedAt(v)
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		v := sysdict.DefaultUpdatedAt()
		sdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdc.mutation.IsActive(); !ok {
		v := sysdict.DefaultIsActive
		sdc.mutation.SetIsActive(v)
	}
	if _, ok := sdc.mutation.Tipe(); !ok {
		v := sysdict.DefaultTipe
		sdc.mutation.SetTipe(v)
	}
	if _, ok := sdc.mutation.ID(); !ok {
		v := sysdict.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SysDictCreate) check() error {
	if _, ok := sdc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysDict.is_del"`)}
	}
	if v, ok := sdc.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "SysDict.memo": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "SysDict.sort"`)}
	}
	if _, ok := sdc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "SysDict.is_active"`)}
	}
	if _, ok := sdc.mutation.NameCn(); !ok {
		return &ValidationError{Name: "name_cn", err: errors.New(`ent: missing required field "SysDict.name_cn"`)}
	}
	if v, ok := sdc.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf(`ent: validator failed for field "SysDict.name_cn": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.NameEn(); !ok {
		return &ValidationError{Name: "name_en", err: errors.New(`ent: missing required field "SysDict.name_en"`)}
	}
	if v, ok := sdc.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "SysDict.name_en": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Tipe(); !ok {
		return &ValidationError{Name: "tipe", err: errors.New(`ent: missing required field "SysDict.tipe"`)}
	}
	if v, ok := sdc.mutation.Tipe(); ok {
		if err := sysdict.TipeValidator(v); err != nil {
			return &ValidationError{Name: "tipe", err: fmt.Errorf(`ent: validator failed for field "SysDict.tipe": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.ID(); ok {
		if err := sysdict.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysDict.id": %w`, err)}
		}
	}
	return nil
}

func (sdc *SysDictCreate) sqlSave(ctx context.Context) (*SysDict, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysDict.ID type: %T", _spec.ID.Value)
		}
	}
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *SysDictCreate) createSpec() (*SysDict, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDict{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(sysdict.Table, sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString))
	)
	_spec.Schema = sdc.schemaConfig.SysDict
	_spec.OnConflict = sdc.conflict
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdc.mutation.IsDel(); ok {
		_spec.SetField(sysdict.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := sdc.mutation.Memo(); ok {
		_spec.SetField(sysdict.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := sdc.mutation.Sort(); ok {
		_spec.SetField(sysdict.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.SetField(sysdict.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdict.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sdc.mutation.DeletedAt(); ok {
		_spec.SetField(sysdict.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := sdc.mutation.IsActive(); ok {
		_spec.SetField(sysdict.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := sdc.mutation.NameCn(); ok {
		_spec.SetField(sysdict.FieldNameCn, field.TypeString, value)
		_node.NameCn = value
	}
	if value, ok := sdc.mutation.NameEn(); ok {
		_spec.SetField(sysdict.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if value, ok := sdc.mutation.Tipe(); ok {
		_spec.SetField(sysdict.FieldTipe, field.TypeEnum, value)
		_node.Tipe = value
	}
	if nodes := sdc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdict.ItemsTable,
			Columns: []string{sysdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString),
			},
		}
		edge.Schema = sdc.schemaConfig.SysDictItem
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.StaffGenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysdict.StaffGenderTable,
			Columns: []string{sysdict.StaffGenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgstaff.FieldID, field.TypeString),
			},
		}
		edge.Schema = sdc.schemaConfig.SysDict
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sys_dict_staff_gender = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.StaffEmpystIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysdict.StaffEmpystTable,
			Columns: []string{sysdict.StaffEmpystColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgstaff.FieldID, field.TypeString),
			},
		}
		edge.Schema = sdc.schemaConfig.SysDict
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sys_dict_staff_empyst = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDict.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sdc *SysDictCreate) OnConflict(opts ...sql.ConflictOption) *SysDictUpsertOne {
	sdc.conflict = opts
	return &SysDictUpsertOne{
		create: sdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDict.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdc *SysDictCreate) OnConflictColumns(columns ...string) *SysDictUpsertOne {
	sdc.conflict = append(sdc.conflict, sql.ConflictColumns(columns...))
	return &SysDictUpsertOne{
		create: sdc,
	}
}

type (
	// SysDictUpsertOne is the builder for "upsert"-ing
	//  one SysDict node.
	SysDictUpsertOne struct {
		create *SysDictCreate
	}

	// SysDictUpsert is the "OnConflict" setter.
	SysDictUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *SysDictUpsert) SetIsDel(v bool) *SysDictUpsert {
	u.Set(sysdict.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateIsDel() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldIsDel)
	return u
}

// SetMemo sets the "memo" field.
func (u *SysDictUpsert) SetMemo(v string) *SysDictUpsert {
	u.Set(sysdict.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateMemo() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictUpsert) ClearMemo() *SysDictUpsert {
	u.SetNull(sysdict.FieldMemo)
	return u
}

// SetSort sets the "sort" field.
func (u *SysDictUpsert) SetSort(v int32) *SysDictUpsert {
	u.Set(sysdict.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateSort() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *SysDictUpsert) AddSort(v int32) *SysDictUpsert {
	u.Add(sysdict.FieldSort, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictUpsert) SetUpdatedAt(v time.Time) *SysDictUpsert {
	u.Set(sysdict.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateUpdatedAt() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictUpsert) ClearUpdatedAt() *SysDictUpsert {
	u.SetNull(sysdict.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictUpsert) SetDeletedAt(v time.Time) *SysDictUpsert {
	u.Set(sysdict.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateDeletedAt() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictUpsert) ClearDeletedAt() *SysDictUpsert {
	u.SetNull(sysdict.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *SysDictUpsert) SetIsActive(v bool) *SysDictUpsert {
	u.Set(sysdict.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateIsActive() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldIsActive)
	return u
}

// SetNameCn sets the "name_cn" field.
func (u *SysDictUpsert) SetNameCn(v string) *SysDictUpsert {
	u.Set(sysdict.FieldNameCn, v)
	return u
}

// UpdateNameCn sets the "name_cn" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateNameCn() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldNameCn)
	return u
}

// SetNameEn sets the "name_en" field.
func (u *SysDictUpsert) SetNameEn(v string) *SysDictUpsert {
	u.Set(sysdict.FieldNameEn, v)
	return u
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateNameEn() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldNameEn)
	return u
}

// SetTipe sets the "tipe" field.
func (u *SysDictUpsert) SetTipe(v sysdict.Tipe) *SysDictUpsert {
	u.Set(sysdict.FieldTipe, v)
	return u
}

// UpdateTipe sets the "tipe" field to the value that was provided on create.
func (u *SysDictUpsert) UpdateTipe() *SysDictUpsert {
	u.SetExcluded(sysdict.FieldTipe)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SysDict.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdict.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictUpsertOne) UpdateNewValues() *SysDictUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sysdict.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sysdict.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDict.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SysDictUpsertOne) Ignore() *SysDictUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictUpsertOne) DoNothing() *SysDictUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictCreate.OnConflict
// documentation for more info.
func (u *SysDictUpsertOne) Update(set func(*SysDictUpsert)) *SysDictUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysDictUpsertOne) SetIsDel(v bool) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateIsDel() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysDictUpsertOne) SetMemo(v string) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateMemo() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictUpsertOne) ClearMemo() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysDictUpsertOne) SetSort(v int32) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysDictUpsertOne) AddSort(v int32) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateSort() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictUpsertOne) SetUpdatedAt(v time.Time) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateUpdatedAt() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictUpsertOne) ClearUpdatedAt() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictUpsertOne) SetDeletedAt(v time.Time) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateDeletedAt() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictUpsertOne) ClearDeletedAt() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysDictUpsertOne) SetIsActive(v bool) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateIsActive() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateIsActive()
	})
}

// SetNameCn sets the "name_cn" field.
func (u *SysDictUpsertOne) SetNameCn(v string) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetNameCn(v)
	})
}

// UpdateNameCn sets the "name_cn" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateNameCn() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateNameCn()
	})
}

// SetNameEn sets the "name_en" field.
func (u *SysDictUpsertOne) SetNameEn(v string) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateNameEn() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateNameEn()
	})
}

// SetTipe sets the "tipe" field.
func (u *SysDictUpsertOne) SetTipe(v sysdict.Tipe) *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.SetTipe(v)
	})
}

// UpdateTipe sets the "tipe" field to the value that was provided on create.
func (u *SysDictUpsertOne) UpdateTipe() *SysDictUpsertOne {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateTipe()
	})
}

// Exec executes the query.
func (u *SysDictUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SysDictUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SysDictUpsertOne.ID is not supported by MySQL driver. Use SysDictUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SysDictUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SysDictCreateBulk is the builder for creating many SysDict entities in bulk.
type SysDictCreateBulk struct {
	config
	builders []*SysDictCreate
	conflict []sql.ConflictOption
}

// Save creates the SysDict entities in the database.
func (sdcb *SysDictCreateBulk) Save(ctx context.Context) ([]*SysDict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SysDict, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SysDictCreateBulk) SaveX(ctx context.Context) []*SysDict {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SysDictCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SysDictCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SysDict.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SysDictUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (sdcb *SysDictCreateBulk) OnConflict(opts ...sql.ConflictOption) *SysDictUpsertBulk {
	sdcb.conflict = opts
	return &SysDictUpsertBulk{
		create: sdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SysDict.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdcb *SysDictCreateBulk) OnConflictColumns(columns ...string) *SysDictUpsertBulk {
	sdcb.conflict = append(sdcb.conflict, sql.ConflictColumns(columns...))
	return &SysDictUpsertBulk{
		create: sdcb,
	}
}

// SysDictUpsertBulk is the builder for "upsert"-ing
// a bulk of SysDict nodes.
type SysDictUpsertBulk struct {
	create *SysDictCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SysDict.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sysdict.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SysDictUpsertBulk) UpdateNewValues() *SysDictUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sysdict.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sysdict.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SysDict.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SysDictUpsertBulk) Ignore() *SysDictUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SysDictUpsertBulk) DoNothing() *SysDictUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SysDictCreateBulk.OnConflict
// documentation for more info.
func (u *SysDictUpsertBulk) Update(set func(*SysDictUpsert)) *SysDictUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SysDictUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *SysDictUpsertBulk) SetIsDel(v bool) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateIsDel() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateIsDel()
	})
}

// SetMemo sets the "memo" field.
func (u *SysDictUpsertBulk) SetMemo(v string) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateMemo() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *SysDictUpsertBulk) ClearMemo() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearMemo()
	})
}

// SetSort sets the "sort" field.
func (u *SysDictUpsertBulk) SetSort(v int32) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *SysDictUpsertBulk) AddSort(v int32) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateSort() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SysDictUpsertBulk) SetUpdatedAt(v time.Time) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateUpdatedAt() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SysDictUpsertBulk) ClearUpdatedAt() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SysDictUpsertBulk) SetDeletedAt(v time.Time) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateDeletedAt() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SysDictUpsertBulk) ClearDeletedAt() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *SysDictUpsertBulk) SetIsActive(v bool) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateIsActive() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateIsActive()
	})
}

// SetNameCn sets the "name_cn" field.
func (u *SysDictUpsertBulk) SetNameCn(v string) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetNameCn(v)
	})
}

// UpdateNameCn sets the "name_cn" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateNameCn() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateNameCn()
	})
}

// SetNameEn sets the "name_en" field.
func (u *SysDictUpsertBulk) SetNameEn(v string) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateNameEn() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateNameEn()
	})
}

// SetTipe sets the "tipe" field.
func (u *SysDictUpsertBulk) SetTipe(v sysdict.Tipe) *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.SetTipe(v)
	})
}

// UpdateTipe sets the "tipe" field to the value that was provided on create.
func (u *SysDictUpsertBulk) UpdateTipe() *SysDictUpsertBulk {
	return u.Update(func(s *SysDictUpsert) {
		s.UpdateTipe()
	})
}

// Exec executes the query.
func (u *SysDictUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SysDictCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SysDictCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SysDictUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
