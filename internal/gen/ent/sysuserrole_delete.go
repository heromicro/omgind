// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRoleDelete is the builder for deleting a SysUserRole entity.
type SysUserRoleDelete struct {
	config
	hooks    []Hook
	mutation *SysUserRoleMutation
}

// Where appends a list predicates to the SysUserRoleDelete builder.
func (surd *SysUserRoleDelete) Where(ps ...predicate.SysUserRole) *SysUserRoleDelete {
	surd.mutation.Where(ps...)
	return surd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (surd *SysUserRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, surd.sqlExec, surd.mutation, surd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (surd *SysUserRoleDelete) ExecX(ctx context.Context) int {
	n, err := surd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (surd *SysUserRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysuserrole.Table, sqlgraph.NewFieldSpec(sysuserrole.FieldID, field.TypeString))
	if ps := surd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, surd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	surd.mutation.done = true
	return affected, err
}

// SysUserRoleDeleteOne is the builder for deleting a single SysUserRole entity.
type SysUserRoleDeleteOne struct {
	surd *SysUserRoleDelete
}

// Where appends a list predicates to the SysUserRoleDelete builder.
func (surdo *SysUserRoleDeleteOne) Where(ps ...predicate.SysUserRole) *SysUserRoleDeleteOne {
	surdo.surd.mutation.Where(ps...)
	return surdo
}

// Exec executes the deletion query.
func (surdo *SysUserRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := surdo.surd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysuserrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (surdo *SysUserRoleDeleteOne) ExecX(ctx context.Context) {
	if err := surdo.Exec(ctx); err != nil {
		panic(err)
	}
}
