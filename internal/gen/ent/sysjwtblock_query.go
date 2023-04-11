// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlockQuery is the builder for querying SysJwtBlock entities.
type SysJwtBlockQuery struct {
	config
	ctx        *QueryContext
	order      []sysjwtblock.Order
	inters     []Interceptor
	predicates []predicate.SysJwtBlock
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysJwtBlockQuery builder.
func (sjbq *SysJwtBlockQuery) Where(ps ...predicate.SysJwtBlock) *SysJwtBlockQuery {
	sjbq.predicates = append(sjbq.predicates, ps...)
	return sjbq
}

// Limit the number of records to be returned by this query.
func (sjbq *SysJwtBlockQuery) Limit(limit int) *SysJwtBlockQuery {
	sjbq.ctx.Limit = &limit
	return sjbq
}

// Offset to start from.
func (sjbq *SysJwtBlockQuery) Offset(offset int) *SysJwtBlockQuery {
	sjbq.ctx.Offset = &offset
	return sjbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sjbq *SysJwtBlockQuery) Unique(unique bool) *SysJwtBlockQuery {
	sjbq.ctx.Unique = &unique
	return sjbq
}

// Order specifies how the records should be ordered.
func (sjbq *SysJwtBlockQuery) Order(o ...sysjwtblock.Order) *SysJwtBlockQuery {
	sjbq.order = append(sjbq.order, o...)
	return sjbq
}

// First returns the first SysJwtBlock entity from the query.
// Returns a *NotFoundError when no SysJwtBlock was found.
func (sjbq *SysJwtBlockQuery) First(ctx context.Context) (*SysJwtBlock, error) {
	nodes, err := sjbq.Limit(1).All(setContextOp(ctx, sjbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysjwtblock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) FirstX(ctx context.Context) *SysJwtBlock {
	node, err := sjbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysJwtBlock ID from the query.
// Returns a *NotFoundError when no SysJwtBlock ID was found.
func (sjbq *SysJwtBlockQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sjbq.Limit(1).IDs(setContextOp(ctx, sjbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysjwtblock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) FirstIDX(ctx context.Context) string {
	id, err := sjbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysJwtBlock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysJwtBlock entity is found.
// Returns a *NotFoundError when no SysJwtBlock entities are found.
func (sjbq *SysJwtBlockQuery) Only(ctx context.Context) (*SysJwtBlock, error) {
	nodes, err := sjbq.Limit(2).All(setContextOp(ctx, sjbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysjwtblock.Label}
	default:
		return nil, &NotSingularError{sysjwtblock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) OnlyX(ctx context.Context) *SysJwtBlock {
	node, err := sjbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysJwtBlock ID in the query.
// Returns a *NotSingularError when more than one SysJwtBlock ID is found.
// Returns a *NotFoundError when no entities are found.
func (sjbq *SysJwtBlockQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sjbq.Limit(2).IDs(setContextOp(ctx, sjbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysjwtblock.Label}
	default:
		err = &NotSingularError{sysjwtblock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) OnlyIDX(ctx context.Context) string {
	id, err := sjbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysJwtBlocks.
func (sjbq *SysJwtBlockQuery) All(ctx context.Context) ([]*SysJwtBlock, error) {
	ctx = setContextOp(ctx, sjbq.ctx, "All")
	if err := sjbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysJwtBlock, *SysJwtBlockQuery]()
	return withInterceptors[[]*SysJwtBlock](ctx, sjbq, qr, sjbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) AllX(ctx context.Context) []*SysJwtBlock {
	nodes, err := sjbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysJwtBlock IDs.
func (sjbq *SysJwtBlockQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sjbq.ctx.Unique == nil && sjbq.path != nil {
		sjbq.Unique(true)
	}
	ctx = setContextOp(ctx, sjbq.ctx, "IDs")
	if err = sjbq.Select(sysjwtblock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) IDsX(ctx context.Context) []string {
	ids, err := sjbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sjbq *SysJwtBlockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sjbq.ctx, "Count")
	if err := sjbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sjbq, querierCount[*SysJwtBlockQuery](), sjbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) CountX(ctx context.Context) int {
	count, err := sjbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sjbq *SysJwtBlockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sjbq.ctx, "Exist")
	switch _, err := sjbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sjbq *SysJwtBlockQuery) ExistX(ctx context.Context) bool {
	exist, err := sjbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysJwtBlockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sjbq *SysJwtBlockQuery) Clone() *SysJwtBlockQuery {
	if sjbq == nil {
		return nil
	}
	return &SysJwtBlockQuery{
		config:     sjbq.config,
		ctx:        sjbq.ctx.Clone(),
		order:      append([]sysjwtblock.Order{}, sjbq.order...),
		inters:     append([]Interceptor{}, sjbq.inters...),
		predicates: append([]predicate.SysJwtBlock{}, sjbq.predicates...),
		// clone intermediate query.
		sql:  sjbq.sql.Clone(),
		path: sjbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsDel bool `json:"is_del,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysJwtBlock.Query().
//		GroupBy(sysjwtblock.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sjbq *SysJwtBlockQuery) GroupBy(field string, fields ...string) *SysJwtBlockGroupBy {
	sjbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysJwtBlockGroupBy{build: sjbq}
	grbuild.flds = &sjbq.ctx.Fields
	grbuild.label = sysjwtblock.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsDel bool `json:"is_del,omitempty"`
//	}
//
//	client.SysJwtBlock.Query().
//		Select(sysjwtblock.FieldIsDel).
//		Scan(ctx, &v)
func (sjbq *SysJwtBlockQuery) Select(fields ...string) *SysJwtBlockSelect {
	sjbq.ctx.Fields = append(sjbq.ctx.Fields, fields...)
	sbuild := &SysJwtBlockSelect{SysJwtBlockQuery: sjbq}
	sbuild.label = sysjwtblock.Label
	sbuild.flds, sbuild.scan = &sjbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysJwtBlockSelect configured with the given aggregations.
func (sjbq *SysJwtBlockQuery) Aggregate(fns ...AggregateFunc) *SysJwtBlockSelect {
	return sjbq.Select().Aggregate(fns...)
}

func (sjbq *SysJwtBlockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sjbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sjbq); err != nil {
				return err
			}
		}
	}
	for _, f := range sjbq.ctx.Fields {
		if !sysjwtblock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sjbq.path != nil {
		prev, err := sjbq.path(ctx)
		if err != nil {
			return err
		}
		sjbq.sql = prev
	}
	return nil
}

func (sjbq *SysJwtBlockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysJwtBlock, error) {
	var (
		nodes = []*SysJwtBlock{}
		_spec = sjbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysJwtBlock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysJwtBlock{config: sjbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(sjbq.modifiers) > 0 {
		_spec.Modifiers = sjbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sjbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sjbq *SysJwtBlockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sjbq.querySpec()
	if len(sjbq.modifiers) > 0 {
		_spec.Modifiers = sjbq.modifiers
	}
	_spec.Node.Columns = sjbq.ctx.Fields
	if len(sjbq.ctx.Fields) > 0 {
		_spec.Unique = sjbq.ctx.Unique != nil && *sjbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sjbq.driver, _spec)
}

func (sjbq *SysJwtBlockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysjwtblock.Table, sysjwtblock.Columns, sqlgraph.NewFieldSpec(sysjwtblock.FieldID, field.TypeString))
	_spec.From = sjbq.sql
	if unique := sjbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sjbq.path != nil {
		_spec.Unique = true
	}
	if fields := sjbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysjwtblock.FieldID)
		for i := range fields {
			if fields[i] != sysjwtblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sjbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sjbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sjbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sjbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sjbq *SysJwtBlockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sjbq.driver.Dialect())
	t1 := builder.Table(sysjwtblock.Table)
	columns := sjbq.ctx.Fields
	if len(columns) == 0 {
		columns = sysjwtblock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sjbq.sql != nil {
		selector = sjbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sjbq.ctx.Unique != nil && *sjbq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sjbq.modifiers {
		m(selector)
	}
	for _, p := range sjbq.predicates {
		p(selector)
	}
	for _, p := range sjbq.order {
		p(selector)
	}
	if offset := sjbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sjbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sjbq *SysJwtBlockQuery) ForUpdate(opts ...sql.LockOption) *SysJwtBlockQuery {
	if sjbq.driver.Dialect() == dialect.Postgres {
		sjbq.Unique(false)
	}
	sjbq.modifiers = append(sjbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sjbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sjbq *SysJwtBlockQuery) ForShare(opts ...sql.LockOption) *SysJwtBlockQuery {
	if sjbq.driver.Dialect() == dialect.Postgres {
		sjbq.Unique(false)
	}
	sjbq.modifiers = append(sjbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sjbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sjbq *SysJwtBlockQuery) Modify(modifiers ...func(s *sql.Selector)) *SysJwtBlockSelect {
	sjbq.modifiers = append(sjbq.modifiers, modifiers...)
	return sjbq.Select()
}

// SysJwtBlockGroupBy is the group-by builder for SysJwtBlock entities.
type SysJwtBlockGroupBy struct {
	selector
	build *SysJwtBlockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sjbgb *SysJwtBlockGroupBy) Aggregate(fns ...AggregateFunc) *SysJwtBlockGroupBy {
	sjbgb.fns = append(sjbgb.fns, fns...)
	return sjbgb
}

// Scan applies the selector query and scans the result into the given value.
func (sjbgb *SysJwtBlockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sjbgb.build.ctx, "GroupBy")
	if err := sjbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysJwtBlockQuery, *SysJwtBlockGroupBy](ctx, sjbgb.build, sjbgb, sjbgb.build.inters, v)
}

func (sjbgb *SysJwtBlockGroupBy) sqlScan(ctx context.Context, root *SysJwtBlockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sjbgb.fns))
	for _, fn := range sjbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sjbgb.flds)+len(sjbgb.fns))
		for _, f := range *sjbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sjbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sjbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysJwtBlockSelect is the builder for selecting fields of SysJwtBlock entities.
type SysJwtBlockSelect struct {
	*SysJwtBlockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sjbs *SysJwtBlockSelect) Aggregate(fns ...AggregateFunc) *SysJwtBlockSelect {
	sjbs.fns = append(sjbs.fns, fns...)
	return sjbs
}

// Scan applies the selector query and scans the result into the given value.
func (sjbs *SysJwtBlockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sjbs.ctx, "Select")
	if err := sjbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysJwtBlockQuery, *SysJwtBlockSelect](ctx, sjbs.SysJwtBlockQuery, sjbs, sjbs.inters, v)
}

func (sjbs *SysJwtBlockSelect) sqlScan(ctx context.Context, root *SysJwtBlockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sjbs.fns))
	for _, fn := range sjbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sjbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sjbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sjbs *SysJwtBlockSelect) Modify(modifiers ...func(s *sql.Selector)) *SysJwtBlockSelect {
	sjbs.modifiers = append(sjbs.modifiers, modifiers...)
	return sjbs
}
