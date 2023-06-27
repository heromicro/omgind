// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/entscheme/predicate"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysmenuactionresource"
)

// SysMenuActionResourceQuery is the builder for querying SysMenuActionResource entities.
type SysMenuActionResourceQuery struct {
	config
	ctx        *QueryContext
	order      []sysmenuactionresource.OrderOption
	inters     []Interceptor
	predicates []predicate.SysMenuActionResource
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuActionResourceQuery builder.
func (smarq *SysMenuActionResourceQuery) Where(ps ...predicate.SysMenuActionResource) *SysMenuActionResourceQuery {
	smarq.predicates = append(smarq.predicates, ps...)
	return smarq
}

// Limit the number of records to be returned by this query.
func (smarq *SysMenuActionResourceQuery) Limit(limit int) *SysMenuActionResourceQuery {
	smarq.ctx.Limit = &limit
	return smarq
}

// Offset to start from.
func (smarq *SysMenuActionResourceQuery) Offset(offset int) *SysMenuActionResourceQuery {
	smarq.ctx.Offset = &offset
	return smarq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smarq *SysMenuActionResourceQuery) Unique(unique bool) *SysMenuActionResourceQuery {
	smarq.ctx.Unique = &unique
	return smarq
}

// Order specifies how the records should be ordered.
func (smarq *SysMenuActionResourceQuery) Order(o ...sysmenuactionresource.OrderOption) *SysMenuActionResourceQuery {
	smarq.order = append(smarq.order, o...)
	return smarq
}

// First returns the first SysMenuActionResource entity from the query.
// Returns a *NotFoundError when no SysMenuActionResource was found.
func (smarq *SysMenuActionResourceQuery) First(ctx context.Context) (*SysMenuActionResource, error) {
	nodes, err := smarq.Limit(1).All(setContextOp(ctx, smarq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysmenuactionresource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) FirstX(ctx context.Context) *SysMenuActionResource {
	node, err := smarq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysMenuActionResource ID from the query.
// Returns a *NotFoundError when no SysMenuActionResource ID was found.
func (smarq *SysMenuActionResourceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smarq.Limit(1).IDs(setContextOp(ctx, smarq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysmenuactionresource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) FirstIDX(ctx context.Context) string {
	id, err := smarq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysMenuActionResource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysMenuActionResource entity is found.
// Returns a *NotFoundError when no SysMenuActionResource entities are found.
func (smarq *SysMenuActionResourceQuery) Only(ctx context.Context) (*SysMenuActionResource, error) {
	nodes, err := smarq.Limit(2).All(setContextOp(ctx, smarq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysmenuactionresource.Label}
	default:
		return nil, &NotSingularError{sysmenuactionresource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) OnlyX(ctx context.Context) *SysMenuActionResource {
	node, err := smarq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysMenuActionResource ID in the query.
// Returns a *NotSingularError when more than one SysMenuActionResource ID is found.
// Returns a *NotFoundError when no entities are found.
func (smarq *SysMenuActionResourceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smarq.Limit(2).IDs(setContextOp(ctx, smarq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = &NotSingularError{sysmenuactionresource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) OnlyIDX(ctx context.Context) string {
	id, err := smarq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysMenuActionResources.
func (smarq *SysMenuActionResourceQuery) All(ctx context.Context) ([]*SysMenuActionResource, error) {
	ctx = setContextOp(ctx, smarq.ctx, "All")
	if err := smarq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysMenuActionResource, *SysMenuActionResourceQuery]()
	return withInterceptors[[]*SysMenuActionResource](ctx, smarq, qr, smarq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) AllX(ctx context.Context) []*SysMenuActionResource {
	nodes, err := smarq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysMenuActionResource IDs.
func (smarq *SysMenuActionResourceQuery) IDs(ctx context.Context) (ids []string, err error) {
	if smarq.ctx.Unique == nil && smarq.path != nil {
		smarq.Unique(true)
	}
	ctx = setContextOp(ctx, smarq.ctx, "IDs")
	if err = smarq.Select(sysmenuactionresource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) IDsX(ctx context.Context) []string {
	ids, err := smarq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smarq *SysMenuActionResourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smarq.ctx, "Count")
	if err := smarq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smarq, querierCount[*SysMenuActionResourceQuery](), smarq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) CountX(ctx context.Context) int {
	count, err := smarq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smarq *SysMenuActionResourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smarq.ctx, "Exist")
	switch _, err := smarq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entscheme: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smarq *SysMenuActionResourceQuery) ExistX(ctx context.Context) bool {
	exist, err := smarq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysMenuActionResourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smarq *SysMenuActionResourceQuery) Clone() *SysMenuActionResourceQuery {
	if smarq == nil {
		return nil
	}
	return &SysMenuActionResourceQuery{
		config:     smarq.config,
		ctx:        smarq.ctx.Clone(),
		order:      append([]sysmenuactionresource.OrderOption{}, smarq.order...),
		inters:     append([]Interceptor{}, smarq.inters...),
		predicates: append([]predicate.SysMenuActionResource{}, smarq.predicates...),
		// clone intermediate query.
		sql:  smarq.sql.Clone(),
		path: smarq.path,
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
//	client.SysMenuActionResource.Query().
//		GroupBy(sysmenuactionresource.FieldIsDel).
//		Aggregate(entscheme.Count()).
//		Scan(ctx, &v)
func (smarq *SysMenuActionResourceQuery) GroupBy(field string, fields ...string) *SysMenuActionResourceGroupBy {
	smarq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysMenuActionResourceGroupBy{build: smarq}
	grbuild.flds = &smarq.ctx.Fields
	grbuild.label = sysmenuactionresource.Label
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
//	client.SysMenuActionResource.Query().
//		Select(sysmenuactionresource.FieldIsDel).
//		Scan(ctx, &v)
func (smarq *SysMenuActionResourceQuery) Select(fields ...string) *SysMenuActionResourceSelect {
	smarq.ctx.Fields = append(smarq.ctx.Fields, fields...)
	sbuild := &SysMenuActionResourceSelect{SysMenuActionResourceQuery: smarq}
	sbuild.label = sysmenuactionresource.Label
	sbuild.flds, sbuild.scan = &smarq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysMenuActionResourceSelect configured with the given aggregations.
func (smarq *SysMenuActionResourceQuery) Aggregate(fns ...AggregateFunc) *SysMenuActionResourceSelect {
	return smarq.Select().Aggregate(fns...)
}

func (smarq *SysMenuActionResourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smarq.inters {
		if inter == nil {
			return fmt.Errorf("entscheme: uninitialized interceptor (forgotten import entscheme/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smarq); err != nil {
				return err
			}
		}
	}
	for _, f := range smarq.ctx.Fields {
		if !sysmenuactionresource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entscheme: invalid field %q for query", f)}
		}
	}
	if smarq.path != nil {
		prev, err := smarq.path(ctx)
		if err != nil {
			return err
		}
		smarq.sql = prev
	}
	return nil
}

func (smarq *SysMenuActionResourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysMenuActionResource, error) {
	var (
		nodes = []*SysMenuActionResource{}
		_spec = smarq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysMenuActionResource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysMenuActionResource{config: smarq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(smarq.modifiers) > 0 {
		_spec.Modifiers = smarq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smarq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (smarq *SysMenuActionResourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smarq.querySpec()
	if len(smarq.modifiers) > 0 {
		_spec.Modifiers = smarq.modifiers
	}
	_spec.Node.Columns = smarq.ctx.Fields
	if len(smarq.ctx.Fields) > 0 {
		_spec.Unique = smarq.ctx.Unique != nil && *smarq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smarq.driver, _spec)
}

func (smarq *SysMenuActionResourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysmenuactionresource.Table, sysmenuactionresource.Columns, sqlgraph.NewFieldSpec(sysmenuactionresource.FieldID, field.TypeString))
	_spec.From = smarq.sql
	if unique := smarq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smarq.path != nil {
		_spec.Unique = true
	}
	if fields := smarq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenuactionresource.FieldID)
		for i := range fields {
			if fields[i] != sysmenuactionresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smarq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smarq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smarq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smarq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smarq *SysMenuActionResourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smarq.driver.Dialect())
	t1 := builder.Table(sysmenuactionresource.Table)
	columns := smarq.ctx.Fields
	if len(columns) == 0 {
		columns = sysmenuactionresource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smarq.sql != nil {
		selector = smarq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smarq.ctx.Unique != nil && *smarq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range smarq.modifiers {
		m(selector)
	}
	for _, p := range smarq.predicates {
		p(selector)
	}
	for _, p := range smarq.order {
		p(selector)
	}
	if offset := smarq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smarq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (smarq *SysMenuActionResourceQuery) ForUpdate(opts ...sql.LockOption) *SysMenuActionResourceQuery {
	if smarq.driver.Dialect() == dialect.Postgres {
		smarq.Unique(false)
	}
	smarq.modifiers = append(smarq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return smarq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (smarq *SysMenuActionResourceQuery) ForShare(opts ...sql.LockOption) *SysMenuActionResourceQuery {
	if smarq.driver.Dialect() == dialect.Postgres {
		smarq.Unique(false)
	}
	smarq.modifiers = append(smarq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return smarq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (smarq *SysMenuActionResourceQuery) Modify(modifiers ...func(s *sql.Selector)) *SysMenuActionResourceSelect {
	smarq.modifiers = append(smarq.modifiers, modifiers...)
	return smarq.Select()
}

// SysMenuActionResourceGroupBy is the group-by builder for SysMenuActionResource entities.
type SysMenuActionResourceGroupBy struct {
	selector
	build *SysMenuActionResourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smargb *SysMenuActionResourceGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuActionResourceGroupBy {
	smargb.fns = append(smargb.fns, fns...)
	return smargb
}

// Scan applies the selector query and scans the result into the given value.
func (smargb *SysMenuActionResourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smargb.build.ctx, "GroupBy")
	if err := smargb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysMenuActionResourceQuery, *SysMenuActionResourceGroupBy](ctx, smargb.build, smargb, smargb.build.inters, v)
}

func (smargb *SysMenuActionResourceGroupBy) sqlScan(ctx context.Context, root *SysMenuActionResourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smargb.fns))
	for _, fn := range smargb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smargb.flds)+len(smargb.fns))
		for _, f := range *smargb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smargb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smargb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysMenuActionResourceSelect is the builder for selecting fields of SysMenuActionResource entities.
type SysMenuActionResourceSelect struct {
	*SysMenuActionResourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (smars *SysMenuActionResourceSelect) Aggregate(fns ...AggregateFunc) *SysMenuActionResourceSelect {
	smars.fns = append(smars.fns, fns...)
	return smars
}

// Scan applies the selector query and scans the result into the given value.
func (smars *SysMenuActionResourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smars.ctx, "Select")
	if err := smars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysMenuActionResourceQuery, *SysMenuActionResourceSelect](ctx, smars.SysMenuActionResourceQuery, smars, smars.inters, v)
}

func (smars *SysMenuActionResourceSelect) sqlScan(ctx context.Context, root *SysMenuActionResourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(smars.fns))
	for _, fn := range smars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*smars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (smars *SysMenuActionResourceSelect) Modify(modifiers ...func(s *sql.Selector)) *SysMenuActionResourceSelect {
	smars.modifiers = append(smars.modifiers, modifiers...)
	return smars
}