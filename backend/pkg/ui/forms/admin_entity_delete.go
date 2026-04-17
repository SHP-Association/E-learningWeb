package forms

import (
	"net/http"

	"github.com/SHP-Association/E-learningWeb/backend/ent/admin"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/routenames"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"
	. "github.com/SHP-Association/E-learningWeb/backend/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AdminEntityDelete(r *ui.Request, entityType admin.EntityType) Node {
	return Form(
		Method(http.MethodPost),
		P(
			Textf("Are you sure you want to delete this %s?", entityType.GetName()),
		),
		ControlGroup(
			FormButton(ColorError, "Delete"),
			ButtonLink(
				ColorNone,
				r.Path(routenames.AdminEntityList(entityType.GetName())),
				"Cancel",
			),
		),
		CSRF(r),
	)
}
