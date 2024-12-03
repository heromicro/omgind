// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenu"
)

// SysMenuQuery is the builder for querying SysMenu entities.
type SysMenuQuery struct {
	config
	ctx          *QueryContext
	order        []sysmenu.OrderOption
	inters       []Interceptor
	predicates   []predicate.SysMenu
	withParent   *SysMenuQuery
	withChildren *SysMenuQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuQuery builder.
func (smq *SysMenuQuery) Where(ps ...predicate.SysMenu) *SysMenuQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit the number of records to be returned by this query.
func (smq *SysMenuQuery) Limit(limit int) *SysMenuQuery {
	smq.ctx.Limit = &limit
	return smq
}

// Offset to start from.
func (smq *SysMenuQuery) Offset(offset int) *SysMenuQuery {
	smq.ctx.Offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *SysMenuQuery) Unique(unique bool) *SysMenuQuery {
	smq.ctx.Unique = &unique
	return smq
}

// Order specifies how the records should be ordered.
func (smq *SysMenuQuery) Order(o ...sysmenu.OrderOption) *SysMenuQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// QueryParent chains the current query on the "parent" edge.
func (smq *SysMenuQuery) QueryParent() *SysMenuQuery {
	query := (&SysMenuClient{config: smq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysmenu.ParentTable, sysmenu.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (smq *SysMenuQuery) QueryChildren() *SysMenuQuery {
	query := (&SysMenuClient{config: smq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysmenu.ChildrenTable, sysmenu.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysMenu entity from the query.
// Returns a *NotFoundError when no SysMenu was found.
func (smq *SysMenuQuery) First(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(1).All(setContextOp(ctx, smq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysmenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *SysMenuQuery) FirstX(ctx context.Context) *SysMenu {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysMenu ID from the query.
// Returns a *NotFoundError when no SysMenu ID was found.
func (smq *SysMenuQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smq.Limit(1).IDs(setContextOp(ctx, smq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysmenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *SysMenuQuery) FirstIDX(ctx context.Context) string {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysMenu entity is found.
// Returns a *NotFoundError when no SysMenu entities are found.
func (smq *SysMenuQuery) Only(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(2).All(setContextOp(ctx, smq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysmenu.Label}
	default:
		return nil, &NotSingularError{sysmenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyX(ctx context.Context) *SysMenu {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysMenu ID in the query.
// Returns a *NotSingularError when more than one SysMenu ID is found.
// Returns a *NotFoundError when no entities are found.
func (smq *SysMenuQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smq.Limit(2).IDs(setContextOp(ctx, smq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = &NotSingularError{sysmenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyIDX(ctx context.Context) string {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysMenus.
func (smq *SysMenuQuery) All(ctx context.Context) ([]*SysMenu, error) {
	ctx = setContextOp(ctx, smq.ctx, ent.OpQueryAll)
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysMenu, *SysMenuQuery]()
	return withInterceptors[[]*SysMenu](ctx, smq, qr, smq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smq *SysMenuQuery) AllX(ctx context.Context) []*SysMenu {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysMenu IDs.
func (smq *SysMenuQuery) IDs(ctx context.Context) (ids []string, err error) {
	if smq.ctx.Unique == nil && smq.path != nil {
		smq.Unique(true)
	}
	ctx = setContextOp(ctx, smq.ctx, ent.OpQueryIDs)
	if err = smq.Select(sysmenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *SysMenuQuery) IDsX(ctx context.Context) []string {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *SysMenuQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smq.ctx, ent.OpQueryCount)
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smq, querierCount[*SysMenuQuery](), smq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smq *SysMenuQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *SysMenuQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smq.ctx, ent.OpQueryExist)
	switch _, err := smq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("mainent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *SysMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *SysMenuQuery) Clone() *SysMenuQuery {
	if smq == nil {
		return nil
	}
	return &SysMenuQuery{
		config:       smq.config,
		ctx:          smq.ctx.Clone(),
		order:        append([]sysmenu.OrderOption{}, smq.order...),
		inters:       append([]Interceptor{}, smq.inters...),
		predicates:   append([]predicate.SysMenu{}, smq.predicates...),
		withParent:   smq.withParent.Clone(),
		withChildren: smq.withChildren.Clone(),
		// clone intermediate query.
		sql:       smq.sql.Clone(),
		path:      smq.path,
		modifiers: append([]func(*sql.Selector){}, smq.modifiers...),
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithParent(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := (&SysMenuClient{config: smq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	smq.withParent = query
	return smq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithChildren(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := (&SysMenuClient{config: smq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	smq.withChildren = query
	return smq
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
//	client.SysMenu.Query().
//		GroupBy(sysmenu.FieldIsDel).
//		Aggregate(mainent.Count()).
//		Scan(ctx, &v)
func (smq *SysMenuQuery) GroupBy(field string, fields ...string) *SysMenuGroupBy {
	smq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysMenuGroupBy{build: smq}
	grbuild.flds = &smq.ctx.Fields
	grbuild.label = sysmenu.Label
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
//	client.SysMenu.Query().
//		Select(sysmenu.FieldIsDel).
//		Scan(ctx, &v)
func (smq *SysMenuQuery) Select(fields ...string) *SysMenuSelect {
	smq.ctx.Fields = append(smq.ctx.Fields, fields...)
	sbuild := &SysMenuSelect{SysMenuQuery: smq}
	sbuild.label = sysmenu.Label
	sbuild.flds, sbuild.scan = &smq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysMenuSelect configured with the given aggregations.
func (smq *SysMenuQuery) Aggregate(fns ...AggregateFunc) *SysMenuSelect {
	return smq.Select().Aggregate(fns...)
}

func (smq *SysMenuQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smq.inters {
		if inter == nil {
			return fmt.Errorf("mainent: uninitialized interceptor (forgotten import mainent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smq); err != nil {
				return err
			}
		}
	}
	for _, f := range smq.ctx.Fields {
		if !sysmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *SysMenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysMenu, error) {
	var (
		nodes       = []*SysMenu{}
		_spec       = smq.querySpec()
		loadedTypes = [2]bool{
			smq.withParent != nil,
			smq.withChildren != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysMenu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysMenu{config: smq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(smq.modifiers) > 0 {
		_spec.Modifiers = smq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := smq.withParent; query != nil {
		if err := smq.loadParent(ctx, query, nodes, nil,
			func(n *SysMenu, e *SysMenu) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := smq.withChildren; query != nil {
		if err := smq.loadChildren(ctx, query, nodes,
			func(n *SysMenu) { n.Edges.Children = []*SysMenu{} },
			func(n *SysMenu, e *SysMenu) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (smq *SysMenuQuery) loadParent(ctx context.Context, query *SysMenuQuery, nodes []*SysMenu, init func(*SysMenu), assign func(*SysMenu, *SysMenu)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*SysMenu)
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
	query.Where(sysmenu.IDIn(ids...))
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
func (smq *SysMenuQuery) loadChildren(ctx context.Context, query *SysMenuQuery, nodes []*SysMenu, init func(*SysMenu), assign func(*SysMenu, *SysMenu)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysMenu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sysmenu.FieldParentID)
	}
	query.Where(predicate.SysMenu(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sysmenu.ChildrenColumn), fks...))
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
			return fmt.Errorf(`unexpected referenced foreign-key "parent_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (smq *SysMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	if len(smq.modifiers) > 0 {
		_spec.Modifiers = smq.modifiers
	}
	_spec.Node.Columns = smq.ctx.Fields
	if len(smq.ctx.Fields) > 0 {
		_spec.Unique = smq.ctx.Unique != nil && *smq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *SysMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysmenu.Table, sysmenu.Columns, sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeString))
	_spec.From = smq.sql
	if unique := smq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smq.path != nil {
		_spec.Unique = true
	}
	if fields := smq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.FieldID)
		for i := range fields {
			if fields[i] != sysmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if smq.withParent != nil {
			_spec.Node.AddColumnOnce(sysmenu.FieldParentID)
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *SysMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(sysmenu.Table)
	columns := smq.ctx.Fields
	if len(columns) == 0 {
		columns = sysmenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.ctx.Unique != nil && *smq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range smq.modifiers {
		m(selector)
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (smq *SysMenuQuery) ForUpdate(opts ...sql.LockOption) *SysMenuQuery {
	if smq.driver.Dialect() == dialect.Postgres {
		smq.Unique(false)
	}
	smq.modifiers = append(smq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return smq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (smq *SysMenuQuery) ForShare(opts ...sql.LockOption) *SysMenuQuery {
	if smq.driver.Dialect() == dialect.Postgres {
		smq.Unique(false)
	}
	smq.modifiers = append(smq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return smq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (smq *SysMenuQuery) Modify(modifiers ...func(s *sql.Selector)) *SysMenuSelect {
	smq.modifiers = append(smq.modifiers, modifiers...)
	return smq.Select()
}

// SysMenuGroupBy is the group-by builder for SysMenu entities.
type SysMenuGroupBy struct {
	selector
	build *SysMenuQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *SysMenuGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the selector query and scans the result into the given value.
func (smgb *SysMenuGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smgb.build.ctx, ent.OpQueryGroupBy)
	if err := smgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysMenuQuery, *SysMenuGroupBy](ctx, smgb.build, smgb, smgb.build.inters, v)
}

func (smgb *SysMenuGroupBy) sqlScan(ctx context.Context, root *SysMenuQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smgb.flds)+len(smgb.fns))
		for _, f := range *smgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysMenuSelect is the builder for selecting fields of SysMenu entities.
type SysMenuSelect struct {
	*SysMenuQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sms *SysMenuSelect) Aggregate(fns ...AggregateFunc) *SysMenuSelect {
	sms.fns = append(sms.fns, fns...)
	return sms
}

// Scan applies the selector query and scans the result into the given value.
func (sms *SysMenuSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sms.ctx, ent.OpQuerySelect)
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysMenuQuery, *SysMenuSelect](ctx, sms.SysMenuQuery, sms, sms.inters, v)
}

func (sms *SysMenuSelect) sqlScan(ctx context.Context, root *SysMenuQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sms.fns))
	for _, fn := range sms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sms *SysMenuSelect) Modify(modifiers ...func(s *sql.Selector)) *SysMenuSelect {
	sms.modifiers = append(sms.modifiers, modifiers...)
	return sms
}
