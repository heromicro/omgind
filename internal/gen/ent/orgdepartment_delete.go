// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/orgdepartment"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
)

// OrgDepartmentDelete is the builder for deleting a OrgDepartment entity.
type OrgDepartmentDelete struct {
	config
	hooks    []Hook
	mutation *OrgDepartmentMutation
}

// Where appends a list predicates to the OrgDepartmentDelete builder.
func (odd *OrgDepartmentDelete) Where(ps ...predicate.OrgDepartment) *OrgDepartmentDelete {
	odd.mutation.Where(ps...)
	return odd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odd *OrgDepartmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrgDepartmentMutation](ctx, odd.sqlExec, odd.mutation, odd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (odd *OrgDepartmentDelete) ExecX(ctx context.Context) int {
	n, err := odd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odd *OrgDepartmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgdepartment.Table, sqlgraph.NewFieldSpec(orgdepartment.FieldID, field.TypeString))
	_spec.Node.Schema = odd.schemaConfig.OrgDepartment
	ctx = internal.NewSchemaConfigContext(ctx, odd.schemaConfig)
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

// OrgDepartmentDeleteOne is the builder for deleting a single OrgDepartment entity.
type OrgDepartmentDeleteOne struct {
	odd *OrgDepartmentDelete
}

// Where appends a list predicates to the OrgDepartmentDelete builder.
func (oddo *OrgDepartmentDeleteOne) Where(ps ...predicate.OrgDepartment) *OrgDepartmentDeleteOne {
	oddo.odd.mutation.Where(ps...)
	return oddo
}

// Exec executes the deletion query.
func (oddo *OrgDepartmentDeleteOne) Exec(ctx context.Context) error {
	n, err := oddo.odd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgdepartment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oddo *OrgDepartmentDeleteOne) ExecX(ctx context.Context) {
	if err := oddo.Exec(ctx); err != nil {
		panic(err)
	}
}
