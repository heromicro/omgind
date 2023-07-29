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
	"github.com/heromicro/omgind/internal/scheme/enumtipe"
)

// SysDictUpdate is the builder for updating SysDict entities.
type SysDictUpdate struct {
	config
	hooks     []Hook
	mutation  *SysDictMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysDictUpdate builder.
func (sdu *SysDictUpdate) Where(ps ...predicate.SysDict) *SysDictUpdate {
	sdu.mutation.Where(ps...)
	return sdu
}

// SetIsDel sets the "is_del" field.
func (sdu *SysDictUpdate) SetIsDel(b bool) *SysDictUpdate {
	sdu.mutation.SetIsDel(b)
	return sdu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableIsDel(b *bool) *SysDictUpdate {
	if b != nil {
		sdu.SetIsDel(*b)
	}
	return sdu
}

// SetMemo sets the "memo" field.
func (sdu *SysDictUpdate) SetMemo(s string) *SysDictUpdate {
	sdu.mutation.SetMemo(s)
	return sdu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableMemo(s *string) *SysDictUpdate {
	if s != nil {
		sdu.SetMemo(*s)
	}
	return sdu
}

// ClearMemo clears the value of the "memo" field.
func (sdu *SysDictUpdate) ClearMemo() *SysDictUpdate {
	sdu.mutation.ClearMemo()
	return sdu
}

// SetSort sets the "sort" field.
func (sdu *SysDictUpdate) SetSort(i int32) *SysDictUpdate {
	sdu.mutation.ResetSort()
	sdu.mutation.SetSort(i)
	return sdu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableSort(i *int32) *SysDictUpdate {
	if i != nil {
		sdu.SetSort(*i)
	}
	return sdu
}

// AddSort adds i to the "sort" field.
func (sdu *SysDictUpdate) AddSort(i int32) *SysDictUpdate {
	sdu.mutation.AddSort(i)
	return sdu
}

// SetUpdatedAt sets the "updated_at" field.
func (sdu *SysDictUpdate) SetUpdatedAt(t time.Time) *SysDictUpdate {
	sdu.mutation.SetUpdatedAt(t)
	return sdu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sdu *SysDictUpdate) ClearUpdatedAt() *SysDictUpdate {
	sdu.mutation.ClearUpdatedAt()
	return sdu
}

// SetDeletedAt sets the "deleted_at" field.
func (sdu *SysDictUpdate) SetDeletedAt(t time.Time) *SysDictUpdate {
	sdu.mutation.SetDeletedAt(t)
	return sdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableDeletedAt(t *time.Time) *SysDictUpdate {
	if t != nil {
		sdu.SetDeletedAt(*t)
	}
	return sdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdu *SysDictUpdate) ClearDeletedAt() *SysDictUpdate {
	sdu.mutation.ClearDeletedAt()
	return sdu
}

// SetIsActive sets the "is_active" field.
func (sdu *SysDictUpdate) SetIsActive(b bool) *SysDictUpdate {
	sdu.mutation.SetIsActive(b)
	return sdu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableIsActive(b *bool) *SysDictUpdate {
	if b != nil {
		sdu.SetIsActive(*b)
	}
	return sdu
}

// SetNameCn sets the "name_cn" field.
func (sdu *SysDictUpdate) SetNameCn(s string) *SysDictUpdate {
	sdu.mutation.SetNameCn(s)
	return sdu
}

// SetNameEn sets the "name_en" field.
func (sdu *SysDictUpdate) SetNameEn(s string) *SysDictUpdate {
	sdu.mutation.SetNameEn(s)
	return sdu
}

// SetValTipe sets the "val_tipe" field.
func (sdu *SysDictUpdate) SetValTipe(evt enumtipe.DictValueTipe) *SysDictUpdate {
	sdu.mutation.SetValTipe(evt)
	return sdu
}

// SetNillableValTipe sets the "val_tipe" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableValTipe(evt *enumtipe.DictValueTipe) *SysDictUpdate {
	if evt != nil {
		sdu.SetValTipe(*evt)
	}
	return sdu
}

// AddItemIDs adds the "items" edge to the SysDictItem entity by IDs.
func (sdu *SysDictUpdate) AddItemIDs(ids ...string) *SysDictUpdate {
	sdu.mutation.AddItemIDs(ids...)
	return sdu
}

// AddItems adds the "items" edges to the SysDictItem entity.
func (sdu *SysDictUpdate) AddItems(s ...*SysDictItem) *SysDictUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdu.AddItemIDs(ids...)
}

// Mutation returns the SysDictMutation object of the builder.
func (sdu *SysDictUpdate) Mutation() *SysDictMutation {
	return sdu.mutation
}

// ClearItems clears all "items" edges to the SysDictItem entity.
func (sdu *SysDictUpdate) ClearItems() *SysDictUpdate {
	sdu.mutation.ClearItems()
	return sdu
}

// RemoveItemIDs removes the "items" edge to SysDictItem entities by IDs.
func (sdu *SysDictUpdate) RemoveItemIDs(ids ...string) *SysDictUpdate {
	sdu.mutation.RemoveItemIDs(ids...)
	return sdu
}

// RemoveItems removes "items" edges to SysDictItem entities.
func (sdu *SysDictUpdate) RemoveItems(s ...*SysDictItem) *SysDictUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *SysDictUpdate) Save(ctx context.Context) (int, error) {
	sdu.defaults()
	return withHooks(ctx, sdu.sqlSave, sdu.mutation, sdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *SysDictUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *SysDictUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *SysDictUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdu *SysDictUpdate) defaults() {
	if _, ok := sdu.mutation.UpdatedAt(); !ok && !sdu.mutation.UpdatedAtCleared() {
		v := sysdict.UpdateDefaultUpdatedAt()
		sdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdu *SysDictUpdate) check() error {
	if v, ok := sdu.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysDict.memo": %w`, err)}
		}
	}
	if v, ok := sdu.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf(`mainent: validator failed for field "SysDict.name_cn": %w`, err)}
		}
	}
	if v, ok := sdu.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`mainent: validator failed for field "SysDict.name_en": %w`, err)}
		}
	}
	if v, ok := sdu.mutation.ValTipe(); ok {
		if err := sysdict.ValTipeValidator(v); err != nil {
			return &ValidationError{Name: "val_tipe", err: fmt.Errorf(`mainent: validator failed for field "SysDict.val_tipe": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sdu *SysDictUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictUpdate {
	sdu.modifiers = append(sdu.modifiers, modifiers...)
	return sdu
}

func (sdu *SysDictUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysdict.Table, sysdict.Columns, sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString))
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.IsDel(); ok {
		_spec.SetField(sysdict.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sdu.mutation.Memo(); ok {
		_spec.SetField(sysdict.FieldMemo, field.TypeString, value)
	}
	if sdu.mutation.MemoCleared() {
		_spec.ClearField(sysdict.FieldMemo, field.TypeString)
	}
	if value, ok := sdu.mutation.Sort(); ok {
		_spec.SetField(sysdict.FieldSort, field.TypeInt32, value)
	}
	if value, ok := sdu.mutation.AddedSort(); ok {
		_spec.AddField(sysdict.FieldSort, field.TypeInt32, value)
	}
	if sdu.mutation.CreatedAtCleared() {
		_spec.ClearField(sysdict.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sdu.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdict.FieldUpdatedAt, field.TypeTime, value)
	}
	if sdu.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysdict.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sdu.mutation.DeletedAt(); ok {
		_spec.SetField(sysdict.FieldDeletedAt, field.TypeTime, value)
	}
	if sdu.mutation.DeletedAtCleared() {
		_spec.ClearField(sysdict.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sdu.mutation.IsActive(); ok {
		_spec.SetField(sysdict.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sdu.mutation.NameCn(); ok {
		_spec.SetField(sysdict.FieldNameCn, field.TypeString, value)
	}
	if value, ok := sdu.mutation.NameEn(); ok {
		_spec.SetField(sysdict.FieldNameEn, field.TypeString, value)
	}
	if value, ok := sdu.mutation.ValTipe(); ok {
		_spec.SetField(sysdict.FieldValTipe, field.TypeEnum, value)
	}
	if sdu.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !sdu.mutation.ItemsCleared() {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.ItemsIDs(); len(nodes) > 0 {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sdu.mutation.done = true
	return n, nil
}

// SysDictUpdateOne is the builder for updating a single SysDict entity.
type SysDictUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysDictMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsDel sets the "is_del" field.
func (sduo *SysDictUpdateOne) SetIsDel(b bool) *SysDictUpdateOne {
	sduo.mutation.SetIsDel(b)
	return sduo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableIsDel(b *bool) *SysDictUpdateOne {
	if b != nil {
		sduo.SetIsDel(*b)
	}
	return sduo
}

// SetMemo sets the "memo" field.
func (sduo *SysDictUpdateOne) SetMemo(s string) *SysDictUpdateOne {
	sduo.mutation.SetMemo(s)
	return sduo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableMemo(s *string) *SysDictUpdateOne {
	if s != nil {
		sduo.SetMemo(*s)
	}
	return sduo
}

// ClearMemo clears the value of the "memo" field.
func (sduo *SysDictUpdateOne) ClearMemo() *SysDictUpdateOne {
	sduo.mutation.ClearMemo()
	return sduo
}

// SetSort sets the "sort" field.
func (sduo *SysDictUpdateOne) SetSort(i int32) *SysDictUpdateOne {
	sduo.mutation.ResetSort()
	sduo.mutation.SetSort(i)
	return sduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableSort(i *int32) *SysDictUpdateOne {
	if i != nil {
		sduo.SetSort(*i)
	}
	return sduo
}

// AddSort adds i to the "sort" field.
func (sduo *SysDictUpdateOne) AddSort(i int32) *SysDictUpdateOne {
	sduo.mutation.AddSort(i)
	return sduo
}

// SetUpdatedAt sets the "updated_at" field.
func (sduo *SysDictUpdateOne) SetUpdatedAt(t time.Time) *SysDictUpdateOne {
	sduo.mutation.SetUpdatedAt(t)
	return sduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sduo *SysDictUpdateOne) ClearUpdatedAt() *SysDictUpdateOne {
	sduo.mutation.ClearUpdatedAt()
	return sduo
}

// SetDeletedAt sets the "deleted_at" field.
func (sduo *SysDictUpdateOne) SetDeletedAt(t time.Time) *SysDictUpdateOne {
	sduo.mutation.SetDeletedAt(t)
	return sduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictUpdateOne {
	if t != nil {
		sduo.SetDeletedAt(*t)
	}
	return sduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sduo *SysDictUpdateOne) ClearDeletedAt() *SysDictUpdateOne {
	sduo.mutation.ClearDeletedAt()
	return sduo
}

// SetIsActive sets the "is_active" field.
func (sduo *SysDictUpdateOne) SetIsActive(b bool) *SysDictUpdateOne {
	sduo.mutation.SetIsActive(b)
	return sduo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableIsActive(b *bool) *SysDictUpdateOne {
	if b != nil {
		sduo.SetIsActive(*b)
	}
	return sduo
}

// SetNameCn sets the "name_cn" field.
func (sduo *SysDictUpdateOne) SetNameCn(s string) *SysDictUpdateOne {
	sduo.mutation.SetNameCn(s)
	return sduo
}

// SetNameEn sets the "name_en" field.
func (sduo *SysDictUpdateOne) SetNameEn(s string) *SysDictUpdateOne {
	sduo.mutation.SetNameEn(s)
	return sduo
}

// SetValTipe sets the "val_tipe" field.
func (sduo *SysDictUpdateOne) SetValTipe(evt enumtipe.DictValueTipe) *SysDictUpdateOne {
	sduo.mutation.SetValTipe(evt)
	return sduo
}

// SetNillableValTipe sets the "val_tipe" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableValTipe(evt *enumtipe.DictValueTipe) *SysDictUpdateOne {
	if evt != nil {
		sduo.SetValTipe(*evt)
	}
	return sduo
}

// AddItemIDs adds the "items" edge to the SysDictItem entity by IDs.
func (sduo *SysDictUpdateOne) AddItemIDs(ids ...string) *SysDictUpdateOne {
	sduo.mutation.AddItemIDs(ids...)
	return sduo
}

// AddItems adds the "items" edges to the SysDictItem entity.
func (sduo *SysDictUpdateOne) AddItems(s ...*SysDictItem) *SysDictUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sduo.AddItemIDs(ids...)
}

// Mutation returns the SysDictMutation object of the builder.
func (sduo *SysDictUpdateOne) Mutation() *SysDictMutation {
	return sduo.mutation
}

// ClearItems clears all "items" edges to the SysDictItem entity.
func (sduo *SysDictUpdateOne) ClearItems() *SysDictUpdateOne {
	sduo.mutation.ClearItems()
	return sduo
}

// RemoveItemIDs removes the "items" edge to SysDictItem entities by IDs.
func (sduo *SysDictUpdateOne) RemoveItemIDs(ids ...string) *SysDictUpdateOne {
	sduo.mutation.RemoveItemIDs(ids...)
	return sduo
}

// RemoveItems removes "items" edges to SysDictItem entities.
func (sduo *SysDictUpdateOne) RemoveItems(s ...*SysDictItem) *SysDictUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sduo.RemoveItemIDs(ids...)
}

// Where appends a list predicates to the SysDictUpdate builder.
func (sduo *SysDictUpdateOne) Where(ps ...predicate.SysDict) *SysDictUpdateOne {
	sduo.mutation.Where(ps...)
	return sduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *SysDictUpdateOne) Select(field string, fields ...string) *SysDictUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated SysDict entity.
func (sduo *SysDictUpdateOne) Save(ctx context.Context) (*SysDict, error) {
	sduo.defaults()
	return withHooks(ctx, sduo.sqlSave, sduo.mutation, sduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *SysDictUpdateOne) SaveX(ctx context.Context) *SysDict {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *SysDictUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *SysDictUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sduo *SysDictUpdateOne) defaults() {
	if _, ok := sduo.mutation.UpdatedAt(); !ok && !sduo.mutation.UpdatedAtCleared() {
		v := sysdict.UpdateDefaultUpdatedAt()
		sduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sduo *SysDictUpdateOne) check() error {
	if v, ok := sduo.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysDict.memo": %w`, err)}
		}
	}
	if v, ok := sduo.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf(`mainent: validator failed for field "SysDict.name_cn": %w`, err)}
		}
	}
	if v, ok := sduo.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`mainent: validator failed for field "SysDict.name_en": %w`, err)}
		}
	}
	if v, ok := sduo.mutation.ValTipe(); ok {
		if err := sysdict.ValTipeValidator(v); err != nil {
			return &ValidationError{Name: "val_tipe", err: fmt.Errorf(`mainent: validator failed for field "SysDict.val_tipe": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sduo *SysDictUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysDictUpdateOne {
	sduo.modifiers = append(sduo.modifiers, modifiers...)
	return sduo
}

func (sduo *SysDictUpdateOne) sqlSave(ctx context.Context) (_node *SysDict, err error) {
	if err := sduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysdict.Table, sysdict.Columns, sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString))
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "SysDict.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdict.FieldID)
		for _, f := range fields {
			if !sysdict.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != sysdict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.IsDel(); ok {
		_spec.SetField(sysdict.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := sduo.mutation.Memo(); ok {
		_spec.SetField(sysdict.FieldMemo, field.TypeString, value)
	}
	if sduo.mutation.MemoCleared() {
		_spec.ClearField(sysdict.FieldMemo, field.TypeString)
	}
	if value, ok := sduo.mutation.Sort(); ok {
		_spec.SetField(sysdict.FieldSort, field.TypeInt32, value)
	}
	if value, ok := sduo.mutation.AddedSort(); ok {
		_spec.AddField(sysdict.FieldSort, field.TypeInt32, value)
	}
	if sduo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysdict.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sduo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdict.FieldUpdatedAt, field.TypeTime, value)
	}
	if sduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysdict.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sduo.mutation.DeletedAt(); ok {
		_spec.SetField(sysdict.FieldDeletedAt, field.TypeTime, value)
	}
	if sduo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysdict.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := sduo.mutation.IsActive(); ok {
		_spec.SetField(sysdict.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := sduo.mutation.NameCn(); ok {
		_spec.SetField(sysdict.FieldNameCn, field.TypeString, value)
	}
	if value, ok := sduo.mutation.NameEn(); ok {
		_spec.SetField(sysdict.FieldNameEn, field.TypeString, value)
	}
	if value, ok := sduo.mutation.ValTipe(); ok {
		_spec.SetField(sysdict.FieldValTipe, field.TypeEnum, value)
	}
	if sduo.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !sduo.mutation.ItemsCleared() {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.ItemsIDs(); len(nodes) > 0 {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sduo.modifiers...)
	_node = &SysDict{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sduo.mutation.done = true
	return _node, nil
}
