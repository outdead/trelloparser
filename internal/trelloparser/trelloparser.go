package trelloparser

import (
	"fmt"
	"github.com/outdead/trelloparser/internal/trelloparser/config"
	"github.com/outdead/trelloparser/internal/trelloparser/usecases"
	"github.com/outdead/trelloparser/internal/utils/logger"
)

type TrelloParser struct {
	name    string
	version string
	config  config.Config
	logger  *logger.Logger
	errors  chan error

	service *usecases.Usecases
}

// New creates new TrelloParser.
func New(name, ver string, cfg *config.Config, log *logger.Logger) *TrelloParser {
	return &TrelloParser{
		name:    name,
		version: ver,
		config:  *cfg,
		logger:  log,
		errors:  make(chan error, cfg.App.ErrorBuffer),
	}
}

func (d *TrelloParser) Run() error {
	if err := d.init(); err != nil {
		return err
	}

	d.logger.Infof("start %s daemon success", d.name)

	if err := d.service.CreateMarkdown(d.config.App.HomeDirectory + "/.tmp/data/source/Trello - todo.json"); err != nil {
		return err
	}

	return nil
}

func (d *TrelloParser) Close() {
	d.close()
}

func (d *TrelloParser) init() error {
	if err := d.logger.Customize(&d.config.Logger); err != nil {
		return fmt.Errorf("customize logger error: %w", err)
	}

	d.service = usecases.New(&d.config, d.logger)

	return nil
}

func (d *TrelloParser) close() {
	d.logger.Debugf("stopping %s daemon...", d.name)

	var errs []error

	if len(errs) != 0 {
		for _, e := range errs {
			d.logger.WithError(e).Errorf("close %s daemon error", d.name)
		}

		d.logger.Errorf("%s daemon closed with %d errors", d.name, len(errs))

		return
	}

	d.logger.Infof("%s daemon closed", d.name)
}
