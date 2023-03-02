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
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemQuery is the builder for querying SysDictItem entities.
type SysDictItemQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.SysDictItem
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDictItemQuery builder.
func (sdiq *SysDictItemQuery) Where(ps ...predicate.SysDictItem) *SysDictItemQuery {
	sdiq.predicates = append(sdiq.predicates, ps...)
	return sdiq
}

// Limit the number of records to be returned by this query.
func (sdiq *SysDictItemQuery) Limit(limit int) *SysDictItemQuery {
	sdiq.ctx.Limit = &limit
	return sdiq
}

// Offset to start from.
func (sdiq *SysDictItemQuery) Offset(offset int) *SysDictItemQuery {
	sdiq.ctx.Offset = &offset
	return sdiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdiq *SysDictItemQuery) Unique(unique bool) *SysDictItemQuery {
	sdiq.ctx.Unique = &unique
	return sdiq
}

// Order specifies how the records should be ordered.
func (sdiq *SysDictItemQuery) Order(o ...OrderFunc) *SysDictItemQuery {
	sdiq.order = append(sdiq.order, o...)
	return sdiq
}

// First returns the first SysDictItem entity from the query.
// Returns a *NotFoundError when no SysDictItem was found.
func (sdiq *SysDictItemQuery) First(ctx context.Context) (*SysDictItem, error) {
	nodes, err := sdiq.Limit(1).All(setContextOp(ctx, sdiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdictitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdiq *SysDictItemQuery) FirstX(ctx context.Context) *SysDictItem {
	node, err := sdiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDictItem ID from the query.
// Returns a *NotFoundError when no SysDictItem ID was found.
func (sdiq *SysDictItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdiq.Limit(1).IDs(setContextOp(ctx, sdiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdictitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdiq *SysDictItemQuery) FirstIDX(ctx context.Context) string {
	id, err := sdiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDictItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysDictItem entity is found.
// Returns a *NotFoundError when no SysDictItem entities are found.
func (sdiq *SysDictItemQuery) Only(ctx context.Context) (*SysDictItem, error) {
	nodes, err := sdiq.Limit(2).All(setContextOp(ctx, sdiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdictitem.Label}
	default:
		return nil, &NotSingularError{sysdictitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdiq *SysDictItemQuery) OnlyX(ctx context.Context) *SysDictItem {
	node, err := sdiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDictItem ID in the query.
// Returns a *NotSingularError when more than one SysDictItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdiq *SysDictItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdiq.Limit(2).IDs(setContextOp(ctx, sdiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = &NotSingularError{sysdictitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdiq *SysDictItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := sdiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDictItems.
func (sdiq *SysDictItemQuery) All(ctx context.Context) ([]*SysDictItem, error) {
	ctx = setContextOp(ctx, sdiq.ctx, "All")
	if err := sdiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysDictItem, *SysDictItemQuery]()
	return withInterceptors[[]*SysDictItem](ctx, sdiq, qr, sdiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sdiq *SysDictItemQuery) AllX(ctx context.Context) []*SysDictItem {
	nodes, err := sdiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDictItem IDs.
func (sdiq *SysDictItemQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sdiq.ctx.Unique == nil && sdiq.path != nil {
		sdiq.Unique(true)
	}
	ctx = setContextOp(ctx, sdiq.ctx, "IDs")
	if err = sdiq.Select(sysdictitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdiq *SysDictItemQuery) IDsX(ctx context.Context) []string {
	ids, err := sdiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdiq *SysDictItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sdiq.ctx, "Count")
	if err := sdiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sdiq, querierCount[*SysDictItemQuery](), sdiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sdiq *SysDictItemQuery) CountX(ctx context.Context) int {
	count, err := sdiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdiq *SysDictItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sdiq.ctx, "Exist")
	switch _, err := sdiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sdiq *SysDictItemQuery) ExistX(ctx context.Context) bool {
	exist, err := sdiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDictItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdiq *SysDictItemQuery) Clone() *SysDictItemQuery {
	if sdiq == nil {
		return nil
	}
	return &SysDictItemQuery{
		config:     sdiq.config,
		ctx:        sdiq.ctx.Clone(),
		order:      append([]OrderFunc{}, sdiq.order...),
		inters:     append([]Interceptor{}, sdiq.inters...),
		predicates: append([]predicate.SysDictItem{}, sdiq.predicates...),
		// clone intermediate query.
		sql:  sdiq.sql.Clone(),
		path: sdiq.path,
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
//	client.SysDictItem.Query().
//		GroupBy(sysdictitem.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdiq *SysDictItemQuery) GroupBy(field string, fields ...string) *SysDictItemGroupBy {
	sdiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysDictItemGroupBy{build: sdiq}
	grbuild.flds = &sdiq.ctx.Fields
	grbuild.label = sysdictitem.Label
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
//	client.SysDictItem.Query().
//		Select(sysdictitem.FieldIsDel).
//		Scan(ctx, &v)
func (sdiq *SysDictItemQuery) Select(fields ...string) *SysDictItemSelect {
	sdiq.ctx.Fields = append(sdiq.ctx.Fields, fields...)
	sbuild := &SysDictItemSelect{SysDictItemQuery: sdiq}
	sbuild.label = sysdictitem.Label
	sbuild.flds, sbuild.scan = &sdiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysDictItemSelect configured with the given aggregations.
func (sdiq *SysDictItemQuery) Aggregate(fns ...AggregateFunc) *SysDictItemSelect {
	return sdiq.Select().Aggregate(fns...)
}

func (sdiq *SysDictItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sdiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sdiq); err != nil {
				return err
			}
		}
	}
	for _, f := range sdiq.ctx.Fields {
		if !sysdictitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdiq.path != nil {
		prev, err := sdiq.path(ctx)
		if err != nil {
			return err
		}
		sdiq.sql = prev
	}
	return nil
}

func (sdiq *SysDictItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysDictItem, error) {
	var (
		nodes = []*SysDictItem{}
		_spec = sdiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysDictItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysDictItem{config: sdiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = sdiq.schemaConfig.SysDictItem
	ctx = internal.NewSchemaConfigContext(ctx, sdiq.schemaConfig)
	if len(sdiq.modifiers) > 0 {
		_spec.Modifiers = sdiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sdiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sdiq *SysDictItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdiq.querySpec()
	_spec.Node.Schema = sdiq.schemaConfig.SysDictItem
	ctx = internal.NewSchemaConfigContext(ctx, sdiq.schemaConfig)
	if len(sdiq.modifiers) > 0 {
		_spec.Modifiers = sdiq.modifiers
	}
	_spec.Node.Columns = sdiq.ctx.Fields
	if len(sdiq.ctx.Fields) > 0 {
		_spec.Unique = sdiq.ctx.Unique != nil && *sdiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sdiq.driver, _spec)
}

func (sdiq *SysDictItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysdictitem.Table, sysdictitem.Columns, sqlgraph.NewFieldSpec(sysdictitem.FieldID, field.TypeString))
	_spec.From = sdiq.sql
	if unique := sdiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sdiq.path != nil {
		_spec.Unique = true
	}
	if fields := sdiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictitem.FieldID)
		for i := range fields {
			if fields[i] != sysdictitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdiq *SysDictItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdiq.driver.Dialect())
	t1 := builder.Table(sysdictitem.Table)
	columns := sdiq.ctx.Fields
	if len(columns) == 0 {
		columns = sysdictitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdiq.sql != nil {
		selector = sdiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdiq.ctx.Unique != nil && *sdiq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(sdiq.schemaConfig.SysDictItem)
	ctx = internal.NewSchemaConfigContext(ctx, sdiq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range sdiq.modifiers {
		m(selector)
	}
	for _, p := range sdiq.predicates {
		p(selector)
	}
	for _, p := range sdiq.order {
		p(selector)
	}
	if offset := sdiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sdiq *SysDictItemQuery) ForUpdate(opts ...sql.LockOption) *SysDictItemQuery {
	if sdiq.driver.Dialect() == dialect.Postgres {
		sdiq.Unique(false)
	}
	sdiq.modifiers = append(sdiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sdiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sdiq *SysDictItemQuery) ForShare(opts ...sql.LockOption) *SysDictItemQuery {
	if sdiq.driver.Dialect() == dialect.Postgres {
		sdiq.Unique(false)
	}
	sdiq.modifiers = append(sdiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sdiq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sdiq *SysDictItemQuery) Modify(modifiers ...func(s *sql.Selector)) *SysDictItemSelect {
	sdiq.modifiers = append(sdiq.modifiers, modifiers...)
	return sdiq.Select()
}

// SysDictItemGroupBy is the group-by builder for SysDictItem entities.
type SysDictItemGroupBy struct {
	selector
	build *SysDictItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdigb *SysDictItemGroupBy) Aggregate(fns ...AggregateFunc) *SysDictItemGroupBy {
	sdigb.fns = append(sdigb.fns, fns...)
	return sdigb
}

// Scan applies the selector query and scans the result into the given value.
func (sdigb *SysDictItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdigb.build.ctx, "GroupBy")
	if err := sdigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictItemQuery, *SysDictItemGroupBy](ctx, sdigb.build, sdigb, sdigb.build.inters, v)
}

func (sdigb *SysDictItemGroupBy) sqlScan(ctx context.Context, root *SysDictItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sdigb.fns))
	for _, fn := range sdigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sdigb.flds)+len(sdigb.fns))
		for _, f := range *sdigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sdigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysDictItemSelect is the builder for selecting fields of SysDictItem entities.
type SysDictItemSelect struct {
	*SysDictItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sdis *SysDictItemSelect) Aggregate(fns ...AggregateFunc) *SysDictItemSelect {
	sdis.fns = append(sdis.fns, fns...)
	return sdis
}

// Scan applies the selector query and scans the result into the given value.
func (sdis *SysDictItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdis.ctx, "Select")
	if err := sdis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictItemQuery, *SysDictItemSelect](ctx, sdis.SysDictItemQuery, sdis, sdis.inters, v)
}

func (sdis *SysDictItemSelect) sqlScan(ctx context.Context, root *SysDictItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sdis.fns))
	for _, fn := range sdis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sdis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sdis *SysDictItemSelect) Modify(modifiers ...func(s *sql.Selector)) *SysDictItemSelect {
	sdis.modifiers = append(sdis.modifiers, modifiers...)
	return sdis
}
