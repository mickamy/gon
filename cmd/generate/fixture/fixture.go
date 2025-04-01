package fixture

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
	Model       string
	BasePackage string
	DomainName  string
}

var Cmd = &cobra.Command{
	Use:   "fixture [name] [fields]",
	Short: "Generate a fixture of model",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Generate(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to generate fixture file: %w", err)
		}
		fmt.Println("✅ Fixture file generated successfully.")
		return nil
	},
}

func Generate(cfg *config.Config, args []string, domain string) error {
	name := caseconv.Capitalize(args[0])

	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}

	data := TemplateData{
		Model:       name,
		BasePackage: cfg.BasePackage,
		DomainName:  caseconv.Uncapitalize(domain),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "fixture", fmt.Sprintf("%s_fixture.go", strings.ToLower(name)))
	if err := templates.Render(cfg.FixtureTemplate, data, outPath); err != nil {
		return err
	}

	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
