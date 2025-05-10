package config

import (
	"os"
	"path/filepath"
)

type Markdown struct {
	SourceFile      string `json:"source_file"       yaml:"source_file"`
	DataDirectory   string `json:"data_directory"    yaml:"data_directory"`
	ResultDirectory string `json:"-"                 yaml:"-"`
	AddDateToCards  bool   `json:"add_date_to_cards" yaml:"add_date_to_cards"`
	Format          string `json:"format"            yaml:"format"`
	Footer          string `json:"footer"            yaml:"footer"`
}

// SetDefaults adds default data to config struct.
func (cfg *Markdown) SetDefaults() error {
	if cfg.DataDirectory == "" {
		home, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}

		cfg.DataDirectory = home + "/trelloparserdata"
	}

	cfg.ResultDirectory = cfg.DataDirectory + "/result"

	if cfg.Format == "" {
		cfg.Format = "markdown"
	}

	return nil
}
