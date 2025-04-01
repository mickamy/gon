package templates

import (
	"embed"
)

//go:embed defaults/*.tmpl
var DefaultFS embed.FS
