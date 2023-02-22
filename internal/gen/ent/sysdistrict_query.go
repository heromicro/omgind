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
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
)

// SysDistrictQuery is the builder for querying SysDistrict entities.
type SysDistrictQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysDistrict
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDistrictQuery builder.
func (sdq *SysDistrictQuery) Where(ps ...predicate.SysDistrict) *SysDistrictQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit adds a limit step to the query.
func (sdq *SysDistrictQuery) Limit(limit int) *SysDistrictQuery {
	sdq.limit = &limit
	return sdq
}

// Offset adds an offset step to the query.
func (sdq *SysDistrictQuery) Offset(offset int) *SysDistrictQuery {
	sdq.offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SysDistrictQuery) Unique(unique bool) *SysDistrictQuery {
	sdq.unique = &unique
	return sdq
}

// Order adds an order step to the query.
func (sdq *SysDistrictQuery) Order(o ...OrderFunc) *SysDistrictQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// First returns the first SysDistrict entity from the query.
// Returns a *NotFoundError when no SysDistrict was found.
func (sdq *SysDistrictQuery) First(ctx context.Context) (*SysDistrict, error) {
	nodes, err := sdq.Limit(1).All(ctx)
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
	if ids, err = sdq.Limit(1).IDs(ctx); err != nil {
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
	nodes, err := sdq.Limit(2).All(ctx)
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
	if ids, err = sdq.Limit(2).IDs(ctx); err != nil {
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
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sdq.sqlAll(ctx)
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
func (sdq *SysDistrictQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sdq.Select(sysdistrict.FieldID).Scan(ctx, &ids); err != nil {
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
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sdq.sqlCount(ctx)
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
	if err := sdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sdq.sqlExist(ctx)
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
		config:     sdq.config,
		limit:      sdq.limit,
		offset:     sdq.offset,
		order:      append([]OrderFunc{}, sdq.order...),
		predicates: append([]predicate.SysDistrict{}, sdq.predicates...),
		// clone intermediate query.
		sql:    sdq.sql.Clone(),
		path:   sdq.path,
		unique: sdq.unique,
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
//	client.SysDistrict.Query().
//		GroupBy(sysdistrict.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sdq *SysDistrictQuery) GroupBy(field string, fields ...string) *SysDistrictGroupBy {
	grbuild := &SysDistrictGroupBy{config: sdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sdq.sqlQuery(ctx), nil
	}
	grbuild.label = sysdistrict.Label
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
//	client.SysDistrict.Query().
//		Select(sysdistrict.FieldIsDel).
//		Scan(ctx, &v)
//
func (sdq *SysDistrictQuery) Select(fields ...string) *SysDistrictSelect {
	sdq.fields = append(sdq.fields, fields...)
	selbuild := &SysDistrictSelect{SysDistrictQuery: sdq}
	selbuild.label = sysdistrict.Label
	selbuild.flds, selbuild.scan = &sdq.fields, selbuild.Scan
	return selbuild
}

func (sdq *SysDistrictQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sdq.fields {
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
		nodes = []*SysDistrict{}
		_spec = sdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SysDistrict).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SysDistrict{config: sdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
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
	return nodes, nil
}

func (sdq *SysDistrictQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	_spec.Node.Columns = sdq.fields
	if len(sdq.fields) > 0 {
		_spec.Unique = sdq.unique != nil && *sdq.unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SysDistrictQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sdq *SysDistrictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdistrict.Table,
			Columns: sysdistrict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdistrict.FieldID,
			},
		},
		From:   sdq.sql,
		Unique: true,
	}
	if unique := sdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sdq.fields; len(fields) > 0 {
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
	if limit := sdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdq.offset; offset != nil {
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
	columns := sdq.fields
	if len(columns) == 0 {
		columns = sysdistrict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.unique != nil && *sdq.unique {
		selector.Distinct()
	}
	for _, p := range sdq.predicates {
		p(selector)
	}
	for _, p := range sdq.order {
		p(selector)
	}
	if offset := sdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysDistrictGroupBy is the group-by builder for SysDistrict entities.
type SysDistrictGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SysDistrictGroupBy) Aggregate(fns ...AggregateFunc) *SysDistrictGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sdgb *SysDistrictGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sdgb.path(ctx)
	if err != nil {
		return err
	}
	sdgb.sql = query
	return sdgb.sqlScan(ctx, v)
}

func (sdgb *SysDistrictGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sdgb.fields {
		if !sysdistrict.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sdgb *SysDistrictGroupBy) sqlQuery() *sql.Selector {
	selector := sdgb.sql.Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sdgb.fields)+len(sdgb.fns))
		for _, f := range sdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sdgb.fields...)...)
}

// SysDistrictSelect is the builder for selecting fields of SysDistrict entities.
type SysDistrictSelect struct {
	*SysDistrictQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SysDistrictSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	sds.sql = sds.SysDistrictQuery.sqlQuery(ctx)
	return sds.sqlScan(ctx, v)
}

func (sds *SysDistrictSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sds.sql.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}