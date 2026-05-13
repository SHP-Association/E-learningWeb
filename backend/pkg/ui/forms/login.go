package forms

import (
	"net/http"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/form"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Login struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	form.Submission
}

func (f *Login) Render(r *ui.Request) Node {
	return Form(
		ID("login"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.Path(routenames.LoginSubmit)),
		FlashMessages(r),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		Div(Class("relative mt-4"),
			InputField(InputFieldParams{
				Form:        f,
				FormField:   "Password",
				Name:        "password",
				InputType:   "password",
				Label:       "Password",
				Placeholder: "******",
			}),
			Div(
				Class("flex justify-end mt-2"),
				A(
					Href(r.Path(routenames.ForgotPassword)),
					Class("text-[13px] font-semibold text-secondary-text hover:text-white transition-colors"),
					Text("Forgot password?"),
				),
			),
		),
		Div(Class("mt-8"),
			FormButton(ColorPrimary, "Sign In"),
		),
		CSRF(r),
	)
}
