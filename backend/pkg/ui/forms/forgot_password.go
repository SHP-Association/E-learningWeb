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
		ControlGroup(
			FormButton(ColorPrimary, "Reset password"),
			ButtonLink(ColorLink, r.Path(routenames.Home), "Cancel"),
		),
		CSRF(r),
	)
}
