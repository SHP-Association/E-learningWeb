package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Quiz holds the schema definition for the Quiz entity.
type Quiz struct {
	ent.Schema
}

// Annotations of the Quiz.
func (Quiz) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Quiz_quiz"},
	}
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").MaxLen(200),
		field.Text("description").Optional(),
		field.Int("passing_score").Default(70),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson", Lesson.Type).
			Ref("quizzes").
			Unique().
			Required(),
		edge.To("attempts", UserQuizAttempt.Type),
		edge.To("questions", Question.Type),
	}
}
