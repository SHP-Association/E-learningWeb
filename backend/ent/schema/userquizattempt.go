package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserQuizAttempt holds the schema definition for the UserQuizAttempt entity.
type UserQuizAttempt struct {
	ent.Schema
}

// Annotations of the UserQuizAttempt.
func (UserQuizAttempt) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Quiz_userquizattempt"},
	}
}

// Fields of the UserQuizAttempt.
func (UserQuizAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Float("score").Default(0.0),
		field.Bool("passed").Default(false),
		field.Int("attempt_number").Default(1),
		field.Time("submitted_at").Default(time.Now).Immutable(),
	}
}

// Edges of the UserQuizAttempt.
func (UserQuizAttempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).
			Ref("quiz_attempts").
			Unique().
			Required(),
		edge.From("quiz", Quiz.Type).
			Ref("attempts").
			Unique().
			Required(),
	}
}
