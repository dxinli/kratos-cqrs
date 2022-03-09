// Code generated by entc, DO NOT EDIT.

package ent

import (
	"kratos-cqrs/app/logger/service/internal/data/ent/sensor"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensordata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   sensor.Table,
			Columns: sensor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensor.FieldID,
			},
		},
		Type: "Sensor",
		Fields: map[string]*sqlgraph.FieldSpec{
			sensor.FieldType:     {Type: field.TypeString, Column: sensor.FieldType},
			sensor.FieldLocation: {Type: field.TypeString, Column: sensor.FieldLocation},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   sensordata.Table,
			Columns: sensordata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sensordata.FieldID,
			},
		},
		Type: "SensorData",
		Fields: map[string]*sqlgraph.FieldSpec{
			sensordata.FieldTime:        {Type: field.TypeInt64, Column: sensordata.FieldTime},
			sensordata.FieldSensorID:    {Type: field.TypeInt, Column: sensordata.FieldSensorID},
			sensordata.FieldTemperature: {Type: field.TypeFloat64, Column: sensordata.FieldTemperature},
			sensordata.FieldCPU:         {Type: field.TypeFloat64, Column: sensordata.FieldCPU},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (sq *SensorQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SensorQuery builder.
func (sq *SensorQuery) Filter() *SensorFilter {
	return &SensorFilter{sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SensorMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SensorMutation builder.
func (m *SensorMutation) Filter() *SensorFilter {
	return &SensorFilter{m}
}

// SensorFilter provides a generic filtering capability at runtime for SensorQuery.
type SensorFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *SensorFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *SensorFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(sensor.FieldID))
}

// WhereType applies the entql string predicate on the type field.
func (f *SensorFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(sensor.FieldType))
}

// WhereLocation applies the entql string predicate on the location field.
func (f *SensorFilter) WhereLocation(p entql.StringP) {
	f.Where(p.Field(sensor.FieldLocation))
}

// addPredicate implements the predicateAdder interface.
func (sdq *SensorDataQuery) addPredicate(pred func(s *sql.Selector)) {
	sdq.predicates = append(sdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SensorDataQuery builder.
func (sdq *SensorDataQuery) Filter() *SensorDataFilter {
	return &SensorDataFilter{sdq}
}

// addPredicate implements the predicateAdder interface.
func (m *SensorDataMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SensorDataMutation builder.
func (m *SensorDataMutation) Filter() *SensorDataFilter {
	return &SensorDataFilter{m}
}

// SensorDataFilter provides a generic filtering capability at runtime for SensorDataQuery.
type SensorDataFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *SensorDataFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *SensorDataFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(sensordata.FieldID))
}

// WhereTime applies the entql int64 predicate on the time field.
func (f *SensorDataFilter) WhereTime(p entql.Int64P) {
	f.Where(p.Field(sensordata.FieldTime))
}

// WhereSensorID applies the entql int predicate on the sensor_id field.
func (f *SensorDataFilter) WhereSensorID(p entql.IntP) {
	f.Where(p.Field(sensordata.FieldSensorID))
}

// WhereTemperature applies the entql float64 predicate on the temperature field.
func (f *SensorDataFilter) WhereTemperature(p entql.Float64P) {
	f.Where(p.Field(sensordata.FieldTemperature))
}

// WhereCPU applies the entql float64 predicate on the cpu field.
func (f *SensorDataFilter) WhereCPU(p entql.Float64P) {
	f.Where(p.Field(sensordata.FieldCPU))
}
