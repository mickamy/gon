package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	BasePackage   string `mapstructure:"basePackage"`
	OutputDir     string `mapstructure:"outputDir"`
	DefaultDriver string `mapstructure:"defaultDriver"`
	DefaultWeb    string `mapstructure:"defaultWeb"`
}

func LoadConfig() (*Config, error) {
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

	return &cfg, nil
}
