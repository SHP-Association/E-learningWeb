package pages

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	"github.com/labstack/echo/v4"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts *models.Posts, stats *models.DashboardStats) error {
	r := ui.NewRequest(ctx)
	r.Title = "Admin Dashboard"
	r.Metatags.Description = "SHP E-learning Platform Admin Dashboard - Manage users, courses, and system updates."

	// Dashboard Header with Admin Stats
	dashboardHeader := func() Node {
		userName := "Administrator"
		if r.IsAuth {
			userName = r.AuthUser.Username
		}

		return Div(
			Class("mb-16"),
			Div(
				Class("flex flex-col gap-2 mb-10"),
				H1(Class("text-5xl font-black text-white tracking-tight"), Text("Welcome, "+userName)),
				P(Class("text-secondary-text font-medium text-lg"), Text("Platform overview and administrative controls.")),
			),
			stats.Render(),
		)
	}

	// Admin Quick Actions
	featuredContent := func() Node {
		return Div(
			Class("grid lg:grid-cols-2 gap-8 mb-16"),
			Card(CardParams{
				Title: "Platform Management",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed mb-4"), Text("Manage your platform's core entities. Add new courses, verify users, and monitor enrollments.")),
					Div(
						Class("flex flex-wrap gap-3 mt-6"),
						A(Href("/admin/entity/course"), Class("px-4 py-2 rounded-full bg-accent/10 border border-accent/20 text-accent text-[11px] font-black uppercase tracking-wider hover:bg-accent/20 transition-all"), Text("Manage Courses")),
						A(Href("/admin/entity/user"), Class("px-4 py-2 rounded-full bg-accent/10 border border-accent/20 text-accent text-[11px] font-black uppercase tracking-wider hover:bg-accent/20 transition-all"), Text("Manage Users")),
					),
				},
			}),
			Card(CardParams{
				Title: "System Operations",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed mb-4"), Text("Monitor background tasks and system health. Ensure all scheduled operations are running smoothly.")),
					Div(
						Class("mt-6 p-4 rounded-2xl bg-accent/5 border border-accent/10 flex items-center gap-4"),
						Icon("ShieldCheck", "w-5 h-5 text-accent"),
						Span(Class("text-xs font-bold text-accent"), Text("All systems operational")),
					),
				},
				Footer: Group{
					ButtonLink(ColorNeutral, "/admin/tasks", "Monitor Tasks"),
				},
			}),
		)
	}

	// Main Section: Platform Updates (Posts)
	mainContent := func() Node {
		return Div(
			Class("space-y-8 flex-1"),
			Div(
				Class("flex items-center justify-between mb-2"),
				H2(Class("text-2xl font-black text-white tracking-tight"), Text("Latest Updates")),
				A(Href("#"), Class("text-[10px] font-black uppercase tracking-ultra text-accent hover:text-white transition-colors"), Text("Internal Announcements")),
			),
			posts.Render(r.Path(routenames.Home)),
		)
	}

	g := Group{
		Iff(r.Htmx.Target != "posts", dashboardHeader),
		Iff(r.Htmx.Target != "posts", featuredContent),
		mainContent(),
	}

	return r.Render(layouts.Primary, g)
}
