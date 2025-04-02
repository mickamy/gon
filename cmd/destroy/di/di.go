package di

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
)

var Cmd = &cobra.Command{
	Use:   "di [domain]",
	Short: "Destroy a domain DI file",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		domain := args[0]
		if err := Destroy(cfg, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to destroy DI: %w", err)
		}
		fmt.Printf("✅ DI %s destroyed successfully.\n", args[0])
		return nil
	},
}

func Destroy(cfg *config.Config, domain string) error {
	outPath := filepath.Join(cfg.OutputDir, domain, "di", fmt.Sprintf("%s_di.go", strings.ToLower(domain)))
	if err := os.Remove(outPath); err != nil {
		return err
	}

	return nil
}
