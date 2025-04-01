package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
)

type TemplateData struct {
	EntityName      string
	LowerEntityName string
	BasePackage     string
	DatabasePackage string
	DomainName      string
}

var Cmd = &cobra.Command{
	Use:   "repository [model] [fields]",
	Short: "Generate a repository to retrieve domain model",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Generate(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed generate repository file: %w", err)
		}
		fmt.Println("✅ Repository file generated successfully.")
		return nil
	},
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}

func Generate(cfg *config.Config, args []string, domain string) error {
	name := gon.Capitalize(args[0])
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}

	data := TemplateData{
		EntityName:      name,
		LowerEntityName: gon.Uncapitalize(name),
		BasePackage:     cfg.BasePackage,
		DatabasePackage: cfg.DatabasePackage,
		DomainName:      gon.Uncapitalize(domain),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name)))
	if err := renderToFile(cfg, data, outPath); err != nil {
		return err
	}

	return nil
}

func renderToFile(cfg *config.Config, data TemplateData, outPath string) error {
	b, err := os.ReadFile(cfg.RepositoryTemplate)
	if err != nil {
		return err
	}
	tmplContent := string(b)

	tmpl, err := template.New("repository").Parse(tmplContent)
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
			fmt.Printf("⚠️ Failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
