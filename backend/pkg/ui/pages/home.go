package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/icons"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts *models.Posts) error {
	r := ui.NewRequest(ctx)
	r.Title = "Student Dashboard"
	r.Metatags.Description = "SHP E-learning Platform Dashboard - Manage your courses and progress."

	// Dashboard Header with High-Fidelity Stats
	dashboardHeader := func() Node {
		userName := "Learner"
		if r.IsAuth {
			userName = r.AuthUser.Username
		}
		
		return Div(
			Class("mb-16"),
			Div(
				Class("flex flex-col gap-2 mb-10"),
				H1(Class("text-5xl font-black text-white tracking-tight"), Text("Welcome, "+userName)),
				P(Class("text-secondary-text font-medium"), Text("Track your academic progress and resume your learning journey.")),
			),
			Stats(
				Stat{
					Title: "Enrolled Courses",
					Value: "12",
					Description: "3 Active now",
					Icon: icons.PencilSquare(),
				},
				Stat{
					Title: "Engagement Score",
					Value: "92%",
					Description: "+5% this week",
					Icon: icons.Info(),
				},
				Stat{
					Title: "Study Hours",
					Value: "48.5",
					Description: "Lifetime total",
					Icon: icons.Clock(),
				},
				Stat{
					Title: "Certificates",
					Value: "04",
					Description: "2 pending review",
					Icon: icons.CheckCircle(),
				},
			),
		)
	}

	// High-Fidelity CTA Section
	featuredContent := func() Node {
		return Div(
			Class("grid lg:grid-cols-2 gap-8 mb-16"),
			Card(CardParams{
				Title: "Resume Progress",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed mb-4"), Text("Continue your progress in 'Advanced Go Backend Development'. You're 65% through the course and have 4 lessons left.")),
					Div(
						Class("w-full bg-white/5 h-2 rounded-full overflow-hidden mt-6"),
						Div(Class("bg-primary h-full w-[65%] shadow-glow-sm")),
					),
				},
				Footer: Group{
					ButtonLink(ColorPrimary, "#", "Jump Back In"),
				},
			}),
			Card(CardParams{
				Title: "Learning Insights",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed"), Text("Your most active study time is between 8 PM and 10 PM. You've completed 12 lessons during this window this month.")),
				},
				Footer: Group{
					ButtonLink(ColorNeutral, "#", "View Full Report"),
				},
			}),
		)
	}

	// Main Section: Recent Activity
	mainContent := func() Node {
		return Div(
			Class("space-y-8 flex-1"),
			Div(
				Class("flex items-center justify-between mb-2"),
				H2(Class("text-2xl font-black text-white tracking-tight"), Text("Platform Updates")),
				A(Href("#"), Class("text-[10px] font-black uppercase tracking-ultra text-primary hover:text-white transition-colors"), Text("View All Announcements")),
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
