// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemDelete is the builder for deleting a SysDictItem entity.
type SysDictItemDelete struct {
	config
	hooks    []Hook
	mutation *SysDictItemMutation
}

// Where appends a list predicates to the SysDictItemDelete builder.
func (sdid *SysDictItemDelete) Where(ps ...predicate.SysDictItem) *SysDictItemDelete {
	sdid.mutation.Where(ps...)
	return sdid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdid *SysDictItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdid.hooks) == 0 {
		affected, err = sdid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdid.mutation = mutation
			affected, err = sdid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdid.hooks) - 1; i >= 0; i-- {
			if sdid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdid *SysDictItemDelete) ExecX(ctx context.Context) int {
	n, err := sdid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdid *SysDictItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysdictitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdictitem.FieldID,
			},
		},
	}
	if ps := sdid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SysDictItemDeleteOne is the builder for deleting a single SysDictItem entity.
type SysDictItemDeleteOne struct {
	sdid *SysDictItemDelete
}

// Exec executes the deletion query.
func (sdido *SysDictItemDeleteOne) Exec(ctx context.Context) error {
	n, err := sdido.sdid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysdictitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdido *SysDictItemDeleteOne) ExecX(ctx context.Context) {
	sdido.sdid.ExecX(ctx)
}
