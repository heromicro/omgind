// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/xxxdemo"
)

// XxxDemoUpdate is the builder for updating XxxDemo entities.
type XxxDemoUpdate struct {
	config
	hooks     []Hook
	mutation  *XxxDemoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the XxxDemoUpdate builder.
func (xdu *XxxDemoUpdate) Where(ps ...predicate.XxxDemo) *XxxDemoUpdate {
	xdu.mutation.Where(ps...)
	return xdu
}

// SetIsDel sets the "is_del" field.
func (xdu *XxxDemoUpdate) SetIsDel(b bool) *XxxDemoUpdate {
	xdu.mutation.SetIsDel(b)
	return xdu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableIsDel(b *bool) *XxxDemoUpdate {
	if b != nil {
		xdu.SetIsDel(*b)
	}
	return xdu
}

// SetMemo sets the "memo" field.
func (xdu *XxxDemoUpdate) SetMemo(s string) *XxxDemoUpdate {
	xdu.mutation.SetMemo(s)
	return xdu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableMemo(s *string) *XxxDemoUpdate {
	if s != nil {
		xdu.SetMemo(*s)
	}
	return xdu
}

// ClearMemo clears the value of the "memo" field.
func (xdu *XxxDemoUpdate) ClearMemo() *XxxDemoUpdate {
	xdu.mutation.ClearMemo()
	return xdu
}

// SetSort sets the "sort" field.
func (xdu *XxxDemoUpdate) SetSort(i int32) *XxxDemoUpdate {
	xdu.mutation.ResetSort()
	xdu.mutation.SetSort(i)
	return xdu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableSort(i *int32) *XxxDemoUpdate {
	if i != nil {
		xdu.SetSort(*i)
	}
	return xdu
}

// AddSort adds i to the "sort" field.
func (xdu *XxxDemoUpdate) AddSort(i int32) *XxxDemoUpdate {
	xdu.mutation.AddSort(i)
	return xdu
}

// SetUpdatedAt sets the "updated_at" field.
func (xdu *XxxDemoUpdate) SetUpdatedAt(t time.Time) *XxxDemoUpdate {
	xdu.mutation.SetUpdatedAt(t)
	return xdu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (xdu *XxxDemoUpdate) ClearUpdatedAt() *XxxDemoUpdate {
	xdu.mutation.ClearUpdatedAt()
	return xdu
}

// SetDeletedAt sets the "deleted_at" field.
func (xdu *XxxDemoUpdate) SetDeletedAt(t time.Time) *XxxDemoUpdate {
	xdu.mutation.SetDeletedAt(t)
	return xdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableDeletedAt(t *time.Time) *XxxDemoUpdate {
	if t != nil {
		xdu.SetDeletedAt(*t)
	}
	return xdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (xdu *XxxDemoUpdate) ClearDeletedAt() *XxxDemoUpdate {
	xdu.mutation.ClearDeletedAt()
	return xdu
}

// SetIsActive sets the "is_active" field.
func (xdu *XxxDemoUpdate) SetIsActive(b bool) *XxxDemoUpdate {
	xdu.mutation.SetIsActive(b)
	return xdu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableIsActive(b *bool) *XxxDemoUpdate {
	if b != nil {
		xdu.SetIsActive(*b)
	}
	return xdu
}

// SetCode sets the "code" field.
func (xdu *XxxDemoUpdate) SetCode(s string) *XxxDemoUpdate {
	xdu.mutation.SetCode(s)
	return xdu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableCode(s *string) *XxxDemoUpdate {
	if s != nil {
		xdu.SetCode(*s)
	}
	return xdu
}

// SetName sets the "name" field.
func (xdu *XxxDemoUpdate) SetName(s string) *XxxDemoUpdate {
	xdu.mutation.SetName(s)
	return xdu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableName(s *string) *XxxDemoUpdate {
	if s != nil {
		xdu.SetName(*s)
	}
	return xdu
}

// Mutation returns the XxxDemoMutation object of the builder.
func (xdu *XxxDemoUpdate) Mutation() *XxxDemoMutation {
	return xdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xdu *XxxDemoUpdate) Save(ctx context.Context) (int, error) {
	xdu.defaults()
	return withHooks(ctx, xdu.sqlSave, xdu.mutation, xdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (xdu *XxxDemoUpdate) SaveX(ctx context.Context) int {
	affected, err := xdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xdu *XxxDemoUpdate) Exec(ctx context.Context) error {
	_, err := xdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xdu *XxxDemoUpdate) ExecX(ctx context.Context) {
	if err := xdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xdu *XxxDemoUpdate) defaults() {
	if _, ok := xdu.mutation.UpdatedAt(); !ok && !xdu.mutation.UpdatedAtCleared() {
		v := xxxdemo.UpdateDefaultUpdatedAt()
		xdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xdu *XxxDemoUpdate) check() error {
	if v, ok := xdu.mutation.Memo(); ok {
		if err := xxxdemo.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.memo": %w`, err)}
		}
	}
	if v, ok := xdu.mutation.Code(); ok {
		if err := xxxdemo.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.code": %w`, err)}
		}
	}
	if v, ok := xdu.mutation.Name(); ok {
		if err := xxxdemo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (xdu *XxxDemoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *XxxDemoUpdate {
	xdu.modifiers = append(xdu.modifiers, modifiers...)
	return xdu
}

func (xdu *XxxDemoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := xdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(xxxdemo.Table, xxxdemo.Columns, sqlgraph.NewFieldSpec(xxxdemo.FieldID, field.TypeString))
	if ps := xdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xdu.mutation.IsDel(); ok {
		_spec.SetField(xxxdemo.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := xdu.mutation.Memo(); ok {
		_spec.SetField(xxxdemo.FieldMemo, field.TypeString, value)
	}
	if xdu.mutation.MemoCleared() {
		_spec.ClearField(xxxdemo.FieldMemo, field.TypeString)
	}
	if value, ok := xdu.mutation.Sort(); ok {
		_spec.SetField(xxxdemo.FieldSort, field.TypeInt32, value)
	}
	if value, ok := xdu.mutation.AddedSort(); ok {
		_spec.AddField(xxxdemo.FieldSort, field.TypeInt32, value)
	}
	if xdu.mutation.CreatedAtCleared() {
		_spec.ClearField(xxxdemo.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := xdu.mutation.UpdatedAt(); ok {
		_spec.SetField(xxxdemo.FieldUpdatedAt, field.TypeTime, value)
	}
	if xdu.mutation.UpdatedAtCleared() {
		_spec.ClearField(xxxdemo.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := xdu.mutation.DeletedAt(); ok {
		_spec.SetField(xxxdemo.FieldDeletedAt, field.TypeTime, value)
	}
	if xdu.mutation.DeletedAtCleared() {
		_spec.ClearField(xxxdemo.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := xdu.mutation.IsActive(); ok {
		_spec.SetField(xxxdemo.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := xdu.mutation.Code(); ok {
		_spec.SetField(xxxdemo.FieldCode, field.TypeString, value)
	}
	if value, ok := xdu.mutation.Name(); ok {
		_spec.SetField(xxxdemo.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(xdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, xdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xxxdemo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	xdu.mutation.done = true
	return n, nil
}

// XxxDemoUpdateOne is the builder for updating a single XxxDemo entity.
type XxxDemoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *XxxDemoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (xduo *XxxDemoUpdateOne) SetIsDel(b bool) *XxxDemoUpdateOne {
	xduo.mutation.SetIsDel(b)
	return xduo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableIsDel(b *bool) *XxxDemoUpdateOne {
	if b != nil {
		xduo.SetIsDel(*b)
	}
	return xduo
}

// SetMemo sets the "memo" field.
func (xduo *XxxDemoUpdateOne) SetMemo(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetMemo(s)
	return xduo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableMemo(s *string) *XxxDemoUpdateOne {
	if s != nil {
		xduo.SetMemo(*s)
	}
	return xduo
}

// ClearMemo clears the value of the "memo" field.
func (xduo *XxxDemoUpdateOne) ClearMemo() *XxxDemoUpdateOne {
	xduo.mutation.ClearMemo()
	return xduo
}

// SetSort sets the "sort" field.
func (xduo *XxxDemoUpdateOne) SetSort(i int32) *XxxDemoUpdateOne {
	xduo.mutation.ResetSort()
	xduo.mutation.SetSort(i)
	return xduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableSort(i *int32) *XxxDemoUpdateOne {
	if i != nil {
		xduo.SetSort(*i)
	}
	return xduo
}

// AddSort adds i to the "sort" field.
func (xduo *XxxDemoUpdateOne) AddSort(i int32) *XxxDemoUpdateOne {
	xduo.mutation.AddSort(i)
	return xduo
}

// SetUpdatedAt sets the "updated_at" field.
func (xduo *XxxDemoUpdateOne) SetUpdatedAt(t time.Time) *XxxDemoUpdateOne {
	xduo.mutation.SetUpdatedAt(t)
	return xduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (xduo *XxxDemoUpdateOne) ClearUpdatedAt() *XxxDemoUpdateOne {
	xduo.mutation.ClearUpdatedAt()
	return xduo
}

// SetDeletedAt sets the "deleted_at" field.
func (xduo *XxxDemoUpdateOne) SetDeletedAt(t time.Time) *XxxDemoUpdateOne {
	xduo.mutation.SetDeletedAt(t)
	return xduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableDeletedAt(t *time.Time) *XxxDemoUpdateOne {
	if t != nil {
		xduo.SetDeletedAt(*t)
	}
	return xduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (xduo *XxxDemoUpdateOne) ClearDeletedAt() *XxxDemoUpdateOne {
	xduo.mutation.ClearDeletedAt()
	return xduo
}

// SetIsActive sets the "is_active" field.
func (xduo *XxxDemoUpdateOne) SetIsActive(b bool) *XxxDemoUpdateOne {
	xduo.mutation.SetIsActive(b)
	return xduo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableIsActive(b *bool) *XxxDemoUpdateOne {
	if b != nil {
		xduo.SetIsActive(*b)
	}
	return xduo
}

// SetCode sets the "code" field.
func (xduo *XxxDemoUpdateOne) SetCode(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetCode(s)
	return xduo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableCode(s *string) *XxxDemoUpdateOne {
	if s != nil {
		xduo.SetCode(*s)
	}
	return xduo
}

// SetName sets the "name" field.
func (xduo *XxxDemoUpdateOne) SetName(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetName(s)
	return xduo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableName(s *string) *XxxDemoUpdateOne {
	if s != nil {
		xduo.SetName(*s)
	}
	return xduo
}

// Mutation returns the XxxDemoMutation object of the builder.
func (xduo *XxxDemoUpdateOne) Mutation() *XxxDemoMutation {
	return xduo.mutation
}

// Where appends a list predicates to the XxxDemoUpdate builder.
func (xduo *XxxDemoUpdateOne) Where(ps ...predicate.XxxDemo) *XxxDemoUpdateOne {
	xduo.mutation.Where(ps...)
	return xduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xduo *XxxDemoUpdateOne) Select(field string, fields ...string) *XxxDemoUpdateOne {
	xduo.fields = append([]string{field}, fields...)
	return xduo
}

// Save executes the query and returns the updated XxxDemo entity.
func (xduo *XxxDemoUpdateOne) Save(ctx context.Context) (*XxxDemo, error) {
	xduo.defaults()
	return withHooks(ctx, xduo.sqlSave, xduo.mutation, xduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (xduo *XxxDemoUpdateOne) SaveX(ctx context.Context) *XxxDemo {
	node, err := xduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xduo *XxxDemoUpdateOne) Exec(ctx context.Context) error {
	_, err := xduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xduo *XxxDemoUpdateOne) ExecX(ctx context.Context) {
	if err := xduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xduo *XxxDemoUpdateOne) defaults() {
	if _, ok := xduo.mutation.UpdatedAt(); !ok && !xduo.mutation.UpdatedAtCleared() {
		v := xxxdemo.UpdateDefaultUpdatedAt()
		xduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xduo *XxxDemoUpdateOne) check() error {
	if v, ok := xduo.mutation.Memo(); ok {
		if err := xxxdemo.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.memo": %w`, err)}
		}
	}
	if v, ok := xduo.mutation.Code(); ok {
		if err := xxxdemo.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.code": %w`, err)}
		}
	}
	if v, ok := xduo.mutation.Name(); ok {
		if err := xxxdemo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`mainent: validator failed for field "XxxDemo.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (xduo *XxxDemoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *XxxDemoUpdateOne {
	xduo.modifiers = append(xduo.modifiers, modifiers...)
	return xduo
}

func (xduo *XxxDemoUpdateOne) sqlSave(ctx context.Context) (_node *XxxDemo, err error) {
	if err := xduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(xxxdemo.Table, xxxdemo.Columns, sqlgraph.NewFieldSpec(xxxdemo.FieldID, field.TypeString))
	id, ok := xduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "XxxDemo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := xduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xxxdemo.FieldID)
		for _, f := range fields {
			if !xxxdemo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != xxxdemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xduo.mutation.IsDel(); ok {
		_spec.SetField(xxxdemo.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := xduo.mutation.Memo(); ok {
		_spec.SetField(xxxdemo.FieldMemo, field.TypeString, value)
	}
	if xduo.mutation.MemoCleared() {
		_spec.ClearField(xxxdemo.FieldMemo, field.TypeString)
	}
	if value, ok := xduo.mutation.Sort(); ok {
		_spec.SetField(xxxdemo.FieldSort, field.TypeInt32, value)
	}
	if value, ok := xduo.mutation.AddedSort(); ok {
		_spec.AddField(xxxdemo.FieldSort, field.TypeInt32, value)
	}
	if xduo.mutation.CreatedAtCleared() {
		_spec.ClearField(xxxdemo.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := xduo.mutation.UpdatedAt(); ok {
		_spec.SetField(xxxdemo.FieldUpdatedAt, field.TypeTime, value)
	}
	if xduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(xxxdemo.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := xduo.mutation.DeletedAt(); ok {
		_spec.SetField(xxxdemo.FieldDeletedAt, field.TypeTime, value)
	}
	if xduo.mutation.DeletedAtCleared() {
		_spec.ClearField(xxxdemo.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := xduo.mutation.IsActive(); ok {
		_spec.SetField(xxxdemo.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := xduo.mutation.Code(); ok {
		_spec.SetField(xxxdemo.FieldCode, field.TypeString, value)
	}
	if value, ok := xduo.mutation.Name(); ok {
		_spec.SetField(xxxdemo.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(xduo.modifiers...)
	_node = &XxxDemo{config: xduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xxxdemo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	xduo.mutation.done = true
	return _node, nil
}
