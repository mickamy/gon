package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Render(tmplFile string, data any, out string) error {
	var bytes []byte
	var err error

	if _, statErr := os.Stat(tmplFile); statErr == nil {
		bytes, err = os.ReadFile(tmplFile)
		if err != nil {
			return fmt.Errorf("⚠️ Failed to read local template %q: %w", tmplFile, err)
		}
	} else {
		bytes, err = FS.ReadFile(tmplFile)
		if err != nil {
			bytes, err = FS.ReadFile(tmplFile)
			if err != nil {
				return fmt.Errorf(
					"⚠️ Failed to read embedded template %q: %w\n💡 You might need to run `gon install` to generate default templates.",
					tmplFile, err,
				)
			}
		}
	}
	tmplContent := string(bytes)

	tmpl, err := template.New(filepath.Base(tmplFile)).Parse(tmplContent)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
		return err
	}

	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("⚠️ Failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
