// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuActionDelete is the builder for deleting a SysMenuAction entity.
type SysMenuActionDelete struct {
	config
	hooks    []Hook
	mutation *SysMenuActionMutation
}

// Where adds a new predicate to the SysMenuActionDelete builder.
func (smad *SysMenuActionDelete) Where(ps ...predicate.SysMenuAction) *SysMenuActionDelete {
	smad.mutation.predicates = append(smad.mutation.predicates, ps...)
	return smad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smad *SysMenuActionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smad.hooks) == 0 {
		affected, err = smad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smad.mutation = mutation
			affected, err = smad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smad.hooks) - 1; i >= 0; i-- {
			mut = smad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smad *SysMenuActionDelete) ExecX(ctx context.Context) int {
	n, err := smad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smad *SysMenuActionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysmenuaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenuaction.FieldID,
			},
		},
	}
	if ps := smad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, smad.driver, _spec)
}

// SysMenuActionDeleteOne is the builder for deleting a single SysMenuAction entity.
type SysMenuActionDeleteOne struct {
	smad *SysMenuActionDelete
}

// Exec executes the deletion query.
func (smado *SysMenuActionDeleteOne) Exec(ctx context.Context) error {
	n, err := smado.smad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysmenuaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smado *SysMenuActionDeleteOne) ExecX(ctx context.Context) {
	smado.smad.ExecX(ctx)
}
