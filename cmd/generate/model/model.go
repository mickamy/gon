package model

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
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
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		return GenerateModel(cfg, args)
	},
}

func GenerateModel(cfg *config.Config, args []string) error {
	name := gon.Capitalize(args[0])
	fields := parseFields(args[1:])

	data := TemplateData{
		EntityName: name,
		Fields:     fields,
	}

	fmt.Println("📄 Generating model file...")
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))
	if err := renderToFile(cfg, data, outPath); err != nil {
		return err
	}

	fmt.Println("✅ Model file generated successfully.")
	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
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

func renderToFile(cfg *config.Config, data TemplateData, outPath string) error {
	tmplFile := cfg.ModelTemplate

	var bytes []byte
	var err error

	if _, statErr := os.Stat(tmplFile); statErr == nil {
		bytes, err = os.ReadFile(tmplFile)
		if err != nil {
			return fmt.Errorf("⚠️ Failed to read local template %q: %w", tmplFile, err)
		}
	} else {
		bytes, err = templates.DefaultFS.ReadFile(tmplFile)
		if err != nil {
			bytes, err = templates.DefaultFS.ReadFile(tmplFile)
			if err != nil {
				return fmt.Errorf(
					"⚠️ Failed to read embedded template %q: %w\n💡 You might need to run `gon install` to generate default templates.",
					tmplFile, err,
				)
			}
		}
	}
	tmplContent := string(bytes)

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
		if err := file.Close(); err != nil {
			fmt.Printf("⚠️ Failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
