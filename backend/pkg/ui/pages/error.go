package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Error(ctx echo.Context, code int) error {
	r := ui.NewRequest(ctx)
	r.Title = http.StatusText(code)
	
	var message string
	var subMessage string

	switch code {
	case http.StatusInternalServerError:
		message = "Unexpected Server Error"
		subMessage = "Something went wrong on our end. Please try again later."
	case http.StatusForbidden, http.StatusUnauthorized:
		message = "Access Denied"
		subMessage = "You don't have permission to view this resource."
	case http.StatusNotFound:
		message = "Page Not Found"
		subMessage = "The page you are looking for might have been removed or is temporarily unavailable."
	default:
		message = "Something Went Wrong"
		subMessage = "An unexpected error has occurred."
	}

	return r.Render(layouts.Primary, Div(
		Class("flex flex-col items-center justify-center py-20 text-center"),
		Div(
			Class("w-24 h-24 rounded-3xl bg-danger/10 flex items-center justify-center text-danger mb-8 shadow-xl shadow-danger/10 animate-bounce"),
			Icon("XCircle", "w-12 h-12"),
		),
		H1(Class("text-6xl font-black text-white mb-4 tracking-tighter"), Textf("%d", code)),
		H2(Class("text-2xl font-bold text-white mb-6"), Text(message)),
		P(Class("text-secondary-text max-w-md mx-auto mb-10 leading-relaxed"), Text(subMessage)),
		A(
			Href(r.Path(routenames.Home)),
			Class("btn btn-teal px-10 h-12 rounded-xl text-xs font-black uppercase tracking-ultra shadow-lg shadow-accent/15"),
			Text("Return to Dashboard"),
		),
	))
}
