// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/entscheme/orgdept"
	"github.com/heromicro/omgind/internal/gen/entscheme/predicate"
)

// OrgDeptDelete is the builder for deleting a OrgDept entity.
type OrgDeptDelete struct {
	config
	hooks    []Hook
	mutation *OrgDeptMutation
}

// Where appends a list predicates to the OrgDeptDelete builder.
func (odd *OrgDeptDelete) Where(ps ...predicate.OrgDept) *OrgDeptDelete {
	odd.mutation.Where(ps...)
	return odd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odd *OrgDeptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, odd.sqlExec, odd.mutation, odd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (odd *OrgDeptDelete) ExecX(ctx context.Context) int {
	n, err := odd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odd *OrgDeptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgdept.Table, sqlgraph.NewFieldSpec(orgdept.FieldID, field.TypeString))
	if ps := odd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, odd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	odd.mutation.done = true
	return affected, err
}

// OrgDeptDeleteOne is the builder for deleting a single OrgDept entity.
type OrgDeptDeleteOne struct {
	odd *OrgDeptDelete
}

// Where appends a list predicates to the OrgDeptDelete builder.
func (oddo *OrgDeptDeleteOne) Where(ps ...predicate.OrgDept) *OrgDeptDeleteOne {
	oddo.odd.mutation.Where(ps...)
	return oddo
}

// Exec executes the deletion query.
func (oddo *OrgDeptDeleteOne) Exec(ctx context.Context) error {
	n, err := oddo.odd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgdept.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oddo *OrgDeptDeleteOne) ExecX(ctx context.Context) {
	if err := oddo.Exec(ctx); err != nil {
		panic(err)
	}
}