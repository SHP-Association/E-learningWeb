package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AnswerChoice holds the schema definition for the AnswerChoice entity.
type AnswerChoice struct {
	ent.Schema
}

// Annotations of the AnswerChoice.
func (AnswerChoice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Question_answerchoice"},
	}
}

// Fields of the AnswerChoice.
func (AnswerChoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("choice_text").MaxLen(500),
		field.Bool("is_correct").Default(false),
	}
}

// Edges of the AnswerChoice.
func (AnswerChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).
			Ref("choices").
			Unique().
			Required(),
	}
}
