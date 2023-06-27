// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/entscheme/orgorgan"
	"github.com/heromicro/omgind/internal/gen/entscheme/orgposition"
	"github.com/heromicro/omgind/internal/gen/entscheme/orgstaff"
)

// OrgPositionCreate is the builder for creating a OrgPosition entity.
type OrgPositionCreate struct {
	config
	mutation *OrgPositionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsDel sets the "is_del" field.
func (opc *OrgPositionCreate) SetIsDel(b bool) *OrgPositionCreate {
	opc.mutation.SetIsDel(b)
	return opc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableIsDel(b *bool) *OrgPositionCreate {
	if b != nil {
		opc.SetIsDel(*b)
	}
	return opc
}

// SetSort sets the "sort" field.
func (opc *OrgPositionCreate) SetSort(i int32) *OrgPositionCreate {
	opc.mutation.SetSort(i)
	return opc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableSort(i *int32) *OrgPositionCreate {
	if i != nil {
		opc.SetSort(*i)
	}
	return opc
}

// SetCreatedAt sets the "created_at" field.
func (opc *OrgPositionCreate) SetCreatedAt(t time.Time) *OrgPositionCreate {
	opc.mutation.SetCreatedAt(t)
	return opc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableCreatedAt(t *time.Time) *OrgPositionCreate {
	if t != nil {
		opc.SetCreatedAt(*t)
	}
	return opc
}

// SetUpdatedAt sets the "updated_at" field.
func (opc *OrgPositionCreate) SetUpdatedAt(t time.Time) *OrgPositionCreate {
	opc.mutation.SetUpdatedAt(t)
	return opc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableUpdatedAt(t *time.Time) *OrgPositionCreate {
	if t != nil {
		opc.SetUpdatedAt(*t)
	}
	return opc
}

// SetDeletedAt sets the "deleted_at" field.
func (opc *OrgPositionCreate) SetDeletedAt(t time.Time) *OrgPositionCreate {
	opc.mutation.SetDeletedAt(t)
	return opc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableDeletedAt(t *time.Time) *OrgPositionCreate {
	if t != nil {
		opc.SetDeletedAt(*t)
	}
	return opc
}

// SetIsActive sets the "is_active" field.
func (opc *OrgPositionCreate) SetIsActive(b bool) *OrgPositionCreate {
	opc.mutation.SetIsActive(b)
	return opc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableIsActive(b *bool) *OrgPositionCreate {
	if b != nil {
		opc.SetIsActive(*b)
	}
	return opc
}

// SetMemo sets the "memo" field.
func (opc *OrgPositionCreate) SetMemo(s string) *OrgPositionCreate {
	opc.mutation.SetMemo(s)
	return opc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableMemo(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetMemo(*s)
	}
	return opc
}

// SetName sets the "name" field.
func (opc *OrgPositionCreate) SetName(s string) *OrgPositionCreate {
	opc.mutation.SetName(s)
	return opc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableName(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetName(*s)
	}
	return opc
}

// SetCode sets the "code" field.
func (opc *OrgPositionCreate) SetCode(s string) *OrgPositionCreate {
	opc.mutation.SetCode(s)
	return opc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableCode(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetCode(*s)
	}
	return opc
}

// SetOrgID sets the "org_id" field.
func (opc *OrgPositionCreate) SetOrgID(s string) *OrgPositionCreate {
	opc.mutation.SetOrgID(s)
	return opc
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableOrgID(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetOrgID(*s)
	}
	return opc
}

// SetCreator sets the "creator" field.
func (opc *OrgPositionCreate) SetCreator(s string) *OrgPositionCreate {
	opc.mutation.SetCreator(s)
	return opc
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableCreator(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetCreator(*s)
	}
	return opc
}

// SetID sets the "id" field.
func (opc *OrgPositionCreate) SetID(s string) *OrgPositionCreate {
	opc.mutation.SetID(s)
	return opc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableID(s *string) *OrgPositionCreate {
	if s != nil {
		opc.SetID(*s)
	}
	return opc
}

// SetOrganID sets the "organ" edge to the OrgOrgan entity by ID.
func (opc *OrgPositionCreate) SetOrganID(id string) *OrgPositionCreate {
	opc.mutation.SetOrganID(id)
	return opc
}

// SetNillableOrganID sets the "organ" edge to the OrgOrgan entity by ID if the given value is not nil.
func (opc *OrgPositionCreate) SetNillableOrganID(id *string) *OrgPositionCreate {
	if id != nil {
		opc = opc.SetOrganID(*id)
	}
	return opc
}

// SetOrgan sets the "organ" edge to the OrgOrgan entity.
func (opc *OrgPositionCreate) SetOrgan(o *OrgOrgan) *OrgPositionCreate {
	return opc.SetOrganID(o.ID)
}

// AddStaffIDs adds the "staffs" edge to the OrgStaff entity by IDs.
func (opc *OrgPositionCreate) AddStaffIDs(ids ...string) *OrgPositionCreate {
	opc.mutation.AddStaffIDs(ids...)
	return opc
}

// AddStaffs adds the "staffs" edges to the OrgStaff entity.
func (opc *OrgPositionCreate) AddStaffs(o ...*OrgStaff) *OrgPositionCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return opc.AddStaffIDs(ids...)
}

// Mutation returns the OrgPositionMutation object of the builder.
func (opc *OrgPositionCreate) Mutation() *OrgPositionMutation {
	return opc.mutation
}

// Save creates the OrgPosition in the database.
func (opc *OrgPositionCreate) Save(ctx context.Context) (*OrgPosition, error) {
	opc.defaults()
	return withHooks(ctx, opc.sqlSave, opc.mutation, opc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrgPositionCreate) SaveX(ctx context.Context) *OrgPosition {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrgPositionCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrgPositionCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrgPositionCreate) defaults() {
	if _, ok := opc.mutation.IsDel(); !ok {
		v := orgposition.DefaultIsDel
		opc.mutation.SetIsDel(v)
	}
	if _, ok := opc.mutation.Sort(); !ok {
		v := orgposition.DefaultSort
		opc.mutation.SetSort(v)
	}
	if _, ok := opc.mutation.CreatedAt(); !ok {
		v := orgposition.DefaultCreatedAt()
		opc.mutation.SetCreatedAt(v)
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		v := orgposition.DefaultUpdatedAt()
		opc.mutation.SetUpdatedAt(v)
	}
	if _, ok := opc.mutation.IsActive(); !ok {
		v := orgposition.DefaultIsActive
		opc.mutation.SetIsActive(v)
	}
	if _, ok := opc.mutation.Memo(); !ok {
		v := orgposition.DefaultMemo
		opc.mutation.SetMemo(v)
	}
	if _, ok := opc.mutation.ID(); !ok {
		v := orgposition.DefaultID()
		opc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrgPositionCreate) check() error {
	if _, ok := opc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`entscheme: missing required field "OrgPosition.is_del"`)}
	}
	if _, ok := opc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`entscheme: missing required field "OrgPosition.sort"`)}
	}
	if _, ok := opc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`entscheme: missing required field "OrgPosition.is_active"`)}
	}
	if v, ok := opc.mutation.Memo(); ok {
		if err := orgposition.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`entscheme: validator failed for field "OrgPosition.memo": %w`, err)}
		}
	}
	if v, ok := opc.mutation.Name(); ok {
		if err := orgposition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entscheme: validator failed for field "OrgPosition.name": %w`, err)}
		}
	}
	if v, ok := opc.mutation.Code(); ok {
		if err := orgposition.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`entscheme: validator failed for field "OrgPosition.code": %w`, err)}
		}
	}
	if v, ok := opc.mutation.OrgID(); ok {
		if err := orgposition.OrgIDValidator(v); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`entscheme: validator failed for field "OrgPosition.org_id": %w`, err)}
		}
	}
	if v, ok := opc.mutation.ID(); ok {
		if err := orgposition.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entscheme: validator failed for field "OrgPosition.id": %w`, err)}
		}
	}
	return nil
}

func (opc *OrgPositionCreate) sqlSave(ctx context.Context) (*OrgPosition, error) {
	if err := opc.check(); err != nil {
		return nil, err
	}
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrgPosition.ID type: %T", _spec.ID.Value)
		}
	}
	opc.mutation.id = &_node.ID
	opc.mutation.done = true
	return _node, nil
}

func (opc *OrgPositionCreate) createSpec() (*OrgPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgPosition{config: opc.config}
		_spec = sqlgraph.NewCreateSpec(orgposition.Table, sqlgraph.NewFieldSpec(orgposition.FieldID, field.TypeString))
	)
	_spec.OnConflict = opc.conflict
	if id, ok := opc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := opc.mutation.IsDel(); ok {
		_spec.SetField(orgposition.FieldIsDel, field.TypeBool, value)
		_node.IsDel = value
	}
	if value, ok := opc.mutation.Sort(); ok {
		_spec.SetField(orgposition.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := opc.mutation.CreatedAt(); ok {
		_spec.SetField(orgposition.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := opc.mutation.UpdatedAt(); ok {
		_spec.SetField(orgposition.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := opc.mutation.DeletedAt(); ok {
		_spec.SetField(orgposition.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := opc.mutation.IsActive(); ok {
		_spec.SetField(orgposition.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := opc.mutation.Memo(); ok {
		_spec.SetField(orgposition.FieldMemo, field.TypeString, value)
		_node.Memo = &value
	}
	if value, ok := opc.mutation.Name(); ok {
		_spec.SetField(orgposition.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := opc.mutation.Code(); ok {
		_spec.SetField(orgposition.FieldCode, field.TypeString, value)
		_node.Code = &value
	}
	if value, ok := opc.mutation.Creator(); ok {
		_spec.SetField(orgposition.FieldCreator, field.TypeString, value)
		_node.Creator = &value
	}
	if nodes := opc.mutation.OrganIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgposition.OrganTable,
			Columns: []string{orgposition.OrganColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgorgan.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := opc.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgposition.StaffsTable,
			Columns: []string{orgposition.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgstaff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgPosition.Create().
//		SetIsDel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgPositionUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (opc *OrgPositionCreate) OnConflict(opts ...sql.ConflictOption) *OrgPositionUpsertOne {
	opc.conflict = opts
	return &OrgPositionUpsertOne{
		create: opc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opc *OrgPositionCreate) OnConflictColumns(columns ...string) *OrgPositionUpsertOne {
	opc.conflict = append(opc.conflict, sql.ConflictColumns(columns...))
	return &OrgPositionUpsertOne{
		create: opc,
	}
}

type (
	// OrgPositionUpsertOne is the builder for "upsert"-ing
	//  one OrgPosition node.
	OrgPositionUpsertOne struct {
		create *OrgPositionCreate
	}

	// OrgPositionUpsert is the "OnConflict" setter.
	OrgPositionUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsDel sets the "is_del" field.
func (u *OrgPositionUpsert) SetIsDel(v bool) *OrgPositionUpsert {
	u.Set(orgposition.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateIsDel() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldIsDel)
	return u
}

// SetSort sets the "sort" field.
func (u *OrgPositionUpsert) SetSort(v int32) *OrgPositionUpsert {
	u.Set(orgposition.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateSort() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *OrgPositionUpsert) AddSort(v int32) *OrgPositionUpsert {
	u.Add(orgposition.FieldSort, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgPositionUpsert) SetUpdatedAt(v time.Time) *OrgPositionUpsert {
	u.Set(orgposition.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateUpdatedAt() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgPositionUpsert) ClearUpdatedAt() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrgPositionUpsert) SetDeletedAt(v time.Time) *OrgPositionUpsert {
	u.Set(orgposition.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateDeletedAt() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *OrgPositionUpsert) ClearDeletedAt() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldDeletedAt)
	return u
}

// SetIsActive sets the "is_active" field.
func (u *OrgPositionUpsert) SetIsActive(v bool) *OrgPositionUpsert {
	u.Set(orgposition.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateIsActive() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldIsActive)
	return u
}

// SetMemo sets the "memo" field.
func (u *OrgPositionUpsert) SetMemo(v string) *OrgPositionUpsert {
	u.Set(orgposition.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateMemo() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldMemo)
	return u
}

// ClearMemo clears the value of the "memo" field.
func (u *OrgPositionUpsert) ClearMemo() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldMemo)
	return u
}

// SetName sets the "name" field.
func (u *OrgPositionUpsert) SetName(v string) *OrgPositionUpsert {
	u.Set(orgposition.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateName() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *OrgPositionUpsert) ClearName() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldName)
	return u
}

// SetCode sets the "code" field.
func (u *OrgPositionUpsert) SetCode(v string) *OrgPositionUpsert {
	u.Set(orgposition.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateCode() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldCode)
	return u
}

// ClearCode clears the value of the "code" field.
func (u *OrgPositionUpsert) ClearCode() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldCode)
	return u
}

// SetOrgID sets the "org_id" field.
func (u *OrgPositionUpsert) SetOrgID(v string) *OrgPositionUpsert {
	u.Set(orgposition.FieldOrgID, v)
	return u
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateOrgID() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldOrgID)
	return u
}

// ClearOrgID clears the value of the "org_id" field.
func (u *OrgPositionUpsert) ClearOrgID() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldOrgID)
	return u
}

// SetCreator sets the "creator" field.
func (u *OrgPositionUpsert) SetCreator(v string) *OrgPositionUpsert {
	u.Set(orgposition.FieldCreator, v)
	return u
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *OrgPositionUpsert) UpdateCreator() *OrgPositionUpsert {
	u.SetExcluded(orgposition.FieldCreator)
	return u
}

// ClearCreator clears the value of the "creator" field.
func (u *OrgPositionUpsert) ClearCreator() *OrgPositionUpsert {
	u.SetNull(orgposition.FieldCreator)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgposition.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgPositionUpsertOne) UpdateNewValues() *OrgPositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orgposition.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(orgposition.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrgPositionUpsertOne) Ignore() *OrgPositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgPositionUpsertOne) DoNothing() *OrgPositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgPositionCreate.OnConflict
// documentation for more info.
func (u *OrgPositionUpsertOne) Update(set func(*OrgPositionUpsert)) *OrgPositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgPositionUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *OrgPositionUpsertOne) SetIsDel(v bool) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateIsDel() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *OrgPositionUpsertOne) SetSort(v int32) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *OrgPositionUpsertOne) AddSort(v int32) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateSort() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgPositionUpsertOne) SetUpdatedAt(v time.Time) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateUpdatedAt() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgPositionUpsertOne) ClearUpdatedAt() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrgPositionUpsertOne) SetDeletedAt(v time.Time) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateDeletedAt() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *OrgPositionUpsertOne) ClearDeletedAt() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *OrgPositionUpsertOne) SetIsActive(v bool) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateIsActive() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateIsActive()
	})
}

// SetMemo sets the "memo" field.
func (u *OrgPositionUpsertOne) SetMemo(v string) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateMemo() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *OrgPositionUpsertOne) ClearMemo() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearMemo()
	})
}

// SetName sets the "name" field.
func (u *OrgPositionUpsertOne) SetName(v string) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateName() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *OrgPositionUpsertOne) ClearName() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearName()
	})
}

// SetCode sets the "code" field.
func (u *OrgPositionUpsertOne) SetCode(v string) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateCode() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *OrgPositionUpsertOne) ClearCode() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearCode()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgPositionUpsertOne) SetOrgID(v string) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateOrgID() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateOrgID()
	})
}

// ClearOrgID clears the value of the "org_id" field.
func (u *OrgPositionUpsertOne) ClearOrgID() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearOrgID()
	})
}

// SetCreator sets the "creator" field.
func (u *OrgPositionUpsertOne) SetCreator(v string) *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *OrgPositionUpsertOne) UpdateCreator() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateCreator()
	})
}

// ClearCreator clears the value of the "creator" field.
func (u *OrgPositionUpsertOne) ClearCreator() *OrgPositionUpsertOne {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearCreator()
	})
}

// Exec executes the query.
func (u *OrgPositionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entscheme: missing options for OrgPositionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgPositionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgPositionUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("entscheme: OrgPositionUpsertOne.ID is not supported by MySQL driver. Use OrgPositionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgPositionUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgPositionCreateBulk is the builder for creating many OrgPosition entities in bulk.
type OrgPositionCreateBulk struct {
	config
	builders []*OrgPositionCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgPosition entities in the database.
func (opcb *OrgPositionCreateBulk) Save(ctx context.Context) ([]*OrgPosition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrgPosition, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgPositionMutation)
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
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = opcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrgPositionCreateBulk) SaveX(ctx context.Context) []*OrgPosition {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrgPositionCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrgPositionCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgPosition.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgPositionUpsert) {
//			SetIsDel(v+v).
//		}).
//		Exec(ctx)
func (opcb *OrgPositionCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgPositionUpsertBulk {
	opcb.conflict = opts
	return &OrgPositionUpsertBulk{
		create: opcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opcb *OrgPositionCreateBulk) OnConflictColumns(columns ...string) *OrgPositionUpsertBulk {
	opcb.conflict = append(opcb.conflict, sql.ConflictColumns(columns...))
	return &OrgPositionUpsertBulk{
		create: opcb,
	}
}

// OrgPositionUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgPosition nodes.
type OrgPositionUpsertBulk struct {
	create *OrgPositionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgposition.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgPositionUpsertBulk) UpdateNewValues() *OrgPositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orgposition.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(orgposition.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgPosition.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrgPositionUpsertBulk) Ignore() *OrgPositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgPositionUpsertBulk) DoNothing() *OrgPositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgPositionCreateBulk.OnConflict
// documentation for more info.
func (u *OrgPositionUpsertBulk) Update(set func(*OrgPositionUpsert)) *OrgPositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgPositionUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsDel sets the "is_del" field.
func (u *OrgPositionUpsertBulk) SetIsDel(v bool) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateIsDel() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateIsDel()
	})
}

// SetSort sets the "sort" field.
func (u *OrgPositionUpsertBulk) SetSort(v int32) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *OrgPositionUpsertBulk) AddSort(v int32) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateSort() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateSort()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgPositionUpsertBulk) SetUpdatedAt(v time.Time) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateUpdatedAt() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgPositionUpsertBulk) ClearUpdatedAt() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrgPositionUpsertBulk) SetDeletedAt(v time.Time) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateDeletedAt() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *OrgPositionUpsertBulk) ClearDeletedAt() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetIsActive sets the "is_active" field.
func (u *OrgPositionUpsertBulk) SetIsActive(v bool) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateIsActive() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateIsActive()
	})
}

// SetMemo sets the "memo" field.
func (u *OrgPositionUpsertBulk) SetMemo(v string) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateMemo() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateMemo()
	})
}

// ClearMemo clears the value of the "memo" field.
func (u *OrgPositionUpsertBulk) ClearMemo() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearMemo()
	})
}

// SetName sets the "name" field.
func (u *OrgPositionUpsertBulk) SetName(v string) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateName() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *OrgPositionUpsertBulk) ClearName() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearName()
	})
}

// SetCode sets the "code" field.
func (u *OrgPositionUpsertBulk) SetCode(v string) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateCode() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *OrgPositionUpsertBulk) ClearCode() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearCode()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgPositionUpsertBulk) SetOrgID(v string) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateOrgID() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateOrgID()
	})
}

// ClearOrgID clears the value of the "org_id" field.
func (u *OrgPositionUpsertBulk) ClearOrgID() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearOrgID()
	})
}

// SetCreator sets the "creator" field.
func (u *OrgPositionUpsertBulk) SetCreator(v string) *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *OrgPositionUpsertBulk) UpdateCreator() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.UpdateCreator()
	})
}

// ClearCreator clears the value of the "creator" field.
func (u *OrgPositionUpsertBulk) ClearCreator() *OrgPositionUpsertBulk {
	return u.Update(func(s *OrgPositionUpsert) {
		s.ClearCreator()
	})
}

// Exec executes the query.
func (u *OrgPositionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entscheme: OnConflict was set for builder %d. Set it on the OrgPositionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entscheme: missing options for OrgPositionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgPositionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}