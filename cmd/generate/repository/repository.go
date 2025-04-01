package repository

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/templates"
)

type TemplateData struct {
	Entity          string
	LowerEntity     string
	BasePackage     string
	DatabasePackage string
	Domain          string
}

var Cmd = &cobra.Command{
	Use:   "repository [model]",
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
	name := caseconv.Capitalize(args[0])
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}

	data := TemplateData{
		Entity:          name,
		LowerEntity:     caseconv.Uncapitalize(name),
		BasePackage:     cfg.BasePackage,
		DatabasePackage: cfg.DatabasePackage,
		Domain:          caseconv.Uncapitalize(domain),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name)))
	if err := templates.Render(cfg.RepositoryTemplate, data, outPath); err != nil {
		return err
	}

	if cfg.RepositoryTestTemplate != "" {
		testOutPath := filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository_test.go", strings.ToLower(name)))
		if err := templates.Render(cfg.RepositoryTestTemplate, data, testOutPath); err != nil {
			return err
		}
	} else {
		fmt.Println("⚠️ No repository test template specified. Skipping test file generation.")
	}

	return nil
}
