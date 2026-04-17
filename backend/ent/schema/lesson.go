package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lesson holds the schema definition for the Lesson entity.
type Lesson struct {
	ent.Schema
}

// Annotations of the Lesson.
func (Lesson) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Lesson_lesson"},
	}
}

// Fields of the Lesson.
func (Lesson) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").MaxLen(200),
		field.String("slug").MaxLen(250),
		field.Text("content").Optional(),
		field.String("video_link").Optional().StorageKey("video_url"),
		field.Int("order"),
		field.Bool("is_preview").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Lesson.
func (Lesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("lessons").
			Unique().
			Required(),
		edge.To("quizzes", Quiz.Type),
	}
}
