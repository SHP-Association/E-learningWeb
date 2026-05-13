package components

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/icons"

	. "maragu.dev/gomponents"
)

// Icon renders a shared SVG icon by name with optional class overrides.
func Icon(name, class string) Node {
	return icons.Icon(name, class)
}
