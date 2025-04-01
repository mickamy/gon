package scaffold

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/mickamy/gon/cmd/generate/handler"
	"github.com/mickamy/gon/cmd/generate/model"
	"github.com/mickamy/gon/cmd/generate/repository"
	"github.com/mickamy/gon/cmd/generate/usecase"
	"github.com/mickamy/gon/internal/config"
)

var domain string

var Cmd = &cobra.Command{
	Use:   "scaffold [name] [fields]",
	Short: "Generate model, repository, usecase, and handler for a domain entity",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fields := args[1:]

		if domain == "" {
			fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
			domain = name
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}

		fmt.Printf("📦 Scaffolding domain: %s (entity: %s)\n\n", domain, name)

		// --- Model ---
		if err := model.Generate(cfg, append([]string{name}, fields...), domain); err != nil {
			return fmt.Errorf("model generation failed: %w", err)
		}
		fmt.Printf("✅ model:      %s\n", filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name))))

		// --- Repository ---
		if err := repository.Generate(cfg, []string{name}, domain); err != nil {
			return fmt.Errorf("repository generation failed: %w", err)
		}
		fmt.Printf("✅ repository: %s\n", filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name))))

		// --- Usecase ---
		actions := []string{"create", "get", "update", "delete"}
		var usecaseFiles []string
		for _, action := range actions {
			caser := cases.Title(language.English)
			ucName := caser.String(action)
			if err := usecase.Generate(cfg, []string{ucName}, domain); err != nil {
				return fmt.Errorf("usecase generation failed: %w", err)
			}
			snake := strings.ToLower(action + "_" + name)
			usecaseFiles = append(usecaseFiles, snake)
		}
		fmt.Printf("✅ usecase:    %s\n", strings.Join(usecaseFiles, ", "))

		// --- Handler ---
		if err := handler.Generate(cfg, append([]string{name}, actions...), domain); err != nil {
			return fmt.Errorf("handler generation failed: %w", err)
		}
		fmt.Printf("✅ handler:    %s\n", filepath.Join(cfg.OutputDir, domain, "handler", fmt.Sprintf("%s_handler.go", strings.ToLower(name))))

		fmt.Println("\n🎉 Done.")
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
