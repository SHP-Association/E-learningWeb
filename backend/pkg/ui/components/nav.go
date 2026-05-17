package components

import (
	"fmt"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/pager"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// isPathActive checks if the current path matches the item path or is a sub-path.
func isPathActive(currentPath, itemPath string) bool {
	c := strings.Trim(strings.ToLower(currentPath), "/")
	i := strings.Trim(strings.ToLower(itemPath), "/")

	if c == i {
		return true
	}

	cParts := strings.Split(c, "/")
	iParts := strings.Split(i, "/")

	if len(cParts) < len(iParts) {
		return false
	}

	for idx := range iParts {
		if cParts[idx] != iParts[idx] {
			return false
		}
	}

	return true
}

func MenuLink(r *ui.Request, icon Node, title, routeName string, routeParams ...any) Node {
	href := r.Path(routeName, routeParams...)
	isActive := isPathActive(r.CurrentPath, href)

	return Li(
		Class("list-none px-2"),
		A(
			Href(href),
			Class("sidebar-link flex items-center gap-3 p-2.5 rounded-lg transition-all active:scale-95"),
			Classes{
				"active": isActive,
				"text-secondary-text hover:text-[var(--color-primary-text)]": !isActive,
			},
			
			Span(Class("flex items-center justify-center w-5"), icon),
			Span(Class("text-[13px] font-medium"), Text(title)),
		),
	)
}

func NavGroup(r *ui.Request, icon Node, title string, isOpen bool, links ...Node) Node {
	return Details(
		Class("group/nav mb-1 px-2"),
		If(isOpen, Attr("open", "true")),
		Summary(
			Class("flex items-center justify-between p-2.5 rounded-lg cursor-pointer hover:bg-white/5 transition-all select-none list-none"),
			Div(
				Class("flex items-center gap-3"),
				Span(Class("flex items-center justify-center w-5 text-secondary-text group-hover:text-accent transition-colors"), icon),
				Span(Class("text-[11px] font-black uppercase tracking-ultra text-secondary-text group-hover:text-primary-text transition-colors"), Text(title)),
			),
			Span(Class("text-secondary-text/30 group-hover:text-secondary-text/60 transition-transform duration-300 group-open/nav:rotate-180 text-[10px] font-bold"), Icon("ChevronDown", "w-3 h-3")),
		),
		Div(
			Class("nav-group-border flex flex-col gap-1 mt-1"),
			Group(links),
		),
	)
}

func Pager(page int, path string, hasNext bool, hxTarget string) Node {
	href := func(p int) string {
		return fmt.Sprintf("%s?%s=%d",
			path,
			pager.QueryKey,
			p,
		)
	}

	btnClass := "btn h-10 min-h-0 bg-divider/10 border border-divider hover:bg-primary/10 hover:border-primary/30 text-secondary-text rounded-xl transition-all active:scale-95 shadow-sm no-underline flex items-center justify-center px-4"

	return Div(
		Class("flex items-center gap-3"),
		// Previous Button
		func() Node {
			if page <= 1 {
				return Span(Class(btnClass+" opacity-20 cursor-not-allowed pointer-events-none"), Text("«"))
			}
			link := A(
				Class(btnClass),
				Href(href(page-1)),
				Text("«"),
			)
			if len(hxTarget) > 0 {
				return Group{link, Attr("hx-get", href(page-1)), Attr("hx-swap", "outerHTML"), Attr("hx-target", hxTarget)}
			}
			return link
		}(),

		// Page Indicator
		Div(
			Class("bg-primary/5 border border-primary/20 text-primary font-black px-6 h-10 flex items-center justify-center rounded-xl text-[10px] font-outfit uppercase tracking-ultra"),
			Textf("Page %d", page),
		),

		// Next Button
		func() Node {
			if !hasNext {
				return Span(Class(btnClass+" opacity-20 cursor-not-allowed pointer-events-none"), Text("»"))
			}
			link := A(
				Class(btnClass),
				Href(href(page+1)),
				Text("»"),
			)
			if len(hxTarget) > 0 {
				return Group{link, Attr("hx-get", href(page+1)), Attr("hx-swap", "outerHTML"), Attr("hx-target", hxTarget)}
			}
			return link
		}(),
	)
}
