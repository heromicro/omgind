// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysrolemenu"
)

// SysRoleMenuQuery is the builder for querying SysRoleMenu entities.
type SysRoleMenuQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysRoleMenu
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysRoleMenuQuery builder.
func (srmq *SysRoleMenuQuery) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuQuery {
	srmq.predicates = append(srmq.predicates, ps...)
	return srmq
}

// Limit adds a limit step to the query.
func (srmq *SysRoleMenuQuery) Limit(limit int) *SysRoleMenuQuery {
	srmq.limit = &limit
	return srmq
}

// Offset adds an offset step to the query.
func (srmq *SysRoleMenuQuery) Offset(offset int) *SysRoleMenuQuery {
	srmq.offset = &offset
	return srmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srmq *SysRoleMenuQuery) Unique(unique bool) *SysRoleMenuQuery {
	srmq.unique = &unique
	return srmq
}

// Order adds an order step to the query.
func (srmq *SysRoleMenuQuery) Order(o ...OrderFunc) *SysRoleMenuQuery {
	srmq.order = append(srmq.order, o...)
	return srmq
}

// First returns the first SysRoleMenu entity from the query.
// Returns a *NotFoundError when no SysRoleMenu was found.
func (srmq *SysRoleMenuQuery) First(ctx context.Context) (*SysRoleMenu, error) {
	nodes, err := srmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysrolemenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) FirstX(ctx context.Context) *SysRoleMenu {
	node, err := srmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysRoleMenu ID from the query.
// Returns a *NotFoundError when no SysRoleMenu ID was found.
func (srmq *SysRoleMenuQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = srmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysrolemenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) FirstIDX(ctx context.Context) string {
	id, err := srmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysRoleMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysRoleMenu entity is found.
// Returns a *NotFoundError when no SysRoleMenu entities are found.
func (srmq *SysRoleMenuQuery) Only(ctx context.Context) (*SysRoleMenu, error) {
	nodes, err := srmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysrolemenu.Label}
	default:
		return nil, &NotSingularError{sysrolemenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) OnlyX(ctx context.Context) *SysRoleMenu {
	node, err := srmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysRoleMenu ID in the query.
// Returns a *NotSingularError when more than one SysRoleMenu ID is found.
// Returns a *NotFoundError when no entities are found.
func (srmq *SysRoleMenuQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = srmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysrolemenu.Label}
	default:
		err = &NotSingularError{sysrolemenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) OnlyIDX(ctx context.Context) string {
	id, err := srmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysRoleMenus.
func (srmq *SysRoleMenuQuery) All(ctx context.Context) ([]*SysRoleMenu, error) {
	if err := srmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return srmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) AllX(ctx context.Context) []*SysRoleMenu {
	nodes, err := srmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysRoleMenu IDs.
func (srmq *SysRoleMenuQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := srmq.Select(sysrolemenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) IDsX(ctx context.Context) []string {
	ids, err := srmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srmq *SysRoleMenuQuery) Count(ctx context.Context) (int, error) {
	if err := srmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return srmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) CountX(ctx context.Context) int {
	count, err := srmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srmq *SysRoleMenuQuery) Exist(ctx context.Context) (bool, error) {
	if err := srmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return srmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (srmq *SysRoleMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := srmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysRoleMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srmq *SysRoleMenuQuery) Clone() *SysRoleMenuQuery {
	if srmq == nil {
		return nil
	}
	return &SysRoleMenuQuery{
		config:     srmq.config,
		limit:      srmq.limit,
		offset:     srmq.offset,
		order:      append([]OrderFunc{}, srmq.order...),
		predicates: append([]predicate.SysRoleMenu{}, srmq.predicates...),
		// clone intermediate query.
		sql:    srmq.sql.Clone(),
		path:   srmq.path,
		unique: srmq.unique,
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
//	client.SysRoleMenu.Query().
//		GroupBy(sysrolemenu.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (srmq *SysRoleMenuQuery) GroupBy(field string, fields ...string) *SysRoleMenuGroupBy {
	grbuild := &SysRoleMenuGroupBy{config: srmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srmq.sqlQuery(ctx), nil
	}
	grbuild.label = sysrolemenu.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.SysRoleMenu.Query().
//		Select(sysrolemenu.FieldIsDel).
//		Scan(ctx, &v)
//
func (srmq *SysRoleMenuQuery) Select(fields ...string) *SysRoleMenuSelect {
	srmq.fields = append(srmq.fields, fields...)
	selbuild := &SysRoleMenuSelect{SysRoleMenuQuery: srmq}
	selbuild.label = sysrolemenu.Label
	selbuild.flds, selbuild.scan = &srmq.fields, selbuild.Scan
	return selbuild
}

func (srmq *SysRoleMenuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range srmq.fields {
		if !sysrolemenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if srmq.path != nil {
		prev, err := srmq.path(ctx)
		if err != nil {
			return err
		}
		srmq.sql = prev
	}
	return nil
}

func (srmq *SysRoleMenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysRoleMenu, error) {
	var (
		nodes = []*SysRoleMenu{}
		_spec = srmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SysRoleMenu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SysRoleMenu{config: srmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (srmq *SysRoleMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srmq.querySpec()
	_spec.Node.Columns = srmq.fields
	if len(srmq.fields) > 0 {
		_spec.Unique = srmq.unique != nil && *srmq.unique
	}
	return sqlgraph.CountNodes(ctx, srmq.driver, _spec)
}

func (srmq *SysRoleMenuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := srmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (srmq *SysRoleMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrolemenu.Table,
			Columns: sysrolemenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrolemenu.FieldID,
			},
		},
		From:   srmq.sql,
		Unique: true,
	}
	if unique := srmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := srmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrolemenu.FieldID)
		for i := range fields {
			if fields[i] != sysrolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srmq *SysRoleMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srmq.driver.Dialect())
	t1 := builder.Table(sysrolemenu.Table)
	columns := srmq.fields
	if len(columns) == 0 {
		columns = sysrolemenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srmq.sql != nil {
		selector = srmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srmq.unique != nil && *srmq.unique {
		selector.Distinct()
	}
	for _, p := range srmq.predicates {
		p(selector)
	}
	for _, p := range srmq.order {
		p(selector)
	}
	if offset := srmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysRoleMenuGroupBy is the group-by builder for SysRoleMenu entities.
type SysRoleMenuGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srmgb *SysRoleMenuGroupBy) Aggregate(fns ...AggregateFunc) *SysRoleMenuGroupBy {
	srmgb.fns = append(srmgb.fns, fns...)
	return srmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (srmgb *SysRoleMenuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := srmgb.path(ctx)
	if err != nil {
		return err
	}
	srmgb.sql = query
	return srmgb.sqlScan(ctx, v)
}

func (srmgb *SysRoleMenuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range srmgb.fields {
		if !sysrolemenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := srmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srmgb *SysRoleMenuGroupBy) sqlQuery() *sql.Selector {
	selector := srmgb.sql.Select()
	aggregation := make([]string, 0, len(srmgb.fns))
	for _, fn := range srmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(srmgb.fields)+len(srmgb.fns))
		for _, f := range srmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(srmgb.fields...)...)
}

// SysRoleMenuSelect is the builder for selecting fields of SysRoleMenu entities.
type SysRoleMenuSelect struct {
	*SysRoleMenuQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (srms *SysRoleMenuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := srms.prepareQuery(ctx); err != nil {
		return err
	}
	srms.sql = srms.SysRoleMenuQuery.sqlQuery(ctx)
	return srms.sqlScan(ctx, v)
}

func (srms *SysRoleMenuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srms.sql.Query()
	if err := srms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
