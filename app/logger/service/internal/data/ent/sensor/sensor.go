// Code generated by ent, DO NOT EDIT.

package sensor

const (
	// Label holds the string label denoting the sensor type in the database.
	Label = "sensor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// Table holds the table name of the sensor in the database.
	Table = "sensors"
)

// Columns holds all SQL columns for sensor fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldLocation,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultLocation holds the default value on creation for the "location" field.
	DefaultLocation string
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
)
