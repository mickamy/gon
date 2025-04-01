package config

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
	BasePackage        string `mapstructure:"basePackage"`
	OutputDir          string `mapstructure:"outputDir"`
	DefaultDriver      Driver `mapstructure:"defaultDriver"`
	DefaultWeb         Web    `mapstructure:"defaultWeb"`
	DatabasePackage    string `mapstructure:"databasePackage"`
	ModelTemplate      string `mapstructure:"modelTemplate"`
	RepositoryTemplate string `mapstructure:"repositoryTemplate"`
}

func (c Config) DatabasePackagePath() string {
	p := strings.TrimPrefix(c.DatabasePackage, c.BasePackage)
	return strings.TrimPrefix(p, "/")
}

// Load reads the config from gon.yaml
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

// New allows injecting a Config manually (e.g. for tests)
func New(c Config) *Config {
	return &c
}

// LoadFromReader allows loading Config from any io.Reader (e.g. for testing)
func LoadFromReader(r io.Reader) (*Config, error) {
	var cfg Config
	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	return &cfg, nil
}
