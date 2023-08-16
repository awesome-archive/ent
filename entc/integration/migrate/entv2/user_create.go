// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/car"
	"entgo.io/ent/entc/integration/migrate/entv2/pet"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetMixedString sets the "mixed_string" field.
func (uc *UserCreate) SetMixedString(s string) *UserCreate {
	uc.mutation.SetMixedString(s)
	return uc
}

// SetNillableMixedString sets the "mixed_string" field if the given value is not nil.
func (uc *UserCreate) SetNillableMixedString(s *string) *UserCreate {
	if s != nil {
		uc.SetMixedString(*s)
	}
	return uc
}

// SetMixedEnum sets the "mixed_enum" field.
func (uc *UserCreate) SetMixedEnum(ue user.MixedEnum) *UserCreate {
	uc.mutation.SetMixedEnum(ue)
	return uc
}

// SetNillableMixedEnum sets the "mixed_enum" field if the given value is not nil.
func (uc *UserCreate) SetNillableMixedEnum(ue *user.MixedEnum) *UserCreate {
	if ue != nil {
		uc.SetMixedEnum(*ue)
	}
	return uc
}

// SetActive sets the "active" field.
func (uc *UserCreate) SetActive(b bool) *UserCreate {
	uc.mutation.SetActive(b)
	return uc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uc *UserCreate) SetNillableActive(b *bool) *UserCreate {
	if b != nil {
		uc.SetActive(*b)
	}
	return uc
}

// SetAge sets the "age" field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetDescription sets the "description" field.
func (uc *UserCreate) SetDescription(s string) *UserCreate {
	uc.mutation.SetDescription(s)
	return uc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uc *UserCreate) SetNillableDescription(s *string) *UserCreate {
	if s != nil {
		uc.SetDescription(*s)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetBuffer sets the "buffer" field.
func (uc *UserCreate) SetBuffer(b []byte) *UserCreate {
	uc.mutation.SetBuffer(b)
	return uc
}

// SetTitle sets the "title" field.
func (uc *UserCreate) SetTitle(s string) *UserCreate {
	uc.mutation.SetTitle(s)
	return uc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uc *UserCreate) SetNillableTitle(s *string) *UserCreate {
	if s != nil {
		uc.SetTitle(*s)
	}
	return uc
}

// SetNewName sets the "new_name" field.
func (uc *UserCreate) SetNewName(s string) *UserCreate {
	uc.mutation.SetNewName(s)
	return uc
}

// SetNillableNewName sets the "new_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableNewName(s *string) *UserCreate {
	if s != nil {
		uc.SetNewName(*s)
	}
	return uc
}

// SetNewToken sets the "new_token" field.
func (uc *UserCreate) SetNewToken(s string) *UserCreate {
	uc.mutation.SetNewToken(s)
	return uc
}

// SetNillableNewToken sets the "new_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableNewToken(s *string) *UserCreate {
	if s != nil {
		uc.SetNewToken(*s)
	}
	return uc
}

// SetBlob sets the "blob" field.
func (uc *UserCreate) SetBlob(b []byte) *UserCreate {
	uc.mutation.SetBlob(b)
	return uc
}

// SetState sets the "state" field.
func (uc *UserCreate) SetState(u user.State) *UserCreate {
	uc.mutation.SetState(u)
	return uc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uc *UserCreate) SetNillableState(u *user.State) *UserCreate {
	if u != nil {
		uc.SetState(*u)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *user.Status) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetWorkplace sets the "workplace" field.
func (uc *UserCreate) SetWorkplace(s string) *UserCreate {
	uc.mutation.SetWorkplace(s)
	return uc
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uc *UserCreate) SetNillableWorkplace(s *string) *UserCreate {
	if s != nil {
		uc.SetWorkplace(*s)
	}
	return uc
}

// SetRoles sets the "roles" field.
func (uc *UserCreate) SetRoles(s []string) *UserCreate {
	uc.mutation.SetRoles(s)
	return uc
}

// SetDefaultExpr sets the "default_expr" field.
func (uc *UserCreate) SetDefaultExpr(s string) *UserCreate {
	uc.mutation.SetDefaultExpr(s)
	return uc
}

// SetNillableDefaultExpr sets the "default_expr" field if the given value is not nil.
func (uc *UserCreate) SetNillableDefaultExpr(s *string) *UserCreate {
	if s != nil {
		uc.SetDefaultExpr(*s)
	}
	return uc
}

// SetDefaultExprs sets the "default_exprs" field.
func (uc *UserCreate) SetDefaultExprs(s string) *UserCreate {
	uc.mutation.SetDefaultExprs(s)
	return uc
}

// SetNillableDefaultExprs sets the "default_exprs" field if the given value is not nil.
func (uc *UserCreate) SetNillableDefaultExprs(s *string) *UserCreate {
	if s != nil {
		uc.SetDefaultExprs(*s)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetDropOptional sets the "drop_optional" field.
func (uc *UserCreate) SetDropOptional(s string) *UserCreate {
	uc.mutation.SetDropOptional(s)
	return uc
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (uc *UserCreate) SetNillableDropOptional(s *string) *UserCreate {
	if s != nil {
		uc.SetDropOptional(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddCarIDs adds the "car" edge to the Car entity by IDs.
func (uc *UserCreate) AddCarIDs(ids ...int) *UserCreate {
	uc.mutation.AddCarIDs(ids...)
	return uc
}

// AddCar adds the "car" edges to the Car entity.
func (uc *UserCreate) AddCar(c ...*Car) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCarIDs(ids...)
}

// SetPetsID sets the "pets" edge to the Pet entity by ID.
func (uc *UserCreate) SetPetsID(id int) *UserCreate {
	uc.mutation.SetPetsID(id)
	return uc
}

// SetNillablePetsID sets the "pets" edge to the Pet entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillablePetsID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetPetsID(*id)
	}
	return uc
}

// SetPets sets the "pets" edge to the Pet entity.
func (uc *UserCreate) SetPets(p *Pet) *UserCreate {
	return uc.SetPetsID(p.ID)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (uc *UserCreate) AddFriendIDs(ids ...int) *UserCreate {
	uc.mutation.AddFriendIDs(ids...)
	return uc
}

// AddFriends adds the "friends" edges to the User entity.
func (uc *UserCreate) AddFriends(u ...*User) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFriendIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.MixedString(); !ok {
		v := user.DefaultMixedString
		uc.mutation.SetMixedString(v)
	}
	if _, ok := uc.mutation.MixedEnum(); !ok {
		v := user.DefaultMixedEnum
		uc.mutation.SetMixedEnum(v)
	}
	if _, ok := uc.mutation.Active(); !ok {
		v := user.DefaultActive
		uc.mutation.SetActive(v)
	}
	if _, ok := uc.mutation.Phone(); !ok {
		v := user.DefaultPhone
		uc.mutation.SetPhone(v)
	}
	if _, ok := uc.mutation.Buffer(); !ok {
		v := user.DefaultBuffer()
		uc.mutation.SetBuffer(v)
	}
	if _, ok := uc.mutation.Title(); !ok {
		v := user.DefaultTitle
		uc.mutation.SetTitle(v)
	}
	if _, ok := uc.mutation.NewToken(); !ok {
		v := user.DefaultNewToken()
		uc.mutation.SetNewToken(v)
	}
	if _, ok := uc.mutation.State(); !ok {
		v := user.DefaultState
		uc.mutation.SetState(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.DropOptional(); !ok {
		v := user.DefaultDropOptional()
		uc.mutation.SetDropOptional(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.MixedString(); !ok {
		return &ValidationError{Name: "mixed_string", err: errors.New(`entv2: missing required field "User.mixed_string"`)}
	}
	if _, ok := uc.mutation.MixedEnum(); !ok {
		return &ValidationError{Name: "mixed_enum", err: errors.New(`entv2: missing required field "User.mixed_enum"`)}
	}
	if v, ok := uc.mutation.MixedEnum(); ok {
		if err := user.MixedEnumValidator(v); err != nil {
			return &ValidationError{Name: "mixed_enum", err: fmt.Errorf(`entv2: validator failed for field "User.mixed_enum": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`entv2: missing required field "User.active"`)}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`entv2: missing required field "User.age"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entv2: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`entv2: missing required field "User.nickname"`)}
	}
	if v, ok := uc.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`entv2: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`entv2: missing required field "User.phone"`)}
	}
	if _, ok := uc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`entv2: missing required field "User.title"`)}
	}
	if _, ok := uc.mutation.NewToken(); !ok {
		return &ValidationError{Name: "new_token", err: errors.New(`entv2: missing required field "User.new_token"`)}
	}
	if v, ok := uc.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv2: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := uc.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv2: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entv2: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entv2: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.DropOptional(); !ok {
		return &ValidationError{Name: "drop_optional", err: errors.New(`entv2: missing required field "User.drop_optional"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.MixedString(); ok {
		_spec.SetField(user.FieldMixedString, field.TypeString, value)
		_node.MixedString = value
	}
	if value, ok := uc.mutation.MixedEnum(); ok {
		_spec.SetField(user.FieldMixedEnum, field.TypeEnum, value)
		_node.MixedEnum = value
	}
	if value, ok := uc.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.Buffer(); ok {
		_spec.SetField(user.FieldBuffer, field.TypeBytes, value)
		_node.Buffer = value
	}
	if value, ok := uc.mutation.Title(); ok {
		_spec.SetField(user.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := uc.mutation.NewName(); ok {
		_spec.SetField(user.FieldNewName, field.TypeString, value)
		_node.NewName = value
	}
	if value, ok := uc.mutation.NewToken(); ok {
		_spec.SetField(user.FieldNewToken, field.TypeString, value)
		_node.NewToken = value
	}
	if value, ok := uc.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
		_node.Blob = value
	}
	if value, ok := uc.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
		_node.Workplace = value
	}
	if value, ok := uc.mutation.Roles(); ok {
		_spec.SetField(user.FieldRoles, field.TypeJSON, value)
		_node.Roles = value
	}
	if value, ok := uc.mutation.DefaultExpr(); ok {
		_spec.SetField(user.FieldDefaultExpr, field.TypeString, value)
		_node.DefaultExpr = value
	}
	if value, ok := uc.mutation.DefaultExprs(); ok {
		_spec.SetField(user.FieldDefaultExprs, field.TypeString, value)
		_node.DefaultExprs = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
		_node.DropOptional = value
	}
	if nodes := uc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
