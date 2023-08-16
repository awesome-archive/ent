// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetPriority sets the "priority" field.
func (tc *TaskCreate) SetPriority(t task.Priority) *TaskCreate {
	tc.mutation.SetPriority(t)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TaskCreate) SetNillablePriority(t *task.Priority) *TaskCreate {
	if t != nil {
		tc.SetPriority(*t)
	}
	return tc
}

// SetPriorities sets the "priorities" field.
func (tc *TaskCreate) SetPriorities(m map[string]task.Priority) *TaskCreate {
	tc.mutation.SetPriorities(m)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tc *TaskCreate) SetNillableName(s *string) *TaskCreate {
	if s != nil {
		tc.SetName(*s)
	}
	return tc
}

// SetOwner sets the "owner" field.
func (tc *TaskCreate) SetOwner(s string) *TaskCreate {
	tc.mutation.SetOwner(s)
	return tc
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOwner(s *string) *TaskCreate {
	if s != nil {
		tc.SetOwner(*s)
	}
	return tc
}

// SetOrder sets the "order" field.
func (tc *TaskCreate) SetOrder(i int) *TaskCreate {
	tc.mutation.SetOrder(i)
	return tc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOrder(i *int) *TaskCreate {
	if i != nil {
		tc.SetOrder(*i)
	}
	return tc
}

// SetOrderOption sets the "order_option" field.
func (tc *TaskCreate) SetOrderOption(i int) *TaskCreate {
	tc.mutation.SetOrderOption(i)
	return tc
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOrderOption(i *int) *TaskCreate {
	if i != nil {
		tc.SetOrderOption(*i)
	}
	return tc
}

// SetOp sets the "op" field.
func (tc *TaskCreate) SetOp(s string) *TaskCreate {
	tc.mutation.SetOpField(s)
	return tc
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOp(s *string) *TaskCreate {
	if s != nil {
		tc.SetOp(*s)
	}
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.gremlinSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.Priority(); !ok {
		v := enttask.DefaultPriority
		tc.mutation.SetPriority(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := enttask.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.GetOp(); !ok {
		v := enttask.DefaultOp
		tc.mutation.SetOpField(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Task.priority"`)}
	}
	if v, ok := tc.mutation.Priority(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.GetOp(); !ok {
		return &ValidationError{Name: "op", err: errors.New(`ent: missing required field "Task.op"`)}
	}
	if v, ok := tc.mutation.GetOp(); ok {
		if err := enttask.OpValidator(v); err != nil {
			return &ValidationError{Name: "op", err: fmt.Errorf(`ent: validator failed for field "Task.op": %w`, err)}
		}
	}
	return nil
}

func (tc *TaskCreate) gremlinSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := tc.gremlin().Query()
	if err := tc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Task{config: tc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	tc.mutation.id = &rnode.ID
	tc.mutation.done = true
	return rnode, nil
}

func (tc *TaskCreate) gremlin() *dsl.Traversal {
	v := g.AddV(enttask.Label)
	if value, ok := tc.mutation.Priority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, value)
	}
	if value, ok := tc.mutation.Priorities(); ok {
		v.Property(dsl.Single, enttask.FieldPriorities, value)
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, enttask.FieldCreatedAt, value)
	}
	if value, ok := tc.mutation.Name(); ok {
		v.Property(dsl.Single, enttask.FieldName, value)
	}
	if value, ok := tc.mutation.Owner(); ok {
		v.Property(dsl.Single, enttask.FieldOwner, value)
	}
	if value, ok := tc.mutation.Order(); ok {
		v.Property(dsl.Single, enttask.FieldOrder, value)
	}
	if value, ok := tc.mutation.OrderOption(); ok {
		v.Property(dsl.Single, enttask.FieldOrderOption, value)
	}
	if value, ok := tc.mutation.GetOp(); ok {
		v.Property(dsl.Single, enttask.FieldOp, value)
	}
	return v.ValueMap(true)
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}
