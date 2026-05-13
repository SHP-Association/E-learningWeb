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

type Contact struct {
	Email      string `form:"email" validate:"required,email"`
	Department string `form:"department" validate:"required,oneof=sales marketing hr"`
	Message    string `form:"message" validate:"required"`
	form.Submission
}

func (f *Contact) Render(r *ui.Request) Node {
	return Form(
		ID("contact"),
		Class("flex flex-col gap-6"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.ContactSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email Address",
			Value:     f.Email,
			Required:  true,
		}),
		CustomSelect(OptionsParams{
			Form:      f,
			FormField: "Department",
			Name:      "department",
			Label:     "Department",
			Value:     f.Department,
			Required:  true,
			Options: []ui.Choice{
				{Value: "sales", Label: "Sales"},
				{Value: "marketing", Label: "Marketing"},
				{Value: "hr", Label: "HR"},
			},
		}),
		TextareaField(TextareaFieldParams{
			Form:      f,
			FormField: "Message",
			Name:      "message",
			Label:     "Your Message",
			Value:     f.Message,
			Required:  true,
		}),
		Div(
			Class("flex justify-end pt-4"),
			Button(
				Type("submit"),
				Class("btn btn-teal px-10 h-12 rounded-xl text-xs font-black uppercase tracking-ultra"),
				Text("Send Message"),
			),
		),
		CSRF(r),
	)
}
