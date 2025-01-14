package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	hcversion "github.com/hashicorp/go-version"
	"github.com/ldez/gomoddirectives"
)

// Config encapsulates the config data specified in the golangci-lint yaml config file.
type Config struct {
	cfgDir string // The directory containing the golangci-lint config file.

	Run Run `mapstructure:"run"`

	Output Output `mapstructure:"output"`

	LintersSettings LintersSettings `mapstructure:"linters-settings"`
	Linters         Linters         `mapstructure:"linters"`
	Issues          Issues          `mapstructure:"issues"`
	Severity        Severity        `mapstructure:"severity"`

	InternalCmdTest bool // Option is used only for testing golangci-lint command, don't use it
	InternalTest    bool // Option is used only for testing golangci-lint code, don't use it
}

// GetConfigDir returns the directory that contains golangci config file.
func (c *Config) GetConfigDir() string {
	return c.cfgDir
}

func (c *Config) Validate() error {
	for i, rule := range c.Issues.ExcludeRules {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("error in exclude rule #%d: %w", i, err)
		}
	}

	if len(c.Severity.Rules) > 0 && c.Severity.Default == "" {
		return errors.New("can't set severity rule option: no default severity defined")
	}
	for i, rule := range c.Severity.Rules {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("error in severity rule #%d: %w", i, err)
		}
	}

	return nil
}

func NewDefault() *Config {
	return &Config{
		LintersSettings: defaultLintersSettings,
	}
}

type Version struct {
	Format string `mapstructure:"format"`
	Debug  bool   `mapstructure:"debug"`
}

func IsGoGreaterThanOrEqual(current, limit string) bool {
	v1, err := hcversion.NewVersion(strings.TrimPrefix(current, "go"))
	if err != nil {
		return false
	}

	l, err := hcversion.NewVersion(limit)
	if err != nil {
		return false
	}

	return v1.GreaterThanOrEqual(l)
}

func detectGoVersion() string {
	file, _ := gomoddirectives.GetModuleFile()

	if file != nil && file.Go != nil && file.Go.Version != "" {
		return file.Go.Version
	}

	v := os.Getenv("GOVERSION")
	if v != "" {
		return v
	}

	return "1.17"
}
