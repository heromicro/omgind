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
	"github.com/heromicro/omgind/internal/gen/mainent/sysdict"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdictitem"
)

// SysDictItemUpdate is the builder for updating SysDictItem entities.
type SysDictItemUpdate struct {
	config
	hooks     []Hook
	mutation  *SysDictItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysDictItemUpdate builder.
func (sdiu *SysDictItemUpdate) Where(ps ...predicate.SysDictItem) *SysDictItemUpdate {
	sdiu.mutation.Where(ps...)
	return sdiu
}

// SetIsDel sets the "is_del" field.
func (sdiu *SysDictItemUpdate) SetIsDel(b bool) *SysDictItemUpdate {
	sdiu.mutation.SetIsDel(b)
	return sdiu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableIsDel(b *bool) *SysDictItemUpdate {
	if b != nil {
		sdiu.SetIsDel(*b)
	}
	return sdiu
}

// SetMemo sets the "memo" field.
func (sdiu *SysDictItemUpdate) SetMemo(s string) *SysDictItemUpdate {
	sdiu.mutation.SetMemo(s)
	return sdiu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableMemo(s *string) *SysDictItemUpdate {
	if s != nil {
		sdiu.SetMemo(*s)
	}
	return sdiu
}

// ClearMemo clears the value of the "memo" field.
func (sdiu *SysDictItemUpdate) ClearMemo() *SysDictItemUpdate {
	sdiu.mutation.ClearMemo()
	return sdiu
}

// SetSort sets the "sort" field.
func (sdiu *SysDictItemUpdate) SetSort(i int32) *SysDictItemUpdate {
	sdiu.mutation.ResetSort()
	sdiu.mutation.SetSort(i)
	return sdiu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableSort(i *int32) *SysDictItemUpdate {
	if i != nil {
		sdiu.SetSort(*i)
	}
	return sdiu
}

// AddSort adds i to the "sort" field.
func (sdiu *SysDictItemUpdate) AddSort(i int32) *SysDictItemUpdate {
	sdiu.mutation.AddSort(i)
	return sdiu
}

// SetUpdatedAt sets the "updated_at" field.
func (sdiu *SysDictItemUpdate) SetUpdatedAt(t time.Time) *SysDictItemUpdate {
	sdiu.mutation.SetUpdatedAt(t)
	return sdiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sdiu *SysDictItemUpdate) ClearUpdatedAt() *SysDictItemUpdate {
	sdiu.mutation.ClearUpdatedAt()
	return sdiu
}

// SetDeletedAt sets the "deleted_at" field.
func (sdiu *SysDictItemUpdate) SetDeletedAt(t time.Time) *SysDictItemUpdate {
	sdiu.mutation.SetDeletedAt(t)
	return sdiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableDeletedAt(t *time.Time) *SysDictItemUpdate {
	if t != nil {
		sdiu.SetDeletedAt(*t)
	}
	return sdiu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdiu *SysDictItemUpdate) ClearDeletedAt() *SysDictItemUpdate {
	sdiu.mutation.ClearDeletedAt()
	return sdiu
}

// SetIsActive sets the "is_active" field.
func (sdiu *SysDictItemUpdate) SetIsActive(b bool) *SysDictItemUpdate {
	sdiu.mutation.SetIsActive(b)
	return sdiu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableIsActive(b *bool) *SysDictItemUpdate {
	if b != nil {
		sdiu.SetIsActive(*b)
	}
	return sdiu
}

// SetLabel sets the "label" field.
func (sdiu *SysDictItemUpdate) SetLabel(s string) *SysDictItemUpdate {
	sdiu.mutation.SetLabel(s)
	return sdiu
}

// SetValue sets the "value" field.
func (sdiu *SysDictItemUpdate) SetValue(i int) *SysDictItemUpdate {
	sdiu.mutation.ResetValue()
	sdiu.mutation.SetValue(i)
	return sdiu
}

// AddValue adds i to the "value" field.
func (sdiu *SysDictItemUpdate) AddValue(i int) *SysDictItemUpdate {
	sdiu.mutation.AddValue(i)
	return sdiu
}

// SetDictID sets the "dict_id" field.
func (sdiu *SysDictItemUpdate) SetDictID(s string) *SysDictItemUpdate {
	sdiu.mutation.SetDictID(s)
	return sdiu
}

// SetNillableDictID sets the "dict_id" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableDictID(s *string) *SysDictItemUpdate {
	if s != nil {
		sdiu.SetDictID(*s)
	}
	return sdiu
}

// ClearDictID clears the value of the "dict_id" field.
func (sdiu *SysDictItemUpdate) ClearDictID() *SysDictItemUpdate {
	sdiu.mutation.ClearDictID()
	return sdiu
}

// SetDict sets the "dict" edge to the SysDict entity.
func (sdiu *SysDictItemUpdate) SetDict(s *SysDict) *SysDictItemUpdate {
	return sdiu.SetDictID(s.ID)
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdiu *SysDictItemUpdate) Mutation() *SysDictItemMutation {
	return sdiu.mutation
}

// ClearDict clears the "dict" edge to the SysDict entity.
func (sdiu *SysDictItemUpdate) ClearDict() *SysDictItemUpdate {
	sdiu.mutation.ClearDict()
	return sdiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdiu *SysDictItemUpdate) Save(ctx context.Context) (int, error) {
	sdiu.defaults()
	return withHooks(ctx, sdiu.sqlSave, sdiu.mutation, sdiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdiu *SysDictItemUpdate) SaveX(ctx context.Context) int {
	affected, err := sdiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdiu *SysDictItemUpdate) Exec(ctx context.Context) error {
	_, err := sdiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdiu *SysDictItemUpdate) ExecX(ctx context.Context) {
	if err := sdiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdiu *SysDictItemUpdate) defaults() {
	if _, ok := sdiu.mutation.UpdatedAt(); !ok && !sdiu.mutation.UpdatedAtCleared() {
		v := sysdictitem.UpdateDefaultUpdatedAt()
		sdiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdiu *SysDictItemUpdate) check() error {
	if v, ok := sdiu.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.memo": %w`, err)}
		}
	}
	if v, ok := sdiu.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.label": %w`, err)}
		}
	}
	if v, ok := sdiu.mutation.DictID(); ok {
		if err := sysdictitem.DictIDValidator(v); err != nil {
			return &ValidationError{Name: "dict_id", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.dict_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sdiu *SysDictItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictItemUpdate {
	sdiu.modifiers = append(sdiu.modifiers, modifiers...)
	return sdiu
}

func (sdiu *SysDictItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sdiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysdictitem.Table, sysdictitem.Columns, sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString))
	if ps := sdiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdiu.mutation.IsDel(); ok {
		_spec.SetField(sysdictitem.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sdiu.mutation.Memo(); ok {
		_spec.SetField(sysdictitem.FieldMemo, field.TypeString, value)
	}
	if sdiu.mutation.MemoCleared() {
		_spec.ClearField(sysdictitem.FieldMemo, field.TypeString)
	}
	if value, ok := sdiu.mutation.Sort(); ok {
		_spec.SetField(sysdictitem.FieldSort, field.TypeInt32, value)
	}
	if value, ok := sdiu.mutation.AddedSort(); ok {
		_spec.AddField(sysdictitem.FieldSort, field.TypeInt32, value)
	}
	if sdiu.mutation.CreatedAtCleared() {
		_spec.ClearField(sysdictitem.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sdiu.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdictitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if sdiu.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysdictitem.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sdiu.mutation.DeletedAt(); ok {
		_spec.SetField(sysdictitem.FieldDeletedAt, field.TypeTime, value)
	}
	if sdiu.mutation.DeletedAtCleared() {
		_spec.ClearField(sysdictitem.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sdiu.mutation.IsActive(); ok {
		_spec.SetField(sysdictitem.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sdiu.mutation.Label(); ok {
		_spec.SetField(sysdictitem.FieldLabel, field.TypeString, value)
	}
	if value, ok := sdiu.mutation.Value(); ok {
		_spec.SetField(sysdictitem.FieldValue, field.TypeInt, value)
	}
	if value, ok := sdiu.mutation.AddedValue(); ok {
		_spec.AddField(sysdictitem.FieldValue, field.TypeInt, value)
	}
	if sdiu.mutation.DictCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.DictTable,
			Columns: []string{sysdictitem.DictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdiu.mutation.DictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.DictTable,
			Columns: []string{sysdictitem.DictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sdiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sdiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sdiu.mutation.done = true
	return n, nil
}

// SysDictItemUpdateOne is the builder for updating a single SysDictItem entity.
type SysDictItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysDictItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (sdiuo *SysDictItemUpdateOne) SetIsDel(b bool) *SysDictItemUpdateOne {
	sdiuo.mutation.SetIsDel(b)
	return sdiuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableIsDel(b *bool) *SysDictItemUpdateOne {
	if b != nil {
		sdiuo.SetIsDel(*b)
	}
	return sdiuo
}

// SetMemo sets the "memo" field.
func (sdiuo *SysDictItemUpdateOne) SetMemo(s string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetMemo(s)
	return sdiuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableMemo(s *string) *SysDictItemUpdateOne {
	if s != nil {
		sdiuo.SetMemo(*s)
	}
	return sdiuo
}

// ClearMemo clears the value of the "memo" field.
func (sdiuo *SysDictItemUpdateOne) ClearMemo() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearMemo()
	return sdiuo
}

// SetSort sets the "sort" field.
func (sdiuo *SysDictItemUpdateOne) SetSort(i int32) *SysDictItemUpdateOne {
	sdiuo.mutation.ResetSort()
	sdiuo.mutation.SetSort(i)
	return sdiuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableSort(i *int32) *SysDictItemUpdateOne {
	if i != nil {
		sdiuo.SetSort(*i)
	}
	return sdiuo
}

// AddSort adds i to the "sort" field.
func (sdiuo *SysDictItemUpdateOne) AddSort(i int32) *SysDictItemUpdateOne {
	sdiuo.mutation.AddSort(i)
	return sdiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sdiuo *SysDictItemUpdateOne) SetUpdatedAt(t time.Time) *SysDictItemUpdateOne {
	sdiuo.mutation.SetUpdatedAt(t)
	return sdiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sdiuo *SysDictItemUpdateOne) ClearUpdatedAt() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearUpdatedAt()
	return sdiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (sdiuo *SysDictItemUpdateOne) SetDeletedAt(t time.Time) *SysDictItemUpdateOne {
	sdiuo.mutation.SetDeletedAt(t)
	return sdiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictItemUpdateOne {
	if t != nil {
		sdiuo.SetDeletedAt(*t)
	}
	return sdiuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdiuo *SysDictItemUpdateOne) ClearDeletedAt() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearDeletedAt()
	return sdiuo
}

// SetIsActive sets the "is_active" field.
func (sdiuo *SysDictItemUpdateOne) SetIsActive(b bool) *SysDictItemUpdateOne {
	sdiuo.mutation.SetIsActive(b)
	return sdiuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableIsActive(b *bool) *SysDictItemUpdateOne {
	if b != nil {
		sdiuo.SetIsActive(*b)
	}
	return sdiuo
}

// SetLabel sets the "label" field.
func (sdiuo *SysDictItemUpdateOne) SetLabel(s string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetLabel(s)
	return sdiuo
}

// SetValue sets the "value" field.
func (sdiuo *SysDictItemUpdateOne) SetValue(i int) *SysDictItemUpdateOne {
	sdiuo.mutation.ResetValue()
	sdiuo.mutation.SetValue(i)
	return sdiuo
}

// AddValue adds i to the "value" field.
func (sdiuo *SysDictItemUpdateOne) AddValue(i int) *SysDictItemUpdateOne {
	sdiuo.mutation.AddValue(i)
	return sdiuo
}

// SetDictID sets the "dict_id" field.
func (sdiuo *SysDictItemUpdateOne) SetDictID(s string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetDictID(s)
	return sdiuo
}

// SetNillableDictID sets the "dict_id" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableDictID(s *string) *SysDictItemUpdateOne {
	if s != nil {
		sdiuo.SetDictID(*s)
	}
	return sdiuo
}

// ClearDictID clears the value of the "dict_id" field.
func (sdiuo *SysDictItemUpdateOne) ClearDictID() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearDictID()
	return sdiuo
}

// SetDict sets the "dict" edge to the SysDict entity.
func (sdiuo *SysDictItemUpdateOne) SetDict(s *SysDict) *SysDictItemUpdateOne {
	return sdiuo.SetDictID(s.ID)
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdiuo *SysDictItemUpdateOne) Mutation() *SysDictItemMutation {
	return sdiuo.mutation
}

// ClearDict clears the "dict" edge to the SysDict entity.
func (sdiuo *SysDictItemUpdateOne) ClearDict() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearDict()
	return sdiuo
}

// Where appends a list predicates to the SysDictItemUpdate builder.
func (sdiuo *SysDictItemUpdateOne) Where(ps ...predicate.SysDictItem) *SysDictItemUpdateOne {
	sdiuo.mutation.Where(ps...)
	return sdiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdiuo *SysDictItemUpdateOne) Select(field string, fields ...string) *SysDictItemUpdateOne {
	sdiuo.fields = append([]string{field}, fields...)
	return sdiuo
}

// Save executes the query and returns the updated SysDictItem entity.
func (sdiuo *SysDictItemUpdateOne) Save(ctx context.Context) (*SysDictItem, error) {
	sdiuo.defaults()
	return withHooks(ctx, sdiuo.sqlSave, sdiuo.mutation, sdiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdiuo *SysDictItemUpdateOne) SaveX(ctx context.Context) *SysDictItem {
	node, err := sdiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdiuo *SysDictItemUpdateOne) Exec(ctx context.Context) error {
	_, err := sdiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdiuo *SysDictItemUpdateOne) ExecX(ctx context.Context) {
	if err := sdiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdiuo *SysDictItemUpdateOne) defaults() {
	if _, ok := sdiuo.mutation.UpdatedAt(); !ok && !sdiuo.mutation.UpdatedAtCleared() {
		v := sysdictitem.UpdateDefaultUpdatedAt()
		sdiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdiuo *SysDictItemUpdateOne) check() error {
	if v, ok := sdiuo.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.memo": %w`, err)}
		}
	}
	if v, ok := sdiuo.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.label": %w`, err)}
		}
	}
	if v, ok := sdiuo.mutation.DictID(); ok {
		if err := sysdictitem.DictIDValidator(v); err != nil {
			return &ValidationError{Name: "dict_id", err: fmt.Errorf(`mainent: validator failed for field "SysDictItem.dict_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sdiuo *SysDictItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictItemUpdateOne {
	sdiuo.modifiers = append(sdiuo.modifiers, modifiers...)
	return sdiuo
}

func (sdiuo *SysDictItemUpdateOne) sqlSave(ctx context.Context) (_node *SysDictItem, err error) {
	if err := sdiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysdictitem.Table, sysdictitem.Columns, sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString))
	id, ok := sdiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "SysDictItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sdiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictitem.FieldID)
		for _, f := range fields {
			if !sysdictitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != sysdictitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdiuo.mutation.IsDel(); ok {
		_spec.SetField(sysdictitem.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sdiuo.mutation.Memo(); ok {
		_spec.SetField(sysdictitem.FieldMemo, field.TypeString, value)
	}
	if sdiuo.mutation.MemoCleared() {
		_spec.ClearField(sysdictitem.FieldMemo, field.TypeString)
	}
	if value, ok := sdiuo.mutation.Sort(); ok {
		_spec.SetField(sysdictitem.FieldSort, field.TypeInt32, value)
	}
	if value, ok := sdiuo.mutation.AddedSort(); ok {
		_spec.AddField(sysdictitem.FieldSort, field.TypeInt32, value)
	}
	if sdiuo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysdictitem.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sdiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdictitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if sdiuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysdictitem.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sdiuo.mutation.DeletedAt(); ok {
		_spec.SetField(sysdictitem.FieldDeletedAt, field.TypeTime, value)
	}
	if sdiuo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysdictitem.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sdiuo.mutation.IsActive(); ok {
		_spec.SetField(sysdictitem.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sdiuo.mutation.Label(); ok {
		_spec.SetField(sysdictitem.FieldLabel, field.TypeString, value)
	}
	if value, ok := sdiuo.mutation.Value(); ok {
		_spec.SetField(sysdictitem.FieldValue, field.TypeInt, value)
	}
	if value, ok := sdiuo.mutation.AddedValue(); ok {
		_spec.AddField(sysdictitem.FieldValue, field.TypeInt, value)
	}
	if sdiuo.mutation.DictCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.DictTable,
			Columns: []string{sysdictitem.DictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdiuo.mutation.DictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.DictTable,
			Columns: []string{sysdictitem.DictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sdiuo.modifiers...)
	_node = &SysDictItem{config: sdiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sdiuo.mutation.done = true
	return _node, nil
}