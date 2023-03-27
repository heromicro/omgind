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
	"github.com/heromicro/omgind/internal/gen/ent/xxxdemo"
)

// XxxDemoQuery is the builder for querying XxxDemo entities.
type XxxDemoQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.XxxDemo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the XxxDemoQuery builder.
func (xdq *XxxDemoQuery) Where(ps ...predicate.XxxDemo) *XxxDemoQuery {
	xdq.predicates = append(xdq.predicates, ps...)
	return xdq
}

// Limit the number of records to be returned by this query.
func (xdq *XxxDemoQuery) Limit(limit int) *XxxDemoQuery {
	xdq.ctx.Limit = &limit
	return xdq
}

// Offset to start from.
func (xdq *XxxDemoQuery) Offset(offset int) *XxxDemoQuery {
	xdq.ctx.Offset = &offset
	return xdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (xdq *XxxDemoQuery) Unique(unique bool) *XxxDemoQuery {
	xdq.ctx.Unique = &unique
	return xdq
}

// Order specifies how the records should be ordered.
func (xdq *XxxDemoQuery) Order(o ...OrderFunc) *XxxDemoQuery {
	xdq.order = append(xdq.order, o...)
	return xdq
}

// First returns the first XxxDemo entity from the query.
// Returns a *NotFoundError when no XxxDemo was found.
func (xdq *XxxDemoQuery) First(ctx context.Context) (*XxxDemo, error) {
	nodes, err := xdq.Limit(1).All(setContextOp(ctx, xdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{xxxdemo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (xdq *XxxDemoQuery) FirstX(ctx context.Context) *XxxDemo {
	node, err := xdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first XxxDemo ID from the query.
// Returns a *NotFoundError when no XxxDemo ID was found.
func (xdq *XxxDemoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = xdq.Limit(1).IDs(setContextOp(ctx, xdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{xxxdemo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (xdq *XxxDemoQuery) FirstIDX(ctx context.Context) string {
	id, err := xdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single XxxDemo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one XxxDemo entity is found.
// Returns a *NotFoundError when no XxxDemo entities are found.
func (xdq *XxxDemoQuery) Only(ctx context.Context) (*XxxDemo, error) {
	nodes, err := xdq.Limit(2).All(setContextOp(ctx, xdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{xxxdemo.Label}
	default:
		return nil, &NotSingularError{xxxdemo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (xdq *XxxDemoQuery) OnlyX(ctx context.Context) *XxxDemo {
	node, err := xdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only XxxDemo ID in the query.
// Returns a *NotSingularError when more than one XxxDemo ID is found.
// Returns a *NotFoundError when no entities are found.
func (xdq *XxxDemoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = xdq.Limit(2).IDs(setContextOp(ctx, xdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{xxxdemo.Label}
	default:
		err = &NotSingularError{xxxdemo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (xdq *XxxDemoQuery) OnlyIDX(ctx context.Context) string {
	id, err := xdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of XxxDemos.
func (xdq *XxxDemoQuery) All(ctx context.Context) ([]*XxxDemo, error) {
	ctx = setContextOp(ctx, xdq.ctx, "All")
	if err := xdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*XxxDemo, *XxxDemoQuery]()
	return withInterceptors[[]*XxxDemo](ctx, xdq, qr, xdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (xdq *XxxDemoQuery) AllX(ctx context.Context) []*XxxDemo {
	nodes, err := xdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of XxxDemo IDs.
func (xdq *XxxDemoQuery) IDs(ctx context.Context) (ids []string, err error) {
	if xdq.ctx.Unique == nil && xdq.path != nil {
		xdq.Unique(true)
	}
	ctx = setContextOp(ctx, xdq.ctx, "IDs")
	if err = xdq.Select(xxxdemo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (xdq *XxxDemoQuery) IDsX(ctx context.Context) []string {
	ids, err := xdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (xdq *XxxDemoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, xdq.ctx, "Count")
	if err := xdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, xdq, querierCount[*XxxDemoQuery](), xdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (xdq *XxxDemoQuery) CountX(ctx context.Context) int {
	count, err := xdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (xdq *XxxDemoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, xdq.ctx, "Exist")
	switch _, err := xdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (xdq *XxxDemoQuery) ExistX(ctx context.Context) bool {
	exist, err := xdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the XxxDemoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (xdq *XxxDemoQuery) Clone() *XxxDemoQuery {
	if xdq == nil {
		return nil
	}
	return &XxxDemoQuery{
		config:     xdq.config,
		ctx:        xdq.ctx.Clone(),
		order:      append([]OrderFunc{}, xdq.order...),
		inters:     append([]Interceptor{}, xdq.inters...),
		predicates: append([]predicate.XxxDemo{}, xdq.predicates...),
		// clone intermediate query.
		sql:  xdq.sql.Clone(),
		path: xdq.path,
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
//	client.XxxDemo.Query().
//		GroupBy(xxxdemo.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (xdq *XxxDemoQuery) GroupBy(field string, fields ...string) *XxxDemoGroupBy {
	xdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &XxxDemoGroupBy{build: xdq}
	grbuild.flds = &xdq.ctx.Fields
	grbuild.label = xxxdemo.Label
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
//	client.XxxDemo.Query().
//		Select(xxxdemo.FieldIsDel).
//		Scan(ctx, &v)
func (xdq *XxxDemoQuery) Select(fields ...string) *XxxDemoSelect {
	xdq.ctx.Fields = append(xdq.ctx.Fields, fields...)
	sbuild := &XxxDemoSelect{XxxDemoQuery: xdq}
	sbuild.label = xxxdemo.Label
	sbuild.flds, sbuild.scan = &xdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a XxxDemoSelect configured with the given aggregations.
func (xdq *XxxDemoQuery) Aggregate(fns ...AggregateFunc) *XxxDemoSelect {
	return xdq.Select().Aggregate(fns...)
}

func (xdq *XxxDemoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range xdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, xdq); err != nil {
				return err
			}
		}
	}
	for _, f := range xdq.ctx.Fields {
		if !xxxdemo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if xdq.path != nil {
		prev, err := xdq.path(ctx)
		if err != nil {
			return err
		}
		xdq.sql = prev
	}
	return nil
}

func (xdq *XxxDemoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*XxxDemo, error) {
	var (
		nodes = []*XxxDemo{}
		_spec = xdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*XxxDemo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &XxxDemo{config: xdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(xdq.modifiers) > 0 {
		_spec.Modifiers = xdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, xdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (xdq *XxxDemoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := xdq.querySpec()
	if len(xdq.modifiers) > 0 {
		_spec.Modifiers = xdq.modifiers
	}
	_spec.Node.Columns = xdq.ctx.Fields
	if len(xdq.ctx.Fields) > 0 {
		_spec.Unique = xdq.ctx.Unique != nil && *xdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, xdq.driver, _spec)
}

func (xdq *XxxDemoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(xxxdemo.Table, xxxdemo.Columns, sqlgraph.NewFieldSpec(xxxdemo.FieldID, field.TypeString))
	_spec.From = xdq.sql
	if unique := xdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if xdq.path != nil {
		_spec.Unique = true
	}
	if fields := xdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xxxdemo.FieldID)
		for i := range fields {
			if fields[i] != xxxdemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := xdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := xdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := xdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := xdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (xdq *XxxDemoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(xdq.driver.Dialect())
	t1 := builder.Table(xxxdemo.Table)
	columns := xdq.ctx.Fields
	if len(columns) == 0 {
		columns = xxxdemo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if xdq.sql != nil {
		selector = xdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if xdq.ctx.Unique != nil && *xdq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range xdq.modifiers {
		m(selector)
	}
	for _, p := range xdq.predicates {
		p(selector)
	}
	for _, p := range xdq.order {
		p(selector)
	}
	if offset := xdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := xdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (xdq *XxxDemoQuery) ForUpdate(opts ...sql.LockOption) *XxxDemoQuery {
	if xdq.driver.Dialect() == dialect.Postgres {
		xdq.Unique(false)
	}
	xdq.modifiers = append(xdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return xdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (xdq *XxxDemoQuery) ForShare(opts ...sql.LockOption) *XxxDemoQuery {
	if xdq.driver.Dialect() == dialect.Postgres {
		xdq.Unique(false)
	}
	xdq.modifiers = append(xdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return xdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (xdq *XxxDemoQuery) Modify(modifiers ...func(s *sql.Selector)) *XxxDemoSelect {
	xdq.modifiers = append(xdq.modifiers, modifiers...)
	return xdq.Select()
}

// XxxDemoGroupBy is the group-by builder for XxxDemo entities.
type XxxDemoGroupBy struct {
	selector
	build *XxxDemoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (xdgb *XxxDemoGroupBy) Aggregate(fns ...AggregateFunc) *XxxDemoGroupBy {
	xdgb.fns = append(xdgb.fns, fns...)
	return xdgb
}

// Scan applies the selector query and scans the result into the given value.
func (xdgb *XxxDemoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xdgb.build.ctx, "GroupBy")
	if err := xdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XxxDemoQuery, *XxxDemoGroupBy](ctx, xdgb.build, xdgb, xdgb.build.inters, v)
}

func (xdgb *XxxDemoGroupBy) sqlScan(ctx context.Context, root *XxxDemoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(xdgb.fns))
	for _, fn := range xdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*xdgb.flds)+len(xdgb.fns))
		for _, f := range *xdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*xdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// XxxDemoSelect is the builder for selecting fields of XxxDemo entities.
type XxxDemoSelect struct {
	*XxxDemoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (xds *XxxDemoSelect) Aggregate(fns ...AggregateFunc) *XxxDemoSelect {
	xds.fns = append(xds.fns, fns...)
	return xds
}

// Scan applies the selector query and scans the result into the given value.
func (xds *XxxDemoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xds.ctx, "Select")
	if err := xds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XxxDemoQuery, *XxxDemoSelect](ctx, xds.XxxDemoQuery, xds, xds.inters, v)
}

func (xds *XxxDemoSelect) sqlScan(ctx context.Context, root *XxxDemoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(xds.fns))
	for _, fn := range xds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*xds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (xds *XxxDemoSelect) Modify(modifiers ...func(s *sql.Selector)) *XxxDemoSelect {
	xds.modifiers = append(xds.modifiers, modifiers...)
	return xds
}
