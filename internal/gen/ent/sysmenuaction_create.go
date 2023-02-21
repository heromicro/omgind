// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuActionCreate is the builder for creating a SysMenuAction entity.
type SysMenuActionCreate struct {
	config
	mutation *SysMenuActionMutation
	hooks    []Hook
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
	var (
		err  error
		node *SysMenuAction
	)
	smac.defaults()
	if len(smac.hooks) == 0 {
		if err = smac.check(); err != nil {
			return nil, err
		}
		node, err = smac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smac.check(); err != nil {
				return nil, err
			}
			smac.mutation = mutation
			if node, err = smac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(smac.hooks) - 1; i >= 0; i-- {
			if smac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, smac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysMenuAction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysMenuActionMutation", v)
		}
		node = nv
	}
	return node, err
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
	if _, ok := smac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysMenuAction.created_at"`)}
	}
	if _, ok := smac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysMenuAction.updated_at"`)}
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
	return _node, nil
}

func (smac *SysMenuActionCreate) createSpec() (*SysMenuAction, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuAction{config: smac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenuaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenuaction.FieldID,
			},
		}
	)
	if id, ok := smac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smac.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenuaction.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := smac.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenuaction.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := smac.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenuaction.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := smac.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldMemo,
		})
		_node.Memo = &value
	}
	if value, ok := smac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := smac.mutation.MenuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldMenuID,
		})
		_node.MenuID = value
	}
	if value, ok := smac.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := smac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// SysMenuActionCreateBulk is the builder for creating many SysMenuAction entities in bulk.
type SysMenuActionCreateBulk struct {
	config
	builders []*SysMenuActionCreate
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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
