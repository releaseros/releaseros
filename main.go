package main

import (
	stdlog "log"
	"os"

	"releaseros/command"

	"github.com/urfave/cli/v2"
)

// Default build time variable.
// This value is orridden via ldflags.
var version = "unknown-version"

func main() {
	app := &cli.App{
		Name:    "releaseros",
		Version: version,
		Usage:   "Generate a release note based on Git repository.",
		Commands: []*cli.Command{
			&command.GenerateCommand,
			&command.InitCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		stdlog.Fatal(err)
	}
}
