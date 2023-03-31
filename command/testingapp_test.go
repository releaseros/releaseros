package command

import (
	"io"

	"github.com/urfave/cli/v2"
)

var testingApp = &cli.App{
	ErrWriter: io.Discard,
	Commands: []*cli.Command{
		&GenerateCommand,
		&InitCommand,
	},
}
