// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/xxxdemo"
)

// XxxDemoDelete is the builder for deleting a XxxDemo entity.
type XxxDemoDelete struct {
	config
	hooks    []Hook
	mutation *XxxDemoMutation
}

// Where appends a list predicates to the XxxDemoDelete builder.
func (xdd *XxxDemoDelete) Where(ps ...predicate.XxxDemo) *XxxDemoDelete {
	xdd.mutation.Where(ps...)
	return xdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (xdd *XxxDemoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, xdd.sqlExec, xdd.mutation, xdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (xdd *XxxDemoDelete) ExecX(ctx context.Context) int {
	n, err := xdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (xdd *XxxDemoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(xxxdemo.Table, sqlgraph.NewFieldSpec(xxxdemo.FieldID, field.TypeString))
	if ps := xdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, xdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	xdd.mutation.done = true
	return affected, err
}

// XxxDemoDeleteOne is the builder for deleting a single XxxDemo entity.
type XxxDemoDeleteOne struct {
	xdd *XxxDemoDelete
}

// Where appends a list predicates to the XxxDemoDelete builder.
func (xddo *XxxDemoDeleteOne) Where(ps ...predicate.XxxDemo) *XxxDemoDeleteOne {
	xddo.xdd.mutation.Where(ps...)
	return xddo
}

// Exec executes the deletion query.
func (xddo *XxxDemoDeleteOne) Exec(ctx context.Context) error {
	n, err := xddo.xdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{xxxdemo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (xddo *XxxDemoDeleteOne) ExecX(ctx context.Context) {
	if err := xddo.Exec(ctx); err != nil {
		panic(err)
	}
}
