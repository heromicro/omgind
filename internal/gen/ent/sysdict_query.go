// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictQuery is the builder for querying SysDict entities.
type SysDictQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.SysDict
	withItems  *SysDictItemQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDictQuery builder.
func (sdq *SysDictQuery) Where(ps ...predicate.SysDict) *SysDictQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit the number of records to be returned by this query.
func (sdq *SysDictQuery) Limit(limit int) *SysDictQuery {
	sdq.ctx.Limit = &limit
	return sdq
}

// Offset to start from.
func (sdq *SysDictQuery) Offset(offset int) *SysDictQuery {
	sdq.ctx.Offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SysDictQuery) Unique(unique bool) *SysDictQuery {
	sdq.ctx.Unique = &unique
	return sdq
}

// Order specifies how the records should be ordered.
func (sdq *SysDictQuery) Order(o ...OrderFunc) *SysDictQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// QueryItems chains the current query on the "items" edge.
func (sdq *SysDictQuery) QueryItems() *SysDictItemQuery {
	query := (&SysDictItemClient{config: sdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysdict.Table, sysdict.FieldID, selector),
			sqlgraph.To(sysdictitem.Table, sysdictitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysdict.ItemsTable, sysdict.ItemsColumn),
		)
		schemaConfig := sdq.schemaConfig
		step.To.Schema = schemaConfig.SysDictItem
		step.Edge.Schema = schemaConfig.SysDictItem
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysDict entity from the query.
// Returns a *NotFoundError when no SysDict was found.
func (sdq *SysDictQuery) First(ctx context.Context) (*SysDict, error) {
	nodes, err := sdq.Limit(1).All(setContextOp(ctx, sdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdict.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SysDictQuery) FirstX(ctx context.Context) *SysDict {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDict ID from the query.
// Returns a *NotFoundError when no SysDict ID was found.
func (sdq *SysDictQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdq.Limit(1).IDs(setContextOp(ctx, sdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdict.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SysDictQuery) FirstIDX(ctx context.Context) string {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDict entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysDict entity is found.
// Returns a *NotFoundError when no SysDict entities are found.
func (sdq *SysDictQuery) Only(ctx context.Context) (*SysDict, error) {
	nodes, err := sdq.Limit(2).All(setContextOp(ctx, sdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdict.Label}
	default:
		return nil, &NotSingularError{sysdict.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SysDictQuery) OnlyX(ctx context.Context) *SysDict {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDict ID in the query.
// Returns a *NotSingularError when more than one SysDict ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SysDictQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdq.Limit(2).IDs(setContextOp(ctx, sdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdict.Label}
	default:
		err = &NotSingularError{sysdict.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SysDictQuery) OnlyIDX(ctx context.Context) string {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDicts.
func (sdq *SysDictQuery) All(ctx context.Context) ([]*SysDict, error) {
	ctx = setContextOp(ctx, sdq.ctx, "All")
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysDict, *SysDictQuery]()
	return withInterceptors[[]*SysDict](ctx, sdq, qr, sdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SysDictQuery) AllX(ctx context.Context) []*SysDict {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDict IDs.
func (sdq *SysDictQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sdq.ctx.Unique == nil && sdq.path != nil {
		sdq.Unique(true)
	}
	ctx = setContextOp(ctx, sdq.ctx, "IDs")
	if err = sdq.Select(sysdict.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SysDictQuery) IDsX(ctx context.Context) []string {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SysDictQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sdq.ctx, "Count")
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sdq, querierCount[*SysDictQuery](), sdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SysDictQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SysDictQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sdq.ctx, "Exist")
	switch _, err := sdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sdq *SysDictQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDictQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SysDictQuery) Clone() *SysDictQuery {
	if sdq == nil {
		return nil
	}
	return &SysDictQuery{
		config:     sdq.config,
		ctx:        sdq.ctx.Clone(),
		order:      append([]OrderFunc{}, sdq.order...),
		inters:     append([]Interceptor{}, sdq.inters...),
		predicates: append([]predicate.SysDict{}, sdq.predicates...),
		withItems:  sdq.withItems.Clone(),
		// clone intermediate query.
		sql:  sdq.sql.Clone(),
		path: sdq.path,
	}
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SysDictQuery) WithItems(opts ...func(*SysDictItemQuery)) *SysDictQuery {
	query := (&SysDictItemClient{config: sdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sdq.withItems = query
	return sdq
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
//	client.SysDict.Query().
//		GroupBy(sysdict.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SysDictQuery) GroupBy(field string, fields ...string) *SysDictGroupBy {
	sdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysDictGroupBy{build: sdq}
	grbuild.flds = &sdq.ctx.Fields
	grbuild.label = sysdict.Label
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
//	client.SysDict.Query().
//		Select(sysdict.FieldIsDel).
//		Scan(ctx, &v)
func (sdq *SysDictQuery) Select(fields ...string) *SysDictSelect {
	sdq.ctx.Fields = append(sdq.ctx.Fields, fields...)
	sbuild := &SysDictSelect{SysDictQuery: sdq}
	sbuild.label = sysdict.Label
	sbuild.flds, sbuild.scan = &sdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysDictSelect configured with the given aggregations.
func (sdq *SysDictQuery) Aggregate(fns ...AggregateFunc) *SysDictSelect {
	return sdq.Select().Aggregate(fns...)
}

func (sdq *SysDictQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sdq); err != nil {
				return err
			}
		}
	}
	for _, f := range sdq.ctx.Fields {
		if !sysdict.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdq.path != nil {
		prev, err := sdq.path(ctx)
		if err != nil {
			return err
		}
		sdq.sql = prev
	}
	return nil
}

func (sdq *SysDictQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysDict, error) {
	var (
		nodes       = []*SysDict{}
		_spec       = sdq.querySpec()
		loadedTypes = [1]bool{
			sdq.withItems != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysDict).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysDict{config: sdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = sdq.schemaConfig.SysDict
	ctx = internal.NewSchemaConfigContext(ctx, sdq.schemaConfig)
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sdq.withItems; query != nil {
		if err := sdq.loadItems(ctx, query, nodes,
			func(n *SysDict) { n.Edges.Items = []*SysDictItem{} },
			func(n *SysDict, e *SysDictItem) { n.Edges.Items = append(n.Edges.Items, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sdq *SysDictQuery) loadItems(ctx context.Context, query *SysDictItemQuery, nodes []*SysDict, init func(*SysDict), assign func(*SysDict, *SysDictItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysDict)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.SysDictItem(func(s *sql.Selector) {
		s.Where(sql.InValues(sysdict.ItemsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DictID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dict_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sdq *SysDictQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	_spec.Node.Schema = sdq.schemaConfig.SysDict
	ctx = internal.NewSchemaConfigContext(ctx, sdq.schemaConfig)
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	_spec.Node.Columns = sdq.ctx.Fields
	if len(sdq.ctx.Fields) > 0 {
		_spec.Unique = sdq.ctx.Unique != nil && *sdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SysDictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysdict.Table, sysdict.Columns, sqlgraph.NewFieldSpec(sysdict.FieldID, field.TypeString))
	_spec.From = sdq.sql
	if unique := sdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sdq.path != nil {
		_spec.Unique = true
	}
	if fields := sdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdict.FieldID)
		for i := range fields {
			if fields[i] != sysdict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdq *SysDictQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(sysdict.Table)
	columns := sdq.ctx.Fields
	if len(columns) == 0 {
		columns = sysdict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.ctx.Unique != nil && *sdq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(sdq.schemaConfig.SysDict)
	ctx = internal.NewSchemaConfigContext(ctx, sdq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range sdq.modifiers {
		m(selector)
	}
	for _, p := range sdq.predicates {
		p(selector)
	}
	for _, p := range sdq.order {
		p(selector)
	}
	if offset := sdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sdq *SysDictQuery) ForUpdate(opts ...sql.LockOption) *SysDictQuery {
	if sdq.driver.Dialect() == dialect.Postgres {
		sdq.Unique(false)
	}
	sdq.modifiers = append(sdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sdq *SysDictQuery) ForShare(opts ...sql.LockOption) *SysDictQuery {
	if sdq.driver.Dialect() == dialect.Postgres {
		sdq.Unique(false)
	}
	sdq.modifiers = append(sdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sdq *SysDictQuery) Modify(modifiers ...func(s *sql.Selector)) *SysDictSelect {
	sdq.modifiers = append(sdq.modifiers, modifiers...)
	return sdq.Select()
}

// SysDictGroupBy is the group-by builder for SysDict entities.
type SysDictGroupBy struct {
	selector
	build *SysDictQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SysDictGroupBy) Aggregate(fns ...AggregateFunc) *SysDictGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the selector query and scans the result into the given value.
func (sdgb *SysDictGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdgb.build.ctx, "GroupBy")
	if err := sdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictQuery, *SysDictGroupBy](ctx, sdgb.build, sdgb, sdgb.build.inters, v)
}

func (sdgb *SysDictGroupBy) sqlScan(ctx context.Context, root *SysDictQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sdgb.flds)+len(sdgb.fns))
		for _, f := range *sdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysDictSelect is the builder for selecting fields of SysDict entities.
type SysDictSelect struct {
	*SysDictQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sds *SysDictSelect) Aggregate(fns ...AggregateFunc) *SysDictSelect {
	sds.fns = append(sds.fns, fns...)
	return sds
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SysDictSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sds.ctx, "Select")
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictQuery, *SysDictSelect](ctx, sds.SysDictQuery, sds, sds.inters, v)
}

func (sds *SysDictSelect) sqlScan(ctx context.Context, root *SysDictQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sds.fns))
	for _, fn := range sds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sds *SysDictSelect) Modify(modifiers ...func(s *sql.Selector)) *SysDictSelect {
	sds.modifiers = append(sds.modifiers, modifiers...)
	return sds
}
