// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-cqrs/app/logger/job/internal/data/ent/predicate"
	"kratos-cqrs/app/logger/job/internal/data/ent/sensordata"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorDataQuery is the builder for querying SensorData entities.
type SensorDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SensorData
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SensorDataQuery builder.
func (sdq *SensorDataQuery) Where(ps ...predicate.SensorData) *SensorDataQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit adds a limit step to the query.
func (sdq *SensorDataQuery) Limit(limit int) *SensorDataQuery {
	sdq.limit = &limit
	return sdq
}

// Offset adds an offset step to the query.
func (sdq *SensorDataQuery) Offset(offset int) *SensorDataQuery {
	sdq.offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SensorDataQuery) Unique(unique bool) *SensorDataQuery {
	sdq.unique = &unique
	return sdq
}

// Order adds an order step to the query.
func (sdq *SensorDataQuery) Order(o ...OrderFunc) *SensorDataQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// First returns the first SensorData entity from the query.
// Returns a *NotFoundError when no SensorData was found.
func (sdq *SensorDataQuery) First(ctx context.Context) (*SensorData, error) {
	nodes, err := sdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sensordata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SensorDataQuery) FirstX(ctx context.Context) *SensorData {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SensorData ID from the query.
// Returns a *NotFoundError when no SensorData ID was found.
func (sdq *SensorDataQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sensordata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SensorDataQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SensorData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SensorData entity is found.
// Returns a *NotFoundError when no SensorData entities are found.
func (sdq *SensorDataQuery) Only(ctx context.Context) (*SensorData, error) {
	nodes, err := sdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sensordata.Label}
	default:
		return nil, &NotSingularError{sensordata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SensorDataQuery) OnlyX(ctx context.Context) *SensorData {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SensorData ID in the query.
// Returns a *NotSingularError when more than one SensorData ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SensorDataQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sensordata.Label}
	default:
		err = &NotSingularError{sensordata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SensorDataQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SensorDataSlice.
func (sdq *SensorDataQuery) All(ctx context.Context) ([]*SensorData, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SensorDataQuery) AllX(ctx context.Context) []*SensorData {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SensorData IDs.
func (sdq *SensorDataQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := sdq.Select(sensordata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SensorDataQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SensorDataQuery) Count(ctx context.Context) (int, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SensorDataQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SensorDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sdq *SensorDataQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SensorDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SensorDataQuery) Clone() *SensorDataQuery {
	if sdq == nil {
		return nil
	}
	return &SensorDataQuery{
		config:     sdq.config,
		limit:      sdq.limit,
		offset:     sdq.offset,
		order:      append([]OrderFunc{}, sdq.order...),
		predicates: append([]predicate.SensorData{}, sdq.predicates...),
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
//		Time int64 `json:"time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SensorData.Query().
//		GroupBy(sensordata.FieldTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SensorDataQuery) GroupBy(field string, fields ...string) *SensorDataGroupBy {
	grbuild := &SensorDataGroupBy{config: sdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sdq.sqlQuery(ctx), nil
	}
	grbuild.label = sensordata.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Time int64 `json:"time,omitempty"`
//	}
//
//	client.SensorData.Query().
//		Select(sensordata.FieldTime).
//		Scan(ctx, &v)
func (sdq *SensorDataQuery) Select(fields ...string) *SensorDataSelect {
	sdq.fields = append(sdq.fields, fields...)
	selbuild := &SensorDataSelect{SensorDataQuery: sdq}
	selbuild.label = sensordata.Label
	selbuild.flds, selbuild.scan = &sdq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SensorDataSelect configured with the given aggregations.
func (sdq *SensorDataQuery) Aggregate(fns ...AggregateFunc) *SensorDataSelect {
	return sdq.Select().Aggregate(fns...)
}

func (sdq *SensorDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sdq.fields {
		if !sensordata.ValidColumn(f) {
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

func (sdq *SensorDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SensorData, error) {
	var (
		nodes = []*SensorData{}
		_spec = sdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SensorData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SensorData{config: sdq.config}
		nodes = append(nodes, node)
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
	return nodes, nil
}

func (sdq *SensorDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	_spec.Node.Columns = sdq.fields
	if len(sdq.fields) > 0 {
		_spec.Unique = sdq.unique != nil && *sdq.unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SensorDataQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sdq *SensorDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sensordata.Table,
			Columns: sensordata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sensordata.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, sensordata.FieldID)
		for i := range fields {
			if fields[i] != sensordata.FieldID {
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

func (sdq *SensorDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(sensordata.Table)
	columns := sdq.fields
	if len(columns) == 0 {
		columns = sensordata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.unique != nil && *sdq.unique {
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

// Modify adds a query modifier for attaching custom logic to queries.
func (sdq *SensorDataQuery) Modify(modifiers ...func(s *sql.Selector)) *SensorDataSelect {
	sdq.modifiers = append(sdq.modifiers, modifiers...)
	return sdq.Select()
}

// SensorDataGroupBy is the group-by builder for SensorData entities.
type SensorDataGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SensorDataGroupBy) Aggregate(fns ...AggregateFunc) *SensorDataGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sdgb *SensorDataGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sdgb.path(ctx)
	if err != nil {
		return err
	}
	sdgb.sql = query
	return sdgb.sqlScan(ctx, v)
}

func (sdgb *SensorDataGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sdgb.fields {
		if !sensordata.ValidColumn(f) {
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

func (sdgb *SensorDataGroupBy) sqlQuery() *sql.Selector {
	selector := sdgb.sql.Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
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

// SensorDataSelect is the builder for selecting fields of SensorData entities.
type SensorDataSelect struct {
	*SensorDataQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sds *SensorDataSelect) Aggregate(fns ...AggregateFunc) *SensorDataSelect {
	sds.fns = append(sds.fns, fns...)
	return sds
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SensorDataSelect) Scan(ctx context.Context, v any) error {
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	sds.sql = sds.SensorDataQuery.sqlQuery(ctx)
	return sds.sqlScan(ctx, v)
}

func (sds *SensorDataSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(sds.fns))
	for _, fn := range sds.fns {
		aggregation = append(aggregation, fn(sds.sql))
	}
	switch n := len(*sds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		sds.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		sds.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := sds.sql.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sds *SensorDataSelect) Modify(modifiers ...func(s *sql.Selector)) *SensorDataSelect {
	sds.modifiers = append(sds.modifiers, modifiers...)
	return sds
}
