package layouts

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Auth(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Data("theme", "dark"),
			Head(
				Metatags(r),
				CSS(),
				JS(),
			),
			Body(
				Class("bg-base-100 font-sans selection:bg-primary/30"),
				Div(
					Class("hero min-h-screen premium-gradient relative overflow-hidden"),
					// Animated background blobs
					Div(Class("absolute top-[-10%] left-[-10%] w-[40%] h-[40%] bg-primary/20 rounded-full blur-[120px]")),
					Div(Class("absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] bg-secondary/20 rounded-full blur-[120px]")),
					
					Div(
						Class("hero-content flex-col w-full max-w-md page-transition"),
						// logo / Brand
						A(
							Href("/"),
							Class("mb-8 transition-transform active:scale-95"),
							Img(
								Class("h-16 w-16 drop-shadow-2xl"),
								Src(ui.StaticFile("logo.png")),
							),
						),
						
						Div(
							Class("glass-card w-full rounded-3xl overflow-hidden"),
							Div(
								Class("p-8 md:p-10"),
								If(len(r.Title) > 0, H1(Class("text-3xl font-serif font-bold tracking-tight text-white mb-2 text-center"), Text(r.Title))),
								P(Class("text-base-content/60 text-sm text-center mb-8"), Text("Welcome back to SHP E-learning Platform")),
								
								FlashMessages(r),
								content,
							),
						),
						
						// Footer Links
						Div(
							Class("mt-8 flex gap-4 text-xs text-base-content/40"),
							A(Href("#"), Class("hover:text-primary transition-colors"), Text("Terms of Service")),
							A(Href("#"), Class("hover:text-primary transition-colors"), Text("Privacy Policy")),
						),
					),
				),
				HtmxListeners(r),
			),
		),
	)
}
