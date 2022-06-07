// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgefield/ent/card"
	"entgo.io/ent/entc/integration/edgefield/ent/metadata"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID int `json:"parent_id,omitempty"`
	// SpouseID holds the value of the "spouse_id" field.
	SpouseID int `json:"spouse_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Pets holds the value of the pets edge.
	Pets []*Pet `json:"pets,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *User `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*User `json:"children,omitempty"`
	// Spouse holds the value of the spouse edge.
	Spouse *User `json:"spouse,omitempty"`
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Info holds the value of the info edge.
	Info []*Info `json:"info,omitempty"`
	// Rentals holds the value of the rentals edge.
	Rentals []*Rental `json:"rentals,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsOrErr() ([]*Pet, error) {
	if e.loadedTypes[0] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.Spouse == nil {
			// The edge spouse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Spouse, nil
	}
	return nil, &NotLoadedError{edge: "spouse"}
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[4] {
		if e.Card == nil {
			// The edge card was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) MetadataOrErr() (*Metadata, error) {
	if e.loadedTypes[5] {
		if e.Metadata == nil {
			// The edge metadata was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metadata.Label}
		}
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// InfoOrErr returns the Info value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InfoOrErr() ([]*Info, error) {
	if e.loadedTypes[6] {
		return e.Info, nil
	}
	return nil, &NotLoadedError{edge: "info"}
}

// RentalsOrErr returns the Rentals value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RentalsOrErr() ([]*Rental, error) {
	if e.loadedTypes[7] {
		return e.Rentals, nil
	}
	return nil, &NotLoadedError{edge: "rentals"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldParentID, user.FieldSpouseID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				u.ParentID = int(value.Int64)
			}
		case user.FieldSpouseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field spouse_id", values[i])
			} else if value.Valid {
				u.SpouseID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{config: u.config}).QueryPets(u)
}

// QueryParent queries the "parent" edge of the User entity.
func (u *User) QueryParent() *UserQuery {
	return (&UserClient{config: u.config}).QueryParent(u)
}

// QueryChildren queries the "children" edge of the User entity.
func (u *User) QueryChildren() *UserQuery {
	return (&UserClient{config: u.config}).QueryChildren(u)
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (u *User) QuerySpouse() *UserQuery {
	return (&UserClient{config: u.config}).QuerySpouse(u)
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return (&UserClient{config: u.config}).QueryCard(u)
}

// QueryMetadata queries the "metadata" edge of the User entity.
func (u *User) QueryMetadata() *MetadataQuery {
	return (&UserClient{config: u.config}).QueryMetadata(u)
}

// QueryInfo queries the "info" edge of the User entity.
func (u *User) QueryInfo() *InfoQuery {
	return (&UserClient{config: u.config}).QueryInfo(u)
}

// QueryRentals queries the "rentals" edge of the User entity.
func (u *User) QueryRentals() *RentalQuery {
	return (&UserClient{config: u.config}).QueryRentals(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ParentID))
	builder.WriteString(", ")
	builder.WriteString("spouse_id=")
	builder.WriteString(fmt.Sprintf("%v", u.SpouseID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
