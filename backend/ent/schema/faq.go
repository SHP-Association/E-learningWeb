package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FAQ holds the schema definition for the FAQ entity.
type FAQ struct {
	ent.Schema
}

// Annotations of the FAQ.
func (FAQ) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "FAQ_faq"},
	}
}

// Fields of the FAQ.
func (FAQ) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Text("question"),
		field.Text("answer"),
		field.Bool("is_published").Default(false),
		field.Int("order").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FAQ.
func (FAQ) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("faqs").
			Unique(),
	}
}
