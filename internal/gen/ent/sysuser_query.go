// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/internal"
	"github.com/heromicro/omgind/internal/gen/ent/predicate"
	"github.com/heromicro/omgind/internal/gen/ent/sysuser"
)

// SysUserQuery is the builder for querying SysUser entities.
type SysUserQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.SysUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysUserQuery builder.
func (suq *SysUserQuery) Where(ps ...predicate.SysUser) *SysUserQuery {
	suq.predicates = append(suq.predicates, ps...)
	return suq
}

// Limit the number of records to be returned by this query.
func (suq *SysUserQuery) Limit(limit int) *SysUserQuery {
	suq.ctx.Limit = &limit
	return suq
}

// Offset to start from.
func (suq *SysUserQuery) Offset(offset int) *SysUserQuery {
	suq.ctx.Offset = &offset
	return suq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (suq *SysUserQuery) Unique(unique bool) *SysUserQuery {
	suq.ctx.Unique = &unique
	return suq
}

// Order specifies how the records should be ordered.
func (suq *SysUserQuery) Order(o ...OrderFunc) *SysUserQuery {
	suq.order = append(suq.order, o...)
	return suq
}

// First returns the first SysUser entity from the query.
// Returns a *NotFoundError when no SysUser was found.
func (suq *SysUserQuery) First(ctx context.Context) (*SysUser, error) {
	nodes, err := suq.Limit(1).All(setContextOp(ctx, suq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (suq *SysUserQuery) FirstX(ctx context.Context) *SysUser {
	node, err := suq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysUser ID from the query.
// Returns a *NotFoundError when no SysUser ID was found.
func (suq *SysUserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = suq.Limit(1).IDs(setContextOp(ctx, suq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (suq *SysUserQuery) FirstIDX(ctx context.Context) string {
	id, err := suq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysUser entity is found.
// Returns a *NotFoundError when no SysUser entities are found.
func (suq *SysUserQuery) Only(ctx context.Context) (*SysUser, error) {
	nodes, err := suq.Limit(2).All(setContextOp(ctx, suq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysuser.Label}
	default:
		return nil, &NotSingularError{sysuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (suq *SysUserQuery) OnlyX(ctx context.Context) *SysUser {
	node, err := suq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysUser ID in the query.
// Returns a *NotSingularError when more than one SysUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (suq *SysUserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = suq.Limit(2).IDs(setContextOp(ctx, suq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysuser.Label}
	default:
		err = &NotSingularError{sysuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (suq *SysUserQuery) OnlyIDX(ctx context.Context) string {
	id, err := suq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysUsers.
func (suq *SysUserQuery) All(ctx context.Context) ([]*SysUser, error) {
	ctx = setContextOp(ctx, suq.ctx, "All")
	if err := suq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysUser, *SysUserQuery]()
	return withInterceptors[[]*SysUser](ctx, suq, qr, suq.inters)
}

// AllX is like All, but panics if an error occurs.
func (suq *SysUserQuery) AllX(ctx context.Context) []*SysUser {
	nodes, err := suq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysUser IDs.
func (suq *SysUserQuery) IDs(ctx context.Context) (ids []string, err error) {
	if suq.ctx.Unique == nil && suq.path != nil {
		suq.Unique(true)
	}
	ctx = setContextOp(ctx, suq.ctx, "IDs")
	if err = suq.Select(sysuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (suq *SysUserQuery) IDsX(ctx context.Context) []string {
	ids, err := suq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (suq *SysUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, suq.ctx, "Count")
	if err := suq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, suq, querierCount[*SysUserQuery](), suq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (suq *SysUserQuery) CountX(ctx context.Context) int {
	count, err := suq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (suq *SysUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, suq.ctx, "Exist")
	switch _, err := suq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (suq *SysUserQuery) ExistX(ctx context.Context) bool {
	exist, err := suq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (suq *SysUserQuery) Clone() *SysUserQuery {
	if suq == nil {
		return nil
	}
	return &SysUserQuery{
		config:     suq.config,
		ctx:        suq.ctx.Clone(),
		order:      append([]OrderFunc{}, suq.order...),
		inters:     append([]Interceptor{}, suq.inters...),
		predicates: append([]predicate.SysUser{}, suq.predicates...),
		// clone intermediate query.
		sql:  suq.sql.Clone(),
		path: suq.path,
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
//	client.SysUser.Query().
//		GroupBy(sysuser.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (suq *SysUserQuery) GroupBy(field string, fields ...string) *SysUserGroupBy {
	suq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysUserGroupBy{build: suq}
	grbuild.flds = &suq.ctx.Fields
	grbuild.label = sysuser.Label
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
//	client.SysUser.Query().
//		Select(sysuser.FieldIsDel).
//		Scan(ctx, &v)
func (suq *SysUserQuery) Select(fields ...string) *SysUserSelect {
	suq.ctx.Fields = append(suq.ctx.Fields, fields...)
	sbuild := &SysUserSelect{SysUserQuery: suq}
	sbuild.label = sysuser.Label
	sbuild.flds, sbuild.scan = &suq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysUserSelect configured with the given aggregations.
func (suq *SysUserQuery) Aggregate(fns ...AggregateFunc) *SysUserSelect {
	return suq.Select().Aggregate(fns...)
}

func (suq *SysUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range suq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, suq); err != nil {
				return err
			}
		}
	}
	for _, f := range suq.ctx.Fields {
		if !sysuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if suq.path != nil {
		prev, err := suq.path(ctx)
		if err != nil {
			return err
		}
		suq.sql = prev
	}
	return nil
}

func (suq *SysUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysUser, error) {
	var (
		nodes = []*SysUser{}
		_spec = suq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysUser{config: suq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = suq.schemaConfig.SysUser
	ctx = internal.NewSchemaConfigContext(ctx, suq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, suq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (suq *SysUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := suq.querySpec()
	_spec.Node.Schema = suq.schemaConfig.SysUser
	ctx = internal.NewSchemaConfigContext(ctx, suq.schemaConfig)
	_spec.Node.Columns = suq.ctx.Fields
	if len(suq.ctx.Fields) > 0 {
		_spec.Unique = suq.ctx.Unique != nil && *suq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, suq.driver, _spec)
}

func (suq *SysUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysuser.Table, sysuser.Columns, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString))
	_spec.From = suq.sql
	if unique := suq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if suq.path != nil {
		_spec.Unique = true
	}
	if fields := suq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuser.FieldID)
		for i := range fields {
			if fields[i] != sysuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := suq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := suq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := suq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := suq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (suq *SysUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(suq.driver.Dialect())
	t1 := builder.Table(sysuser.Table)
	columns := suq.ctx.Fields
	if len(columns) == 0 {
		columns = sysuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if suq.sql != nil {
		selector = suq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if suq.ctx.Unique != nil && *suq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(suq.schemaConfig.SysUser)
	ctx = internal.NewSchemaConfigContext(ctx, suq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range suq.predicates {
		p(selector)
	}
	for _, p := range suq.order {
		p(selector)
	}
	if offset := suq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := suq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysUserGroupBy is the group-by builder for SysUser entities.
type SysUserGroupBy struct {
	selector
	build *SysUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sugb *SysUserGroupBy) Aggregate(fns ...AggregateFunc) *SysUserGroupBy {
	sugb.fns = append(sugb.fns, fns...)
	return sugb
}

// Scan applies the selector query and scans the result into the given value.
func (sugb *SysUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sugb.build.ctx, "GroupBy")
	if err := sugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysUserQuery, *SysUserGroupBy](ctx, sugb.build, sugb, sugb.build.inters, v)
}

func (sugb *SysUserGroupBy) sqlScan(ctx context.Context, root *SysUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sugb.fns))
	for _, fn := range sugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sugb.flds)+len(sugb.fns))
		for _, f := range *sugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysUserSelect is the builder for selecting fields of SysUser entities.
type SysUserSelect struct {
	*SysUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sus *SysUserSelect) Aggregate(fns ...AggregateFunc) *SysUserSelect {
	sus.fns = append(sus.fns, fns...)
	return sus
}

// Scan applies the selector query and scans the result into the given value.
func (sus *SysUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sus.ctx, "Select")
	if err := sus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysUserQuery, *SysUserSelect](ctx, sus.SysUserQuery, sus, sus.inters, v)
}

func (sus *SysUserSelect) sqlScan(ctx context.Context, root *SysUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sus.fns))
	for _, fn := range sus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
