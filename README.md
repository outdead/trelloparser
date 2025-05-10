# trelloparser
Parses Trello json files and converts them to selected formats.

Trelloparser was created to solve my specific needs and wasn't designed to be a general-purpose tool. However, if you find it useful, you're welcome to use it, create issues, make forks, and submit pull requests.

## Install
Download the binary for your platform from the [latest releases](https://github.com/outdead/trelloparser/releases/latest)

See [Changelog](CHANGELOG.md) for release details.

## Usage
```text
NAME:
   trelloparser - A new cli application

USAGE:
   trelloparser [global options] command [command options]

VERSION:
   0.0.0

DESCRIPTION:
   Parses Trello json files and converts them to selected formats

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  Path to config file (default: "./config.yaml")
   --print, -p               Print config file and exit (default: false)
   --noenv                   Ignore env variables (default: false)
   --help, -h                show help
   --version, -v             print the version
```

## Configuration
For more convenient use, the ability to create the `config.yaml` configuration file provided. See [dist.yaml](build/config/trelloparser/dist.yaml).

## Usage as lib

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/outdead/trelloparser/libs/trello"
)

func main() {
	parser := trello.NewParser(trello.Config{})

	dash, err := parser.Parse("trello.json")
	if err != nil {
		log.Fatal(err)
	}

	js, _ := json.MarshalIndent(dash, "", "  ")

	fmt.Println(string(js))
}
```

## License
MIT License, see [LICENSE](LICENSE)
