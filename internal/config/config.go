package config

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type DBDriver string

func (d DBDriver) String() string {
	return string(d)
}

func IsValidDBDriver(driver string) bool {
	drivers := []DBDriver{DBDriverGorm}
	for _, d := range drivers {
		if driver == d.String() {
			return true
		}
	}
	return false
}

const (
	DBDriverGorm DBDriver = "gorm"
)

type WebFramework string

func IsValidWebFramework(web string) bool {
	fs := []WebFramework{WebFrameworkEcho}
	for _, w := range fs {
		if web == w.String() {
			return true
		}
	}
	return false
}

func (w WebFramework) String() string {
	return string(w)
}

const (
	WebFrameworkEcho WebFramework = "echo"
)

type DIFramework string

func IsValidDIFramework(di string) bool {
	fs := []DIFramework{DIFrameworkWire}
	for _, d := range fs {
		if di == d.String() {
			return true
		}
	}
	return false
}

func (d DIFramework) String() string {
	return string(d)
}

func (d DIFramework) InstallPackage() string {
	switch d {
	case DIFrameworkWire:
		return "github.com/google/wire/cmd/wire"
	}
	return ""
}

const (
	DIFrameworkWire DIFramework = "wire"
)

type Config struct {
	BasePackage            string       `mapstructure:"basePackage"`
	OutputDir              string       `mapstructure:"outputDir"`
	TestUtilDir            string       `mapstructure:"testUtilDir"`
	DBDriver               DBDriver     `mapstructure:"dbDriver"`
	WebFramework           WebFramework `mapstructure:"webFramework"`
	DIFramework            DIFramework  `mapstructure:"diFramework"`
	DatabasePackage        string       `mapstructure:"databasePackage"`
	ModelTemplate          string       `mapstructure:"modelTemplate"`
	ModelTestTemplate      string       `mapstructure:"modelTestTemplate"`
	FixtureTemplate        string       `mapstructure:"fixtureTemplate"`
	RepositoryTemplate     string       `mapstructure:"repositoryTemplate"`
	RepositoryTestTemplate string       `mapstructure:"repositoryTestTemplate"`
	UsecaseTemplate        string       `mapstructure:"usecaseTemplate"`
	UsecaseTestTemplate    string       `mapstructure:"usecaseTestTemplate"`
	HandlerTemplate        string       `mapstructure:"handlerTemplate"`
	HandlerTestTemplate    string       `mapstructure:"handlerTestTemplate"`
	DiTemplate             string       `mapstructure:"diTemplate"`
}

func (c Config) DatabasePackagePath() string {
	p := strings.TrimPrefix(c.DatabasePackage, c.BasePackage)
	return strings.TrimPrefix(p, "/")
}

func (c Config) DomainPackage(domain string) string {
	return c.BasePackage + "/" + c.OutputDir + "/" + domain
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
