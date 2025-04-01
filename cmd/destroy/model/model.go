package model

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
	Use:   "model",
	Short: "Destroy a model",
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

	fmt.Println("📄 Destroying model file...")
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))
	if err := os.Remove(outPath); err != nil {
		return fmt.Errorf("⚠️ Failed to remove model file %q: %w", outPath, err)
	}

	fmt.Println("✅ Model file destroyed successfully.")
	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
