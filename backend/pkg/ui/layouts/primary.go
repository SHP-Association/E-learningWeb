package layouts

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/cache"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Primary(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Data("theme", "dark"),
			Head(
				ThemeInitScript(),
				Metatags(r),
				CSS(),
				JS(),
				ThemeStyles(),
				PremiumStyles(),
			),
			Body(
				Class("bg-[var(--color-page-bg)] text-[var(--color-primary-text)] selection:bg-primary/30 transition-colors duration-400"),
				Div(
					Class("drawer lg:drawer-open"),
					Input(
						ID("sidebar"),
						Type("checkbox"),
						Class("drawer-toggle"),
					),
					Div(
						Class("drawer-content flex flex-col min-h-screen"),
						// Fixed Navbar
						navbar(r),
						
						// Main Content responsive container
						Main(
							Class("flex-1 p-8 sm:p-12 page-transition prose-base flex flex-col"),
							If(len(r.Title) > 0, H1(
								Class("text-4xl font-black tracking-tight mb-10 text-[var(--color-primary-text)] border-b border-divider pb-6"), 
								Text(r.Title),
							)),
							content,
						),
						
						// Mobile Overlay button
						Label(
							For("sidebar"),
							Class("btn btn-brand fixed bottom-6 right-6 lg:hidden shadow-2xl z-50 px-6"),
							Group{
								Icon("Bars3", "w-5 h-5"),
								Span(Class("ml-2"), Text("Menu")),
							},
						),
					),
					sidebarMenu(r),
				),
				AlertContainer(),
				AlertJS(),
				FormScripts(),
				FlashMessages(r),
				HtmxListeners(r),
				searchModal(r),
				
				// Global Modal Body for Admin Forms
				Div(
					ID("admin-modal-container"),
					Class("modal z-50"),
					Div(
						Class("modal-box max-w-2xl bg-card-bg border border-divider p-0 overflow-hidden rounded-3xl shadow-2xl"),
						Div(
							ID("modal-form-body"),
							Class("max-h-[85vh] overflow-hidden flex flex-col"),
							Div(Class("flex justify-center py-10"), Span(Class("loading loading-spinner loading-lg text-primary"))),
						),
					),
					Div(Class("modal-backdrop backdrop-blur-md bg-page-bg/40"), Attr("onclick", "document.getElementById('admin-modal-container').classList.remove('modal-open')")),
				),
			),
		),
	)
}

func navbar(r *ui.Request) Node {
	userName := "Administrator"
	if r.IsAuth && r.AuthUser != nil {
		userName = r.AuthUser.Username
	}

	return Nav(
		Class("sticky top-0 z-30 flex h-20 w-full justify-end items-center bg-transparent px-8"),
		Div(
			Class("flex items-center gap-4"),
			// Search Pill
			Div(
				Class("search-pill flex items-center gap-3 w-64"),
				icons.Icon("MagnifyingGlass", "w-4 h-4 opacity-40"),
				Input(
					Type("text"),
					Placeholder("Type to search..."),
					Class("bg-transparent border-none outline-none flex-1 text-sm"),
					Attr("@click", "search_modal.showModal();"),
				),
				Span(Class("text-[10px] font-bold opacity-30 px-1.5 py-0.5 rounded border border-white/10 uppercase"), Text("⌘K")),
			),
			// Profile Pill
			If(r.IsAuth, Div(
				Class("dropdown dropdown-end"),
				Div(
					TabIndex("0"),
					Role("button"),
					Class("profile-pill"),
					Div(
						Class("w-7 h-7 rounded-lg bg-accent/20 p-0.5 border border-accent/30"),
						Img(Src("https://api.dicebear.com/7.x/avataaars/svg?seed="+userName), Class("w-full h-full rounded")),
					),
					Span(Class("text-[12px] font-bold text-[var(--color-primary-text)]"), Text(userName)),
				),
				Ul(
					TabIndex("0"),
					Class("mt-4 z-[1] p-2 shadow-2xl menu menu-sm dropdown-content bg-card-bg rounded-2xl w-52 border border-white/5"),
					Li(A(Class("rounded-xl"), Href(r.Path(routenames.Home)), Text("Account Settings"))),
					Li(A(Class("rounded-xl text-danger"), Href(r.Path(routenames.Logout)), Text("Logout"))),
				),
			)),
		),
	)
}

func searchModal(r *ui.Request) Node {
	return cache.SetIfNotExists("layout.searchModal", func() Node {
		return Dialog(
			ID("search_modal"),
			Class("modal"),
			Div(
				Class("modal-box bg-base-200 border border-white/10 shadow-3xl p-8 rounded-3xl"),
				Form(
					Method("dialog"),
					Button(
						Class("btn btn-sm btn-circle btn-ghost absolute right-4 top-4"),
						Text("✕"),
					),
				),
				H3(
					Class("text-xs font-black uppercase tracking-ultra text-primary mb-6"),
					Text("Global Search"),
				),
				Input(
					Attr("hx-get", r.Path(routenames.Search)),
					Attr("hx-trigger", "keyup changed delay:500ms"),
					Attr("hx-target", "#results"),
					Name("query"),
					Class("input input-bordered w-full bg-base-100 rounded-xl border-white/10 focus:border-primary transition-all"),
					Type("search"),
					Placeholder("Search courses, lessons, topics..."),
				),
				Ul(
					ID("results"),
					Class("menu w-full mt-6 p-0 gap-1"),
				),
			),
			Form(
				Method("dialog"),
				Class("modal-backdrop"),
				Button(Text("close")),
			),
		)
	})
}

func sidebarMenu(r *ui.Request) Node {
	return Div(
		Class("drawer-side z-40"),
		Label(
			For("sidebar"),
			Aria("label", "close sidebar"),
			Class("drawer-overlay"),
		),
		Div(
			Class("flex flex-col h-full w-80 glass-modern p-8"),
			// Brand Section
			Div(
				Class("flex items-center justify-between px-2 mb-12"),
				A(
					Href(r.Path(routenames.Home)),
					Class("flex items-center gap-3 transition-transform active:scale-95 group"),
					Img(
						Class("h-10 w-auto drop-shadow-glow transition-transform group-hover:rotate-6"),
						Src(ui.StaticFile("logo.png")),
					),
					Span(Class("text-2xl font-black tracking-tighter text-[var(--color-primary-text)]"), Text("SHP")),
				),
				ThemeToggle(),
			),
			
			// Main Navigation
			Div(
				Class("flex-1 overflow-y-auto custom-scrollbar flex flex-col gap-6"),
				
				// Dashboard - Direct Link
				Div(
					Class("flex flex-col gap-1"),
					MenuLink(r, icons.Home(), "Main Dashboard", routenames.Home),
				),

				// Academy Hub Group
				NavGroup(r, icons.AcademicCap(), "Academy Hub", true,
					MenuLink(r, icons.BookOpen(), "Course Builder", routenames.AdminEntityList("Course")),
					MenuLink(r, icons.Document(), "Lesson Library", routenames.AdminEntityList("Lesson")),
					MenuLink(r, icons.Folder(), "Categories", routenames.AdminEntityList("Category")),
					MenuLink(r, icons.Star(), "Student Reviews", routenames.AdminEntityList("Review")),
				),

				// Student Center Group
				NavGroup(r, icons.Users(), "Student Center", false,
					MenuLink(r, icons.DocumentCheck(), "Enrollments", routenames.AdminEntityList("Enrollment")),
					MenuLink(r, icons.UserCircle(), "User Directory", routenames.AdminEntityList("User")),
					MenuLink(r, icons.AcademicCap(), "Certificate Registry", routenames.AdminEntityList("Certificate")),
				),

				// Exam Center Group
				NavGroup(r, icons.DocumentCheck(), "Exam Center", false,
					MenuLink(r, icons.PencilSquare(), "Quiz Designer", routenames.AdminEntityList("Quiz")),
					MenuLink(r, icons.QuestionCircle(), "Question Bank", routenames.AdminEntityList("Question")),
					MenuLink(r, icons.Clock(), "Quiz Attempts", routenames.AdminEntityList("UserQuizAttempt")),
				),

				// Operations Group
				NavGroup(r, icons.Cog6Tooth(), "System Operations", false,
					MenuLink(r, icons.CircleStack(), "Task Monitor", routenames.AdminTasks),
					MenuLink(r, icons.Folder(), "File Assets", routenames.Files),
					MenuLink(r, icons.CircleStack(), "Cache Manager", routenames.Cache),
					MenuLink(r, icons.Info(), "Support & FAQ", routenames.AdminEntityList("Faq")),
				),
			),
			
			// Footer Section with User Profile & Logout
			Div(
				Class("mt-auto pt-8 border-t border-white/5 flex flex-col gap-6"),
				If(r.IsAuth, Div(
					Class("flex flex-col gap-4"),
					// Exact Logout Matching
					A(
						Href("/user/logout"),
						Attr("onclick", "return confirm('Are you sure you want to logout?')"),
						Class("logout-btn flex items-center gap-3 px-4 py-3 rounded-lg font-bold text-[13px] active:scale-95"),
						icons.Icon("ArrowLeftOnRectangle", "w-5 h-5"),
						Span(Text("Logout")),
					),
				)),
				If(!r.IsAuth, Group{
					ButtonLink(ColorPrimary, r.Path(routenames.Login), "Admin Access"),
				}),
			),
		),
	)
}
