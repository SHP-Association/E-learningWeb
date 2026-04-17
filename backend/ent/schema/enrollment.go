package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Enrollment holds the schema definition for the Enrollment entity.
type Enrollment struct {
	ent.Schema
}

// Annotations of the Enrollment.
func (Enrollment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Enrollment_enrollment"},
	}
}

// Fields of the Enrollment.
func (Enrollment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Time("enrolled_at").Default(time.Now).Immutable(),
		field.Time("completed_at").Optional().Nillable(),
		field.Float("progress").Default(0.0),
		field.Bool("is_completed").Default(false),
	}
}

// Edges of the Enrollment.
func (Enrollment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).
			Ref("enrollments").
			Unique().
			Required(),
		edge.From("course", Course.Type).
			Ref("enrollments").
			Unique().
			Required(),
		edge.To("certificate", Certificate.Type).
			Unique(),
	}
}
