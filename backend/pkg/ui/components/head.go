package components

import (
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func JS() Node {
	return Group{
		Script(Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"), Defer()),
		Script(Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), Defer()),
	}
}

func CSS() Node {
	return Group{
		Link(Rel("preconnect"), Href("https://fonts.googleapis.com")),
		Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), Attr("crossorigin", "true")),
		Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Outfit:wght@400;600;700&display=swap")),
		Link(
			Href(ui.StaticFile("main.css")),
			Rel("stylesheet"),
			Type("text/css"),
		),
	}
}

func Metatags(r *ui.Request) Node {
	appName := "SHP"
	if r.Config != nil && r.Config.App.Name != "" {
		appName = r.Config.App.Name
	}
	
	title := appName
	if r.Title != "" {
		title = r.Title + " | " + appName
	}
	
	return Group{
		Meta(Charset("utf-8")),
		Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		Link(Rel("icon"), Href(ui.StaticFile("favicon.ico"))),
		TitleEl(Text(title)),
		If(r.Metatags.Description != "", Meta(Name("description"), Content(r.Metatags.Description))),
		If(len(r.Metatags.Keywords) > 0, Meta(Name("keywords"), Content(strings.Join(r.Metatags.Keywords, ", ")))),
	}
}
