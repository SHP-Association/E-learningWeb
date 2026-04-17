package schema

import (
	"context"
	"net/mail"
	"strings"
	"time"

	ge "github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/hook"
	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Account_customuser"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.String("username").
			NotEmpty().
			Unique(),
		field.String("first_name").
			Optional(),
		field.String("last_name").
			Optional(),
		field.String("email").
			NotEmpty().
			Unique().
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				return err
			}),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.String("role").
			Default("student"),
		field.Text("bio").
			Optional(),
		field.String("profile_picture").
			Optional(),
		field.Time("date_of_birth").
			Optional().
			Nillable(),
		field.String("gender").
			Optional(),
		field.String("contact_number").
			Optional(),
		field.Text("address").
			Optional(),
		field.String("country").
			Optional(),
		field.Bool("is_email_verified").
			Default(false),
		field.String("highest_qualification").
			Optional(),
		field.String("institution").
			Optional(),
		field.Text("skills").
			Optional(),
		field.String("linkedin_profile").
			Optional(),
		field.String("github_profile").
			Optional(),
		field.Float("instructor_rating").
			Default(0.0),
		field.Int("total_reviews").
			Default(0),
		field.Bool("is_active_user").
			Default(true),
		field.Time("last_activity").
			Optional().
			Nillable(),
		field.String("login_ip").
			Optional(),
		field.Bool("two_factor_enabled").
			Default(false),
		field.Bool("is_staff").
			Default(false),
		field.Bool("is_active").
			Default(true),
		field.Bool("is_superuser").
			Default(false),
		field.Time("last_login").
			Optional().
			Nillable(),
		field.Time("date_joined").
			Default(time.Now).
			Immutable(),
		// Standard Pagoda fields for compatibility with existing logic
		field.Bool("verified").
			Default(false),
		field.Bool("admin").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", PasswordToken.Type).
			Ref("user"),
		edge.To("taught_courses", Course.Type),
		edge.To("quiz_attempts", UserQuizAttempt.Type),
		edge.To("enrollments", Enrollment.Type),
		edge.To("reviews_given", Review.Type),
		edge.To("enrolled_courses", Course.Type),
		edge.To("completed_courses", Course.Type),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ge.UserMutation) (ent.Value, error) {
					if v, exists := m.Email(); exists {
						m.SetEmail(strings.ToLower(v))
					}

					if v, exists := m.Password(); exists {
						hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
						if err != nil {
							return "", err
						}
						m.SetPassword(string(hash))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
