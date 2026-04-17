package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Certificate holds the schema definition for the Certificate entity.
type Certificate struct {
	ent.Schema
}

// Annotations of the Certificate.
func (Certificate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Certificate_certificate"},
	}
}

// Fields of the Certificate.
func (Certificate) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("unique_id").Unique(),
		field.Time("issue_date").Default(time.Now).Immutable(),
	}
}

// Edges of the Certificate.
func (Certificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("enrollment", Enrollment.Type).
			Ref("certificate").
			Unique().
			Required(),
	}
}
