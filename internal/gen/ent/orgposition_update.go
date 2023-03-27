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
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/orgposition"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// OrgPositionUpdate is the builder for updating OrgPosition entities.
type OrgPositionUpdate struct {
	config
	hooks     []Hook
	mutation  *OrgPositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrgPositionUpdate builder.
func (opu *OrgPositionUpdate) Where(ps ...predicate.OrgPosition) *OrgPositionUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetIsDel sets the "is_del" field.
func (opu *OrgPositionUpdate) SetIsDel(b bool) *OrgPositionUpdate {
	opu.mutation.SetIsDel(b)
	return opu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableIsDel(b *bool) *OrgPositionUpdate {
	if b != nil {
		opu.SetIsDel(*b)
	}
	return opu
}

// SetSort sets the "sort" field.
func (opu *OrgPositionUpdate) SetSort(i int32) *OrgPositionUpdate {
	opu.mutation.ResetSort()
	opu.mutation.SetSort(i)
	return opu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableSort(i *int32) *OrgPositionUpdate {
	if i != nil {
		opu.SetSort(*i)
	}
	return opu
}

// AddSort adds i to the "sort" field.
func (opu *OrgPositionUpdate) AddSort(i int32) *OrgPositionUpdate {
	opu.mutation.AddSort(i)
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OrgPositionUpdate) SetUpdatedAt(t time.Time) *OrgPositionUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opu *OrgPositionUpdate) ClearUpdatedAt() *OrgPositionUpdate {
	opu.mutation.ClearUpdatedAt()
	return opu
}

// SetDeletedAt sets the "deleted_at" field.
func (opu *OrgPositionUpdate) SetDeletedAt(t time.Time) *OrgPositionUpdate {
	opu.mutation.SetDeletedAt(t)
	return opu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableDeletedAt(t *time.Time) *OrgPositionUpdate {
	if t != nil {
		opu.SetDeletedAt(*t)
	}
	return opu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (opu *OrgPositionUpdate) ClearDeletedAt() *OrgPositionUpdate {
	opu.mutation.ClearDeletedAt()
	return opu
}

// SetIsActive sets the "is_active" field.
func (opu *OrgPositionUpdate) SetIsActive(b bool) *OrgPositionUpdate {
	opu.mutation.SetIsActive(b)
	return opu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableIsActive(b *bool) *OrgPositionUpdate {
	if b != nil {
		opu.SetIsActive(*b)
	}
	return opu
}

// SetMemo sets the "memo" field.
func (opu *OrgPositionUpdate) SetMemo(s string) *OrgPositionUpdate {
	opu.mutation.SetMemo(s)
	return opu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableMemo(s *string) *OrgPositionUpdate {
	if s != nil {
		opu.SetMemo(*s)
	}
	return opu
}

// ClearMemo clears the value of the "memo" field.
func (opu *OrgPositionUpdate) ClearMemo() *OrgPositionUpdate {
	opu.mutation.ClearMemo()
	return opu
}

// SetName sets the "name" field.
func (opu *OrgPositionUpdate) SetName(s string) *OrgPositionUpdate {
	opu.mutation.SetName(s)
	return opu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableName(s *string) *OrgPositionUpdate {
	if s != nil {
		opu.SetName(*s)
	}
	return opu
}

// ClearName clears the value of the "name" field.
func (opu *OrgPositionUpdate) ClearName() *OrgPositionUpdate {
	opu.mutation.ClearName()
	return opu
}

// SetCode sets the "code" field.
func (opu *OrgPositionUpdate) SetCode(s string) *OrgPositionUpdate {
	opu.mutation.SetCode(s)
	return opu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableCode(s *string) *OrgPositionUpdate {
	if s != nil {
		opu.SetCode(*s)
	}
	return opu
}

// ClearCode clears the value of the "code" field.
func (opu *OrgPositionUpdate) ClearCode() *OrgPositionUpdate {
	opu.mutation.ClearCode()
	return opu
}

// SetOrgID sets the "org_id" field.
func (opu *OrgPositionUpdate) SetOrgID(s string) *OrgPositionUpdate {
	opu.mutation.SetOrgID(s)
	return opu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableOrgID(s *string) *OrgPositionUpdate {
	if s != nil {
		opu.SetOrgID(*s)
	}
	return opu
}

// ClearOrgID clears the value of the "org_id" field.
func (opu *OrgPositionUpdate) ClearOrgID() *OrgPositionUpdate {
	opu.mutation.ClearOrgID()
	return opu
}

// SetCreator sets the "creator" field.
func (opu *OrgPositionUpdate) SetCreator(s string) *OrgPositionUpdate {
	opu.mutation.SetCreator(s)
	return opu
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableCreator(s *string) *OrgPositionUpdate {
	if s != nil {
		opu.SetCreator(*s)
	}
	return opu
}

// ClearCreator clears the value of the "creator" field.
func (opu *OrgPositionUpdate) ClearCreator() *OrgPositionUpdate {
	opu.mutation.ClearCreator()
	return opu
}

// SetOrganID sets the "organ" edge to the OrgOrgan entity by ID.
func (opu *OrgPositionUpdate) SetOrganID(id string) *OrgPositionUpdate {
	opu.mutation.SetOrganID(id)
	return opu
}

// SetNillableOrganID sets the "organ" edge to the OrgOrgan entity by ID if the given value is not nil.
func (opu *OrgPositionUpdate) SetNillableOrganID(id *string) *OrgPositionUpdate {
	if id != nil {
		opu = opu.SetOrganID(*id)
	}
	return opu
}

// SetOrgan sets the "organ" edge to the OrgOrgan entity.
func (opu *OrgPositionUpdate) SetOrgan(o *OrgOrgan) *OrgPositionUpdate {
	return opu.SetOrganID(o.ID)
}

// Mutation returns the OrgPositionMutation object of the builder.
func (opu *OrgPositionUpdate) Mutation() *OrgPositionMutation {
	return opu.mutation
}

// ClearOrgan clears the "organ" edge to the OrgOrgan entity.
func (opu *OrgPositionUpdate) ClearOrgan() *OrgPositionUpdate {
	opu.mutation.ClearOrgan()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrgPositionUpdate) Save(ctx context.Context) (int, error) {
	opu.defaults()
	return withHooks[int, OrgPositionMutation](ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrgPositionUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrgPositionUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrgPositionUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OrgPositionUpdate) defaults() {
	if _, ok := opu.mutation.UpdatedAt(); !ok && !opu.mutation.UpdatedAtCleared() {
		v := orgposition.UpdateDefaultUpdatedAt()
		opu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opu *OrgPositionUpdate) check() error {
	if v, ok := opu.mutation.Memo(); ok {
		if err := orgposition.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.memo": %w`, err)}
		}
	}
	if v, ok := opu.mutation.Name(); ok {
		if err := orgposition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.name": %w`, err)}
		}
	}
	if v, ok := opu.mutation.Code(); ok {
		if err := orgposition.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.code": %w`, err)}
		}
	}
	if v, ok := opu.mutation.OrgID(); ok {
		if err := orgposition.OrgIDValidator(v); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.org_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opu *OrgPositionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrgPositionUpdate {
	opu.modifiers = append(opu.modifiers, modifiers...)
	return opu
}

func (opu *OrgPositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := opu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgposition.Table, orgposition.Columns, sqlgraph.NewFieldSpec(orgposition.FieldID, field.TypeString))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.IsDel(); ok {
		_spec.SetField(orgposition.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := opu.mutation.Sort(); ok {
		_spec.SetField(orgposition.FieldSort, field.TypeInt32, value)
	}
	if value, ok := opu.mutation.AddedSort(); ok {
		_spec.AddField(orgposition.FieldSort, field.TypeInt32, value)
	}
	if opu.mutation.CreatedAtCleared() {
		_spec.ClearField(orgposition.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(orgposition.FieldUpdatedAt, field.TypeTime, value)
	}
	if opu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgposition.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.DeletedAt(); ok {
		_spec.SetField(orgposition.FieldDeletedAt, field.TypeTime, value)
	}
	if opu.mutation.DeletedAtCleared() {
		_spec.ClearField(orgposition.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.IsActive(); ok {
		_spec.SetField(orgposition.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := opu.mutation.Memo(); ok {
		_spec.SetField(orgposition.FieldMemo, field.TypeString, value)
	}
	if opu.mutation.MemoCleared() {
		_spec.ClearField(orgposition.FieldMemo, field.TypeString)
	}
	if value, ok := opu.mutation.Name(); ok {
		_spec.SetField(orgposition.FieldName, field.TypeString, value)
	}
	if opu.mutation.NameCleared() {
		_spec.ClearField(orgposition.FieldName, field.TypeString)
	}
	if value, ok := opu.mutation.Code(); ok {
		_spec.SetField(orgposition.FieldCode, field.TypeString, value)
	}
	if opu.mutation.CodeCleared() {
		_spec.ClearField(orgposition.FieldCode, field.TypeString)
	}
	if value, ok := opu.mutation.Creator(); ok {
		_spec.SetField(orgposition.FieldCreator, field.TypeString, value)
	}
	if opu.mutation.CreatorCleared() {
		_spec.ClearField(orgposition.FieldCreator, field.TypeString)
	}
	if opu.mutation.OrganCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OrganIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(opu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OrgPositionUpdateOne is the builder for updating a single OrgPosition entity.
type OrgPositionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrgPositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (opuo *OrgPositionUpdateOne) SetIsDel(b bool) *OrgPositionUpdateOne {
	opuo.mutation.SetIsDel(b)
	return opuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableIsDel(b *bool) *OrgPositionUpdateOne {
	if b != nil {
		opuo.SetIsDel(*b)
	}
	return opuo
}

// SetSort sets the "sort" field.
func (opuo *OrgPositionUpdateOne) SetSort(i int32) *OrgPositionUpdateOne {
	opuo.mutation.ResetSort()
	opuo.mutation.SetSort(i)
	return opuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableSort(i *int32) *OrgPositionUpdateOne {
	if i != nil {
		opuo.SetSort(*i)
	}
	return opuo
}

// AddSort adds i to the "sort" field.
func (opuo *OrgPositionUpdateOne) AddSort(i int32) *OrgPositionUpdateOne {
	opuo.mutation.AddSort(i)
	return opuo
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OrgPositionUpdateOne) SetUpdatedAt(t time.Time) *OrgPositionUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opuo *OrgPositionUpdateOne) ClearUpdatedAt() *OrgPositionUpdateOne {
	opuo.mutation.ClearUpdatedAt()
	return opuo
}

// SetDeletedAt sets the "deleted_at" field.
func (opuo *OrgPositionUpdateOne) SetDeletedAt(t time.Time) *OrgPositionUpdateOne {
	opuo.mutation.SetDeletedAt(t)
	return opuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableDeletedAt(t *time.Time) *OrgPositionUpdateOne {
	if t != nil {
		opuo.SetDeletedAt(*t)
	}
	return opuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (opuo *OrgPositionUpdateOne) ClearDeletedAt() *OrgPositionUpdateOne {
	opuo.mutation.ClearDeletedAt()
	return opuo
}

// SetIsActive sets the "is_active" field.
func (opuo *OrgPositionUpdateOne) SetIsActive(b bool) *OrgPositionUpdateOne {
	opuo.mutation.SetIsActive(b)
	return opuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableIsActive(b *bool) *OrgPositionUpdateOne {
	if b != nil {
		opuo.SetIsActive(*b)
	}
	return opuo
}

// SetMemo sets the "memo" field.
func (opuo *OrgPositionUpdateOne) SetMemo(s string) *OrgPositionUpdateOne {
	opuo.mutation.SetMemo(s)
	return opuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableMemo(s *string) *OrgPositionUpdateOne {
	if s != nil {
		opuo.SetMemo(*s)
	}
	return opuo
}

// ClearMemo clears the value of the "memo" field.
func (opuo *OrgPositionUpdateOne) ClearMemo() *OrgPositionUpdateOne {
	opuo.mutation.ClearMemo()
	return opuo
}

// SetName sets the "name" field.
func (opuo *OrgPositionUpdateOne) SetName(s string) *OrgPositionUpdateOne {
	opuo.mutation.SetName(s)
	return opuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableName(s *string) *OrgPositionUpdateOne {
	if s != nil {
		opuo.SetName(*s)
	}
	return opuo
}

// ClearName clears the value of the "name" field.
func (opuo *OrgPositionUpdateOne) ClearName() *OrgPositionUpdateOne {
	opuo.mutation.ClearName()
	return opuo
}

// SetCode sets the "code" field.
func (opuo *OrgPositionUpdateOne) SetCode(s string) *OrgPositionUpdateOne {
	opuo.mutation.SetCode(s)
	return opuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableCode(s *string) *OrgPositionUpdateOne {
	if s != nil {
		opuo.SetCode(*s)
	}
	return opuo
}

// ClearCode clears the value of the "code" field.
func (opuo *OrgPositionUpdateOne) ClearCode() *OrgPositionUpdateOne {
	opuo.mutation.ClearCode()
	return opuo
}

// SetOrgID sets the "org_id" field.
func (opuo *OrgPositionUpdateOne) SetOrgID(s string) *OrgPositionUpdateOne {
	opuo.mutation.SetOrgID(s)
	return opuo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableOrgID(s *string) *OrgPositionUpdateOne {
	if s != nil {
		opuo.SetOrgID(*s)
	}
	return opuo
}

// ClearOrgID clears the value of the "org_id" field.
func (opuo *OrgPositionUpdateOne) ClearOrgID() *OrgPositionUpdateOne {
	opuo.mutation.ClearOrgID()
	return opuo
}

// SetCreator sets the "creator" field.
func (opuo *OrgPositionUpdateOne) SetCreator(s string) *OrgPositionUpdateOne {
	opuo.mutation.SetCreator(s)
	return opuo
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableCreator(s *string) *OrgPositionUpdateOne {
	if s != nil {
		opuo.SetCreator(*s)
	}
	return opuo
}

// ClearCreator clears the value of the "creator" field.
func (opuo *OrgPositionUpdateOne) ClearCreator() *OrgPositionUpdateOne {
	opuo.mutation.ClearCreator()
	return opuo
}

// SetOrganID sets the "organ" edge to the OrgOrgan entity by ID.
func (opuo *OrgPositionUpdateOne) SetOrganID(id string) *OrgPositionUpdateOne {
	opuo.mutation.SetOrganID(id)
	return opuo
}

// SetNillableOrganID sets the "organ" edge to the OrgOrgan entity by ID if the given value is not nil.
func (opuo *OrgPositionUpdateOne) SetNillableOrganID(id *string) *OrgPositionUpdateOne {
	if id != nil {
		opuo = opuo.SetOrganID(*id)
	}
	return opuo
}

// SetOrgan sets the "organ" edge to the OrgOrgan entity.
func (opuo *OrgPositionUpdateOne) SetOrgan(o *OrgOrgan) *OrgPositionUpdateOne {
	return opuo.SetOrganID(o.ID)
}

// Mutation returns the OrgPositionMutation object of the builder.
func (opuo *OrgPositionUpdateOne) Mutation() *OrgPositionMutation {
	return opuo.mutation
}

// ClearOrgan clears the "organ" edge to the OrgOrgan entity.
func (opuo *OrgPositionUpdateOne) ClearOrgan() *OrgPositionUpdateOne {
	opuo.mutation.ClearOrgan()
	return opuo
}

// Where appends a list predicates to the OrgPositionUpdate builder.
func (opuo *OrgPositionUpdateOne) Where(ps ...predicate.OrgPosition) *OrgPositionUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrgPositionUpdateOne) Select(field string, fields ...string) *OrgPositionUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrgPosition entity.
func (opuo *OrgPositionUpdateOne) Save(ctx context.Context) (*OrgPosition, error) {
	opuo.defaults()
	return withHooks[*OrgPosition, OrgPositionMutation](ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrgPositionUpdateOne) SaveX(ctx context.Context) *OrgPosition {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrgPositionUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrgPositionUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OrgPositionUpdateOne) defaults() {
	if _, ok := opuo.mutation.UpdatedAt(); !ok && !opuo.mutation.UpdatedAtCleared() {
		v := orgposition.UpdateDefaultUpdatedAt()
		opuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opuo *OrgPositionUpdateOne) check() error {
	if v, ok := opuo.mutation.Memo(); ok {
		if err := orgposition.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.memo": %w`, err)}
		}
	}
	if v, ok := opuo.mutation.Name(); ok {
		if err := orgposition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.name": %w`, err)}
		}
	}
	if v, ok := opuo.mutation.Code(); ok {
		if err := orgposition.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.code": %w`, err)}
		}
	}
	if v, ok := opuo.mutation.OrgID(); ok {
		if err := orgposition.OrgIDValidator(v); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`ent: validator failed for field "OrgPosition.org_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opuo *OrgPositionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrgPositionUpdateOne {
	opuo.modifiers = append(opuo.modifiers, modifiers...)
	return opuo
}

func (opuo *OrgPositionUpdateOne) sqlSave(ctx context.Context) (_node *OrgPosition, err error) {
	if err := opuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgposition.Table, orgposition.Columns, sqlgraph.NewFieldSpec(orgposition.FieldID, field.TypeString))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgPosition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgposition.FieldID)
		for _, f := range fields {
			if !orgposition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orgposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.IsDel(); ok {
		_spec.SetField(orgposition.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := opuo.mutation.Sort(); ok {
		_spec.SetField(orgposition.FieldSort, field.TypeInt32, value)
	}
	if value, ok := opuo.mutation.AddedSort(); ok {
		_spec.AddField(orgposition.FieldSort, field.TypeInt32, value)
	}
	if opuo.mutation.CreatedAtCleared() {
		_spec.ClearField(orgposition.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orgposition.FieldUpdatedAt, field.TypeTime, value)
	}
	if opuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgposition.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.DeletedAt(); ok {
		_spec.SetField(orgposition.FieldDeletedAt, field.TypeTime, value)
	}
	if opuo.mutation.DeletedAtCleared() {
		_spec.ClearField(orgposition.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.IsActive(); ok {
		_spec.SetField(orgposition.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := opuo.mutation.Memo(); ok {
		_spec.SetField(orgposition.FieldMemo, field.TypeString, value)
	}
	if opuo.mutation.MemoCleared() {
		_spec.ClearField(orgposition.FieldMemo, field.TypeString)
	}
	if value, ok := opuo.mutation.Name(); ok {
		_spec.SetField(orgposition.FieldName, field.TypeString, value)
	}
	if opuo.mutation.NameCleared() {
		_spec.ClearField(orgposition.FieldName, field.TypeString)
	}
	if value, ok := opuo.mutation.Code(); ok {
		_spec.SetField(orgposition.FieldCode, field.TypeString, value)
	}
	if opuo.mutation.CodeCleared() {
		_spec.ClearField(orgposition.FieldCode, field.TypeString)
	}
	if value, ok := opuo.mutation.Creator(); ok {
		_spec.SetField(orgposition.FieldCreator, field.TypeString, value)
	}
	if opuo.mutation.CreatorCleared() {
		_spec.ClearField(orgposition.FieldCreator, field.TypeString)
	}
	if opuo.mutation.OrganCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OrganIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(opuo.modifiers...)
	_node = &OrgPosition{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
