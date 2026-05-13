package models

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/pager"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	Posts struct {
		Posts []Post
		Pager pager.Pager
	}

	Post struct {
		ID          int
		Title, Body string
		Author      string
		Date        string
	}
)

func (p *Posts) Render(path string) Node {
	g := make(Group, len(p.Posts))
	for i, post := range p.Posts {
		g[i] = post.Render()
	}

	return Div(
		ID("posts"),
		Class("flex flex-col gap-4"),
		g,
		Div(Class("mt-8")),
		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#posts"),
	)
}

func (p *Post) Render() Node {
	return Div(
		Class("bg-card-bg/30 border border-divider/40 rounded-3xl p-6 hover:bg-card-bg/50 transition-all group"),
		Div(
			Class("flex items-start gap-6"),
			Div(
				Class("w-12 h-12 rounded-2xl bg-accent/5 flex items-center justify-center text-accent group-hover:scale-110 transition-transform"),
				Icon("Megaphone", "w-6 h-6"),
			),
			Div(
				Class("flex-1 min-w-0"),
				Div(
					Class("flex items-center justify-between mb-1"),
					H3(Class("text-lg font-bold text-white truncate"), Text(p.Title)),
					Span(Class("text-[10px] font-black uppercase tracking-widest text-secondary-text opacity-40"), Text(p.Date)),
				),
				P(Class("text-secondary-text text-sm leading-relaxed line-clamp-2"), Text(p.Body)),
				Div(
					Class("mt-4 flex items-center gap-3"),
					Div(Class("w-5 h-5 rounded-full bg-divider/40")),
					Span(Class("text-[11px] font-bold text-secondary-text/60"), Text(p.Author)),
				),
			),
		),
	)
}
