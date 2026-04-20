package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PasswordToken holds the schema definition for the PasswordToken entity.
type PasswordToken struct {
	ent.Schema
}

// Fields of the PasswordToken.
func (PasswordToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Sensitive().
			NotEmpty(),
		field.Int("user_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the PasswordToken.
func (PasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Required().
			Unique(),
	}
}

// Hooks of the PasswordToken.
func (PasswordToken) Hooks() []ent.Hook {
    // Hooks are temporarily disabled to avoid import cycles during generation.
    // They should be applied in the service layer or after successful generation.
    return nil
}
