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
			Head(
				Metatags(r),
				CSS(),
				JS(),
				ThemeStyles(),
				PremiumStyles(),
				ThemeInitScript(),
			),
			Body(
				Class("bg-[#030712] font-main text-primary-text selection:bg-accent/30 selection:text-white min-h-screen relative overflow-hidden"),
				AlertContainer(),
				
				// Layered Premium Background
				Div(Class("absolute inset-0 z-0 pointer-events-none"),
					// High-Fidelity Shell Gradient
					Div(Class("absolute inset-0 bg-gradient-to-br from-[#0b0f1a] via-[#030712] to-[#1e1b4b] opacity-95")),
					
					// Background Noise/Texture
					Div(Class("absolute inset-0 opacity-[0.03] bg-[url('https://grainy-gradients.vercel.app/noise.svg')]")),
					
					// Dynamic Mesh Gradient Blobs
					Div(Class("absolute top-[-10%] left-[-10%] w-[50%] h-[50%] bg-accent/10 rounded-full blur-[100px] animate-pulse-subtle")),
					Div(Class("absolute bottom-[-15%] right-[-10%] w-[60%] h-[60%] bg-indigo-500/10 rounded-full blur-[120px]")),
				),
				
				Div(
					Class("min-h-screen flex flex-col items-center justify-center p-6 relative z-10"),
					
					// Ultra-Premium Compact Auth Card
					Div(
						Class("admin-card glass-modern relative p-10 md:p-12 overflow-hidden w-full max-w-[400px] mx-auto"),
						// Subtle inner glow
						Div(Class("absolute top-0 left-0 w-full h-[1px] bg-gradient-to-r from-transparent via-white/10 to-transparent")),
						
						// Minimalist Branding Section (Logo Only)
						Div(
							Class("flex flex-col items-center mb-8 group"),
							Div(
								Class("relative"),
								Div(Class("absolute inset-0 bg-accent/20 blur-2xl rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-1000")),
								Img(
									Class("h-16 w-16 relative z-10 drop-shadow-glow transform group-hover:scale-105 transition-all duration-700 ease-out"),
									Src(ui.StaticFile("logo.png")),
								),
							),
						),
						
						FlashMessages(r),
						Div(Class("relative z-10"), content),
					),
				),
				AlertJS(),
				HtmxListeners(r),
			),
		),
	)
}
