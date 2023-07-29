// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysannex"
)

// SysAnnexQuery is the builder for querying SysAnnex entities.
type SysAnnexQuery struct {
	config
	ctx        *QueryContext
	order      []sysannex.OrderOption
	inters     []Interceptor
	predicates []predicate.SysAnnex
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysAnnexQuery builder.
func (saq *SysAnnexQuery) Where(ps ...predicate.SysAnnex) *SysAnnexQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit the number of records to be returned by this query.
func (saq *SysAnnexQuery) Limit(limit int) *SysAnnexQuery {
	saq.ctx.Limit = &limit
	return saq
}

// Offset to start from.
func (saq *SysAnnexQuery) Offset(offset int) *SysAnnexQuery {
	saq.ctx.Offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SysAnnexQuery) Unique(unique bool) *SysAnnexQuery {
	saq.ctx.Unique = &unique
	return saq
}

// Order specifies how the records should be ordered.
func (saq *SysAnnexQuery) Order(o ...sysannex.OrderOption) *SysAnnexQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// First returns the first SysAnnex entity from the query.
// Returns a *NotFoundError when no SysAnnex was found.
func (saq *SysAnnexQuery) First(ctx context.Context) (*SysAnnex, error) {
	nodes, err := saq.Limit(1).All(setContextOp(ctx, saq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysannex.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SysAnnexQuery) FirstX(ctx context.Context) *SysAnnex {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysAnnex ID from the query.
// Returns a *NotFoundError when no SysAnnex ID was found.
func (saq *SysAnnexQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(1).IDs(setContextOp(ctx, saq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysannex.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SysAnnexQuery) FirstIDX(ctx context.Context) string {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysAnnex entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysAnnex entity is found.
// Returns a *NotFoundError when no SysAnnex entities are found.
func (saq *SysAnnexQuery) Only(ctx context.Context) (*SysAnnex, error) {
	nodes, err := saq.Limit(2).All(setContextOp(ctx, saq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysannex.Label}
	default:
		return nil, &NotSingularError{sysannex.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SysAnnexQuery) OnlyX(ctx context.Context) *SysAnnex {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysAnnex ID in the query.
// Returns a *NotSingularError when more than one SysAnnex ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SysAnnexQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(2).IDs(setContextOp(ctx, saq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysannex.Label}
	default:
		err = &NotSingularError{sysannex.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SysAnnexQuery) OnlyIDX(ctx context.Context) string {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysAnnexes.
func (saq *SysAnnexQuery) All(ctx context.Context) ([]*SysAnnex, error) {
	ctx = setContextOp(ctx, saq.ctx, "All")
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysAnnex, *SysAnnexQuery]()
	return withInterceptors[[]*SysAnnex](ctx, saq, qr, saq.inters)
}

// AllX is like All, but panics if an error occurs.
func (saq *SysAnnexQuery) AllX(ctx context.Context) []*SysAnnex {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysAnnex IDs.
func (saq *SysAnnexQuery) IDs(ctx context.Context) (ids []string, err error) {
	if saq.ctx.Unique == nil && saq.path != nil {
		saq.Unique(true)
	}
	ctx = setContextOp(ctx, saq.ctx, "IDs")
	if err = saq.Select(sysannex.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SysAnnexQuery) IDsX(ctx context.Context) []string {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SysAnnexQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, saq.ctx, "Count")
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, saq, querierCount[*SysAnnexQuery](), saq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SysAnnexQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SysAnnexQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, saq.ctx, "Exist")
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("mainent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *SysAnnexQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysAnnexQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SysAnnexQuery) Clone() *SysAnnexQuery {
	if saq == nil {
		return nil
	}
	return &SysAnnexQuery{
		config:     saq.config,
		ctx:        saq.ctx.Clone(),
		order:      append([]sysannex.OrderOption{}, saq.order...),
		inters:     append([]Interceptor{}, saq.inters...),
		predicates: append([]predicate.SysAnnex{}, saq.predicates...),
		// clone intermediate query.
		sql:  saq.sql.Clone(),
		path: saq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Memo string `json:"memo,omitempty" sql:"memo"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysAnnex.Query().
//		GroupBy(sysannex.FieldMemo).
//		Aggregate(mainent.Count()).
//		Scan(ctx, &v)
func (saq *SysAnnexQuery) GroupBy(field string, fields ...string) *SysAnnexGroupBy {
	saq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysAnnexGroupBy{build: saq}
	grbuild.flds = &saq.ctx.Fields
	grbuild.label = sysannex.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Memo string `json:"memo,omitempty" sql:"memo"`
//	}
//
//	client.SysAnnex.Query().
//		Select(sysannex.FieldMemo).
//		Scan(ctx, &v)
func (saq *SysAnnexQuery) Select(fields ...string) *SysAnnexSelect {
	saq.ctx.Fields = append(saq.ctx.Fields, fields...)
	sbuild := &SysAnnexSelect{SysAnnexQuery: saq}
	sbuild.label = sysannex.Label
	sbuild.flds, sbuild.scan = &saq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysAnnexSelect configured with the given aggregations.
func (saq *SysAnnexQuery) Aggregate(fns ...AggregateFunc) *SysAnnexSelect {
	return saq.Select().Aggregate(fns...)
}

func (saq *SysAnnexQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range saq.inters {
		if inter == nil {
			return fmt.Errorf("mainent: uninitialized interceptor (forgotten import mainent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, saq); err != nil {
				return err
			}
		}
	}
	for _, f := range saq.ctx.Fields {
		if !sysannex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *SysAnnexQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysAnnex, error) {
	var (
		nodes = []*SysAnnex{}
		_spec = saq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysAnnex).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysAnnex{config: saq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (saq *SysAnnexQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	_spec.Node.Columns = saq.ctx.Fields
	if len(saq.ctx.Fields) > 0 {
		_spec.Unique = saq.ctx.Unique != nil && *saq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *SysAnnexQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysannex.Table, sysannex.Columns, sqlgraph.NewFieldSpec(sysannex.FieldID, field.TypeString))
	_spec.From = saq.sql
	if unique := saq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if saq.path != nil {
		_spec.Unique = true
	}
	if fields := saq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysannex.FieldID)
		for i := range fields {
			if fields[i] != sysannex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *SysAnnexQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(sysannex.Table)
	columns := saq.ctx.Fields
	if len(columns) == 0 {
		columns = sysannex.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.ctx.Unique != nil && *saq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range saq.modifiers {
		m(selector)
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (saq *SysAnnexQuery) ForUpdate(opts ...sql.LockOption) *SysAnnexQuery {
	if saq.driver.Dialect() == dialect.Postgres {
		saq.Unique(false)
	}
	saq.modifiers = append(saq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return saq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (saq *SysAnnexQuery) ForShare(opts ...sql.LockOption) *SysAnnexQuery {
	if saq.driver.Dialect() == dialect.Postgres {
		saq.Unique(false)
	}
	saq.modifiers = append(saq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return saq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (saq *SysAnnexQuery) Modify(modifiers ...func(s *sql.Selector)) *SysAnnexSelect {
	saq.modifiers = append(saq.modifiers, modifiers...)
	return saq.Select()
}

// SysAnnexGroupBy is the group-by builder for SysAnnex entities.
type SysAnnexGroupBy struct {
	selector
	build *SysAnnexQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SysAnnexGroupBy) Aggregate(fns ...AggregateFunc) *SysAnnexGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the selector query and scans the result into the given value.
func (sagb *SysAnnexGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sagb.build.ctx, "GroupBy")
	if err := sagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysAnnexQuery, *SysAnnexGroupBy](ctx, sagb.build, sagb, sagb.build.inters, v)
}

func (sagb *SysAnnexGroupBy) sqlScan(ctx context.Context, root *SysAnnexQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sagb.flds)+len(sagb.fns))
		for _, f := range *sagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysAnnexSelect is the builder for selecting fields of SysAnnex entities.
type SysAnnexSelect struct {
	*SysAnnexQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sas *SysAnnexSelect) Aggregate(fns ...AggregateFunc) *SysAnnexSelect {
	sas.fns = append(sas.fns, fns...)
	return sas
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SysAnnexSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sas.ctx, "Select")
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysAnnexQuery, *SysAnnexSelect](ctx, sas.SysAnnexQuery, sas, sas.inters, v)
}

func (sas *SysAnnexSelect) sqlScan(ctx context.Context, root *SysAnnexQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sas.fns))
	for _, fn := range sas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sas *SysAnnexSelect) Modify(modifiers ...func(s *sql.Selector)) *SysAnnexSelect {
	sas.modifiers = append(sas.modifiers, modifiers...)
	return sas
}
