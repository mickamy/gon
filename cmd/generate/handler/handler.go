package handler

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/templates"
)

type Action struct {
	Name   string
	Entity string
	Method string
	Path   string
}

type TemplateData struct {
	Actions         []Action
	DomainPackage   string
	TestUtilPackage string
}

var Cmd = &cobra.Command{
	Use:   "handler [entity] [actions]",
	Short: "Generate a handler to handle HTTP requests",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Generate(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to generate handler file: %w", err)
		}
		fmt.Println("✅ Handler file generated successfully.")
		return nil
	},
}

func Generate(cfg *config.Config, args []string, domain string) error {
	entity := caseconv.Capitalize(args[0])
	actions := args[1:]

	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", entity)
		domain = caseconv.Uncapitalize(entity)
	}

	data := TemplateData{
		Actions:         parseAction(entity, actions),
		DomainPackage:   cfg.DomainPackage(domain),
		TestUtilPackage: filepath.Join(cfg.BasePackage, cfg.TestUtilDir),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "handler", fmt.Sprintf("%s_handler.go", strings.ToLower(entity)))
	if err := templates.Render(cfg.HandlerTemplate, data, outPath); err != nil {
		return err
	}

	if cfg.HandlerTestTemplate != "" {
		testOutPath := filepath.Join(cfg.OutputDir, domain, "handler", fmt.Sprintf("%s_handler_test.go", strings.ToLower(entity)))
		if err := templates.Render(cfg.HandlerTestTemplate, data, testOutPath); err != nil {
			return err
		}
	} else {
		fmt.Println("⚠️ No handler test template specified. Skipping test file generation.")
	}

	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}

func parseAction(entity string, raw []string) []Action {
	var actions []Action
	for _, item := range raw {
		actions = append(actions, Action{
			Name:   caseconv.Capitalize(item),
			Entity: entity,
			Method: "Get",
			Path:   caseconv.SnakeCase(entity),
		})
	}
	return actions
}
