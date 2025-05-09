package main

import (
	"fmt"
	"os"

	"github.com/outdead/trelloparser/internal/trelloparser"
	"github.com/outdead/trelloparser/internal/trelloparser/config"
	"github.com/outdead/trelloparser/internal/utils/logger"
	"github.com/urfave/cli/v2"
)

// Name contains the name of the service. Displayed in logs and when help
// command is called.
const Name = "trelloparser"

// Description contains description of the service.
const Description = "Parses Trello json files and converts them to selected formats"

// Version contains the service version number in the semantic versioning
// format (http://semver.org/). Displayed in logs and when help command is
// called. During service compilation, you can pass the version value using the
// flag `-ldflags "-X main.Version=${VERSION}"`.
var Version = "0.0.0"

// ConfigPath contains path to config if no flag -c value was passed.
// Can be replaced while compiling with flag
// `-ldflags "-X main.ConfigPath=./build/config/trelloparser/local.yaml"`.
var ConfigPath = "./config.yaml"

func main() {
	log := logger.New()

	app := NewApp(log)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func NewApp(log *logger.Logger) *cli.App {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Description = Description
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Path to config file",
			Value:   ConfigPath,
		},
		&cli.BoolFlag{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Print config file and exit",
		},
		&cli.BoolFlag{
			Name:  "noenv",
			Usage: "Ignore env variables",
		},
	}

	app.Action = action(log)

	return app
}

// action returns cli action function.
func action(log *logger.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		cfg, err := config.NewConfig(c.String("config"), c.Bool("noenv"))
		if err != nil {
			return fmt.Errorf("new config: %w", err)
		}

		if c.Bool("print") {
			log.Info("got -p flag - print config and terminate")

			return cfg.Print(os.Stdout)
		}

		if err := cfg.Validate(); err != nil {
			return fmt.Errorf("validate config: %w", err)
		}

		o := trelloparser.New(Name, Version, cfg, log)
		defer o.Close()

		return o.Run()
	}
}
