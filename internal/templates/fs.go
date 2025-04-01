package templates

import (
	"embed"
)

//go:embed defaults/**/*
var DefaultFS embed.FS
