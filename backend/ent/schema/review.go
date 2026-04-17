package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Annotations of the Review.
func (Review) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Review_review"},
	}
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("rating"),
		field.Text("comment").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Bool("is_approved").Default(false),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("reviews").
			Unique().
			Required(),
		edge.From("student", User.Type).
			Ref("reviews_given").
			Unique().
			Required(),
	}
}
