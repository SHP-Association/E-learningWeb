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

func AddTask(ctx echo.Context, form *forms.Task) error {
	r := ui.NewRequest(ctx)
	r.Title = "Queue Task"
	r.Metatags.Description = "Create a background task for processing."

	g := Group{
		Iff(r.Htmx.Target != "task", func() Node {
			return Card(CardParams{
				Title: "Task Automation",
				Body: Group{
					P(Class("text-secondary-text leading-relaxed"), Raw("Submitting this form will create an <code class='text-accent'>ExampleTask</code> in the distributed task queue. The message will be processed after the specified delay.")),
				},
			})
		}),
		form.Render(r),
		Iff(r.Htmx.Target != "task", func() Node {
			var text string
			if r.IsAdmin {
				text = "You can monitor the queue status in the Tasks management section."
			} else {
				text = "Administrative privileges are required to monitor the task queue."
			}
			return Div(
				Class("mt-8"),
				AlertBox(ColorWarning, text),
			)
		}),
	}

	return r.Render(layouts.Primary, g)
}
