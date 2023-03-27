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
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
)

// SysDistrictQuery is the builder for querying SysDistrict entities.
type SysDistrictQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.SysDistrict
	withParent   *SysDistrictQuery
	withChildren *SysDistrictQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDistrictQuery builder.
func (sdq *SysDistrictQuery) Where(ps ...predicate.SysDistrict) *SysDistrictQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit the number of records to be returned by this query.
func (sdq *SysDistrictQuery) Limit(limit int) *SysDistrictQuery {
	sdq.ctx.Limit = &limit
	return sdq
}

// Offset to start from.
func (sdq *SysDistrictQuery) Offset(offset int) *SysDistrictQuery {
	sdq.ctx.Offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SysDistrictQuery) Unique(unique bool) *SysDistrictQuery {
	sdq.ctx.Unique = &unique
	return sdq
}

// Order specifies how the records should be ordered.
func (sdq *SysDistrictQuery) Order(o ...OrderFunc) *SysDistrictQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// QueryParent chains the current query on the "parent" edge.
func (sdq *SysDistrictQuery) QueryParent() *SysDistrictQuery {
	query := (&SysDistrictClient{config: sdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysdistrict.Table, sysdistrict.FieldID, selector),
			sqlgraph.To(sysdistrict.Table, sysdistrict.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysdistrict.ParentTable, sysdistrict.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (sdq *SysDistrictQuery) QueryChildren() *SysDistrictQuery {
	query := (&SysDistrictClient{config: sdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysdistrict.Table, sysdistrict.FieldID, selector),
			sqlgraph.To(sysdistrict.Table, sysdistrict.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysdistrict.ChildrenTable, sysdistrict.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysDistrict entity from the query.
// Returns a *NotFoundError when no SysDistrict was found.
func (sdq *SysDistrictQuery) First(ctx context.Context) (*SysDistrict, error) {
	nodes, err := sdq.Limit(1).All(setContextOp(ctx, sdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdistrict.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SysDistrictQuery) FirstX(ctx context.Context) *SysDistrict {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDistrict ID from the query.
// Returns a *NotFoundError when no SysDistrict ID was found.
func (sdq *SysDistrictQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdq.Limit(1).IDs(setContextOp(ctx, sdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdistrict.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SysDistrictQuery) FirstIDX(ctx context.Context) string {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDistrict entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysDistrict entity is found.
// Returns a *NotFoundError when no SysDistrict entities are found.
func (sdq *SysDistrictQuery) Only(ctx context.Context) (*SysDistrict, error) {
	nodes, err := sdq.Limit(2).All(setContextOp(ctx, sdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdistrict.Label}
	default:
		return nil, &NotSingularError{sysdistrict.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SysDistrictQuery) OnlyX(ctx context.Context) *SysDistrict {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDistrict ID in the query.
// Returns a *NotSingularError when more than one SysDistrict ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SysDistrictQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdq.Limit(2).IDs(setContextOp(ctx, sdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdistrict.Label}
	default:
		err = &NotSingularError{sysdistrict.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SysDistrictQuery) OnlyIDX(ctx context.Context) string {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDistricts.
func (sdq *SysDistrictQuery) All(ctx context.Context) ([]*SysDistrict, error) {
	ctx = setContextOp(ctx, sdq.ctx, "All")
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysDistrict, *SysDistrictQuery]()
	return withInterceptors[[]*SysDistrict](ctx, sdq, qr, sdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SysDistrictQuery) AllX(ctx context.Context) []*SysDistrict {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDistrict IDs.
func (sdq *SysDistrictQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sdq.ctx.Unique == nil && sdq.path != nil {
		sdq.Unique(true)
	}
	ctx = setContextOp(ctx, sdq.ctx, "IDs")
	if err = sdq.Select(sysdistrict.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SysDistrictQuery) IDsX(ctx context.Context) []string {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SysDistrictQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sdq.ctx, "Count")
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sdq, querierCount[*SysDistrictQuery](), sdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SysDistrictQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SysDistrictQuery) Exist(ctx context.Context) (bool, error) {
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
func (sdq *SysDistrictQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDistrictQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SysDistrictQuery) Clone() *SysDistrictQuery {
	if sdq == nil {
		return nil
	}
	return &SysDistrictQuery{
		config:       sdq.config,
		ctx:          sdq.ctx.Clone(),
		order:        append([]OrderFunc{}, sdq.order...),
		inters:       append([]Interceptor{}, sdq.inters...),
		predicates:   append([]predicate.SysDistrict{}, sdq.predicates...),
		withParent:   sdq.withParent.Clone(),
		withChildren: sdq.withChildren.Clone(),
		// clone intermediate query.
		sql:  sdq.sql.Clone(),
		path: sdq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SysDistrictQuery) WithParent(opts ...func(*SysDistrictQuery)) *SysDistrictQuery {
	query := (&SysDistrictClient{config: sdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sdq.withParent = query
	return sdq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SysDistrictQuery) WithChildren(opts ...func(*SysDistrictQuery)) *SysDistrictQuery {
	query := (&SysDistrictClient{config: sdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sdq.withChildren = query
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
//	client.SysDistrict.Query().
//		GroupBy(sysdistrict.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SysDistrictQuery) GroupBy(field string, fields ...string) *SysDistrictGroupBy {
	sdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysDistrictGroupBy{build: sdq}
	grbuild.flds = &sdq.ctx.Fields
	grbuild.label = sysdistrict.Label
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
//	client.SysDistrict.Query().
//		Select(sysdistrict.FieldIsDel).
//		Scan(ctx, &v)
func (sdq *SysDistrictQuery) Select(fields ...string) *SysDistrictSelect {
	sdq.ctx.Fields = append(sdq.ctx.Fields, fields...)
	sbuild := &SysDistrictSelect{SysDistrictQuery: sdq}
	sbuild.label = sysdistrict.Label
	sbuild.flds, sbuild.scan = &sdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysDistrictSelect configured with the given aggregations.
func (sdq *SysDistrictQuery) Aggregate(fns ...AggregateFunc) *SysDistrictSelect {
	return sdq.Select().Aggregate(fns...)
}

func (sdq *SysDistrictQuery) prepareQuery(ctx context.Context) error {
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
		if !sysdistrict.ValidColumn(f) {
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

func (sdq *SysDistrictQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysDistrict, error) {
	var (
		nodes       = []*SysDistrict{}
		_spec       = sdq.querySpec()
		loadedTypes = [2]bool{
			sdq.withParent != nil,
			sdq.withChildren != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysDistrict).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysDistrict{config: sdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
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
	if query := sdq.withParent; query != nil {
		if err := sdq.loadParent(ctx, query, nodes, nil,
			func(n *SysDistrict, e *SysDistrict) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := sdq.withChildren; query != nil {
		if err := sdq.loadChildren(ctx, query, nodes,
			func(n *SysDistrict) { n.Edges.Children = []*SysDistrict{} },
			func(n *SysDistrict, e *SysDistrict) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sdq *SysDistrictQuery) loadParent(ctx context.Context, query *SysDistrictQuery, nodes []*SysDistrict, init func(*SysDistrict), assign func(*SysDistrict, *SysDistrict)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*SysDistrict)
	for i := range nodes {
		if nodes[i].ParentID == nil {
			continue
		}
		fk := *nodes[i].ParentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sysdistrict.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sdq *SysDistrictQuery) loadChildren(ctx context.Context, query *SysDistrictQuery, nodes []*SysDistrict, init func(*SysDistrict), assign func(*SysDistrict, *SysDistrict)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysDistrict)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.SysDistrict(func(s *sql.Selector) {
		s.Where(sql.InValues(sysdistrict.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		if fk == nil {
			return fmt.Errorf(`foreign-key "parent_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sdq *SysDistrictQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	_spec.Node.Columns = sdq.ctx.Fields
	if len(sdq.ctx.Fields) > 0 {
		_spec.Unique = sdq.ctx.Unique != nil && *sdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SysDistrictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysdistrict.Table, sysdistrict.Columns, sqlgraph.NewFieldSpec(sysdistrict.FieldID, field.TypeString))
	_spec.From = sdq.sql
	if unique := sdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sdq.path != nil {
		_spec.Unique = true
	}
	if fields := sdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdistrict.FieldID)
		for i := range fields {
			if fields[i] != sysdistrict.FieldID {
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

func (sdq *SysDistrictQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(sysdistrict.Table)
	columns := sdq.ctx.Fields
	if len(columns) == 0 {
		columns = sysdistrict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.ctx.Unique != nil && *sdq.ctx.Unique {
		selector.Distinct()
	}
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
func (sdq *SysDistrictQuery) ForUpdate(opts ...sql.LockOption) *SysDistrictQuery {
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
func (sdq *SysDistrictQuery) ForShare(opts ...sql.LockOption) *SysDistrictQuery {
	if sdq.driver.Dialect() == dialect.Postgres {
		sdq.Unique(false)
	}
	sdq.modifiers = append(sdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sdq *SysDistrictQuery) Modify(modifiers ...func(s *sql.Selector)) *SysDistrictSelect {
	sdq.modifiers = append(sdq.modifiers, modifiers...)
	return sdq.Select()
}

// SysDistrictGroupBy is the group-by builder for SysDistrict entities.
type SysDistrictGroupBy struct {
	selector
	build *SysDistrictQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SysDistrictGroupBy) Aggregate(fns ...AggregateFunc) *SysDistrictGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the selector query and scans the result into the given value.
func (sdgb *SysDistrictGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdgb.build.ctx, "GroupBy")
	if err := sdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDistrictQuery, *SysDistrictGroupBy](ctx, sdgb.build, sdgb, sdgb.build.inters, v)
}

func (sdgb *SysDistrictGroupBy) sqlScan(ctx context.Context, root *SysDistrictQuery, v any) error {
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

// SysDistrictSelect is the builder for selecting fields of SysDistrict entities.
type SysDistrictSelect struct {
	*SysDistrictQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sds *SysDistrictSelect) Aggregate(fns ...AggregateFunc) *SysDistrictSelect {
	sds.fns = append(sds.fns, fns...)
	return sds
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SysDistrictSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sds.ctx, "Select")
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDistrictQuery, *SysDistrictSelect](ctx, sds.SysDistrictQuery, sds, sds.inters, v)
}

func (sds *SysDistrictSelect) sqlScan(ctx context.Context, root *SysDistrictQuery, v any) error {
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
func (sds *SysDistrictSelect) Modify(modifiers ...func(s *sql.Selector)) *SysDistrictSelect {
	sds.modifiers = append(sds.modifiers, modifiers...)
	return sds
}
