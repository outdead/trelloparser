package usecases

import (
	"github.com/outdead/golibs/logger"
	"github.com/outdead/trelloparser/internal/trelloparser/config"
	"github.com/outdead/trelloparser/internal/trelloparser/usecases/markdown"
)

type Usecases struct {
	*markdown.Markdown
}

func New(cfg *config.Config, log *logger.Logger) *Usecases {
	return &Usecases{
		Markdown: markdown.New(cfg, log),
	}
}
