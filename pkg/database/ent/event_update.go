// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/crowdsecurity/crowdsec/pkg/database/ent/alert"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/event"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks      []Hook
	mutation   *EventMutation
	predicates []predicate.Event
}

// Where adds a new predicate for the builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetCreatedAt sets the created_at field.
func (eu *EventUpdate) SetCreatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetUpdatedAt sets the updated_at field.
func (eu *EventUpdate) SetUpdatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (eu *EventUpdate) SetNillableUpdatedAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetUpdatedAt(*t)
	}
	return eu
}

// SetTime sets the time field.
func (eu *EventUpdate) SetTime(t time.Time) *EventUpdate {
	eu.mutation.SetTime(t)
	return eu
}

// SetSerialized sets the serialized field.
func (eu *EventUpdate) SetSerialized(s string) *EventUpdate {
	eu.mutation.SetSerialized(s)
	return eu
}

// SetOwnerID sets the owner edge to Alert by id.
func (eu *EventUpdate) SetOwnerID(id int) *EventUpdate {
	eu.mutation.SetOwnerID(id)
	return eu
}

// SetNillableOwnerID sets the owner edge to Alert by id if the given value is not nil.
func (eu *EventUpdate) SetNillableOwnerID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetOwnerID(*id)
	}
	return eu
}

// SetOwner sets the owner edge to Alert.
func (eu *EventUpdate) SetOwner(a *Alert) *EventUpdate {
	return eu.SetOwnerID(a.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearOwner clears the "owner" edge to type Alert.
func (eu *EventUpdate) ClearOwner() *EventUpdate {
	eu.mutation.ClearOwner()
	return eu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.Serialized(); ok {
		if err := event.SerializedValidator(v); err != nil {
			return &ValidationError{Name: "serialized", err: fmt.Errorf("ent: validator failed for field \"serialized\": %w", err)}
		}
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := eu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldTime,
		})
	}
	if value, ok := eu.mutation.Serialized(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSerialized,
		})
	}
	if eu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerTable,
			Columns: []string{event.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerTable,
			Columns: []string{event.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// SetCreatedAt sets the created_at field.
func (euo *EventUpdateOne) SetCreatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetUpdatedAt sets the updated_at field.
func (euo *EventUpdateOne) SetUpdatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableUpdatedAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetUpdatedAt(*t)
	}
	return euo
}

// SetTime sets the time field.
func (euo *EventUpdateOne) SetTime(t time.Time) *EventUpdateOne {
	euo.mutation.SetTime(t)
	return euo
}

// SetSerialized sets the serialized field.
func (euo *EventUpdateOne) SetSerialized(s string) *EventUpdateOne {
	euo.mutation.SetSerialized(s)
	return euo
}

// SetOwnerID sets the owner edge to Alert by id.
func (euo *EventUpdateOne) SetOwnerID(id int) *EventUpdateOne {
	euo.mutation.SetOwnerID(id)
	return euo
}

// SetNillableOwnerID sets the owner edge to Alert by id if the given value is not nil.
func (euo *EventUpdateOne) SetNillableOwnerID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetOwnerID(*id)
	}
	return euo
}

// SetOwner sets the owner edge to Alert.
func (euo *EventUpdateOne) SetOwner(a *Alert) *EventUpdateOne {
	return euo.SetOwnerID(a.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearOwner clears the "owner" edge to type Alert.
func (euo *EventUpdateOne) ClearOwner() *EventUpdateOne {
	euo.mutation.ClearOwner()
	return euo
}

// Save executes the query and returns the updated entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.Serialized(); ok {
		if err := event.SerializedValidator(v); err != nil {
			return &ValidationError{Name: "serialized", err: fmt.Errorf("ent: validator failed for field \"serialized\": %w", err)}
		}
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Event.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := euo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldTime,
		})
	}
	if value, ok := euo.mutation.Serialized(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSerialized,
		})
	}
	if euo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerTable,
			Columns: []string{event.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerTable,
			Columns: []string{event.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
