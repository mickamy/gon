package model

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/templates"
)

type Field struct {
	Name string
	Type string
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
		if err := Generate(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to generate model file: %w", err)
		}
		fmt.Println("✅ Model file generated successfully.")
		return nil
	},
}

func Generate(cfg *config.Config, args []string, domain string) error {
	name := caseconv.Capitalize(args[0])
	fields := parseFields(args[1:])

	data := TemplateData{
		EntityName: name,
		Fields:     fields,
	}

	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))
	if err := templates.Render(cfg.ModelTemplate, data, outPath); err != nil {
		return err
	}

	if cfg.ModelTestTemplate != "" {
		testOutPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model_test.go", strings.ToLower(name)))
		if err := templates.Render(cfg.ModelTestTemplate, data, testOutPath); err != nil {
			return err
		}
	} else {
		fmt.Printf("⚠️ No model test template specified. Skipping test file generation.\n")
	}

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
			Name: caseconv.Capitalize(parts[0]),
			Type: parts[1],
		})
	}
	return fields
}
