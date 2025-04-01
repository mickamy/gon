package templates

import (
	"embed"
)

//go:embed defaults/**/* defaults/*
var FS embed.FS
