package pages

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/ent/admin"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/layouts"
	"entgo.io/ent/schema/field"
	"github.com/labstack/echo/v4"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AdminEntityDelete(ctx echo.Context, entityType admin.EntityType) error {
	r := ui.NewRequest(ctx)
	r.Title = fmt.Sprintf("Delete %s", entityType.GetName())

	return r.Render(
		layouts.Primary,
		Card(CardParams{
			Title: "Confirm Deletion",
			Body: Group{
				P(Class("mb-6 text-secondary-text"), Textf("Are you sure you want to delete this %s? This action cannot be undone.", entityType.GetName())),
				Form(
					Method("POST"),
					Action(r.CurrentPath),
					CSRF(r),
					Div(
						Class("flex gap-3"),
						Button(Type("submit"), Class("btn btn-danger px-8 h-12 rounded-xl"), Text("Delete")),
						A(Href(r.Path(routenames.AdminEntityList(entityType.GetName()))), Class("btn btn-neutral px-8 h-12 rounded-xl"), Text("Cancel")),
					),
				),
			},
		}),
	)
}

func AdminEntityInput(ctx echo.Context, entityType admin.EntityType, values url.Values) error {
	r := ui.NewRequest(ctx)
	title := entityType.GetName()
	if values == nil {
		r.Title = "Add " + title
	} else {
		r.Title = "Edit " + title
	}

	formNode := Form(
		ID("entity-form"),
		Method(http.MethodPost),
		Attr("hx-post", ctx.Request().URL.String()),
		Attr("hx-target", "#modal-form-body"),
		Attr("hx-swap", "innerHTML"),
		Class("p-6 flex flex-col"),
		
		// Modal Header
		Div(
			Class("flex items-center justify-between pb-4 mb-4 border-b border-divider"),
			H2(
				Class("text-lg font-black text-[var(--color-primary-text)] tracking-tight"), 
				Text(r.Title),
			),
			Button(
				Type("button"),
				Class("btn btn-sm btn-circle btn-ghost text-secondary-text hover:text-[var(--color-primary-text)] transition-colors"),
				Attr("onclick", "document.getElementById('admin-modal-container').classList.remove('modal-open')"),
				Text("✕"),
			),
		),

		// Scrollable form fields
		Div(
			Class("flex flex-col gap-5 overflow-y-auto px-1 pr-2 custom-scrollbar"),
			Style("max-height: calc(100vh - 350px);"),
			Group(func() []Node {
				fields := entityType.GetSchema()
				nodes := make([]Node, len(fields))
				for i, f := range fields {
					nodes[i] = renderField(r, f, values)
				}
				return nodes
			}()),
		),

		// Modal Footer / Buttons
		Div(
			Class("flex justify-end gap-3 mt-4 pt-4 border-t border-divider"),
			Button(
				Type("button"),
				Class("btn bg-divider/20 hover:bg-divider/30 text-secondary-text border-none rounded-xl px-6"),
				Attr("onclick", "document.getElementById('admin-modal-container').classList.remove('modal-open')"),
				Text("Cancel"),
			),
			Button(
				Type("submit"),
				Class("btn btn-brand px-10 h-12 rounded-xl shadow-lg shadow-accent/10"),
				Text("Save Changes"),
			),
		),
		CSRF(r),
	)

	if r.Htmx.Enabled {
		return formNode.Render(ctx.Response().Writer)
	}

	return r.Render(layouts.Primary, Group{
		formNode,
	})
}

func formatLabel(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	parts := strings.Fields(s)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func renderField(r *ui.Request, f *admin.FieldSchema, values url.Values) Node {
	val := ""
	if values != nil {
		val = values.Get(f.Name)
	}

	labelName := formatLabel(f.Name)

	common := InputFieldParams{
		Form:      nil,
		FormField: f.Name,
		Name:      f.Name,
		Label:     labelName,
		Value:     val,
		Required:  !f.Optional,
		Help:      "",
	}

	switch {
	case f.Type == field.TypeBool:
		return Switch(CheckboxParams{
			Form:      nil,
			FormField: f.Name,
			Name:      f.Name,
			Label:     labelName,
			Checked:   val == "true",
			Required:  !f.Optional,
		})
	case strings.Contains(strings.ToLower(f.Name), "image") || strings.Contains(strings.ToLower(f.Name), "thumbnail") || strings.Contains(strings.ToLower(f.Name), "picture"):
		return ImagePicker(ImagePickerParams{
			Name:     f.Name,
			Label:    labelName,
			Value:    val,
			Required: !f.Optional,
		})
	case len(f.Enums) > 0:
		choices := make([]ui.Choice, len(f.Enums))
		for i, e := range f.Enums {
			choices[i] = ui.Choice{Label: e, Value: e}
		}
		return CustomSelect(OptionsParams{
			Form:      nil,
			FormField: f.Name,
			Name:      f.Name,
			Label:     labelName,
			Value:     val,
			Options:   choices,
			Required:  !f.Optional,
		})
	case f.Type == field.TypeString && (strings.Contains(f.Name, "description") || strings.Contains(f.Name, "content") || strings.Contains(f.Name, "learn") || strings.Contains(f.Name, "requirement") || strings.Contains(f.Name, "audience") || strings.Contains(f.Name, "bio") || strings.Contains(f.Name, "comment") || strings.Contains(f.Name, "body") || strings.Contains(f.Name, "answer")):
		return TextareaField(TextareaFieldParams{
			Form:      nil,
			FormField: f.Name,
			Name:      f.Name,
			Label:     labelName,
			Value:     val,
			Help:      "",
			Required:  !f.Optional,
		})
	case f.Type == field.TypeInt || f.Type == field.TypeInt8 || f.Type == field.TypeInt16 || f.Type == field.TypeInt32 || f.Type == field.TypeInt64 || f.Type == field.TypeUint || f.Type == field.TypeUint8 || f.Type == field.TypeUint16 || f.Type == field.TypeUint32 || f.Type == field.TypeUint64 || f.Type == field.TypeFloat32 || f.Type == field.TypeFloat64:
		return NumberField(common)
	case f.Type == field.TypeTime:
		return DateTimeField(common)
	default:
		return InputField(common)
	}
}

func AdminEntityList(
	ctx echo.Context,
	entityType admin.EntityType,
	entityList *admin.EntityList,
) error {
	r := ui.NewRequest(ctx)
	title := entityType.GetName()
	r.Title = title

	addURL := r.Path(routenames.AdminEntityAdd(title))

	genHeader := func() Node {
		g := make(Group, 0, len(entityList.Columns)+2)
		g = append(g, Th(Class("sticky-column-left z-10"), Text("ID")))
		for _, h := range entityList.Columns {
			g = append(g, Th(Text(h)))
		}
		g = append(g, Th(Class("sticky-column-right z-10"), Text("Actions")))
		return g
	}

	genRow := func(row admin.EntityValues) Node {
		g := make(Group, 0, len(row.Values)+3)
		g = append(g, Td(Class("sticky-column-left font-black text-accent/80"), Text(fmt.Sprint(row.ID))))
		for _, v := range row.Values {
			// Basic formatting for common types
			display := v
			if len(display) > 50 {
				display = display[:47] + "..."
			}
			g = append(g, Td(Class("text-sm max-w-[200px] truncate"), Text(display)))
		}
		
		editURL := r.Path(routenames.AdminEntityEdit(title), row.ID)
		deleteURL := r.Path(routenames.AdminEntityDelete(title), row.ID)

		g = append(g,
			Td(
				Class("sticky-column-right"),
				Div(
					Class("flex gap-2"),
					Button(
						Type("button"),
						Class("btn btn-sm bg-white/5 border-divider/40 hover:bg-accent/10 hover:border-accent/40 text-primary-text px-3"),
						Attr("hx-get", editURL),
						Attr("hx-target", "#modal-form-body"),
						Attr("onclick", "document.getElementById('admin-modal-container').classList.add('modal-open')"),
						Icon("Pencil", "w-4 h-4"),
					),
					A(
						Href(deleteURL),
						Class("btn btn-sm bg-white/5 border-divider/40 hover:bg-danger/10 hover:border-danger/40 text-danger px-3"),
						Icon("Trash", "w-4 h-4"),
					),
				),
			),
		)
		return g
	}

	genRows := func() Node {
		g := make(Group, 0, len(entityList.Entities))
		for _, row := range entityList.Entities {
			g = append(g, Tr(Class("hover:bg-accent/5 transition-colors"), genRow(row)))
		}
		return g
	}

	return r.Render(layouts.Primary, Group{
		Div(
			Class("flex flex-col gap-6"),
			
			// Toolbar
			Div(
				Class("sticky-glass flex flex-col md:flex-row items-stretch md:items-center justify-between gap-4 p-4 px-6 rounded-3xl mb-6 shadow-sm"),
				Div(
					Class("flex flex-col sm:flex-row items-stretch sm:items-center gap-4 w-full md:w-auto"),
					Div(
						Class("relative w-full md:w-72 group"),
						Div(
							Class("absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none z-10"),
							Icon("MagnifyingGlass", "w-5 h-5 text-accent"),
						),
						Input(
							Type("text"),
							Placeholder("Search records..."),
							Class("admin-form-input h-11 w-full pl-11 !bg-card-bg/30 border-divider/50 rounded-2xl focus:border-accent focus:!bg-card-bg/60 transition-all text-sm font-medium focus:ring-4 focus:ring-accent/5"),
							ID("entity-search"),
							Attr("onkeyup", "filterTable()"),
						),
					),
				),
				Div(
					Class("flex items-center justify-end gap-4 w-full md:w-auto"),
					Button(
						Type("button"),
						Class("btn btn-brand w-full md:w-auto px-6 h-11 rounded-2xl shadow-xl shadow-accent/10 text-sm transition-all hover:scale-105 active:scale-95"),
						Attr("hx-get", addURL),
						Attr("hx-target", "#modal-form-body"),
						Attr("onclick", "document.getElementById('admin-modal-container').classList.add('modal-open')"),
						Text("+ ADD "+strings.ToUpper(title)),
					),
				),
			),

			// Table
			Div(
				Class("relative group/table"),
				Button(
					Type("button"),
					Class("table-scroll-btn table-scroll-btn-left"),
					ID("scroll-left"),
					Attr("onclick", "scrollTable(-300)"),
					Icon("ChevronLeft", "w-4 h-4"),
				),
				Button(
					Type("button"),
					Class("table-scroll-btn table-scroll-btn-right"),
					ID("scroll-right"),
					Attr("onclick", "scrollTable(300)"),
					Icon("ChevronRight", "w-4 h-4"),
				),
				Div(
					Class("admin-table-container"),
					ID("table-container"),
					Table(
						Class("table w-full admin-table"),
						ID("entity-table"),
						THead(
							Tr(Class("border-b border-divider"), genHeader()),
						),
						TBody(genRows()),
					),
				),
			),

			// Footer / Pagination
			Div(
				Class("flex items-center justify-center mt-8"),
				Pager(
					entityList.Page,
					r.CurrentPath,
					entityList.HasNextPage,
					"",
				),
			),
		),
		Script(Raw(`
			function filterTable() {
				const input = document.getElementById("entity-search");
				const filter = input.value.toUpperCase();
				const table = document.getElementById("entity-table");
				const tr = table.getElementsByTagName("tr");
				for (let i = 1; i < tr.length; i++) {
					let found = false;
					const td = tr[i].getElementsByTagName("td");
					for (let j = 0; j < td.length; j++) {
						if (td[j]) {
							const text = td[j].textContent || td[j].innerText;
							if (text.toUpperCase().indexOf(filter) > -1) {
								found = true;
								break;
							}
						}
					}
					tr[i].style.display = found ? "" : "none";
				}
			}

			function scrollTable(amount) {
				const container = document.getElementById('table-container');
				container.scrollBy({ left: amount, behavior: 'smooth' });
			}

			function updateScrollButtons() {
				const container = document.getElementById('table-container');
				const leftBtn = document.getElementById('scroll-left');
				const rightBtn = document.getElementById('scroll-right');
				
				if (!container || !leftBtn || !rightBtn) return;

				const showLeft = container.scrollLeft > 10;
				const showRight = container.scrollLeft < (container.scrollWidth - container.clientWidth - 10);
				const hasScroll = container.scrollWidth > container.clientWidth;

				if (showLeft && hasScroll) leftBtn.classList.add('visible');
				else leftBtn.classList.remove('visible');

				if (showRight && hasScroll) rightBtn.classList.add('visible');
				else rightBtn.classList.remove('visible');
			}

			document.addEventListener('DOMContentLoaded', () => {
				const container = document.getElementById('table-container');
				if (container) {
					container.addEventListener('scroll', updateScrollButtons);
					window.addEventListener('resize', updateScrollButtons);
					setTimeout(updateScrollButtons, 500);
				}
			});

			document.body.addEventListener('htmx:afterOnLoad', () => {
				setTimeout(updateScrollButtons, 100);
			});
		`)),
	})
}
