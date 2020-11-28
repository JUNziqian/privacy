// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/kallydev/privacy/ent/predicate"
	"github.com/kallydev/privacy/ent/qqmodel"
)

// QQModelQuery is the builder for querying QQModel entities.
type QQModelQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.QQModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (qmq *QQModelQuery) Where(ps ...predicate.QQModel) *QQModelQuery {
	qmq.predicates = append(qmq.predicates, ps...)
	return qmq
}

// Limit adds a limit step to the query.
func (qmq *QQModelQuery) Limit(limit int) *QQModelQuery {
	qmq.limit = &limit
	return qmq
}

// Offset adds an offset step to the query.
func (qmq *QQModelQuery) Offset(offset int) *QQModelQuery {
	qmq.offset = &offset
	return qmq
}

// Order adds an order step to the query.
func (qmq *QQModelQuery) Order(o ...OrderFunc) *QQModelQuery {
	qmq.order = append(qmq.order, o...)
	return qmq
}

// First returns the first QQModel entity in the query. Returns *NotFoundError when no qqmodel was found.
func (qmq *QQModelQuery) First(ctx context.Context) (*QQModel, error) {
	nodes, err := qmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{qqmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qmq *QQModelQuery) FirstX(ctx context.Context) *QQModel {
	node, err := qmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first QQModel id in the query. Returns *NotFoundError when no id was found.
func (qmq *QQModelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{qqmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qmq *QQModelQuery) FirstIDX(ctx context.Context) int {
	id, err := qmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only QQModel entity in the query, returns an error if not exactly one entity was returned.
func (qmq *QQModelQuery) Only(ctx context.Context) (*QQModel, error) {
	nodes, err := qmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{qqmodel.Label}
	default:
		return nil, &NotSingularError{qqmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qmq *QQModelQuery) OnlyX(ctx context.Context) *QQModel {
	node, err := qmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only QQModel id in the query, returns an error if not exactly one id was returned.
func (qmq *QQModelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = &NotSingularError{qqmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qmq *QQModelQuery) OnlyIDX(ctx context.Context) int {
	id, err := qmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QQModels.
func (qmq *QQModelQuery) All(ctx context.Context) ([]*QQModel, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qmq *QQModelQuery) AllX(ctx context.Context) []*QQModel {
	nodes, err := qmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of QQModel ids.
func (qmq *QQModelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := qmq.Select(qqmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qmq *QQModelQuery) IDsX(ctx context.Context) []int {
	ids, err := qmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qmq *QQModelQuery) Count(ctx context.Context) (int, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qmq *QQModelQuery) CountX(ctx context.Context) int {
	count, err := qmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qmq *QQModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qmq *QQModelQuery) ExistX(ctx context.Context) bool {
	exist, err := qmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qmq *QQModelQuery) Clone() *QQModelQuery {
	if qmq == nil {
		return nil
	}
	return &QQModelQuery{
		config:     qmq.config,
		limit:      qmq.limit,
		offset:     qmq.offset,
		order:      append([]OrderFunc{}, qmq.order...),
		unique:     append([]string{}, qmq.unique...),
		predicates: append([]predicate.QQModel{}, qmq.predicates...),
		// clone intermediate query.
		sql:  qmq.sql.Clone(),
		path: qmq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		QqNumber int64 `json:"qq_number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QQModel.Query().
//		GroupBy(qqmodel.FieldQqNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (qmq *QQModelQuery) GroupBy(field string, fields ...string) *QQModelGroupBy {
	group := &QQModelGroupBy{config: qmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qmq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		QqNumber int64 `json:"qq_number,omitempty"`
//	}
//
//	client.QQModel.Query().
//		Select(qqmodel.FieldQqNumber).
//		Scan(ctx, &v)
//
func (qmq *QQModelQuery) Select(field string, fields ...string) *QQModelSelect {
	selector := &QQModelSelect{config: qmq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qmq.sqlQuery(), nil
	}
	return selector
}

func (qmq *QQModelQuery) prepareQuery(ctx context.Context) error {
	if qmq.path != nil {
		prev, err := qmq.path(ctx)
		if err != nil {
			return err
		}
		qmq.sql = prev
	}
	return nil
}

func (qmq *QQModelQuery) sqlAll(ctx context.Context) ([]*QQModel, error) {
	var (
		nodes = []*QQModel{}
		_spec = qmq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &QQModel{config: qmq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, qmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (qmq *QQModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qmq.querySpec()
	return sqlgraph.CountNodes(ctx, qmq.driver, _spec)
}

func (qmq *QQModelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := qmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (qmq *QQModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qqmodel.Table,
			Columns: qqmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qqmodel.FieldID,
			},
		},
		From:   qmq.sql,
		Unique: true,
	}
	if ps := qmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, qqmodel.ValidColumn)
			}
		}
	}
	return _spec
}

func (qmq *QQModelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(qmq.driver.Dialect())
	t1 := builder.Table(qqmodel.Table)
	selector := builder.Select(t1.Columns(qqmodel.Columns...)...).From(t1)
	if qmq.sql != nil {
		selector = qmq.sql
		selector.Select(selector.Columns(qqmodel.Columns...)...)
	}
	for _, p := range qmq.predicates {
		p(selector)
	}
	for _, p := range qmq.order {
		p(selector, qqmodel.ValidColumn)
	}
	if offset := qmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QQModelGroupBy is the builder for group-by QQModel entities.
type QQModelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qmgb *QQModelGroupBy) Aggregate(fns ...AggregateFunc) *QQModelGroupBy {
	qmgb.fns = append(qmgb.fns, fns...)
	return qmgb
}

// Scan applies the group-by query and scan the result into the given value.
func (qmgb *QQModelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := qmgb.path(ctx)
	if err != nil {
		return err
	}
	qmgb.sql = query
	return qmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qmgb *QQModelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := qmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(qmgb.fields) > 1 {
		return nil, errors.New("ent: QQModelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := qmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qmgb *QQModelGroupBy) StringsX(ctx context.Context) []string {
	v, err := qmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qmgb *QQModelGroupBy) StringX(ctx context.Context) string {
	v, err := qmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(qmgb.fields) > 1 {
		return nil, errors.New("ent: QQModelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := qmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qmgb *QQModelGroupBy) IntsX(ctx context.Context) []int {
	v, err := qmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qmgb *QQModelGroupBy) IntX(ctx context.Context) int {
	v, err := qmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(qmgb.fields) > 1 {
		return nil, errors.New("ent: QQModelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := qmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qmgb *QQModelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := qmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qmgb *QQModelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := qmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(qmgb.fields) > 1 {
		return nil, errors.New("ent: QQModelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := qmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qmgb *QQModelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := qmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (qmgb *QQModelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qmgb *QQModelGroupBy) BoolX(ctx context.Context) bool {
	v, err := qmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qmgb *QQModelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range qmgb.fields {
		if !qqmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qmgb *QQModelGroupBy) sqlQuery() *sql.Selector {
	selector := qmgb.sql
	columns := make([]string, 0, len(qmgb.fields)+len(qmgb.fns))
	columns = append(columns, qmgb.fields...)
	for _, fn := range qmgb.fns {
		columns = append(columns, fn(selector, qqmodel.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(qmgb.fields...)
}

// QQModelSelect is the builder for select fields of QQModel entities.
type QQModelSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (qms *QQModelSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := qms.path(ctx)
	if err != nil {
		return err
	}
	qms.sql = query
	return qms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qms *QQModelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := qms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(qms.fields) > 1 {
		return nil, errors.New("ent: QQModelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := qms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qms *QQModelSelect) StringsX(ctx context.Context) []string {
	v, err := qms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qms *QQModelSelect) StringX(ctx context.Context) string {
	v, err := qms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(qms.fields) > 1 {
		return nil, errors.New("ent: QQModelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := qms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qms *QQModelSelect) IntsX(ctx context.Context) []int {
	v, err := qms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qms *QQModelSelect) IntX(ctx context.Context) int {
	v, err := qms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(qms.fields) > 1 {
		return nil, errors.New("ent: QQModelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := qms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qms *QQModelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := qms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qms *QQModelSelect) Float64X(ctx context.Context) float64 {
	v, err := qms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(qms.fields) > 1 {
		return nil, errors.New("ent: QQModelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := qms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qms *QQModelSelect) BoolsX(ctx context.Context) []bool {
	v, err := qms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (qms *QQModelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qqmodel.Label}
	default:
		err = fmt.Errorf("ent: QQModelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qms *QQModelSelect) BoolX(ctx context.Context) bool {
	v, err := qms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qms *QQModelSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range qms.fields {
		if !qqmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := qms.sqlQuery().Query()
	if err := qms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qms *QQModelSelect) sqlQuery() sql.Querier {
	selector := qms.sql
	selector.Select(selector.Columns(qms.fields...)...)
	return selector
}