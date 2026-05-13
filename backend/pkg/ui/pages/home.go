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

	userEmail := "admin@slingsight.com"
	if r.IsAuth && r.AuthUser != nil {
		userEmail = r.AuthUser.Email
	}

	// Dashboard Header
	dashboardHeader := func() Node {
		return Div(
			Class("mb-16"),
			
			Div(
				Class("flex flex-col gap-8"),
				Div(
					Class("flex flex-col gap-1"),
					H2(Class("text-sm font-black uppercase tracking-ultra text-white/40"), Text("Admin Dashboard")),
					Div(Class("h-px w-full bg-white/5 mt-4")),
				),
				Div(
					Class("flex flex-col gap-2"),
					H1(Class("text-5xl font-black text-white tracking-tight"), Text("Welcome, "+userEmail)),
					P(Class("text-secondary-text font-medium text-sm mt-4"), Text("Platform overview and administrative controls.")),
				),
			),
		)
	}

	// Quick Management Actions
	managementSection := func() Node {
		return Div(
			Class("grid lg:grid-cols-2 gap-8 mb-16"),
			Card(CardParams{
				Title: "Platform Management",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed mb-6"), Text("Manage your platform's core entities. Add new courses, verify users, and monitor enrollments.")),
					Div(
						Class("flex flex-wrap gap-4"),
						A(Href("/admin/entity/course"), Class("px-6 py-3 rounded-xl bg-accent/10 border border-accent/20 text-accent text-[11px] font-black uppercase tracking-wider hover:bg-accent/20 transition-all active:scale-95"), Text("Manage Courses")),
						A(Href("/admin/entity/user"), Class("px-6 py-3 rounded-xl bg-white/[0.03] border border-white/10 text-white/70 text-[11px] font-black uppercase tracking-wider hover:bg-white/[0.05] hover:text-white transition-all active:scale-95"), Text("Manage Users")),
					),
				},
			}),
			Card(CardParams{
				Title: "System Operations",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed mb-6"), Text("Monitor background tasks and system health. Ensure all scheduled operations are running smoothly.")),
					Div(
						Class("p-5 rounded-2xl bg-white/[0.02] border border-white/5 flex items-center justify-between"),
						Div(
							Class("flex items-center gap-3"),
							Div(Class("w-2 h-2 rounded-full bg-accent animate-pulse")),
							Span(Class("text-xs font-bold text-accent/80"), Text("All systems operational")),
						),
						A(
							Href("/admin/tasks"),
							Class("px-5 py-2.5 rounded-lg bg-white/[0.05] text-white text-[10px] font-black uppercase tracking-widest hover:bg-white/[0.1] transition-all"),
							Text("Monitor Tasks"),
						),
					),
				},
			}),
		)
	}

	// Latest Updates Section
	updatesSection := func() Node {
		return Div(
			Class("space-y-8 flex-1"),
			Div(
				Class("flex items-center justify-between"),
				H2(Class("text-2xl font-black text-white tracking-tight"), Text("Latest Updates")),
				A(Href("#"), Class("text-[10px] font-black uppercase tracking-ultra text-accent hover:text-white transition-colors"), Text("Internal Announcements")),
			),
			posts.Render(r.Path(routenames.Home)),
		)
	}

	g := Group{
		Iff(r.Htmx.Target != "posts", dashboardHeader),
		Iff(r.Htmx.Target != "posts", func() Node { return stats.Render() }),
		Iff(r.Htmx.Target != "posts", managementSection),
		updatesSection(),
	}

	return r.Render(layouts.Primary, g)
}
