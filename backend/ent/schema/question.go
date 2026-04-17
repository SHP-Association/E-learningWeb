package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Annotations of the Question.
func (Question) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Question_question"},
	}
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Text("question_text"),
		field.String("question_type").Default("mcq"),
		field.Int("order").Default(0),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("quiz", Quiz.Type).
			Ref("questions").
			Unique().
			Required(),
		edge.To("choices", AnswerChoice.Type),
	}
}
