// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// FieldTypeQuery is the builder for querying FieldType entities.
type FieldTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.FieldType
	// intermediate query.
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (ftq *FieldTypeQuery) Where(ps ...predicate.FieldType) *FieldTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FieldTypeQuery) Limit(limit int) *FieldTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FieldTypeQuery) Offset(offset int) *FieldTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Order adds an order step to the query.
func (ftq *FieldTypeQuery) Order(o ...Order) *FieldTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// First returns the first FieldType entity in the query. Returns *ErrNotFound when no fieldtype was found.
func (ftq *FieldTypeQuery) First(ctx context.Context) (*FieldType, error) {
	fts, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fts) == 0 {
		return nil, &ErrNotFound{fieldtype.Label}
	}
	return fts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstX(ctx context.Context) *FieldType {
	ft, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ft
}

// FirstID returns the first FieldType id in the query. Returns *ErrNotFound when no id was found.
func (ftq *FieldTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{fieldtype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstXID(ctx context.Context) string {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FieldType entity in the query, returns an error if not exactly one entity was returned.
func (ftq *FieldTypeQuery) Only(ctx context.Context) (*FieldType, error) {
	fts, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fts) {
	case 1:
		return fts[0], nil
	case 0:
		return nil, &ErrNotFound{fieldtype.Label}
	default:
		return nil, &ErrNotSingular{fieldtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyX(ctx context.Context) *FieldType {
	ft, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// OnlyID returns the only FieldType id in the query, returns an error if not exactly one id was returned.
func (ftq *FieldTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{fieldtype.Label}
	default:
		err = &ErrNotSingular{fieldtype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyXID(ctx context.Context) string {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FieldTypes.
func (ftq *FieldTypeQuery) All(ctx context.Context) ([]*FieldType, error) {
	return ftq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FieldTypeQuery) AllX(ctx context.Context) []*FieldType {
	fts, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fts
}

// IDs executes the query and returns a list of FieldType ids.
func (ftq *FieldTypeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := ftq.Select(fieldtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FieldTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FieldTypeQuery) Count(ctx context.Context) (int, error) {
	return ftq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FieldTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FieldTypeQuery) Exist(ctx context.Context) (bool, error) {
	return ftq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FieldTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FieldTypeQuery) Clone() *FieldTypeQuery {
	return &FieldTypeQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]Order{}, ftq.order...),
		unique:     append([]string{}, ftq.unique...),
		predicates: append([]predicate.FieldType{}, ftq.predicates...),
		// clone intermediate query.
		gremlin: ftq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FieldType.Query().
//		GroupBy(fieldtype.FieldInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FieldTypeQuery) GroupBy(field string, fields ...string) *FieldTypeGroupBy {
	group := &FieldTypeGroupBy{config: ftq.config}
	group.fields = append([]string{field}, fields...)
	group.gremlin = ftq.gremlinQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.FieldType.Query().
//		Select(fieldtype.FieldInt).
//		Scan(ctx, &v)
//
func (ftq *FieldTypeQuery) Select(field string, fields ...string) *FieldTypeSelect {
	selector := &FieldTypeSelect{config: ftq.config}
	selector.fields = append([]string{field}, fields...)
	selector.gremlin = ftq.gremlinQuery()
	return selector
}

func (ftq *FieldTypeQuery) gremlinAll(ctx context.Context) ([]*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().ValueMap(true).Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fts FieldTypes
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	fts.config(ftq.config)
	return fts, nil
}

func (ftq *FieldTypeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().Count().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftq *FieldTypeQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().HasNext().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (ftq *FieldTypeQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(fieldtype.Label)
	if ftq.gremlin != nil {
		v = ftq.gremlin.Clone()
	}
	for _, p := range ftq.predicates {
		p(v)
	}
	if len(ftq.order) > 0 {
		v.Order()
		for _, p := range ftq.order {
			p(v)
		}
	}
	switch limit, offset := ftq.limit, ftq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ftq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// FieldTypeGroupBy is the builder for group-by FieldType entities.
type FieldTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FieldTypeGroupBy) Aggregate(fns ...Aggregate) *FieldTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ftgb *FieldTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ftgb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftgb *FieldTypeGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ftgb.gremlinQuery().Query()
	if err := ftgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ftgb.fields)+len(ftgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ftgb *FieldTypeGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ftgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ftgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ftgb.gremlin.Group().
		By(__.Values(ftgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// FieldTypeSelect is the builder for select fields of FieldType entities.
type FieldTypeSelect struct {
	config
	fields []string
	// intermediate queries.
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scan the result into the given value.
func (fts *FieldTypeSelect) Scan(ctx context.Context, v interface{}) error {
	return fts.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fts *FieldTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fts *FieldTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fts *FieldTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fts *FieldTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fts *FieldTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fts *FieldTypeSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(fts.fields) == 1 {
		if fts.fields[0] != fieldtype.FieldID {
			traversal = fts.gremlin.Values(fts.fields...)
		} else {
			traversal = fts.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(fts.fields))
		for i, f := range fts.fields {
			fields[i] = f
		}
		traversal = fts.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fts.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(fts.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}