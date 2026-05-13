package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SearchResults(ctx echo.Context, results []*models.SearchResult) error {
	r := ui.NewRequest(ctx)
	r.Title = "Search Results"

	g := make(Group, len(results))
	for i, result := range results {
		g[i] = result.Render()
	}

	return r.Render(layouts.Primary, Div(
		Class("flex flex-col gap-6"),
		Div(
			Class("bg-card-bg/30 p-6 rounded-3xl border border-divider/40 mb-4"),
			P(Class("text-sm font-medium text-secondary-text"), Textf("Found %d records matching your search.", len(results))),
		),
		Div(
			Class("grid grid-cols-1 gap-6"),
			g,
		),
	))
}
