package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	CardParams struct {
		Title  string
		Body   Group
		Footer Group
		Color  Color
		Size   Size
	}

	Stat struct {
		Title       string
		Value       string
		Description string
		Icon        Node
	}
)

func Badge(color Color, text string) Node {
	var class string

	switch color {
	case ColorSuccess:
		class = "bg-emerald-500/10 text-emerald-500 border-emerald-500/20"
	case ColorWarning:
		class = "bg-amber-500/10 text-amber-500 border-amber-500/20"
	case ColorPrimary:
		class = "bg-accent/10 text-accent border-accent/20"
	default:
		class = "bg-white/5 text-white border-white/10"
	}

	return Span(
		Class("px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-widest border "+class),
		Text(text),
	)
}

func Divider(text string) Node {
	return Div(
		Class("flex items-center gap-4 my-8"),
		Div(Class("h-px flex-1 bg-white/5")),
		Span(Class("text-[10px] font-black uppercase tracking-ultra text-secondary-text"), Text(text)),
		Div(Class("h-px flex-1 bg-white/5")),
	)
}

func Card(params CardParams) Node {
	return Div(
		Class("admin-card p-8 flex flex-col gap-8 h-full"),
		If(len(params.Title) > 0, Div(
			Class("flex items-center justify-between"),
			H3(
				Class("text-[11px] font-black uppercase tracking-ultra text-accent"),
				Text(params.Title),
			),
		)),
		Div(Class("flex-1 text-sm text-secondary-text leading-relaxed font-medium"), params.Body),
		If(params.Footer != nil, Div(
			Class("pt-6 flex justify-end gap-3"),
			params.Footer,
		)),
	)
}

func Stats(stats ...Stat) Node {
	g := make(Group, 0, len(stats))
	for _, stat := range stats {
		g = append(g, Div(
			Class("admin-card p-6 flex items-center justify-between group hover:teal-lume"),
			Div(
				Class("flex flex-col"),
				P(Class("text-[10px] font-black uppercase tracking-ultra text-secondary-text/70"), Text(stat.Title)),
				H3(Class("text-3xl font-black text-[var(--color-primary-text)] tracking-tight mt-1"), Text(stat.Value)),
				If(stat.Description != "", P(Class("text-[11px] text-secondary-text/50 font-medium mt-1"), Text(stat.Description))),
			),
			Iff(stat.Icon != nil, func() Node {
				return Div(
					Class("w-12 h-12 rounded-xl bg-divider/10 border border-divider/25 flex items-center justify-center text-accent transition-all group-hover:scale-105 group-hover:bg-accent/15 group-hover:text-accent"),
					stat.Icon,
				)
			}),
		))
	}
	return Div(
		Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-16"),
		g,
	)
}

func TopBar(userName string) Node {
	return Div(
		Class("flex items-center justify-between mb-12"),
		// Search Pill
		Div(
			Class("search-pill flex items-center gap-3 w-80"),
			Icon("MagnifyingGlass", "w-4 h-4 opacity-40"),
			Input(
				Type("text"),
				Placeholder("Type to search..."),
				Class("bg-transparent border-none outline-none flex-1 text-sm"),
			),
			Span(Class("text-[10px] font-bold opacity-30 px-1.5 py-0.5 rounded border border-white/10 uppercase"), Text("⌘K")),
		),
		// Profile Pill
		Div(
			Class("profile-pill"),
			Div(
				Class("w-7 h-7 rounded-lg bg-accent/20 p-0.5 border border-accent/30"),
				Img(Src("https://api.dicebear.com/7.x/avataaars/svg?seed="+userName), Class("w-full h-full rounded")),
			),
			Span(Class("text-[12px] font-bold text-white"), Text(userName)),
		),
	)
}

func AlertBox(color Color, text string) Node {
	var class string
	switch color {
	case ColorWarning:
		class = "bg-amber-500/10 text-amber-500 border-amber-500/20"
	case ColorError:
		class = "bg-rose-500/10 text-rose-500 border-rose-500/20"
	case ColorSuccess:
		class = "bg-emerald-500/10 text-emerald-500 border-emerald-500/20"
	default:
		class = "bg-blue-500/10 text-blue-500 border-blue-500/20"
	}

	return Div(
		Class("p-4 rounded-xl border flex gap-4 items-center "+class),
		Icon("Info", "w-5 h-5 flex-shrink-0"),
		Span(Class("text-[13px] font-medium"), Text(text)),
	)
}
