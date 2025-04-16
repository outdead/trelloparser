package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/kelseyhightower/envconfig"
	"github.com/outdead/golibs/logger"
	"github.com/outdead/trelloparser/libs/trello"
	"gopkg.in/yaml.v3"
)

var ErrInvalidConfig = errors.New("invalid config")

// Config contains config data structure.
type Config struct {
	App      App           `json:"app"      yaml:"app"`
	Logger   logger.Config `json:"logger"   yaml:"logger"`
	Trello   trello.Config `json:"trello"   yaml:"trello"`
	Markdown Markdown      `json:"markdown" yaml:"markdown"`
}

// NewConfig parses config from file.
func NewConfig(name string, noenv bool) (*Config, error) {
	cfg := new(Config)
	if err := cfg.ParseFromFile(name); err != nil {
		return nil, err
	}

	if !noenv {
		if err := cfg.ParseFromEnv(); err != nil {
			return nil, err
		}
	}

	if err := cfg.SetDefaults(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// ParseFromFile reads config data from file and binds to config struct.
func (cfg *Config) ParseFromFile(name string) error {
	file, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	switch ext := path.Ext(name); ext {
	case ".yaml":
		err = yaml.Unmarshal(file, cfg)
	case ".json":
		err = json.Unmarshal(file, cfg)
	default:
		err = fmt.Errorf("%w extension: %s", ErrInvalidConfig, ext)
	}

	return err
}

// ParseFromEnv reads config data from ENV and binds to config struct.
func (cfg *Config) ParseFromEnv() error {
	return envconfig.Process("", cfg)
}

// Validate checks config to required fields.
func (cfg *Config) Validate() error {
	return nil
}

// SetDefaults adds default data to config struct.
func (cfg *Config) SetDefaults() error {
	if err := cfg.App.SetDefaults(); err != nil {
		return err
	}

	if err := cfg.Markdown.SetDefaults(); err != nil {
		return err
	}

	return nil
}

// Print print config to console.
func (cfg *Config) Print(w io.Writer) error {
	js, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("print config: %w", err)
	}

	_, err = fmt.Fprintln(w, string(js))

	return err
}
