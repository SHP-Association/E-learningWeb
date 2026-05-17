package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Annotations of the Course.
func (Course) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "courses_course"},
	}
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").MaxLen(200),
		field.String("slug").MaxLen(200).Unique(),
		field.String("short_description").MaxLen(500),
		field.Text("description"),
		field.Text("what_you_will_learn").Optional(),
		field.Text("requirements").Optional(),
		field.Text("target_audience").Optional(),
		field.String("thumbnail").Optional(),
		field.String("promo_video_link").Optional().StorageKey("promo_video_url"),
		field.Float("price").Default(0.0),
		field.Bool("is_free").Default(false),
		field.Bool("is_published").Default(false),
		field.String("level").Default("beginner"),
		field.String("duration").Optional(),
		field.Int("total_lectures").Default(0),
		field.Float("average_rating").Default(0.0),
		field.Int("number_of_reviews").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instructor", User.Type).
			Ref("taught_courses").
			Unique(),
		edge.From("category", Category.Type).
			Ref("courses").
			Unique(),
		edge.To("lessons", Lesson.Type),
		edge.To("enrollments", Enrollment.Type).
			StorageKey(edge.Column("course_id")),
		edge.To("reviews", Review.Type),
		edge.From("enrolled_students", User.Type).
			Ref("enrolled_courses"),
		edge.From("completed_students", User.Type).
			Ref("completed_courses"),
	}
}
