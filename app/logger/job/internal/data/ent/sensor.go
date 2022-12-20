// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-cqrs/app/logger/job/internal/data/ent/sensor"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Sensor is the model entity for the Sensor schema.
type Sensor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 传感器类型
	Type string `json:"type,omitempty"`
	// 所在位置
	Location string `json:"location,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sensor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sensor.FieldID:
			values[i] = new(sql.NullInt64)
		case sensor.FieldType, sensor.FieldLocation:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sensor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sensor fields.
func (s *Sensor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sensor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sensor.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = value.String
			}
		case sensor.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				s.Location = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Sensor.
// Note that you need to call Sensor.Unwrap() before calling this method if this Sensor
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sensor) Update() *SensorUpdateOne {
	return (&SensorClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sensor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sensor) Unwrap() *Sensor {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sensor is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sensor) String() string {
	var builder strings.Builder
	builder.WriteString("Sensor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("type=")
	builder.WriteString(s.Type)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(s.Location)
	builder.WriteByte(')')
	return builder.String()
}

// Sensors is a parsable slice of Sensor.
type Sensors []*Sensor

func (s Sensors) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
