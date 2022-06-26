// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/telegram-bot-template/pkg/db/ent/example"
)

// Example is the model entity for the Example schema.
type Example struct {
	config
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Example) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case example.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Example", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Example fields.
func (e *Example) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case example.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int64(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Example.
// Note that you need to call Example.Unwrap() before calling this method if this Example
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Example) Update() *ExampleUpdateOne {
	return (&ExampleClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Example entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Example) Unwrap() *Example {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Example is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Example) String() string {
	var builder strings.Builder
	builder.WriteString("Example(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Examples is a parsable slice of Example.
type Examples []*Example

func (e Examples) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
