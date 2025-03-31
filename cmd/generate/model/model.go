package model

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/gon"
	"github.com/mickamy/gon/templates"
)

type Field struct {
	Name     string
	Type     string
	JSONName string
}

type TemplateData struct {
	EntityName string
	Fields     []Field
}

var Cmd = &cobra.Command{
	Use:   "model [name] [fields]",
	Short: "Generate a domain model",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fields := parseFields(args[1:])

		data := TemplateData{
			EntityName: name,
			Fields:     fields,
		}

		fmt.Println("Generating model...")
		if err := renderToFile(data, filepath.Join("internal", "domain", "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))); err != nil {
			return err
		}

		fmt.Println("Model generated successfully.")
		return nil
	},
}

func parseFields(raw []string) []Field {
	var fields []Field
	for _, item := range raw {
		parts := strings.Split(item, ":")
		if len(parts) != 2 {
			continue
		}
		fields = append(fields, Field{
			Name:     gon.Capitalize(parts[0]),
			Type:     parts[1],
			JSONName: parts[0],
		})
	}
	return fields
}

func renderToFile(data TemplateData, outPath string) error {
	b, err := templates.FS.ReadFile("model.tmpl")
	if err != nil {
		return err
	}
	tmplContent := string(b)

	tmpl, err := template.New("model").Parse(tmplContent)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
