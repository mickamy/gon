package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
)

var Cmd = &cobra.Command{
	Use:   "usecase",
	Short: "Destroy a usecase",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Destroy(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to destroy usecase: %w", err)
		}
		fmt.Printf("✅ Usecase %s destroyed successfully.\n", args[0])
		return nil
	},
}

func Destroy(cfg *config.Config, args []string, domain string) error {
	name := gon.Capitalize(args[0])

	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "usecase", fmt.Sprintf("%s_use_case.go", gon.SnakeCase(name)))
	if err := os.Remove(outPath); err != nil {
		return fmt.Errorf("⚠️ Failed to remove usecase file %q: %w", outPath, err)
	}

	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
