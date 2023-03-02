// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
)

// SysDistrictDelete is the builder for deleting a SysDistrict entity.
type SysDistrictDelete struct {
	config
	hooks    []Hook
	mutation *SysDistrictMutation
}

// Where appends a list predicates to the SysDistrictDelete builder.
func (sdd *SysDistrictDelete) Where(ps ...predicate.SysDistrict) *SysDistrictDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SysDistrictDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SysDistrictMutation](ctx, sdd.sqlExec, sdd.mutation, sdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SysDistrictDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SysDistrictDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysdistrict.Table, sqlgraph.NewFieldSpec(sysdistrict.FieldID, field.TypeString))
	_spec.Node.Schema = sdd.schemaConfig.SysDistrict
	ctx = internal.NewSchemaConfigContext(ctx, sdd.schemaConfig)
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdd.mutation.done = true
	return affected, err
}

// SysDistrictDeleteOne is the builder for deleting a single SysDistrict entity.
type SysDistrictDeleteOne struct {
	sdd *SysDistrictDelete
}

// Where appends a list predicates to the SysDistrictDelete builder.
func (sddo *SysDistrictDeleteOne) Where(ps ...predicate.SysDistrict) *SysDistrictDeleteOne {
	sddo.sdd.mutation.Where(ps...)
	return sddo
}

// Exec executes the deletion query.
func (sddo *SysDistrictDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysdistrict.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SysDistrictDeleteOne) ExecX(ctx context.Context) {
	if err := sddo.Exec(ctx); err != nil {
		panic(err)
	}
}
