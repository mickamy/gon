package usecase

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/templates"
)

type TemplateData struct {
	Name              string
	UncapitalizedName string
	DomainPackage     string
}

var Cmd = &cobra.Command{
	Use:   "usecase [name]",
	Short: "Generate a usecase to execute business logic",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Generate(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed generate usecase file: %w", err)
		}
		fmt.Println("✅ Usecase file generated successfully.")
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
		domain = caseconv.SnakeCase(name)
	}

	data := TemplateData{
		Name:              name,
		UncapitalizedName: caseconv.Uncapitalize(name),
		DomainPackage:     cfg.DomainPackage(domain),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "usecase", fmt.Sprintf("%s_use_case.go", caseconv.SnakeCase(name)))
	if err := templates.Render(cfg.UsecaseTemplate, data, outPath); err != nil {
		return err
	}

	if cfg.UsecaseTestTemplate != "" {
		testOutPath := filepath.Join(cfg.OutputDir, domain, "usecase", fmt.Sprintf("%s_use_case_test.go", caseconv.SnakeCase(name)))
		if err := templates.Render(cfg.UsecaseTestTemplate, data, testOutPath); err != nil {
			return err
		}
	} else {
		fmt.Println("⚠️ No usecase test template provided. Skipping test file generation.")
	}

	return nil
}
