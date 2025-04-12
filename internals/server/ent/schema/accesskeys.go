package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessKeys holds the schema definition for the AccessKeys entity.
type AccessKeys struct {
	ent.Schema
}

// Fields of the AccessKeys.
func (AccessKeys) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("user").Unique(),
	}
}

// Edges of the AccessKeys.
func (AccessKeys) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("categories", Category.Type),
		edge.To("logs", Logs.Type),
	}
}
