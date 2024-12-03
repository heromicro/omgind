// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/orgdept"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgposition"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysaddress"
)

// OrgStaffQuery is the builder for querying OrgStaff entities.
type OrgStaffQuery struct {
	config
	ctx          *QueryContext
	order        []orgstaff.OrderOption
	inters       []Interceptor
	predicates   []predicate.OrgStaff
	withOrgan    *OrgOrganQuery
	withIdenAddr *SysAddressQuery
	withResiAddr *SysAddressQuery
	withDept     *OrgDeptQuery
	withPosi     *OrgPositionQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgStaffQuery builder.
func (osq *OrgStaffQuery) Where(ps ...predicate.OrgStaff) *OrgStaffQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit the number of records to be returned by this query.
func (osq *OrgStaffQuery) Limit(limit int) *OrgStaffQuery {
	osq.ctx.Limit = &limit
	return osq
}

// Offset to start from.
func (osq *OrgStaffQuery) Offset(offset int) *OrgStaffQuery {
	osq.ctx.Offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrgStaffQuery) Unique(unique bool) *OrgStaffQuery {
	osq.ctx.Unique = &unique
	return osq
}

// Order specifies how the records should be ordered.
func (osq *OrgStaffQuery) Order(o ...orgstaff.OrderOption) *OrgStaffQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryOrgan chains the current query on the "organ" edge.
func (osq *OrgStaffQuery) QueryOrgan() *OrgOrganQuery {
	query := (&OrgOrganClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgstaff.Table, orgstaff.FieldID, selector),
			sqlgraph.To(orgorgan.Table, orgorgan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgstaff.OrganTable, orgstaff.OrganColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIdenAddr chains the current query on the "iden_addr" edge.
func (osq *OrgStaffQuery) QueryIdenAddr() *SysAddressQuery {
	query := (&SysAddressClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgstaff.Table, orgstaff.FieldID, selector),
			sqlgraph.To(sysaddress.Table, sysaddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, orgstaff.IdenAddrTable, orgstaff.IdenAddrColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResiAddr chains the current query on the "resi_addr" edge.
func (osq *OrgStaffQuery) QueryResiAddr() *SysAddressQuery {
	query := (&SysAddressClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgstaff.Table, orgstaff.FieldID, selector),
			sqlgraph.To(sysaddress.Table, sysaddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, orgstaff.ResiAddrTable, orgstaff.ResiAddrColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDept chains the current query on the "dept" edge.
func (osq *OrgStaffQuery) QueryDept() *OrgDeptQuery {
	query := (&OrgDeptClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgstaff.Table, orgstaff.FieldID, selector),
			sqlgraph.To(orgdept.Table, orgdept.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgstaff.DeptTable, orgstaff.DeptColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosi chains the current query on the "posi" edge.
func (osq *OrgStaffQuery) QueryPosi() *OrgPositionQuery {
	query := (&OrgPositionClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgstaff.Table, orgstaff.FieldID, selector),
			sqlgraph.To(orgposition.Table, orgposition.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgstaff.PosiTable, orgstaff.PosiColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgStaff entity from the query.
// Returns a *NotFoundError when no OrgStaff was found.
func (osq *OrgStaffQuery) First(ctx context.Context) (*OrgStaff, error) {
	nodes, err := osq.Limit(1).All(setContextOp(ctx, osq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgstaff.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrgStaffQuery) FirstX(ctx context.Context) *OrgStaff {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgStaff ID from the query.
// Returns a *NotFoundError when no OrgStaff ID was found.
func (osq *OrgStaffQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(1).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgstaff.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrgStaffQuery) FirstIDX(ctx context.Context) string {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgStaff entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgStaff entity is found.
// Returns a *NotFoundError when no OrgStaff entities are found.
func (osq *OrgStaffQuery) Only(ctx context.Context) (*OrgStaff, error) {
	nodes, err := osq.Limit(2).All(setContextOp(ctx, osq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgstaff.Label}
	default:
		return nil, &NotSingularError{orgstaff.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrgStaffQuery) OnlyX(ctx context.Context) *OrgStaff {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgStaff ID in the query.
// Returns a *NotSingularError when more than one OrgStaff ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrgStaffQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = osq.Limit(2).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgstaff.Label}
	default:
		err = &NotSingularError{orgstaff.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrgStaffQuery) OnlyIDX(ctx context.Context) string {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgStaffs.
func (osq *OrgStaffQuery) All(ctx context.Context) ([]*OrgStaff, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryAll)
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgStaff, *OrgStaffQuery]()
	return withInterceptors[[]*OrgStaff](ctx, osq, qr, osq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrgStaffQuery) AllX(ctx context.Context) []*OrgStaff {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgStaff IDs.
func (osq *OrgStaffQuery) IDs(ctx context.Context) (ids []string, err error) {
	if osq.ctx.Unique == nil && osq.path != nil {
		osq.Unique(true)
	}
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryIDs)
	if err = osq.Select(orgstaff.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrgStaffQuery) IDsX(ctx context.Context) []string {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrgStaffQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryCount)
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osq, querierCount[*OrgStaffQuery](), osq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrgStaffQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrgStaffQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryExist)
	switch _, err := osq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("mainent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrgStaffQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgStaffQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrgStaffQuery) Clone() *OrgStaffQuery {
	if osq == nil {
		return nil
	}
	return &OrgStaffQuery{
		config:       osq.config,
		ctx:          osq.ctx.Clone(),
		order:        append([]orgstaff.OrderOption{}, osq.order...),
		inters:       append([]Interceptor{}, osq.inters...),
		predicates:   append([]predicate.OrgStaff{}, osq.predicates...),
		withOrgan:    osq.withOrgan.Clone(),
		withIdenAddr: osq.withIdenAddr.Clone(),
		withResiAddr: osq.withResiAddr.Clone(),
		withDept:     osq.withDept.Clone(),
		withPosi:     osq.withPosi.Clone(),
		// clone intermediate query.
		sql:       osq.sql.Clone(),
		path:      osq.path,
		modifiers: append([]func(*sql.Selector){}, osq.modifiers...),
	}
}

// WithOrgan tells the query-builder to eager-load the nodes that are connected to
// the "organ" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrgStaffQuery) WithOrgan(opts ...func(*OrgOrganQuery)) *OrgStaffQuery {
	query := (&OrgOrganClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrgan = query
	return osq
}

// WithIdenAddr tells the query-builder to eager-load the nodes that are connected to
// the "iden_addr" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrgStaffQuery) WithIdenAddr(opts ...func(*SysAddressQuery)) *OrgStaffQuery {
	query := (&SysAddressClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withIdenAddr = query
	return osq
}

// WithResiAddr tells the query-builder to eager-load the nodes that are connected to
// the "resi_addr" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrgStaffQuery) WithResiAddr(opts ...func(*SysAddressQuery)) *OrgStaffQuery {
	query := (&SysAddressClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withResiAddr = query
	return osq
}

// WithDept tells the query-builder to eager-load the nodes that are connected to
// the "dept" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrgStaffQuery) WithDept(opts ...func(*OrgDeptQuery)) *OrgStaffQuery {
	query := (&OrgDeptClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withDept = query
	return osq
}

// WithPosi tells the query-builder to eager-load the nodes that are connected to
// the "posi" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrgStaffQuery) WithPosi(opts ...func(*OrgPositionQuery)) *OrgStaffQuery {
	query := (&OrgPositionClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withPosi = query
	return osq
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
//	client.OrgStaff.Query().
//		GroupBy(orgstaff.FieldIsDel).
//		Aggregate(mainent.Count()).
//		Scan(ctx, &v)
func (osq *OrgStaffQuery) GroupBy(field string, fields ...string) *OrgStaffGroupBy {
	osq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgStaffGroupBy{build: osq}
	grbuild.flds = &osq.ctx.Fields
	grbuild.label = orgstaff.Label
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
//	client.OrgStaff.Query().
//		Select(orgstaff.FieldIsDel).
//		Scan(ctx, &v)
func (osq *OrgStaffQuery) Select(fields ...string) *OrgStaffSelect {
	osq.ctx.Fields = append(osq.ctx.Fields, fields...)
	sbuild := &OrgStaffSelect{OrgStaffQuery: osq}
	sbuild.label = orgstaff.Label
	sbuild.flds, sbuild.scan = &osq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgStaffSelect configured with the given aggregations.
func (osq *OrgStaffQuery) Aggregate(fns ...AggregateFunc) *OrgStaffSelect {
	return osq.Select().Aggregate(fns...)
}

func (osq *OrgStaffQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osq.inters {
		if inter == nil {
			return fmt.Errorf("mainent: uninitialized interceptor (forgotten import mainent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osq); err != nil {
				return err
			}
		}
	}
	for _, f := range osq.ctx.Fields {
		if !orgstaff.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OrgStaffQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgStaff, error) {
	var (
		nodes       = []*OrgStaff{}
		_spec       = osq.querySpec()
		loadedTypes = [5]bool{
			osq.withOrgan != nil,
			osq.withIdenAddr != nil,
			osq.withResiAddr != nil,
			osq.withDept != nil,
			osq.withPosi != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgStaff).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgStaff{config: osq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := osq.withOrgan; query != nil {
		if err := osq.loadOrgan(ctx, query, nodes, nil,
			func(n *OrgStaff, e *OrgOrgan) { n.Edges.Organ = e }); err != nil {
			return nil, err
		}
	}
	if query := osq.withIdenAddr; query != nil {
		if err := osq.loadIdenAddr(ctx, query, nodes, nil,
			func(n *OrgStaff, e *SysAddress) { n.Edges.IdenAddr = e }); err != nil {
			return nil, err
		}
	}
	if query := osq.withResiAddr; query != nil {
		if err := osq.loadResiAddr(ctx, query, nodes, nil,
			func(n *OrgStaff, e *SysAddress) { n.Edges.ResiAddr = e }); err != nil {
			return nil, err
		}
	}
	if query := osq.withDept; query != nil {
		if err := osq.loadDept(ctx, query, nodes, nil,
			func(n *OrgStaff, e *OrgDept) { n.Edges.Dept = e }); err != nil {
			return nil, err
		}
	}
	if query := osq.withPosi; query != nil {
		if err := osq.loadPosi(ctx, query, nodes, nil,
			func(n *OrgStaff, e *OrgPosition) { n.Edges.Posi = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (osq *OrgStaffQuery) loadOrgan(ctx context.Context, query *OrgOrganQuery, nodes []*OrgStaff, init func(*OrgStaff), assign func(*OrgStaff, *OrgOrgan)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgStaff)
	for i := range nodes {
		if nodes[i].OrgID == nil {
			continue
		}
		fk := *nodes[i].OrgID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(orgorgan.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "org_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (osq *OrgStaffQuery) loadIdenAddr(ctx context.Context, query *SysAddressQuery, nodes []*OrgStaff, init func(*OrgStaff), assign func(*OrgStaff, *SysAddress)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgStaff)
	for i := range nodes {
		if nodes[i].IdenAddrID == nil {
			continue
		}
		fk := *nodes[i].IdenAddrID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sysaddress.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "iden_addr_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (osq *OrgStaffQuery) loadResiAddr(ctx context.Context, query *SysAddressQuery, nodes []*OrgStaff, init func(*OrgStaff), assign func(*OrgStaff, *SysAddress)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgStaff)
	for i := range nodes {
		if nodes[i].ResiAddrID == nil {
			continue
		}
		fk := *nodes[i].ResiAddrID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sysaddress.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "resi_addr_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (osq *OrgStaffQuery) loadDept(ctx context.Context, query *OrgDeptQuery, nodes []*OrgStaff, init func(*OrgStaff), assign func(*OrgStaff, *OrgDept)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgStaff)
	for i := range nodes {
		if nodes[i].DeptID == nil {
			continue
		}
		fk := *nodes[i].DeptID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(orgdept.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dept_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (osq *OrgStaffQuery) loadPosi(ctx context.Context, query *OrgPositionQuery, nodes []*OrgStaff, init func(*OrgStaff), assign func(*OrgStaff, *OrgPosition)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OrgStaff)
	for i := range nodes {
		if nodes[i].PosiID == nil {
			continue
		}
		fk := *nodes[i].PosiID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(orgposition.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "posi_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (osq *OrgStaffQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	_spec.Node.Columns = osq.ctx.Fields
	if len(osq.ctx.Fields) > 0 {
		_spec.Unique = osq.ctx.Unique != nil && *osq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrgStaffQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orgstaff.Table, orgstaff.Columns, sqlgraph.NewFieldSpec(orgstaff.FieldID, field.TypeString))
	_spec.From = osq.sql
	if unique := osq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osq.path != nil {
		_spec.Unique = true
	}
	if fields := osq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgstaff.FieldID)
		for i := range fields {
			if fields[i] != orgstaff.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if osq.withOrgan != nil {
			_spec.Node.AddColumnOnce(orgstaff.FieldOrgID)
		}
		if osq.withIdenAddr != nil {
			_spec.Node.AddColumnOnce(orgstaff.FieldIdenAddrID)
		}
		if osq.withResiAddr != nil {
			_spec.Node.AddColumnOnce(orgstaff.FieldResiAddrID)
		}
		if osq.withDept != nil {
			_spec.Node.AddColumnOnce(orgstaff.FieldDeptID)
		}
		if osq.withPosi != nil {
			_spec.Node.AddColumnOnce(orgstaff.FieldPosiID)
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrgStaffQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(orgstaff.Table)
	columns := osq.ctx.Fields
	if len(columns) == 0 {
		columns = orgstaff.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.ctx.Unique != nil && *osq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range osq.modifiers {
		m(selector)
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (osq *OrgStaffQuery) ForUpdate(opts ...sql.LockOption) *OrgStaffQuery {
	if osq.driver.Dialect() == dialect.Postgres {
		osq.Unique(false)
	}
	osq.modifiers = append(osq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return osq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (osq *OrgStaffQuery) ForShare(opts ...sql.LockOption) *OrgStaffQuery {
	if osq.driver.Dialect() == dialect.Postgres {
		osq.Unique(false)
	}
	osq.modifiers = append(osq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return osq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osq *OrgStaffQuery) Modify(modifiers ...func(s *sql.Selector)) *OrgStaffSelect {
	osq.modifiers = append(osq.modifiers, modifiers...)
	return osq.Select()
}

// OrgStaffGroupBy is the group-by builder for OrgStaff entities.
type OrgStaffGroupBy struct {
	selector
	build *OrgStaffQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrgStaffGroupBy) Aggregate(fns ...AggregateFunc) *OrgStaffGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the selector query and scans the result into the given value.
func (osgb *OrgStaffGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osgb.build.ctx, ent.OpQueryGroupBy)
	if err := osgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgStaffQuery, *OrgStaffGroupBy](ctx, osgb.build, osgb, osgb.build.inters, v)
}

func (osgb *OrgStaffGroupBy) sqlScan(ctx context.Context, root *OrgStaffQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osgb.flds)+len(osgb.fns))
		for _, f := range *osgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgStaffSelect is the builder for selecting fields of OrgStaff entities.
type OrgStaffSelect struct {
	*OrgStaffQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oss *OrgStaffSelect) Aggregate(fns ...AggregateFunc) *OrgStaffSelect {
	oss.fns = append(oss.fns, fns...)
	return oss
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrgStaffSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oss.ctx, ent.OpQuerySelect)
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgStaffQuery, *OrgStaffSelect](ctx, oss.OrgStaffQuery, oss, oss.inters, v)
}

func (oss *OrgStaffSelect) sqlScan(ctx context.Context, root *OrgStaffQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oss.fns))
	for _, fn := range oss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oss *OrgStaffSelect) Modify(modifiers ...func(s *sql.Selector)) *OrgStaffSelect {
	oss.modifiers = append(oss.modifiers, modifiers...)
	return oss
}
