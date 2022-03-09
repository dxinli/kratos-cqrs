// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-cqrs/app/logger/service/internal/data/ent/predicate"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensor"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorUpdate is the builder for updating Sensor entities.
type SensorUpdate struct {
	config
	hooks    []Hook
	mutation *SensorMutation
}

// Where appends a list predicates to the SensorUpdate builder.
func (su *SensorUpdate) Where(ps ...predicate.Sensor) *SensorUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetType sets the "type" field.
func (su *SensorUpdate) SetType(s string) *SensorUpdate {
	su.mutation.SetType(s)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *SensorUpdate) SetNillableType(s *string) *SensorUpdate {
	if s != nil {
		su.SetType(*s)
	}
	return su
}

// SetLocation sets the "location" field.
func (su *SensorUpdate) SetLocation(s string) *SensorUpdate {
	su.mutation.SetLocation(s)
	return su
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (su *SensorUpdate) SetNillableLocation(s *string) *SensorUpdate {
	if s != nil {
		su.SetLocation(*s)
	}
	return su
}

// Mutation returns the SensorMutation object of the builder.
func (su *SensorUpdate) Mutation() *SensorMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SensorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SensorUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SensorUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SensorUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SensorUpdate) check() error {
	if v, ok := su.mutation.GetType(); ok {
		if err := sensor.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := su.mutation.Location(); ok {
		if err := sensor.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	return nil
}

func (su *SensorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sensor.Table,
			Columns: sensor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensor.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sensor.FieldType,
		})
	}
	if value, ok := su.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sensor.FieldLocation,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sensor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SensorUpdateOne is the builder for updating a single Sensor entity.
type SensorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SensorMutation
}

// SetType sets the "type" field.
func (suo *SensorUpdateOne) SetType(s string) *SensorUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *SensorUpdateOne) SetNillableType(s *string) *SensorUpdateOne {
	if s != nil {
		suo.SetType(*s)
	}
	return suo
}

// SetLocation sets the "location" field.
func (suo *SensorUpdateOne) SetLocation(s string) *SensorUpdateOne {
	suo.mutation.SetLocation(s)
	return suo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (suo *SensorUpdateOne) SetNillableLocation(s *string) *SensorUpdateOne {
	if s != nil {
		suo.SetLocation(*s)
	}
	return suo
}

// Mutation returns the SensorMutation object of the builder.
func (suo *SensorUpdateOne) Mutation() *SensorMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SensorUpdateOne) Select(field string, fields ...string) *SensorUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sensor entity.
func (suo *SensorUpdateOne) Save(ctx context.Context) (*Sensor, error) {
	var (
		err  error
		node *Sensor
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SensorUpdateOne) SaveX(ctx context.Context) *Sensor {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SensorUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SensorUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SensorUpdateOne) check() error {
	if v, ok := suo.mutation.GetType(); ok {
		if err := sensor.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Location(); ok {
		if err := sensor.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	return nil
}

func (suo *SensorUpdateOne) sqlSave(ctx context.Context) (_node *Sensor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sensor.Table,
			Columns: sensor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensor.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Sensor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sensor.FieldID)
		for _, f := range fields {
			if !sensor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sensor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sensor.FieldType,
		})
	}
	if value, ok := suo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sensor.FieldLocation,
		})
	}
	_node = &Sensor{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sensor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
