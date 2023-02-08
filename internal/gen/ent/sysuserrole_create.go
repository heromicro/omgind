// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRoleCreate is the builder for creating a SysUserRole entity.
type SysUserRoleCreate struct {
	config
	mutation *SysUserRoleMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (surc *SysUserRoleCreate) SetIsDel(b bool) *SysUserRoleCreate {
	surc.mutation.SetIsDel(b)
	return surc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableIsDel(b *bool) *SysUserRoleCreate {
	if b != nil {
		surc.SetIsDel(*b)
	}
	return surc
}

// SetCreatedAt sets the "created_at" field.
func (surc *SysUserRoleCreate) SetCreatedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetCreatedAt(t)
	return surc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableCreatedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetCreatedAt(*t)
	}
	return surc
}

// SetUpdatedAt sets the "updated_at" field.
func (surc *SysUserRoleCreate) SetUpdatedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetUpdatedAt(t)
	return surc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableUpdatedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetUpdatedAt(*t)
	}
	return surc
}

// SetDeletedAt sets the "deleted_at" field.
func (surc *SysUserRoleCreate) SetDeletedAt(t time.Time) *SysUserRoleCreate {
	surc.mutation.SetDeletedAt(t)
	return surc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableDeletedAt(t *time.Time) *SysUserRoleCreate {
	if t != nil {
		surc.SetDeletedAt(*t)
	}
	return surc
}

// SetUserID sets the "user_id" field.
func (surc *SysUserRoleCreate) SetUserID(s string) *SysUserRoleCreate {
	surc.mutation.SetUserID(s)
	return surc
}

// SetRoleID sets the "role_id" field.
func (surc *SysUserRoleCreate) SetRoleID(s string) *SysUserRoleCreate {
	surc.mutation.SetRoleID(s)
	return surc
}

// SetID sets the "id" field.
func (surc *SysUserRoleCreate) SetID(s string) *SysUserRoleCreate {
	surc.mutation.SetID(s)
	return surc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (surc *SysUserRoleCreate) SetNillableID(s *string) *SysUserRoleCreate {
	if s != nil {
		surc.SetID(*s)
	}
	return surc
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (surc *SysUserRoleCreate) Mutation() *SysUserRoleMutation {
	return surc.mutation
}

// Save creates the SysUserRole in the database.
func (surc *SysUserRoleCreate) Save(ctx context.Context) (*SysUserRole, error) {
	var (
		err  error
		node *SysUserRole
	)
	surc.defaults()
	if len(surc.hooks) == 0 {
		if err = surc.check(); err != nil {
			return nil, err
		}
		node, err = surc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = surc.check(); err != nil {
				return nil, err
			}
			surc.mutation = mutation
			node, err = surc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(surc.hooks) - 1; i >= 0; i-- {
			mut = surc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, surc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (surc *SysUserRoleCreate) SaveX(ctx context.Context) *SysUserRole {
	v, err := surc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (surc *SysUserRoleCreate) defaults() {
	if _, ok := surc.mutation.IsDel(); !ok {
		v := sysuserrole.DefaultIsDel
		surc.mutation.SetIsDel(v)
	}
	if _, ok := surc.mutation.CreatedAt(); !ok {
		v := sysuserrole.DefaultCreatedAt()
		surc.mutation.SetCreatedAt(v)
	}
	if _, ok := surc.mutation.UpdatedAt(); !ok {
		v := sysuserrole.DefaultUpdatedAt()
		surc.mutation.SetUpdatedAt(v)
	}
	if _, ok := surc.mutation.ID(); !ok {
		v := sysuserrole.DefaultID()
		surc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (surc *SysUserRoleCreate) check() error {
	if _, ok := surc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := surc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := surc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := surc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if v, ok := surc.mutation.UserID(); ok {
		if err := sysuserrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf("ent: validator failed for field \"user_id\": %w", err)}
		}
	}
	if _, ok := surc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New("ent: missing required field \"role_id\"")}
	}
	if v, ok := surc.mutation.RoleID(); ok {
		if err := sysuserrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf("ent: validator failed for field \"role_id\": %w", err)}
		}
	}
	if v, ok := surc.mutation.ID(); ok {
		if err := sysuserrole.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (surc *SysUserRoleCreate) sqlSave(ctx context.Context) (*SysUserRole, error) {
	_node, _spec := surc.createSpec()
	if err := sqlgraph.CreateNode(ctx, surc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (surc *SysUserRoleCreate) createSpec() (*SysUserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUserRole{config: surc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysuserrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysuserrole.FieldID,
			},
		}
	)
	if id, ok := surc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := surc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysuserrole.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := surc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuserrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := surc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuserrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := surc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuserrole.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := surc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuserrole.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := surc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuserrole.FieldRoleID,
		})
		_node.RoleID = value
	}
	return _node, _spec
}

// SysUserRoleCreateBulk is the builder for creating many SysUserRole entities in bulk.
type SysUserRoleCreateBulk struct {
	config
	builders []*SysUserRoleCreate
}

// Save creates the SysUserRole entities in the database.
func (surcb *SysUserRoleCreateBulk) Save(ctx context.Context) ([]*SysUserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(surcb.builders))
	nodes := make([]*SysUserRole, len(surcb.builders))
	mutators := make([]Mutator, len(surcb.builders))
	for i := range surcb.builders {
		func(i int, root context.Context) {
			builder := surcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, surcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, surcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, surcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (surcb *SysUserRoleCreateBulk) SaveX(ctx context.Context) []*SysUserRole {
	v, err := surcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
