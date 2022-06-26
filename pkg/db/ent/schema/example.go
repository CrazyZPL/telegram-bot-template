package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
	}
}

// Edges of the Example.
func (Example) Edges() []ent.Edge {
	return nil
}
