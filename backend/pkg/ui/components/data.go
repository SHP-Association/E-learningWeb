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
		Class("admin-card p-8 flex flex-col gap-6"),
		If(len(params.Title) > 0, Div(
			Class("flex items-center justify-between"),
			H3(
				Class("text-xs font-black uppercase tracking-ultra text-primary"),
				Text(params.Title),
			),
		)),
		Div(Class("flex-1 text-sm text-secondary-text leading-relaxed"), params.Body),
		If(params.Footer != nil, Div(
			Class("pt-6 border-t border-white/5 flex justify-end gap-3"),
			params.Footer,
		)),
	)
}

func Stats(stats ...Stat) Node {
	g := make(Group, 0, len(stats))
	for _, stat := range stats {
		g = append(g, Div(
			Class("admin-card p-8 flex items-center justify-between group hover:teal-lume"),
			Div(
				Class("flex flex-col gap-1"),
				P(Class("text-[10px] font-black uppercase tracking-ultra text-secondary-text opacity-60"), Text(stat.Title)),
				H3(Class("text-3xl font-black text-white tracking-tight"), Text(stat.Value)),
				If(stat.Description != "", P(Class("text-[11px] text-secondary-text/60 mt-1"), Text(stat.Description))),
			),
			Iff(stat.Icon != nil, func() Node {
				return Div(
					Class("w-14 h-14 rounded-2xl bg-primary/5 flex items-center justify-center text-primary transition-all group-hover:scale-110 group-hover:bg-primary/10"),
					stat.Icon,
				)
			}),
		))
	}
	return Div(
		Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 mb-12"),
		g,
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
