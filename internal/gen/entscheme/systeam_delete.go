// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/entscheme/predicate"
	"github.com/heromicro/omgind/internal/gen/entscheme/systeam"
)

// SysTeamDelete is the builder for deleting a SysTeam entity.
type SysTeamDelete struct {
	config
	hooks    []Hook
	mutation *SysTeamMutation
}

// Where appends a list predicates to the SysTeamDelete builder.
func (std *SysTeamDelete) Where(ps ...predicate.SysTeam) *SysTeamDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *SysTeamDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *SysTeamDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *SysTeamDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(systeam.Table, sqlgraph.NewFieldSpec(systeam.FieldID, field.TypeString))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// SysTeamDeleteOne is the builder for deleting a single SysTeam entity.
type SysTeamDeleteOne struct {
	std *SysTeamDelete
}

// Where appends a list predicates to the SysTeamDelete builder.
func (stdo *SysTeamDeleteOne) Where(ps ...predicate.SysTeam) *SysTeamDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *SysTeamDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{systeam.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *SysTeamDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}
