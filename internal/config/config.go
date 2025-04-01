package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Driver string

func (d Driver) String() string {
	return string(d)
}

const (
	DriverGorm Driver = "gorm"
)

type Web string

func (w Web) String() string {
	return string(w)
}

const (
	WebEcho Web = "echo"
)

type Config struct {
	BasePackage     string `mapstructure:"basePackage"`
	OutputDir       string `mapstructure:"outputDir"`
	DefaultDriver   Driver `mapstructure:"defaultDriver"`
	DefaultWeb      Web    `mapstructure:"defaultWeb"`
	DatabasePackage string `mapstructure:"databasePackage"`
}

func (c Config) DatabasePackagePath() string {
	p := strings.TrimPrefix(c.DatabasePackage, c.BasePackage)
	return strings.TrimPrefix(p, "/")
}

func Load() (*Config, error) {
	viper.SetConfigName("gon")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read gon.yaml: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	fmt.Printf("Driver: %s\n", cfg.DefaultDriver.String())
	fmt.Printf("Web framework: %s\n", cfg.DefaultWeb.String())

	return &cfg, nil
}
