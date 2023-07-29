// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuaction"
)

// SysMenuActionDelete is the builder for deleting a SysMenuAction entity.
type SysMenuActionDelete struct {
	config
	hooks    []Hook
	mutation *SysMenuActionMutation
}

// Where appends a list predicates to the SysMenuActionDelete builder.
func (smad *SysMenuActionDelete) Where(ps ...predicate.SysMenuAction) *SysMenuActionDelete {
	smad.mutation.Where(ps...)
	return smad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smad *SysMenuActionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, smad.sqlExec, smad.mutation, smad.hooks)
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
	_spec := sqlgraph.NewDeleteSpec(sysmenuaction.Table, sqlgraph.NewFieldSpec(sysmenuaction.FieldID, field.TypeString))
	if ps := smad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, smad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	smad.mutation.done = true
	return affected, err
}

// SysMenuActionDeleteOne is the builder for deleting a single SysMenuAction entity.
type SysMenuActionDeleteOne struct {
	smad *SysMenuActionDelete
}

// Where appends a list predicates to the SysMenuActionDelete builder.
func (smado *SysMenuActionDeleteOne) Where(ps ...predicate.SysMenuAction) *SysMenuActionDeleteOne {
	smado.smad.mutation.Where(ps...)
	return smado
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
	if err := smado.Exec(ctx); err != nil {
		panic(err)
	}
}