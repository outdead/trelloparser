package config

import (
	"os"
	"path/filepath"
)

type App struct {
	ProfilerAddr  string `json:"profiler_addr" yaml:"profiler_addr"`
	HomeDirectory string `json:"home_directory" yaml:"home_directory"`
}

// SetDefaults adds default data to config struct.
func (cfg *App) SetDefaults() error {
	if cfg.HomeDirectory == "" {
		home, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}

		cfg.HomeDirectory = home
	}

	return nil
}
