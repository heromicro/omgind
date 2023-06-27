// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/entscheme/predicate"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysrole"
)

// SysRoleDelete is the builder for deleting a SysRole entity.
type SysRoleDelete struct {
	config
	hooks    []Hook
	mutation *SysRoleMutation
}

// Where appends a list predicates to the SysRoleDelete builder.
func (srd *SysRoleDelete) Where(ps ...predicate.SysRole) *SysRoleDelete {
	srd.mutation.Where(ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *SysRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srd.sqlExec, srd.mutation, srd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *SysRoleDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *SysRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysrole.Table, sqlgraph.NewFieldSpec(sysrole.FieldID, field.TypeString))
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srd.mutation.done = true
	return affected, err
}

// SysRoleDeleteOne is the builder for deleting a single SysRole entity.
type SysRoleDeleteOne struct {
	srd *SysRoleDelete
}

// Where appends a list predicates to the SysRoleDelete builder.
func (srdo *SysRoleDeleteOne) Where(ps ...predicate.SysRole) *SysRoleDeleteOne {
	srdo.srd.mutation.Where(ps...)
	return srdo
}

// Exec executes the deletion query.
func (srdo *SysRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *SysRoleDeleteOne) ExecX(ctx context.Context) {
	if err := srdo.Exec(ctx); err != nil {
		panic(err)
	}
}