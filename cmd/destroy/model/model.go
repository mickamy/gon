package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
)

var Cmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Destroy a model",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		if err := Destroy(cfg, args, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to destroy model: %w", err)
		}
		fmt.Printf("✅ Model %s destroyed successfully.\n", args[0])
		return nil
	},
}

func Destroy(cfg *config.Config, args []string, domain string) error {
	name := caseconv.Capitalize(args[0])

	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = name
	}
	outPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model.go", strings.ToLower(name)))
	if err := os.Remove(outPath); err != nil {
		return err
	}

	// Remove the test file if it exists
	testPath := filepath.Join(cfg.OutputDir, domain, "model", fmt.Sprintf("%s_model_test.go", strings.ToLower(name)))
	if _, err := os.Stat(testPath); err == nil {
		if err := os.Remove(testPath); err != nil {
			return err
		}
	}

	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
