// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebook/ent/entc/integration/gremlin/ent/group"
	"github.com/facebook/ent/entc/integration/gremlin/ent/user"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetActive sets the active field.
func (gc *GroupCreate) SetActive(b bool) *GroupCreate {
	gc.mutation.SetActive(b)
	return gc
}

// SetNillableActive sets the active field if the given value is not nil.
func (gc *GroupCreate) SetNillableActive(b *bool) *GroupCreate {
	if b != nil {
		gc.SetActive(*b)
	}
	return gc
}

// SetExpire sets the expire field.
func (gc *GroupCreate) SetExpire(t time.Time) *GroupCreate {
	gc.mutation.SetExpire(t)
	return gc
}

// SetType sets the type field.
func (gc *GroupCreate) SetType(s string) *GroupCreate {
	gc.mutation.SetType(s)
	return gc
}

// SetNillableType sets the type field if the given value is not nil.
func (gc *GroupCreate) SetNillableType(s *string) *GroupCreate {
	if s != nil {
		gc.SetType(*s)
	}
	return gc
}

// SetMaxUsers sets the max_users field.
func (gc *GroupCreate) SetMaxUsers(i int) *GroupCreate {
	gc.mutation.SetMaxUsers(i)
	return gc
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gc *GroupCreate) SetNillableMaxUsers(i *int) *GroupCreate {
	if i != nil {
		gc.SetMaxUsers(*i)
	}
	return gc
}

// SetName sets the name field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// AddFileIDs adds the files edge to File by ids.
func (gc *GroupCreate) AddFileIDs(ids ...string) *GroupCreate {
	gc.mutation.AddFileIDs(ids...)
	return gc
}

// AddFiles adds the files edges to File.
func (gc *GroupCreate) AddFiles(f ...*File) *GroupCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gc.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (gc *GroupCreate) AddBlockedIDs(ids ...string) *GroupCreate {
	gc.mutation.AddBlockedIDs(ids...)
	return gc
}

// AddBlocked adds the blocked edges to User.
func (gc *GroupCreate) AddBlocked(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (gc *GroupCreate) AddUserIDs(ids ...string) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the users edges to User.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (gc *GroupCreate) SetInfoID(id string) *GroupCreate {
	gc.mutation.SetInfoID(id)
	return gc
}

// SetInfo sets the info edge to GroupInfo.
func (gc *GroupCreate) SetInfo(g *GroupInfo) *GroupCreate {
	return gc.SetInfoID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() {
	if _, ok := gc.mutation.Active(); !ok {
		v := group.DefaultActive
		gc.mutation.SetActive(v)
	}
	if _, ok := gc.mutation.MaxUsers(); !ok {
		v := group.DefaultMaxUsers
		gc.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New("ent: missing required field \"active\"")}
	}
	if _, ok := gc.mutation.Expire(); !ok {
		return &ValidationError{Name: "expire", err: errors.New("ent: missing required field \"expire\"")}
	}
	if v, ok := gc.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := gc.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return &ValidationError{Name: "max_users", err: fmt.Errorf("ent: validator failed for field \"max_users\": %w", err)}
		}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := gc.mutation.InfoID(); !ok {
		return &ValidationError{Name: "info", err: errors.New("ent: missing required edge \"info\"")}
	}
	return nil
}

func (gc *GroupCreate) gremlinSave(ctx context.Context) (*Group, error) {
	res := &gremlin.Response{}
	query, bindings := gc.gremlin().Query()
	if err := gc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gr := &Group{config: gc.config}
	if err := gr.FromResponse(res); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gc *GroupCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(group.Label)
	if value, ok := gc.mutation.Active(); ok {
		v.Property(dsl.Single, group.FieldActive, value)
	}
	if value, ok := gc.mutation.Expire(); ok {
		v.Property(dsl.Single, group.FieldExpire, value)
	}
	if value, ok := gc.mutation.GetType(); ok {
		v.Property(dsl.Single, group.FieldType, value)
	}
	if value, ok := gc.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, value)
	}
	if value, ok := gc.mutation.Name(); ok {
		v.Property(dsl.Single, group.FieldName, value)
	}
	for _, id := range gc.mutation.FilesIDs() {
		v.AddE(group.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.FilesLabel, id)),
		})
	}
	for _, id := range gc.mutation.BlockedIDs() {
		v.AddE(group.BlockedLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.BlockedLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.BlockedLabel, id)),
		})
	}
	for _, id := range gc.mutation.UsersIDs() {
		v.AddE(user.GroupsLabel).From(g.V(id)).InV()
	}
	for _, id := range gc.mutation.InfoIDs() {
		v.AddE(group.InfoLabel).To(g.V(id)).OutV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// GroupCreateBulk is the builder for creating a bulk of Group entities.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
}
