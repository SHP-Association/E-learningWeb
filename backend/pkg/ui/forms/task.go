package forms

import (
	"fmt"
	"net/http"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/form"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Task struct {
	Delay   int    `form:"delay" validate:"gte=0"`
	Message string `form:"message" validate:"required"`
	form.Submission
}

func (f *Task) Render(r *ui.Request) Node {
	return Form(
		ID("task"),
		Class("flex flex-col gap-6 mt-8"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.TaskSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Delay",
			Name:      "delay",
			InputType: "number",
			Label:     "Execution Delay",
			Help:      "Seconds to wait before processing",
			Value:     fmt.Sprint(f.Delay),
			Required:  true,
		}),
		TextareaField(TextareaFieldParams{
			Form:      f,
			FormField: "Message",
			Name:      "message",
			Label:     "Task Payload / Message",
			Value:     f.Message,
			Help:      "Data to be logged by the worker",
			Required:  true,
		}),
		Div(
			Class("flex justify-end pt-4"),
			Button(
				Type("submit"),
				Class("btn btn-teal px-10 h-12 rounded-xl text-xs font-black uppercase tracking-ultra"),
				Text("Queue Background Task"),
			),
		),
		CSRF(r),
	)
}
