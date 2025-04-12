package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Logs holds the schema definition for the Logs entity.
type Logs struct {
	ent.Schema
}

// Fields of the Logs.
func (Logs) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Int("access_key_id").Optional().Nillable(),
		field.String("route").Optional().Nillable(),
		field.String("method").Optional().Nillable(),
	}
}

// Edges of the Logs.
func (Logs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accesskeys", AccessKeys.Type).Ref("logs").Field("access_key_id").Unique(),
	}
}
