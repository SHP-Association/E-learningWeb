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

type ForgotPassword struct {
	Email string `form:"email" validate:"required,email"`
	form.Submission
}

func (f *ForgotPassword) Render(r *ui.Request) Node {
	return Form(
		ID("forgot-password"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.Path(routenames.ForgotPasswordSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		Div(
			Class("mt-2 mb-6"),
			A(
				Href(r.Path(routenames.Login)),
				Class("text-[11px] font-bold text-accent hover:text-white transition-colors"),
				Text("← Back to login"),
			),
		),
		FormButton(ColorPrimary, "Send Reset Link"),
		CSRF(r),
	)
}
