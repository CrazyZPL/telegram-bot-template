// Code generated by entc, DO NOT EDIT.

package example

const (
	// Label holds the string label denoting the example type in the database.
	Label = "example"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the example in the database.
	Table = "examples"
)

// Columns holds all SQL columns for example fields.
var Columns = []string{
	FieldID,
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
