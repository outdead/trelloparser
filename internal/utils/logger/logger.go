package logger

import (
	"fmt"
	"io"
	"time"

	"github.com/outdead/golibs/files"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Level string `json:"level" yaml:"level"`
	File  struct {
		Path   string `json:"path"   yaml:"path"`
		Layout string `json:"layout" yaml:"layout"`
	} `json:"file" yaml:"file"`
}

type Logger struct {
	config Config
	*logrus.Logger
}

func New() *Logger {
	logger := &Logger{
		Logger: logrus.New(),
	}
	logger.Formatter = new(logrus.JSONFormatter)

	return logger
}

func (logger *Logger) AddOutput(w io.Writer) {
	logger.Out = io.MultiWriter(logger.Out, w)
}

func (logger *Logger) Customize() error {
	if logger.config.File.Layout != "" {
		file, err := files.CreateAndOpenFile(logger.config.File.Path, time.Now().Format(logger.config.File.Layout))
		if err != nil {
			return fmt.Errorf("create logger file hook: %w", err)
		}

		logger.AddOutput(file)
	}

	return nil
}

func (logger *Logger) Writer() io.Writer {
	return logger.Logger.Writer()
}

func (logger *Logger) Close() error {
	return nil
}
