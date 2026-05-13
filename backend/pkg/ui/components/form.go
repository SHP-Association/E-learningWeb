package components

import (
	"fmt"

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
		Form        form.Form
		FormField   string
		Name        string
		Label       string
		Value       string
		Options     []ui.Choice
		Placeholder string
		Help        string
		Required    bool
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
	ImagePickerParams struct {
		Name        string
		Label       string
		Value       string // Current image URL
		Help        string
		Required    bool
		Placeholder string
	}

	RelationSelectParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Options   []ui.Choice
		Help      string
		Required  bool
	}
)

func ControlGroup(controls ...Node) Node {
	return Div(
		Class("mt-2 flex gap-2"),
		Group(controls),
	)
}

func TextareaField(el TextareaFieldParams) Node {
	return FormGroup(el.Label, el.Required,
		Textarea(
			Class("admin-form-input min-h-[120px] resize-y "+formFieldStatusClass(el.Form, el.FormField)),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
			If(el.Help != "", Placeholder(el.Help)),
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
			Class("mb-2"),
			Input(
				ID(id),
				Type("radio"),
				Name(el.Name),
				Value(opt.Value),
				Class("radio radio-accent mr-1 "+formFieldStatusClass(el.Form, el.FormField)),
				If(el.Value == opt.Value, Checked()),
			),
			Label(
				Class("text-sm cursor-pointer"),
				Text(opt.Label),
				For(id),
			),
		)
	}

	return FormGroup(el.Label, el.Required,
		Div(Class("flex flex-wrap gap-4"), buttons),
		formFieldErrors(el.Form, el.FormField),
	)
}

func SelectList(el OptionsParams) Node {
	return CustomSelect(el)
}

func RelationSelect(el RelationSelectParams) Node {
	return CustomSelect(OptionsParams{
		Form:        el.Form,
		FormField:   el.FormField,
		Name:        el.Name,
		Label:       el.Label,
		Value:       el.Value,
		Options:     el.Options,
		Help:        el.Help,
		Required:    el.Required,
		Placeholder: "Select " + el.Label + "...",
	})
}

func CustomSelect(el OptionsParams) Node {
	id := "custom-select-" + el.Name
	buttonID := id + "-button"
	optionsID := id + "-options"
	selectedLabelID := id + "-selected-label"
	hiddenInputName := el.Name

	var selectedLabel string
	if el.Value != "" {
		for _, opt := range el.Options {
			if opt.Value == el.Value {
				selectedLabel = opt.Label
				break
			}
		}
	}
	if selectedLabel == "" && el.Placeholder != "" {
		selectedLabel = el.Placeholder
	}

	return FormGroup(el.Label, el.Required,
		Div(
			Class("relative"),
			ID(id),
			Input(Type("hidden"), Name(hiddenInputName), Value(el.Value)),
			Button(
				ID(buttonID),
				Type("button"),
				Class("admin-form-input text-left w-full flex justify-between items-center pr-4"),
				Attr("onclick", fmt.Sprintf("toggleCustomSelect('%s')", id)),
				Attr("aria-haspopup", "listbox"),
				Attr("aria-expanded", "false"),
				Span(ID(selectedLabelID), Text(selectedLabel)),
				Icon("ChevronDown", "w-4 h-4 text-secondary-text/60"),
			),
			Ul(
				ID(optionsID),
				Class("absolute z-10 w-full bg-card-bg border border-divider rounded-xl shadow-lg mt-1 hidden max-h-60 overflow-y-auto custom-scrollbar"),
				If(el.Placeholder != "" && !el.Required, Li(
					Class("px-4 py-2 text-sm text-secondary-text cursor-pointer hover:bg-accent hover:text-page-bg"),
					Attr("data-value", ""),
					Attr("onclick", fmt.Sprintf("selectCustomOption(this, '%s')", id)),
					Text(el.Placeholder),
				)),
				Group(func() []Node {
					nodes := make([]Node, len(el.Options))
					for i, opt := range el.Options {
						nodes[i] = Li(
							Class("px-4 py-2 text-sm text-primary-text cursor-pointer hover:bg-accent hover:text-page-bg"),
							Attr("data-value", opt.Value),
							Attr("onclick", fmt.Sprintf("selectCustomOption(this, '%s')", id)),
							Text(opt.Label),
						)
					}
					return nodes
				}()),
			),
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
				Class("checkbox checkbox-accent rounded-lg transition-all active:scale-90 group-hover:shadow-lg group-hover:shadow-accent/10"),
				Type("checkbox"),
				Name(el.Name),
				If(el.Checked, Checked()),
				Value("true"),
			),
			Span(Class("text-[13px] font-medium text-secondary-text group-hover:text-primary-text transition-colors"), Text(el.Label)),
			If(el.Required, Span(Class("text-danger"), Text(" *"))),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Switch(el CheckboxParams) Node {
	return FormGroup(el.Label, el.Required,
		Label(
			Class("relative inline-flex items-center cursor-pointer group"),
			Input(
				Type("checkbox"),
				Name(el.Name),
				Class("sr-only peer"),
				If(el.Checked, Checked()),
				Value("true"),
			),
			Div(Class("w-11 h-6 bg-divider/40 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent group-hover:after:scale-110 shadow-inner")),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func ImagePicker(el ImagePickerParams) Node {
	id := "img-picker-" + el.Name
	previewID := id + "-preview"

	return FormGroup(el.Label, el.Required,
		Div(
			Class("flex flex-col gap-4 p-5 bg-page-bg/30 border border-divider/50 rounded-2xl transition-all hover:border-accent/30 group"),
			Div(
				Class("flex items-center gap-6"),
				Div(
					Class("relative w-24 h-24 rounded-2xl overflow-hidden bg-card-bg border-2 border-divider/50 shadow-inner flex items-center justify-center shrink-0 group-hover:border-accent/20 transition-colors"),
					If(el.Value != "", Img(
						ID(previewID),
						Src(el.Value),
						Class("w-full h-full object-cover"),
						Alt("Preview"),
					)),
					If(el.Value == "", Div(
						ID(previewID),
						Class("flex flex-col items-center gap-2 opacity-30"),
						Icon("UserCircle", "w-8 h-8"),
					)),
				),
				Div(
					Class("flex-1 flex flex-col gap-3"),
					Div(
						Class("relative"),
						Input(
							Type("file"),
							ID(id),
							Name(el.Name),
							Class("hidden"),
							Accept("image/*"),
							Attr("onchange", fmt.Sprintf(`
								(function(input) {
									if (input.files && input.files[0]) {
										var reader = new FileReader();
										reader.onload = function(e) {
											var preview = document.getElementById('%s');
											if (preview.tagName === 'IMG') {
												preview.src = e.target.result;
											} else {
												var img = document.createElement('img');
												img.id = '%s';
												img.src = e.target.result;
												img.className = 'w-full h-full object-cover';
												preview.parentNode.replaceChild(img, preview);
											}
										};
										reader.readAsDataURL(input.files[0]);
									}
								})(this)
							`, previewID, previewID)),
						),
						Label(
							For(id),
							Class("btn btn-teal btn-sm h-10 px-5 rounded-xl cursor-pointer inline-flex items-center gap-2 normal-case"),
							Icon("PencilSquare", "w-4 h-4"),
							Text("Choose Image"),
						),
					),
					P(Class("text-[11px] text-secondary-text leading-relaxed"), Text("Recommended: Square JPG, PNG or WebP. Max 5MB.")),
				),
			),
			If(el.Help != "", P(Class("text-[12px] text-secondary-text opacity-60 italic"), Text(el.Help))),
		),
	)
}

func InputField(el InputFieldParams) Node {
	attrs := make(Group, len(el.Attributes)/2)
	for i := 0; i < len(el.Attributes); i += 2 {
		if i+1 < len(el.Attributes) {
			attrs[i/2] = Attr(el.Attributes[i], el.Attributes[i+1])
		}
	}

	return FormGroup(el.Label, el.Required,
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			Class("admin-form-input "+formFieldStatusClass(el.Form, el.FormField)),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
			If(el.Autocomplete != "", Attr("autocomplete", el.Autocomplete)),
			attrs,
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func EmailField(el InputFieldParams) Node {
	el.InputType = "email"
	if el.Autocomplete == "" {
		el.Autocomplete = "email"
	}
	return InputField(el)
}

func PasswordField(el InputFieldParams) Node {
	el.InputType = "password"
	return InputField(el)
}

func NumberField(el InputFieldParams) Node {
	el.InputType = "number"
	return InputField(el)
}

func DateTimeField(el InputFieldParams) Node {
	el.InputType = "datetime-local"
	return InputField(el)
}

func DateField(el InputFieldParams) Node {
	el.InputType = "date"
	return InputField(el)
}

func Help(text string) Node {
	return If(len(text) > 0, Div(
		Class("label"),
		Text(text),
	))
}

func FormGroup(label string, required bool, els ...Node) Node {
	return Div(
		Class("flex flex-col gap-2"),
		If(len(label) > 0, Label(
			Class("flex items-center text-[13px] font-[600] text-secondary-text"),
			Text(label),
			If(required, Span(Class("inline-flex items-center ml-2 px-1.5 py-0.5 rounded-md bg-danger/10 text-danger text-[9px] font-black uppercase tracking-wider"), Text("Required"))),
		)),
		Group(els),
	)
}

func FileField(el FileFieldParams) Node {
	return FormGroup(el.Label, false,
		Input(
			Type("file"),
			Class("file-input"),
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
	class := "btn btn-brand w-full mt-4 h-12"
	if color == ColorError {
		class = "btn btn-danger w-full mt-4 h-12"
	}
	return Button(
		Type("submit"),
		Class(class),
		Text(label),
	)
}

func FormScripts() Node {
	return Script(Raw(`
		function toggleCustomSelect(id) {
			const options = document.getElementById(id + '-options');
			const button = document.getElementById(id + '-button');
			const isExpanded = options.classList.toggle('hidden');
			button.setAttribute('aria-expanded', !isExpanded);
		}

		function selectCustomOption(option, id) {
			const container = document.getElementById(id);
			if (!container) return;
			const hiddenInput = container.querySelector('input[type="hidden"]');
			const selectedLabel = document.getElementById(id + '-selected-label');
			if (!hiddenInput || !selectedLabel) return;

			hiddenInput.value = option.dataset.value;
			selectedLabel.textContent = option.textContent;
			hiddenInput.dispatchEvent(new Event('change', { bubbles: true }));

			toggleCustomSelect(id);
		}

		window.addEventListener('click', function(e) {
			const customSelects = document.querySelectorAll('.relative[id^="custom-select-"]');
			for (const select of customSelects) {
				if (!select.contains(e.target)) {
					const options = select.querySelector('ul');
					const button = select.querySelector('button');
					if (options && !options.classList.contains('hidden')) {
						options.classList.add('hidden');
						button.setAttribute('aria-expanded', 'false');
					}
				}
			}
		});

		function togglePasswordVisibility(checkbox) {
			const container = checkbox.closest('.admin-form-input-container');
			if (!container) return;
			const input = container.querySelector('input:not([type="checkbox"])');
			if(input) input.type = checkbox.checked ? 'text' : 'password';
		}

		function updatePasswordStrength(input) {
			const container = input.closest('.admin-form-input-container');
			if (!container) return;
			const strengthMeter = container.querySelector('.password-strength-meter');
			if (!strengthMeter) return;
			
			const password = input.value;
			let strength = 0;
			if (password.length >= 8) strength++;
			if (/[A-Z]/.test(password)) strength++;
			if (/[a-z]/.test(password)) strength++;
			if (/[0-9]/.test(password)) strength++;
			if (/[^A-Za-z0-9]/.test(password)) strength++;
			
			const colors = ['bg-danger', 'bg-warning', 'bg-warning', 'bg-info', 'bg-success', 'bg-success'];
			const widths = ['0%', '20%', '40%', '60%', '80%', '100%'];
			
			const inner = strengthMeter.querySelector('.password-strength-inner');
			if(inner) {
				inner.className = 'password-strength-inner h-full transition-all duration-300 ' + colors[strength];
				inner.style.width = widths[strength];
			}
		}
	`))
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
