package components

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/form"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	InputFieldParams struct {
		Form        form.Form
		FormField   string
		Name        string
		InputType   string
		Label       string
		Value       string
		Placeholder string
		Help        string
		Autocomplete string
		Required     bool
		Attributes   []string
	}

	FileFieldParams struct {
		Name  string
		Label string
		Help  string
	}

	OptionsParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Options   []ui.Choice
		Help      string
		Required  bool
	}

	TextareaFieldParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Help      string
		Required  bool
	}

	CheckboxParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Checked   bool
		Required  bool
	}
)

func ControlGroup(controls ...Node) Node {
	return Div(
		Class("mt-4 flex gap-3"),
		Group(controls),
	)
}

func TextareaField(el TextareaFieldParams) Node {
	return Fieldset(
		el.Label,
		el.Required,
		Textarea(
			Class("textarea h-32 w-full bg-white/5 border-white/10 rounded-2xl transition-all duration-300 focus:border-primary/50 focus:ring-4 focus:ring-primary/5 "+formFieldStatusClass(el.Form, el.FormField)),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Radios(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))
	for i, opt := range el.Options {
		id := "radio-" + el.Name + "-" + opt.Value
		buttons[i] = Div(
			Class("mb-3 flex items-center gap-3"),
			Input(
				ID(id),
				Type("radio"),
				Name(el.Name),
				Value(opt.Value),
				Class("radio radio-primary bg-white/5 border-white/20 "+formFieldStatusClass(el.Form, el.FormField)),
				If(el.Value == opt.Value, Checked()),
			),
			Label(
				Class("text-sm font-semibold text-secondary-text cursor-pointer"),
				Text(opt.Label),
				For(id),
			),
		)
	}

	return Fieldset(
		el.Label,
		el.Required,
		buttons,
		formFieldErrors(el.Form, el.FormField),
	)
}

func SelectList(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))
	for i, opt := range el.Options {
		buttons[i] = Option(
			Text(opt.Label),
			Value(opt.Value),
			If(opt.Value == el.Value, Attr("selected")),
		)
	}

	return Fieldset(
		el.Label,
		el.Required,
		Select(
			Class("select w-full bg-white/5 border-white/10 rounded-2xl transition-all duration-300 focus:border-primary/50 focus:ring-4 focus:ring-primary/5 "+formFieldStatusClass(el.Form, el.FormField)),
			Name(el.Name),
			buttons,
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Checkbox(el CheckboxParams) Node {
	return Div(
		Label(
			Class("label cursor-pointer flex items-center gap-3 p-1 group"),
			Input(
				Class("checkbox checkbox-primary rounded-lg transition-all active:scale-90 bg-white/5 border-white/20"),
				Type("checkbox"),
				Name(el.Name),
				If(el.Checked, Checked()),
				Value("true"),
			),
			Span(Class("text-sm font-bold text-secondary-text group-hover:text-white transition-colors"), Text(el.Label)),
			If(el.Required, Span(Class("text-error ml-[-8px]"), Text("*"))),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func InputField(el InputFieldParams) Node {
	attrs := make(Group, len(el.Attributes)/2)
	for i := 0; i < len(el.Attributes); i += 2 {
		if i+1 < len(el.Attributes) {
			attrs[i/2] = Attr(el.Attributes[i], el.Attributes[i+1])
		}
	}

	return Fieldset(
		el.Label,
		el.Required,
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			Class("input w-full bg-white/5 border-white/10 rounded-2xl transition-all duration-300 focus:border-primary/50 focus:ring-4 focus:ring-primary/5 text-sm font-medium "+formFieldStatusClass(el.Form, el.FormField)),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
			If(el.Autocomplete != "", Attr("autocomplete", el.Autocomplete)),
			attrs,
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Help(text string) Node {
	return If(len(text) > 0, Div(
		Class("label px-1 py-1.5 opacity-50"),
		Span(Class("label-text-alt text-[10px] uppercase font-black tracking-widest"), Text(text)),
	))
}

func Fieldset(label string, required bool, els ...Node) Node {
	return FieldSet(
		Class("fieldset p-0 border-0 flex flex-col gap-2 mb-6"),
		If(len(label) > 0, Legend(
			Class("fieldset-legend text-[11px] font-black text-secondary-text/80 mb-1 flex items-center gap-1.5 font-outfit uppercase tracking-ultra"),
			Text(label),
			If(required, Span(Class("text-error"), Text("*"))),
		)),
		Group(els),
	)
}

func FileField(el FileFieldParams) Node {
	return Fieldset(
		el.Label,
		false,
		Input(
			Type("file"),
			Class("file-input file-input-bordered bg-white/5 rounded-2xl w-full"),
			Name(el.Name),
		),
		Help(el.Help),
	)
}

func formFieldStatusClass(fm form.Form, formField string) string {
	switch {
	case fm == nil:
		return ""
	case !fm.IsSubmitted():
		return ""
	case fm.FieldHasErrors(formField):
		return "border-error focus:border-error focus:ring-error/10"
	default:
		return "border-success/30 focus:border-success/50 focus:ring-success/5"
	}
}

func formFieldErrors(fm form.Form, field string) Node {
	if fm == nil {
		return nil
	}

	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return nil
	}

	g := make(Group, len(errs))
	for i, err := range errs {
		g[i] = Div(
			Class("text-error text-[10px] font-black uppercase tracking-widest mt-1 px-1"),
			Text(err),
		)
	}

	return g
}

func CSRF(r *ui.Request) Node {
	return Input(
		Type("hidden"),
		Name("csrf"),
		Value(r.CSRF),
	)
}

func FormButton(color Color, label string) Node {
	return Button(
		Class("btn btn-teal w-full mt-4"),
		Text(label),
	)
}

func ButtonLink(color Color, href, label string) Node {
	return A(
		Href(href),
		Class("btn "+buttonColor(color)+" rounded-2xl font-black uppercase tracking-widest text-[11px]"),
		Text(label),
	)
}

func buttonColor(color Color) string {
	switch color {
	case ColorPrimary:
		return "btn-primary hover:shadow-lg hover:shadow-primary/20"
	case ColorInfo:
		return "btn-info"
	case ColorAccent:
		return "btn-accent"
	case ColorError:
		return "btn-error"
	case ColorLink:
		return "btn-link text-primary no-underline hover:text-white"
	case ColorNeutral:
		return "bg-white/5 border-white/10 text-white hover:bg-white/10"
	default:
		return ""
	}
}
