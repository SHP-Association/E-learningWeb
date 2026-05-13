package pages

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts *models.Posts, stats *models.DashboardStats) error {
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
				P(Class("text-secondary-text font-medium text-lg"), Text("Track your academic progress and resume your learning journey.")),
			),
			stats.Render(),
		)
	}

	// High-Fidelity CTA Section
	featuredContent := func() Node {
		resumeTitle := "Resume Progress"
		resumeBody := Group{
			P(Class("text-secondary-text text-sm leading-relaxed mb-4"), Text("You haven't started any courses yet. Explore our catalog to begin your learning journey.")),
		}

		if stats.RecentProgress != nil {
			resumeTitle = "Continue: " + stats.RecentProgress.CourseTitle
			resumeBody = Group{
				P(Class("text-secondary-text text-sm leading-relaxed mb-4"), Textf("You're %.0f%% through the course. You have about %d lessons left to complete.", stats.RecentProgress.Progress, stats.RecentProgress.Remaining)),
				Div(
					Class("w-full bg-divider/40 h-3 rounded-full overflow-hidden mt-6 border border-divider/20"),
					Div(Class("bg-accent h-full shadow-[0_0_15px_rgba(0,196,160,0.5)]"), Style(fmt.Sprintf("width: %.0f%%", stats.RecentProgress.Progress))),
				),
			}
		}

		return Div(
			Class("grid lg:grid-cols-2 gap-8 mb-16"),
			Card(CardParams{
				Title: resumeTitle,
				Body:  resumeBody,
				Footer: Group{
					ButtonLink(ColorPrimary, "#", "Jump Back In"),
				},
			}),
			Card(CardParams{
				Title: "Learning Insights",
				Body: Group{
					P(Class("text-secondary-text text-sm leading-relaxed"), Textf("You've spent %.1f hours learning this month. Keep up the great work!", stats.StudyHours)),
					Div(
						Class("mt-6 p-4 rounded-2xl bg-accent/5 border border-accent/10 flex items-center gap-4"),
						Icon("Clock", "w-5 h-5 text-accent"),
						Span(Class("text-xs font-bold text-accent"), Text("Keep the momentum going!")),
					),
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
				A(Href("#"), Class("text-[10px] font-black uppercase tracking-ultra text-accent hover:text-white transition-colors"), Text("View All Announcements")),
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
