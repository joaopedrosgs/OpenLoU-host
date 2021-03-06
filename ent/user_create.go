// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/joaopedrosgs/openlou/ent/city"
	"github.com/joaopedrosgs/openlou/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPasswordHash sets the password_hash field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.mutation.SetPasswordHash(s)
	return uc
}

// SetGold sets the gold field.
func (uc *UserCreate) SetGold(i int) *UserCreate {
	uc.mutation.SetGold(i)
	return uc
}

// SetNillableGold sets the gold field if the given value is not nil.
func (uc *UserCreate) SetNillableGold(i *int) *UserCreate {
	if i != nil {
		uc.SetGold(*i)
	}
	return uc
}

// SetDiamonds sets the diamonds field.
func (uc *UserCreate) SetDiamonds(i int) *UserCreate {
	uc.mutation.SetDiamonds(i)
	return uc
}

// SetNillableDiamonds sets the diamonds field if the given value is not nil.
func (uc *UserCreate) SetNillableDiamonds(i *int) *UserCreate {
	if i != nil {
		uc.SetDiamonds(*i)
	}
	return uc
}

// SetDarkwood sets the darkwood field.
func (uc *UserCreate) SetDarkwood(i int) *UserCreate {
	uc.mutation.SetDarkwood(i)
	return uc
}

// SetNillableDarkwood sets the darkwood field if the given value is not nil.
func (uc *UserCreate) SetNillableDarkwood(i *int) *UserCreate {
	if i != nil {
		uc.SetDarkwood(*i)
	}
	return uc
}

// SetRunestone sets the runestone field.
func (uc *UserCreate) SetRunestone(i int) *UserCreate {
	uc.mutation.SetRunestone(i)
	return uc
}

// SetNillableRunestone sets the runestone field if the given value is not nil.
func (uc *UserCreate) SetNillableRunestone(i *int) *UserCreate {
	if i != nil {
		uc.SetRunestone(*i)
	}
	return uc
}

// SetVeritium sets the veritium field.
func (uc *UserCreate) SetVeritium(i int) *UserCreate {
	uc.mutation.SetVeritium(i)
	return uc
}

// SetNillableVeritium sets the veritium field if the given value is not nil.
func (uc *UserCreate) SetNillableVeritium(i *int) *UserCreate {
	if i != nil {
		uc.SetVeritium(*i)
	}
	return uc
}

// SetTrueseed sets the trueseed field.
func (uc *UserCreate) SetTrueseed(i int) *UserCreate {
	uc.mutation.SetTrueseed(i)
	return uc
}

// SetNillableTrueseed sets the trueseed field if the given value is not nil.
func (uc *UserCreate) SetNillableTrueseed(i *int) *UserCreate {
	if i != nil {
		uc.SetTrueseed(*i)
	}
	return uc
}

// SetRank sets the rank field.
func (uc *UserCreate) SetRank(i int) *UserCreate {
	uc.mutation.SetRank(i)
	return uc
}

// SetNillableRank sets the rank field if the given value is not nil.
func (uc *UserCreate) SetNillableRank(i *int) *UserCreate {
	if i != nil {
		uc.SetRank(*i)
	}
	return uc
}

// SetAllianceRank sets the alliance_rank field.
func (uc *UserCreate) SetAllianceRank(i int) *UserCreate {
	uc.mutation.SetAllianceRank(i)
	return uc
}

// SetNillableAllianceRank sets the alliance_rank field if the given value is not nil.
func (uc *UserCreate) SetNillableAllianceRank(i *int) *UserCreate {
	if i != nil {
		uc.SetAllianceRank(*i)
	}
	return uc
}

// AddCityIDs adds the cities edge to City by ids.
func (uc *UserCreate) AddCityIDs(ids ...int) *UserCreate {
	uc.mutation.AddCityIDs(ids...)
	return uc
}

// AddCities adds the cities edges to City.
func (uc *UserCreate) AddCities(c ...*City) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCityIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
		}
	}
	if _, ok := uc.mutation.PasswordHash(); !ok {
		return nil, errors.New("ent: missing required field \"password_hash\"")
	}
	if v, ok := uc.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"password_hash\": %v", err)
		}
	}
	if _, ok := uc.mutation.Gold(); !ok {
		v := user.DefaultGold
		uc.mutation.SetGold(v)
	}
	if _, ok := uc.mutation.Diamonds(); !ok {
		v := user.DefaultDiamonds
		uc.mutation.SetDiamonds(v)
	}
	if _, ok := uc.mutation.Darkwood(); !ok {
		v := user.DefaultDarkwood
		uc.mutation.SetDarkwood(v)
	}
	if _, ok := uc.mutation.Runestone(); !ok {
		v := user.DefaultRunestone
		uc.mutation.SetRunestone(v)
	}
	if _, ok := uc.mutation.Veritium(); !ok {
		v := user.DefaultVeritium
		uc.mutation.SetVeritium(v)
	}
	if _, ok := uc.mutation.Trueseed(); !ok {
		v := user.DefaultTrueseed
		uc.mutation.SetTrueseed(v)
	}
	if _, ok := uc.mutation.Rank(); !ok {
		v := user.DefaultRank
		uc.mutation.SetRank(v)
	}
	if _, ok := uc.mutation.AllianceRank(); !ok {
		v := user.DefaultAllianceRank
		uc.mutation.SetAllianceRank(v)
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		u.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = value
	}
	if value, ok := uc.mutation.PasswordHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
		u.PasswordHash = value
	}
	if value, ok := uc.mutation.Gold(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldGold,
		})
		u.Gold = value
	}
	if value, ok := uc.mutation.Diamonds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldDiamonds,
		})
		u.Diamonds = value
	}
	if value, ok := uc.mutation.Darkwood(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldDarkwood,
		})
		u.Darkwood = value
	}
	if value, ok := uc.mutation.Runestone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldRunestone,
		})
		u.Runestone = value
	}
	if value, ok := uc.mutation.Veritium(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldVeritium,
		})
		u.Veritium = value
	}
	if value, ok := uc.mutation.Trueseed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldTrueseed,
		})
		u.Trueseed = value
	}
	if value, ok := uc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldRank,
		})
		u.Rank = value
	}
	if value, ok := uc.mutation.AllianceRank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAllianceRank,
		})
		u.AllianceRank = value
	}
	if nodes := uc.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
