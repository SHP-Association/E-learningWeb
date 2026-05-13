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
					
					Div(
						Class("w-full max-w-[460px] animate-fadeIn"),
						
						// Brand Identity
						A(
							Href("/"),
							Class("flex flex-col items-center gap-6 mb-12 group"),
							Div(
								Class("relative"),
								Div(Class("absolute inset-0 bg-accent/30 blur-2xl rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-700")),
								Img(
									Class("h-20 w-20 relative z-10 drop-shadow-glow transform group-hover:scale-110 group-active:scale-95 transition-all duration-500 ease-out"),
									Src(ui.StaticFile("logo.png")),
								),
							),
							Div(
								Class("text-center"),
								H2(Class("text-2xl font-black uppercase tracking-[0.3em] text-white"), Text("SHP")),
								P(Class("text-[10px] font-black uppercase tracking-[0.5em] text-accent mt-1 opacity-80"), Text("Admin Control Center")),
							),
						),
						
						// Ultra-Premium Auth Card
						Div(
							Class("admin-card glass-modern relative p-10 md:p-14 overflow-hidden"),
							// Subtle inner glow
							Div(Class("absolute top-0 left-0 w-full h-[1px] bg-gradient-to-r from-transparent via-white/20 to-transparent")),
							
							If(len(r.Title) > 0, H1(Class("text-3xl font-black text-white tracking-tight mb-3 text-center"), Text(r.Title))),
							P(Class("text-secondary-text text-sm text-center mb-10 font-medium"), Text("Authorized personnel only. Please sign in.")),
							
							FlashMessages(r),
							Div(Class("relative z-10"), content),
						),
						
						// Footer Navigation
						Div(
							Class("mt-12 flex flex-col items-center gap-6"),
							Div(
								Class("flex justify-center gap-10"),
								A(Href("#"), Class("text-[11px] font-bold uppercase tracking-widest text-secondary-text hover:text-white transition-all hover:translate-y-[-1px]"), Text("Support")),
								A(Href("#"), Class("text-[11px] font-bold uppercase tracking-widest text-secondary-text hover:text-white transition-all hover:translate-y-[-1px]"), Text("Privacy Policy")),
								A(Href("#"), Class("text-[11px] font-bold uppercase tracking-widest text-secondary-text hover:text-white transition-all hover:translate-y-[-1px]"), Text("Terms")),
							),
							Div(
								Class("h-px w-12 bg-white/10"),
							),
							P(Class("text-[10px] font-medium text-tertiary-text uppercase tracking-widest"), Text("© 2024 SHP Association. All rights reserved.")),
						),
					),
				),
				AlertJS(),
				HtmxListeners(r),
			),
		),
	)
}
