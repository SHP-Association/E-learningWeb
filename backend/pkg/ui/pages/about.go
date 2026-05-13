package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/cache"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func About(ctx echo.Context) error {
	r := ui.NewRequest(ctx)
	r.Title = "About SHP LMS"
	r.Metatags.Description = "Learn about the technology powering the SHP E-learning platform."

	content := cache.SetIfNotExists("pages.about.Content", func() Node {
		return Div(
			Class("flex flex-col gap-12"),
			Card(CardParams{
				Title: "Our Mission",
				Body: Group{
					P(Class("text-secondary-text leading-relaxed"), Text("The SHP E-learning Platform is dedicated to providing high-quality, accessible education through a modern, high-performance web experience. Our goal is to empower learners with the tools they need to succeed in their academic and professional journeys.")),
				},
			}),
			Div(
				Class("grid lg:grid-cols-2 gap-8"),
				Card(CardParams{
					Title: "Frontend Excellence",
					Body: Group{
						P(Class("text-secondary-text text-sm mb-6"), Text("We use cutting-edge technologies to deliver a zero-refresh, fluid interface.")),
						Tabs([]Tab{
							{Title: "HTMX", Body: "Powers our dynamic, server-driven UI updates without the complexity of heavy JS frameworks."},
							{Title: "Alpine.js", Body: "Provides lightweight client-side interactions and state management."},
							{Title: "Vanilla CSS", Body: "Custom-crafted OKLAB-based design system for maximum performance and visual fidelity."},
						}),
					},
				}),
				Card(CardParams{
					Title: "Backend Reliability",
					Body: Group{
						P(Class("text-secondary-text text-sm mb-6"), Text("Our core is built on Go, ensuring top-tier performance and safety.")),
						Tabs([]Tab{
							{Title: "Echo", Body: "A high-performance, minimalist Go web framework that handles our routing and middleware."},
							{Title: "Ent", Body: "An entity framework that provides a type-safe and powerful way to interact with our database."},
							{Title: "Gomponents", Body: "Allows us to build our UI entirely in Go, ensuring type safety from backend to frontend."},
						}),
					},
				}),
			),
		)
	})

	return r.Render(layouts.Primary, content)
}
