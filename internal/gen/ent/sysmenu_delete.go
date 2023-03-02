// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

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
	return withHooks[int, SysMenuMutation](ctx, smd.sqlExec, smd.mutation, smd.hooks)
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
	_spec := sqlgraph.NewDeleteSpec(sysmenu.Table, sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeString))
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
	smd.mutation.done = true
	return affected, err
}

// SysMenuDeleteOne is the builder for deleting a single SysMenu entity.
type SysMenuDeleteOne struct {
	smd *SysMenuDelete
}

// Where appends a list predicates to the SysMenuDelete builder.
func (smdo *SysMenuDeleteOne) Where(ps ...predicate.SysMenu) *SysMenuDeleteOne {
	smdo.smd.mutation.Where(ps...)
	return smdo
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
	if err := smdo.Exec(ctx); err != nil {
		panic(err)
	}
}
