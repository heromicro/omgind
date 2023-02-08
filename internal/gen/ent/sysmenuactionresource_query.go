// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuactionresource"
)

// SysMenuActionResourceQuery is the builder for querying SysMenuActionResource entities.
type SysMenuActionResourceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysMenuActionResource
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuActionResourceQuery builder.
func (smarq *SysMenuActionResourceQuery) Where(ps ...predicate.SysMenuActionResource) *SysMenuActionResourceQuery {
	smarq.predicates = append(smarq.predicates, ps...)
	return smarq
}

// Limit adds a limit step to the query.
func (smarq *SysMenuActionResourceQuery) Limit(limit int) *SysMenuActionResourceQuery {
	smarq.limit = &limit
	return smarq
}

// Offset adds an offset step to the query.
func (smarq *SysMenuActionResourceQuery) Offset(offset int) *SysMenuActionResourceQuery {
	smarq.offset = &offset
	return smarq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smarq *SysMenuActionResourceQuery) Unique(unique bool) *SysMenuActionResourceQuery {
	smarq.unique = &unique
	return smarq
}

// Order adds an order step to the query.
func (smarq *SysMenuActionResourceQuery) Order(o ...OrderFunc) *SysMenuActionResourceQuery {
	smarq.order = append(smarq.order, o...)
	return smarq
}

// First returns the first SysMenuActionResource entity from the query.
// Returns a *NotFoundError when no SysMenuActionResource was found.
func (smarq *SysMenuActionResourceQuery) First(ctx context.Context) (*SysMenuActionResource, error) {
	nodes, err := smarq.Limit(1).All(ctx)
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
	if ids, err = smarq.Limit(1).IDs(ctx); err != nil {
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
// Returns a *NotSingularError when exactly one SysMenuActionResource entity is not found.
// Returns a *NotFoundError when no SysMenuActionResource entities are found.
func (smarq *SysMenuActionResourceQuery) Only(ctx context.Context) (*SysMenuActionResource, error) {
	nodes, err := smarq.Limit(2).All(ctx)
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
// Returns a *NotSingularError when exactly one SysMenuActionResource ID is not found.
// Returns a *NotFoundError when no entities are found.
func (smarq *SysMenuActionResourceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smarq.Limit(2).IDs(ctx); err != nil {
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
	if err := smarq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smarq.sqlAll(ctx)
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
func (smarq *SysMenuActionResourceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := smarq.Select(sysmenuactionresource.FieldID).Scan(ctx, &ids); err != nil {
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
	if err := smarq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smarq.sqlCount(ctx)
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
	if err := smarq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smarq.sqlExist(ctx)
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
		limit:      smarq.limit,
		offset:     smarq.offset,
		order:      append([]OrderFunc{}, smarq.order...),
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
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (smarq *SysMenuActionResourceQuery) GroupBy(field string, fields ...string) *SysMenuActionResourceGroupBy {
	group := &SysMenuActionResourceGroupBy{config: smarq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smarq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smarq.sqlQuery(ctx), nil
	}
	return group
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
//
func (smarq *SysMenuActionResourceQuery) Select(field string, fields ...string) *SysMenuActionResourceSelect {
	smarq.fields = append([]string{field}, fields...)
	return &SysMenuActionResourceSelect{SysMenuActionResourceQuery: smarq}
}

func (smarq *SysMenuActionResourceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smarq.fields {
		if !sysmenuactionresource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
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

func (smarq *SysMenuActionResourceQuery) sqlAll(ctx context.Context) ([]*SysMenuActionResource, error) {
	var (
		nodes = []*SysMenuActionResource{}
		_spec = smarq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysMenuActionResource{config: smarq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
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
	return sqlgraph.CountNodes(ctx, smarq.driver, _spec)
}

func (smarq *SysMenuActionResourceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := smarq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (smarq *SysMenuActionResourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenuactionresource.Table,
			Columns: sysmenuactionresource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenuactionresource.FieldID,
			},
		},
		From:   smarq.sql,
		Unique: true,
	}
	if unique := smarq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := smarq.fields; len(fields) > 0 {
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
	if limit := smarq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smarq.offset; offset != nil {
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
	selector := builder.Select(t1.Columns(sysmenuactionresource.Columns...)...).From(t1)
	if smarq.sql != nil {
		selector = smarq.sql
		selector.Select(selector.Columns(sysmenuactionresource.Columns...)...)
	}
	for _, p := range smarq.predicates {
		p(selector)
	}
	for _, p := range smarq.order {
		p(selector)
	}
	if offset := smarq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smarq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysMenuActionResourceGroupBy is the group-by builder for SysMenuActionResource entities.
type SysMenuActionResourceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smargb *SysMenuActionResourceGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuActionResourceGroupBy {
	smargb.fns = append(smargb.fns, fns...)
	return smargb
}

// Scan applies the group-by query and scans the result into the given value.
func (smargb *SysMenuActionResourceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := smargb.path(ctx)
	if err != nil {
		return err
	}
	smargb.sql = query
	return smargb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := smargb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(smargb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := smargb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) StringsX(ctx context.Context) []string {
	v, err := smargb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smargb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) StringX(ctx context.Context) string {
	v, err := smargb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(smargb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := smargb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) IntsX(ctx context.Context) []int {
	v, err := smargb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smargb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) IntX(ctx context.Context) int {
	v, err := smargb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(smargb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := smargb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := smargb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smargb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := smargb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(smargb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := smargb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := smargb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smargb *SysMenuActionResourceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smargb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smargb *SysMenuActionResourceGroupBy) BoolX(ctx context.Context) bool {
	v, err := smargb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smargb *SysMenuActionResourceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range smargb.fields {
		if !sysmenuactionresource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := smargb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smargb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smargb *SysMenuActionResourceGroupBy) sqlQuery() *sql.Selector {
	selector := smargb.sql
	columns := make([]string, 0, len(smargb.fields)+len(smargb.fns))
	columns = append(columns, smargb.fields...)
	for _, fn := range smargb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(smargb.fields...)
}

// SysMenuActionResourceSelect is the builder for selecting fields of SysMenuActionResource entities.
type SysMenuActionResourceSelect struct {
	*SysMenuActionResourceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (smars *SysMenuActionResourceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := smars.prepareQuery(ctx); err != nil {
		return err
	}
	smars.sql = smars.SysMenuActionResourceQuery.sqlQuery(ctx)
	return smars.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := smars.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(smars.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := smars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) StringsX(ctx context.Context) []string {
	v, err := smars.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smars.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) StringX(ctx context.Context) string {
	v, err := smars.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(smars.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := smars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) IntsX(ctx context.Context) []int {
	v, err := smars.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smars.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) IntX(ctx context.Context) int {
	v, err := smars.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(smars.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := smars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := smars.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smars.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) Float64X(ctx context.Context) float64 {
	v, err := smars.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(smars.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionResourceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := smars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) BoolsX(ctx context.Context) []bool {
	v, err := smars.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (smars *SysMenuActionResourceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smars.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuactionresource.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionResourceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smars *SysMenuActionResourceSelect) BoolX(ctx context.Context) bool {
	v, err := smars.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smars *SysMenuActionResourceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := smars.sqlQuery().Query()
	if err := smars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smars *SysMenuActionResourceSelect) sqlQuery() sql.Querier {
	selector := smars.sql
	selector.Select(selector.Columns(smars.fields...)...)
	return selector
}
