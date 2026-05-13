package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/forms"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ContactUs(ctx echo.Context, form *forms.Contact) error {
	r := ui.NewRequest(ctx)
	r.Title = "Get in Touch"
	r.Metatags.Description = "Have questions? We're here to help."

	g := Group{
		Iff(r.Htmx.Target != "contact", func() Node {
			return Card(CardParams{
				Title: "Contact Information",
				Body: Group{
					P(Class("text-secondary-text mb-4"), Text("Use the form below to send us a message. Our team typically responds within 24 hours.")),
					Div(
						Class("flex flex-col gap-4 mt-8"),
						Div(
							Class("flex items-center gap-4"),
							Div(Class("w-10 h-10 rounded-full bg-accent/10 flex items-center justify-center text-accent"), Icon("Mail", "w-5 h-5")),
							Div(
								Class("flex flex-col"),
								Span(Class("text-xs font-black uppercase tracking-widest text-secondary-text"), Text("Email")),
								Span(Class("text-sm font-bold"), Text("support@shpassociation.org")),
							),
						),
					),
				},
			})
		}),
		Iff(form.IsDone(), func() Node {
			return Card(CardParams{
				Title: "Message Received",
				Body: Group{
					P(Class("text-secondary-text"), Text("Thank you for your message! We've received your inquiry and will get back to you shortly.")),
				},
			})
		}),
		Iff(!form.IsDone(), func() Node {
			return form.Render(r)
		}),
	}

	return r.Render(layouts.Primary, g)
}
