// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenu"
)

// SysMenuDelete is the builder for deleting a SysMenu entity.
type SysMenuDelete struct {
	config
	hooks    []Hook
	mutation *SysMenuMutation
}

// Where appends a list predicates to the SysMenuDelete builder.
func (smd *SysMenuDelete) Where(ps ...predicate.SysMenu) *SysMenuDelete {
	smd.mutation.Where(ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *SysMenuDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smd.hooks) == 0 {
		affected, err = smd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smd.mutation = mutation
			affected, err = smd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smd.hooks) - 1; i >= 0; i-- {
			if smd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *SysMenuDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *SysMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenu.FieldID,
			},
		},
	}
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SysMenuDeleteOne is the builder for deleting a single SysMenu entity.
type SysMenuDeleteOne struct {
	smd *SysMenuDelete
}

// Exec executes the deletion query.
func (smdo *SysMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysmenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *SysMenuDeleteOne) ExecX(ctx context.Context) {
	smdo.smd.ExecX(ctx)
}
