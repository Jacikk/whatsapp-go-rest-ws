package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Float("price"),
		field.String("description"),
		field.String("image_url"),
		field.Int("category_id").Optional().Nillable(),
		field.Int("access_key_id").Optional().Nillable(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("products").Field("category_id").Unique(),
		edge.From("accesskeys", AccessKeys.Type).Ref("products").Field("access_key_id").Unique(),
	}
}
