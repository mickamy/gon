package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
)

var Cmd = &cobra.Command{
	Use:   "repository",
	Short: "Destroy a repository",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		return Destroy(cfg, args)
	},
}

func Destroy(cfg *config.Config, args []string) error {
	name := gon.Capitalize(args[0])

	fmt.Println("📄 Destroying repository file...")
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name)))
	if err := os.Remove(outPath); err != nil {
		return fmt.Errorf("⚠️ Failed to remove repository file %q: %w", outPath, err)
	}

	fmt.Println("✅ Repository file destroyed successfully.")
	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
