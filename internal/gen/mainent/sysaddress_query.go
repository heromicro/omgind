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
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/sysaddress"
)

// SysAddressQuery is the builder for querying SysAddress entities.
type SysAddressQuery struct {
	config
	ctx           *QueryContext
	order         []sysaddress.OrderOption
	inters        []Interceptor
	predicates    []predicate.SysAddress
	withOrgan     *OrgOrganQuery
	withStaffResi *OrgStaffQuery
	withStaffIden *OrgStaffQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysAddressQuery builder.
func (saq *SysAddressQuery) Where(ps ...predicate.SysAddress) *SysAddressQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit the number of records to be returned by this query.
func (saq *SysAddressQuery) Limit(limit int) *SysAddressQuery {
	saq.ctx.Limit = &limit
	return saq
}

// Offset to start from.
func (saq *SysAddressQuery) Offset(offset int) *SysAddressQuery {
	saq.ctx.Offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SysAddressQuery) Unique(unique bool) *SysAddressQuery {
	saq.ctx.Unique = &unique
	return saq
}

// Order specifies how the records should be ordered.
func (saq *SysAddressQuery) Order(o ...sysaddress.OrderOption) *SysAddressQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// QueryOrgan chains the current query on the "organ" edge.
func (saq *SysAddressQuery) QueryOrgan() *OrgOrganQuery {
	query := (&OrgOrganClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysaddress.Table, sysaddress.FieldID, selector),
			sqlgraph.To(orgorgan.Table, orgorgan.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, sysaddress.OrganTable, sysaddress.OrganColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStaffResi chains the current query on the "staff_resi" edge.
func (saq *SysAddressQuery) QueryStaffResi() *OrgStaffQuery {
	query := (&OrgStaffClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysaddress.Table, sysaddress.FieldID, selector),
			sqlgraph.To(orgstaff.Table, orgstaff.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, sysaddress.StaffResiTable, sysaddress.StaffResiColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStaffIden chains the current query on the "staff_iden" edge.
func (saq *SysAddressQuery) QueryStaffIden() *OrgStaffQuery {
	query := (&OrgStaffClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysaddress.Table, sysaddress.FieldID, selector),
			sqlgraph.To(orgstaff.Table, orgstaff.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, sysaddress.StaffIdenTable, sysaddress.StaffIdenColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysAddress entity from the query.
// Returns a *NotFoundError when no SysAddress was found.
func (saq *SysAddressQuery) First(ctx context.Context) (*SysAddress, error) {
	nodes, err := saq.Limit(1).All(setContextOp(ctx, saq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SysAddressQuery) FirstX(ctx context.Context) *SysAddress {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysAddress ID from the query.
// Returns a *NotFoundError when no SysAddress ID was found.
func (saq *SysAddressQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(1).IDs(setContextOp(ctx, saq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SysAddressQuery) FirstIDX(ctx context.Context) string {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysAddress entity is found.
// Returns a *NotFoundError when no SysAddress entities are found.
func (saq *SysAddressQuery) Only(ctx context.Context) (*SysAddress, error) {
	nodes, err := saq.Limit(2).All(setContextOp(ctx, saq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysaddress.Label}
	default:
		return nil, &NotSingularError{sysaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SysAddressQuery) OnlyX(ctx context.Context) *SysAddress {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysAddress ID in the query.
// Returns a *NotSingularError when more than one SysAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SysAddressQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = saq.Limit(2).IDs(setContextOp(ctx, saq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysaddress.Label}
	default:
		err = &NotSingularError{sysaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SysAddressQuery) OnlyIDX(ctx context.Context) string {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysAddresses.
func (saq *SysAddressQuery) All(ctx context.Context) ([]*SysAddress, error) {
	ctx = setContextOp(ctx, saq.ctx, ent.OpQueryAll)
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysAddress, *SysAddressQuery]()
	return withInterceptors[[]*SysAddress](ctx, saq, qr, saq.inters)
}

// AllX is like All, but panics if an error occurs.
func (saq *SysAddressQuery) AllX(ctx context.Context) []*SysAddress {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysAddress IDs.
func (saq *SysAddressQuery) IDs(ctx context.Context) (ids []string, err error) {
	if saq.ctx.Unique == nil && saq.path != nil {
		saq.Unique(true)
	}
	ctx = setContextOp(ctx, saq.ctx, ent.OpQueryIDs)
	if err = saq.Select(sysaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SysAddressQuery) IDsX(ctx context.Context) []string {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SysAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, saq.ctx, ent.OpQueryCount)
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, saq, querierCount[*SysAddressQuery](), saq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SysAddressQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SysAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, saq.ctx, ent.OpQueryExist)
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
func (saq *SysAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SysAddressQuery) Clone() *SysAddressQuery {
	if saq == nil {
		return nil
	}
	return &SysAddressQuery{
		config:        saq.config,
		ctx:           saq.ctx.Clone(),
		order:         append([]sysaddress.OrderOption{}, saq.order...),
		inters:        append([]Interceptor{}, saq.inters...),
		predicates:    append([]predicate.SysAddress{}, saq.predicates...),
		withOrgan:     saq.withOrgan.Clone(),
		withStaffResi: saq.withStaffResi.Clone(),
		withStaffIden: saq.withStaffIden.Clone(),
		// clone intermediate query.
		sql:       saq.sql.Clone(),
		path:      saq.path,
		modifiers: append([]func(*sql.Selector){}, saq.modifiers...),
	}
}

// WithOrgan tells the query-builder to eager-load the nodes that are connected to
// the "organ" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SysAddressQuery) WithOrgan(opts ...func(*OrgOrganQuery)) *SysAddressQuery {
	query := (&OrgOrganClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withOrgan = query
	return saq
}

// WithStaffResi tells the query-builder to eager-load the nodes that are connected to
// the "staff_resi" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SysAddressQuery) WithStaffResi(opts ...func(*OrgStaffQuery)) *SysAddressQuery {
	query := (&OrgStaffClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withStaffResi = query
	return saq
}

// WithStaffIden tells the query-builder to eager-load the nodes that are connected to
// the "staff_iden" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SysAddressQuery) WithStaffIden(opts ...func(*OrgStaffQuery)) *SysAddressQuery {
	query := (&OrgStaffClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withStaffIden = query
	return saq
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
//	client.SysAddress.Query().
//		GroupBy(sysaddress.FieldIsDel).
//		Aggregate(mainent.Count()).
//		Scan(ctx, &v)
func (saq *SysAddressQuery) GroupBy(field string, fields ...string) *SysAddressGroupBy {
	saq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysAddressGroupBy{build: saq}
	grbuild.flds = &saq.ctx.Fields
	grbuild.label = sysaddress.Label
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
//	client.SysAddress.Query().
//		Select(sysaddress.FieldIsDel).
//		Scan(ctx, &v)
func (saq *SysAddressQuery) Select(fields ...string) *SysAddressSelect {
	saq.ctx.Fields = append(saq.ctx.Fields, fields...)
	sbuild := &SysAddressSelect{SysAddressQuery: saq}
	sbuild.label = sysaddress.Label
	sbuild.flds, sbuild.scan = &saq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysAddressSelect configured with the given aggregations.
func (saq *SysAddressQuery) Aggregate(fns ...AggregateFunc) *SysAddressSelect {
	return saq.Select().Aggregate(fns...)
}

func (saq *SysAddressQuery) prepareQuery(ctx context.Context) error {
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
		if !sysaddress.ValidColumn(f) {
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

func (saq *SysAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysAddress, error) {
	var (
		nodes       = []*SysAddress{}
		_spec       = saq.querySpec()
		loadedTypes = [3]bool{
			saq.withOrgan != nil,
			saq.withStaffResi != nil,
			saq.withStaffIden != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysAddress{config: saq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := saq.withOrgan; query != nil {
		if err := saq.loadOrgan(ctx, query, nodes, nil,
			func(n *SysAddress, e *OrgOrgan) { n.Edges.Organ = e }); err != nil {
			return nil, err
		}
	}
	if query := saq.withStaffResi; query != nil {
		if err := saq.loadStaffResi(ctx, query, nodes, nil,
			func(n *SysAddress, e *OrgStaff) { n.Edges.StaffResi = e }); err != nil {
			return nil, err
		}
	}
	if query := saq.withStaffIden; query != nil {
		if err := saq.loadStaffIden(ctx, query, nodes, nil,
			func(n *SysAddress, e *OrgStaff) { n.Edges.StaffIden = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (saq *SysAddressQuery) loadOrgan(ctx context.Context, query *OrgOrganQuery, nodes []*SysAddress, init func(*SysAddress), assign func(*SysAddress, *OrgOrgan)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysAddress)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orgorgan.FieldHaddrID)
	}
	query.Where(predicate.OrgOrgan(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sysaddress.OrganColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.HaddrID
		if fk == nil {
			return fmt.Errorf(`foreign-key "haddr_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "haddr_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (saq *SysAddressQuery) loadStaffResi(ctx context.Context, query *OrgStaffQuery, nodes []*SysAddress, init func(*SysAddress), assign func(*SysAddress, *OrgStaff)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysAddress)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orgstaff.FieldResiAddrID)
	}
	query.Where(predicate.OrgStaff(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sysaddress.StaffResiColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ResiAddrID
		if fk == nil {
			return fmt.Errorf(`foreign-key "resi_addr_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "resi_addr_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (saq *SysAddressQuery) loadStaffIden(ctx context.Context, query *OrgStaffQuery, nodes []*SysAddress, init func(*SysAddress), assign func(*SysAddress, *OrgStaff)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SysAddress)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orgstaff.FieldIdenAddrID)
	}
	query.Where(predicate.OrgStaff(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sysaddress.StaffIdenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.IdenAddrID
		if fk == nil {
			return fmt.Errorf(`foreign-key "iden_addr_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "iden_addr_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (saq *SysAddressQuery) sqlCount(ctx context.Context) (int, error) {
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

func (saq *SysAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysaddress.Table, sysaddress.Columns, sqlgraph.NewFieldSpec(sysaddress.FieldID, field.TypeString))
	_spec.From = saq.sql
	if unique := saq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if saq.path != nil {
		_spec.Unique = true
	}
	if fields := saq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysaddress.FieldID)
		for i := range fields {
			if fields[i] != sysaddress.FieldID {
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

func (saq *SysAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(sysaddress.Table)
	columns := saq.ctx.Fields
	if len(columns) == 0 {
		columns = sysaddress.Columns
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
func (saq *SysAddressQuery) ForUpdate(opts ...sql.LockOption) *SysAddressQuery {
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
func (saq *SysAddressQuery) ForShare(opts ...sql.LockOption) *SysAddressQuery {
	if saq.driver.Dialect() == dialect.Postgres {
		saq.Unique(false)
	}
	saq.modifiers = append(saq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return saq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (saq *SysAddressQuery) Modify(modifiers ...func(s *sql.Selector)) *SysAddressSelect {
	saq.modifiers = append(saq.modifiers, modifiers...)
	return saq.Select()
}

// SysAddressGroupBy is the group-by builder for SysAddress entities.
type SysAddressGroupBy struct {
	selector
	build *SysAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SysAddressGroupBy) Aggregate(fns ...AggregateFunc) *SysAddressGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the selector query and scans the result into the given value.
func (sagb *SysAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sagb.build.ctx, ent.OpQueryGroupBy)
	if err := sagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysAddressQuery, *SysAddressGroupBy](ctx, sagb.build, sagb, sagb.build.inters, v)
}

func (sagb *SysAddressGroupBy) sqlScan(ctx context.Context, root *SysAddressQuery, v any) error {
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

// SysAddressSelect is the builder for selecting fields of SysAddress entities.
type SysAddressSelect struct {
	*SysAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sas *SysAddressSelect) Aggregate(fns ...AggregateFunc) *SysAddressSelect {
	sas.fns = append(sas.fns, fns...)
	return sas
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SysAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sas.ctx, ent.OpQuerySelect)
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysAddressQuery, *SysAddressSelect](ctx, sas.SysAddressQuery, sas, sas.inters, v)
}

func (sas *SysAddressSelect) sqlScan(ctx context.Context, root *SysAddressQuery, v any) error {
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
func (sas *SysAddressSelect) Modify(modifiers ...func(s *sql.Selector)) *SysAddressSelect {
	sas.modifiers = append(sas.modifiers, modifiers...)
	return sas
}
