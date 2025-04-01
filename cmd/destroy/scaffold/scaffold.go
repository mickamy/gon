package scaffold

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/destroy/handler"
	"github.com/mickamy/gon/cmd/destroy/model"
	"github.com/mickamy/gon/cmd/destroy/repository"
	"github.com/mickamy/gon/cmd/destroy/usecase"
	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
)

var domain string

var Cmd = &cobra.Command{
	Use:   "scaffold [entity]",
	Short: "Destroy model, repository, usecase, and handler for a domain entity",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if domain == "" {
			fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
			domain = name
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}

		fmt.Printf("🧨 Destroying domain: %s (entity: %s)\n\n", domain, name)

		// --- Model ---
		if err := model.Destroy(cfg, []string{name}, domain); err != nil {
			return fmt.Errorf("model destroy failed: %w", err)
		}
		modelPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))
		fmt.Printf("✅ model:      %s\n", modelPath)

		// --- Repository ---
		if err := repository.Destroy(cfg, []string{name}, domain); err != nil {
			return fmt.Errorf("repository destroy failed: %w", err)
		}
		repoPath := filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name)))
		fmt.Printf("✅ repository: %s\n", repoPath)

		// --- Usecase ---
		actions := []string{"list", "get", "create", "update", "delete"}
		var usecaseFiles []string
		for _, action := range actions {
			pascal := caseconv.PascalCase(action + "_" + name)
			if err := usecase.Destroy(cfg, []string{pascal}, domain); err != nil {
				return fmt.Errorf("usecase destroy failed: %w", err)
			}
			usecaseFiles = append(usecaseFiles, caseconv.SnakeCase(action+"_"+name))
		}
		fmt.Printf("✅ usecase:    %s\n", strings.Join(usecaseFiles, ", "))

		// --- Handler ---
		if err := handler.Destroy(cfg, []string{name}, domain); err != nil {
			return fmt.Errorf("handler destroy failed: %w", err)
		}
		handlerPath := filepath.Join(cfg.OutputDir, domain, "handler", fmt.Sprintf("%s_handler.go", strings.ToLower(name)))
		fmt.Printf("✅ handler:    %s\n", handlerPath)

		fmt.Println("\n🎉 Done.")
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
