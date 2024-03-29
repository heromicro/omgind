// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdictitem"
)

// SysDictItemDelete is the builder for deleting a SysDictItem entity.
type SysDictItemDelete struct {
	config
	hooks    []Hook
	mutation *SysDictItemMutation
}

// Where appends a list predicates to the SysDictItemDelete builder.
func (sdid *SysDictItemDelete) Where(ps ...predicate.SysDictItem) *SysDictItemDelete {
	sdid.mutation.Where(ps...)
	return sdid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdid *SysDictItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sdid.sqlExec, sdid.mutation, sdid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdid *SysDictItemDelete) ExecX(ctx context.Context) int {
	n, err := sdid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdid *SysDictItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysdictitem.Table, sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString))
	if ps := sdid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdid.mutation.done = true
	return affected, err
}

// SysDictItemDeleteOne is the builder for deleting a single SysDictItem entity.
type SysDictItemDeleteOne struct {
	sdid *SysDictItemDelete
}

// Where appends a list predicates to the SysDictItemDelete builder.
func (sdido *SysDictItemDeleteOne) Where(ps ...predicate.SysDictItem) *SysDictItemDeleteOne {
	sdido.sdid.mutation.Where(ps...)
	return sdido
}

// Exec executes the deletion query.
func (sdido *SysDictItemDeleteOne) Exec(ctx context.Context) error {
	n, err := sdido.sdid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysdictitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdido *SysDictItemDeleteOne) ExecX(ctx context.Context) {
	if err := sdido.Exec(ctx); err != nil {
		panic(err)
	}
}
