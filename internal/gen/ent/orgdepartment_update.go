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
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/orgdepartment"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// OrgDepartmentUpdate is the builder for updating OrgDepartment entities.
type OrgDepartmentUpdate struct {
	config
	hooks     []Hook
	mutation  *OrgDepartmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrgDepartmentUpdate builder.
func (odu *OrgDepartmentUpdate) Where(ps ...predicate.OrgDepartment) *OrgDepartmentUpdate {
	odu.mutation.Where(ps...)
	return odu
}

// SetIsDel sets the "is_del" field.
func (odu *OrgDepartmentUpdate) SetIsDel(b bool) *OrgDepartmentUpdate {
	odu.mutation.SetIsDel(b)
	return odu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableIsDel(b *bool) *OrgDepartmentUpdate {
	if b != nil {
		odu.SetIsDel(*b)
	}
	return odu
}

// SetSort sets the "sort" field.
func (odu *OrgDepartmentUpdate) SetSort(i int32) *OrgDepartmentUpdate {
	odu.mutation.ResetSort()
	odu.mutation.SetSort(i)
	return odu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableSort(i *int32) *OrgDepartmentUpdate {
	if i != nil {
		odu.SetSort(*i)
	}
	return odu
}

// AddSort adds i to the "sort" field.
func (odu *OrgDepartmentUpdate) AddSort(i int32) *OrgDepartmentUpdate {
	odu.mutation.AddSort(i)
	return odu
}

// SetUpdatedAt sets the "updated_at" field.
func (odu *OrgDepartmentUpdate) SetUpdatedAt(t time.Time) *OrgDepartmentUpdate {
	odu.mutation.SetUpdatedAt(t)
	return odu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (odu *OrgDepartmentUpdate) ClearUpdatedAt() *OrgDepartmentUpdate {
	odu.mutation.ClearUpdatedAt()
	return odu
}

// SetDeletedAt sets the "deleted_at" field.
func (odu *OrgDepartmentUpdate) SetDeletedAt(t time.Time) *OrgDepartmentUpdate {
	odu.mutation.SetDeletedAt(t)
	return odu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableDeletedAt(t *time.Time) *OrgDepartmentUpdate {
	if t != nil {
		odu.SetDeletedAt(*t)
	}
	return odu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (odu *OrgDepartmentUpdate) ClearDeletedAt() *OrgDepartmentUpdate {
	odu.mutation.ClearDeletedAt()
	return odu
}

// SetIsActive sets the "is_active" field.
func (odu *OrgDepartmentUpdate) SetIsActive(b bool) *OrgDepartmentUpdate {
	odu.mutation.SetIsActive(b)
	return odu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableIsActive(b *bool) *OrgDepartmentUpdate {
	if b != nil {
		odu.SetIsActive(*b)
	}
	return odu
}

// SetMemo sets the "memo" field.
func (odu *OrgDepartmentUpdate) SetMemo(s string) *OrgDepartmentUpdate {
	odu.mutation.SetMemo(s)
	return odu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableMemo(s *string) *OrgDepartmentUpdate {
	if s != nil {
		odu.SetMemo(*s)
	}
	return odu
}

// ClearMemo clears the value of the "memo" field.
func (odu *OrgDepartmentUpdate) ClearMemo() *OrgDepartmentUpdate {
	odu.mutation.ClearMemo()
	return odu
}

// SetName sets the "name" field.
func (odu *OrgDepartmentUpdate) SetName(s string) *OrgDepartmentUpdate {
	odu.mutation.SetName(s)
	return odu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableName(s *string) *OrgDepartmentUpdate {
	if s != nil {
		odu.SetName(*s)
	}
	return odu
}

// ClearName clears the value of the "name" field.
func (odu *OrgDepartmentUpdate) ClearName() *OrgDepartmentUpdate {
	odu.mutation.ClearName()
	return odu
}

// SetCode sets the "code" field.
func (odu *OrgDepartmentUpdate) SetCode(s string) *OrgDepartmentUpdate {
	odu.mutation.SetCode(s)
	return odu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableCode(s *string) *OrgDepartmentUpdate {
	if s != nil {
		odu.SetCode(*s)
	}
	return odu
}

// ClearCode clears the value of the "code" field.
func (odu *OrgDepartmentUpdate) ClearCode() *OrgDepartmentUpdate {
	odu.mutation.ClearCode()
	return odu
}

// SetOrgID sets the "org_id" field.
func (odu *OrgDepartmentUpdate) SetOrgID(s string) *OrgDepartmentUpdate {
	odu.mutation.SetOrgID(s)
	return odu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableOrgID(s *string) *OrgDepartmentUpdate {
	if s != nil {
		odu.SetOrgID(*s)
	}
	return odu
}

// ClearOrgID clears the value of the "org_id" field.
func (odu *OrgDepartmentUpdate) ClearOrgID() *OrgDepartmentUpdate {
	odu.mutation.ClearOrgID()
	return odu
}

// SetCreator sets the "creator" field.
func (odu *OrgDepartmentUpdate) SetCreator(s string) *OrgDepartmentUpdate {
	odu.mutation.SetCreator(s)
	return odu
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableCreator(s *string) *OrgDepartmentUpdate {
	if s != nil {
		odu.SetCreator(*s)
	}
	return odu
}

// ClearCreator clears the value of the "creator" field.
func (odu *OrgDepartmentUpdate) ClearCreator() *OrgDepartmentUpdate {
	odu.mutation.ClearCreator()
	return odu
}

// SetOrganID sets the "organ" edge to the OrgOrgan entity by ID.
func (odu *OrgDepartmentUpdate) SetOrganID(id string) *OrgDepartmentUpdate {
	odu.mutation.SetOrganID(id)
	return odu
}

// SetNillableOrganID sets the "organ" edge to the OrgOrgan entity by ID if the given value is not nil.
func (odu *OrgDepartmentUpdate) SetNillableOrganID(id *string) *OrgDepartmentUpdate {
	if id != nil {
		odu = odu.SetOrganID(*id)
	}
	return odu
}

// SetOrgan sets the "organ" edge to the OrgOrgan entity.
func (odu *OrgDepartmentUpdate) SetOrgan(o *OrgOrgan) *OrgDepartmentUpdate {
	return odu.SetOrganID(o.ID)
}

// Mutation returns the OrgDepartmentMutation object of the builder.
func (odu *OrgDepartmentUpdate) Mutation() *OrgDepartmentMutation {
	return odu.mutation
}

// ClearOrgan clears the "organ" edge to the OrgOrgan entity.
func (odu *OrgDepartmentUpdate) ClearOrgan() *OrgDepartmentUpdate {
	odu.mutation.ClearOrgan()
	return odu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odu *OrgDepartmentUpdate) Save(ctx context.Context) (int, error) {
	odu.defaults()
	return withHooks[int, OrgDepartmentMutation](ctx, odu.sqlSave, odu.mutation, odu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (odu *OrgDepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := odu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odu *OrgDepartmentUpdate) Exec(ctx context.Context) error {
	_, err := odu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odu *OrgDepartmentUpdate) ExecX(ctx context.Context) {
	if err := odu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (odu *OrgDepartmentUpdate) defaults() {
	if _, ok := odu.mutation.UpdatedAt(); !ok && !odu.mutation.UpdatedAtCleared() {
		v := orgdepartment.UpdateDefaultUpdatedAt()
		odu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odu *OrgDepartmentUpdate) check() error {
	if v, ok := odu.mutation.Memo(); ok {
		if err := orgdepartment.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.memo": %w`, err)}
		}
	}
	if v, ok := odu.mutation.Name(); ok {
		if err := orgdepartment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.name": %w`, err)}
		}
	}
	if v, ok := odu.mutation.Code(); ok {
		if err := orgdepartment.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.code": %w`, err)}
		}
	}
	if v, ok := odu.mutation.OrgID(); ok {
		if err := orgdepartment.OrgIDValidator(v); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.org_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (odu *OrgDepartmentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrgDepartmentUpdate {
	odu.modifiers = append(odu.modifiers, modifiers...)
	return odu
}

func (odu *OrgDepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := odu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgdepartment.Table, orgdepartment.Columns, sqlgraph.NewFieldSpec(orgdepartment.FieldID, field.TypeString))
	if ps := odu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odu.mutation.IsDel(); ok {
		_spec.SetField(orgdepartment.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := odu.mutation.Sort(); ok {
		_spec.SetField(orgdepartment.FieldSort, field.TypeInt32, value)
	}
	if value, ok := odu.mutation.AddedSort(); ok {
		_spec.AddField(orgdepartment.FieldSort, field.TypeInt32, value)
	}
	if odu.mutation.CreatedAtCleared() {
		_spec.ClearField(orgdepartment.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := odu.mutation.UpdatedAt(); ok {
		_spec.SetField(orgdepartment.FieldUpdatedAt, field.TypeTime, value)
	}
	if odu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgdepartment.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := odu.mutation.DeletedAt(); ok {
		_spec.SetField(orgdepartment.FieldDeletedAt, field.TypeTime, value)
	}
	if odu.mutation.DeletedAtCleared() {
		_spec.ClearField(orgdepartment.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := odu.mutation.IsActive(); ok {
		_spec.SetField(orgdepartment.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := odu.mutation.Memo(); ok {
		_spec.SetField(orgdepartment.FieldMemo, field.TypeString, value)
	}
	if odu.mutation.MemoCleared() {
		_spec.ClearField(orgdepartment.FieldMemo, field.TypeString)
	}
	if value, ok := odu.mutation.Name(); ok {
		_spec.SetField(orgdepartment.FieldName, field.TypeString, value)
	}
	if odu.mutation.NameCleared() {
		_spec.ClearField(orgdepartment.FieldName, field.TypeString)
	}
	if value, ok := odu.mutation.Code(); ok {
		_spec.SetField(orgdepartment.FieldCode, field.TypeString, value)
	}
	if odu.mutation.CodeCleared() {
		_spec.ClearField(orgdepartment.FieldCode, field.TypeString)
	}
	if value, ok := odu.mutation.Creator(); ok {
		_spec.SetField(orgdepartment.FieldCreator, field.TypeString, value)
	}
	if odu.mutation.CreatorCleared() {
		_spec.ClearField(orgdepartment.FieldCreator, field.TypeString)
	}
	if odu.mutation.OrganCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgdepartment.OrganTable,
			Columns: []string{orgdepartment.OrganColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgorgan.FieldID, field.TypeString),
			},
		}
		edge.Schema = odu.schemaConfig.OrgDepartment
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.OrganIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgdepartment.OrganTable,
			Columns: []string{orgdepartment.OrganColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgorgan.FieldID, field.TypeString),
			},
		}
		edge.Schema = odu.schemaConfig.OrgDepartment
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = odu.schemaConfig.OrgDepartment
	ctx = internal.NewSchemaConfigContext(ctx, odu.schemaConfig)
	_spec.AddModifiers(odu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, odu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgdepartment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	odu.mutation.done = true
	return n, nil
}

// OrgDepartmentUpdateOne is the builder for updating a single OrgDepartment entity.
type OrgDepartmentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrgDepartmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (oduo *OrgDepartmentUpdateOne) SetIsDel(b bool) *OrgDepartmentUpdateOne {
	oduo.mutation.SetIsDel(b)
	return oduo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableIsDel(b *bool) *OrgDepartmentUpdateOne {
	if b != nil {
		oduo.SetIsDel(*b)
	}
	return oduo
}

// SetSort sets the "sort" field.
func (oduo *OrgDepartmentUpdateOne) SetSort(i int32) *OrgDepartmentUpdateOne {
	oduo.mutation.ResetSort()
	oduo.mutation.SetSort(i)
	return oduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableSort(i *int32) *OrgDepartmentUpdateOne {
	if i != nil {
		oduo.SetSort(*i)
	}
	return oduo
}

// AddSort adds i to the "sort" field.
func (oduo *OrgDepartmentUpdateOne) AddSort(i int32) *OrgDepartmentUpdateOne {
	oduo.mutation.AddSort(i)
	return oduo
}

// SetUpdatedAt sets the "updated_at" field.
func (oduo *OrgDepartmentUpdateOne) SetUpdatedAt(t time.Time) *OrgDepartmentUpdateOne {
	oduo.mutation.SetUpdatedAt(t)
	return oduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oduo *OrgDepartmentUpdateOne) ClearUpdatedAt() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearUpdatedAt()
	return oduo
}

// SetDeletedAt sets the "deleted_at" field.
func (oduo *OrgDepartmentUpdateOne) SetDeletedAt(t time.Time) *OrgDepartmentUpdateOne {
	oduo.mutation.SetDeletedAt(t)
	return oduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableDeletedAt(t *time.Time) *OrgDepartmentUpdateOne {
	if t != nil {
		oduo.SetDeletedAt(*t)
	}
	return oduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (oduo *OrgDepartmentUpdateOne) ClearDeletedAt() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearDeletedAt()
	return oduo
}

// SetIsActive sets the "is_active" field.
func (oduo *OrgDepartmentUpdateOne) SetIsActive(b bool) *OrgDepartmentUpdateOne {
	oduo.mutation.SetIsActive(b)
	return oduo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableIsActive(b *bool) *OrgDepartmentUpdateOne {
	if b != nil {
		oduo.SetIsActive(*b)
	}
	return oduo
}

// SetMemo sets the "memo" field.
func (oduo *OrgDepartmentUpdateOne) SetMemo(s string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetMemo(s)
	return oduo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableMemo(s *string) *OrgDepartmentUpdateOne {
	if s != nil {
		oduo.SetMemo(*s)
	}
	return oduo
}

// ClearMemo clears the value of the "memo" field.
func (oduo *OrgDepartmentUpdateOne) ClearMemo() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearMemo()
	return oduo
}

// SetName sets the "name" field.
func (oduo *OrgDepartmentUpdateOne) SetName(s string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetName(s)
	return oduo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableName(s *string) *OrgDepartmentUpdateOne {
	if s != nil {
		oduo.SetName(*s)
	}
	return oduo
}

// ClearName clears the value of the "name" field.
func (oduo *OrgDepartmentUpdateOne) ClearName() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearName()
	return oduo
}

// SetCode sets the "code" field.
func (oduo *OrgDepartmentUpdateOne) SetCode(s string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetCode(s)
	return oduo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableCode(s *string) *OrgDepartmentUpdateOne {
	if s != nil {
		oduo.SetCode(*s)
	}
	return oduo
}

// ClearCode clears the value of the "code" field.
func (oduo *OrgDepartmentUpdateOne) ClearCode() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearCode()
	return oduo
}

// SetOrgID sets the "org_id" field.
func (oduo *OrgDepartmentUpdateOne) SetOrgID(s string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetOrgID(s)
	return oduo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableOrgID(s *string) *OrgDepartmentUpdateOne {
	if s != nil {
		oduo.SetOrgID(*s)
	}
	return oduo
}

// ClearOrgID clears the value of the "org_id" field.
func (oduo *OrgDepartmentUpdateOne) ClearOrgID() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearOrgID()
	return oduo
}

// SetCreator sets the "creator" field.
func (oduo *OrgDepartmentUpdateOne) SetCreator(s string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetCreator(s)
	return oduo
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableCreator(s *string) *OrgDepartmentUpdateOne {
	if s != nil {
		oduo.SetCreator(*s)
	}
	return oduo
}

// ClearCreator clears the value of the "creator" field.
func (oduo *OrgDepartmentUpdateOne) ClearCreator() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearCreator()
	return oduo
}

// SetOrganID sets the "organ" edge to the OrgOrgan entity by ID.
func (oduo *OrgDepartmentUpdateOne) SetOrganID(id string) *OrgDepartmentUpdateOne {
	oduo.mutation.SetOrganID(id)
	return oduo
}

// SetNillableOrganID sets the "organ" edge to the OrgOrgan entity by ID if the given value is not nil.
func (oduo *OrgDepartmentUpdateOne) SetNillableOrganID(id *string) *OrgDepartmentUpdateOne {
	if id != nil {
		oduo = oduo.SetOrganID(*id)
	}
	return oduo
}

// SetOrgan sets the "organ" edge to the OrgOrgan entity.
func (oduo *OrgDepartmentUpdateOne) SetOrgan(o *OrgOrgan) *OrgDepartmentUpdateOne {
	return oduo.SetOrganID(o.ID)
}

// Mutation returns the OrgDepartmentMutation object of the builder.
func (oduo *OrgDepartmentUpdateOne) Mutation() *OrgDepartmentMutation {
	return oduo.mutation
}

// ClearOrgan clears the "organ" edge to the OrgOrgan entity.
func (oduo *OrgDepartmentUpdateOne) ClearOrgan() *OrgDepartmentUpdateOne {
	oduo.mutation.ClearOrgan()
	return oduo
}

// Where appends a list predicates to the OrgDepartmentUpdate builder.
func (oduo *OrgDepartmentUpdateOne) Where(ps ...predicate.OrgDepartment) *OrgDepartmentUpdateOne {
	oduo.mutation.Where(ps...)
	return oduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oduo *OrgDepartmentUpdateOne) Select(field string, fields ...string) *OrgDepartmentUpdateOne {
	oduo.fields = append([]string{field}, fields...)
	return oduo
}

// Save executes the query and returns the updated OrgDepartment entity.
func (oduo *OrgDepartmentUpdateOne) Save(ctx context.Context) (*OrgDepartment, error) {
	oduo.defaults()
	return withHooks[*OrgDepartment, OrgDepartmentMutation](ctx, oduo.sqlSave, oduo.mutation, oduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oduo *OrgDepartmentUpdateOne) SaveX(ctx context.Context) *OrgDepartment {
	node, err := oduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oduo *OrgDepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := oduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oduo *OrgDepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := oduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oduo *OrgDepartmentUpdateOne) defaults() {
	if _, ok := oduo.mutation.UpdatedAt(); !ok && !oduo.mutation.UpdatedAtCleared() {
		v := orgdepartment.UpdateDefaultUpdatedAt()
		oduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oduo *OrgDepartmentUpdateOne) check() error {
	if v, ok := oduo.mutation.Memo(); ok {
		if err := orgdepartment.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.memo": %w`, err)}
		}
	}
	if v, ok := oduo.mutation.Name(); ok {
		if err := orgdepartment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.name": %w`, err)}
		}
	}
	if v, ok := oduo.mutation.Code(); ok {
		if err := orgdepartment.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.code": %w`, err)}
		}
	}
	if v, ok := oduo.mutation.OrgID(); ok {
		if err := orgdepartment.OrgIDValidator(v); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`ent: validator failed for field "OrgDepartment.org_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oduo *OrgDepartmentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrgDepartmentUpdateOne {
	oduo.modifiers = append(oduo.modifiers, modifiers...)
	return oduo
}

func (oduo *OrgDepartmentUpdateOne) sqlSave(ctx context.Context) (_node *OrgDepartment, err error) {
	if err := oduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgdepartment.Table, orgdepartment.Columns, sqlgraph.NewFieldSpec(orgdepartment.FieldID, field.TypeString))
	id, ok := oduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgDepartment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgdepartment.FieldID)
		for _, f := range fields {
			if !orgdepartment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orgdepartment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oduo.mutation.IsDel(); ok {
		_spec.SetField(orgdepartment.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := oduo.mutation.Sort(); ok {
		_spec.SetField(orgdepartment.FieldSort, field.TypeInt32, value)
	}
	if value, ok := oduo.mutation.AddedSort(); ok {
		_spec.AddField(orgdepartment.FieldSort, field.TypeInt32, value)
	}
	if oduo.mutation.CreatedAtCleared() {
		_spec.ClearField(orgdepartment.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := oduo.mutation.UpdatedAt(); ok {
		_spec.SetField(orgdepartment.FieldUpdatedAt, field.TypeTime, value)
	}
	if oduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgdepartment.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oduo.mutation.DeletedAt(); ok {
		_spec.SetField(orgdepartment.FieldDeletedAt, field.TypeTime, value)
	}
	if oduo.mutation.DeletedAtCleared() {
		_spec.ClearField(orgdepartment.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := oduo.mutation.IsActive(); ok {
		_spec.SetField(orgdepartment.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := oduo.mutation.Memo(); ok {
		_spec.SetField(orgdepartment.FieldMemo, field.TypeString, value)
	}
	if oduo.mutation.MemoCleared() {
		_spec.ClearField(orgdepartment.FieldMemo, field.TypeString)
	}
	if value, ok := oduo.mutation.Name(); ok {
		_spec.SetField(orgdepartment.FieldName, field.TypeString, value)
	}
	if oduo.mutation.NameCleared() {
		_spec.ClearField(orgdepartment.FieldName, field.TypeString)
	}
	if value, ok := oduo.mutation.Code(); ok {
		_spec.SetField(orgdepartment.FieldCode, field.TypeString, value)
	}
	if oduo.mutation.CodeCleared() {
		_spec.ClearField(orgdepartment.FieldCode, field.TypeString)
	}
	if value, ok := oduo.mutation.Creator(); ok {
		_spec.SetField(orgdepartment.FieldCreator, field.TypeString, value)
	}
	if oduo.mutation.CreatorCleared() {
		_spec.ClearField(orgdepartment.FieldCreator, field.TypeString)
	}
	if oduo.mutation.OrganCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgdepartment.OrganTable,
			Columns: []string{orgdepartment.OrganColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgorgan.FieldID, field.TypeString),
			},
		}
		edge.Schema = oduo.schemaConfig.OrgDepartment
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.OrganIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgdepartment.OrganTable,
			Columns: []string{orgdepartment.OrganColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgorgan.FieldID, field.TypeString),
			},
		}
		edge.Schema = oduo.schemaConfig.OrgDepartment
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = oduo.schemaConfig.OrgDepartment
	ctx = internal.NewSchemaConfigContext(ctx, oduo.schemaConfig)
	_spec.AddModifiers(oduo.modifiers...)
	_node = &OrgDepartment{config: oduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgdepartment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oduo.mutation.done = true
	return _node, nil
}
