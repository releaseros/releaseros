package command

import (
	"context"
	"fmt"

	"releaseros/internal/config"
	"releaseros/internal/releasenote"

	commandline "releaseros/internal/cli"

	"github.com/urfave/cli/v2"
)

var GenerateCommand cli.Command

func init() {
	GenerateCommand = cli.Command{
		Name:  "generate",
		Usage: "Generate the release note.",
		Description: `Generate the release note between the previous and latest tags.

For the first tag/release, if the "initial_release_message" is configured then
this message will be the content of the release note.
Otherwise, it will be generated between the first tag and the initial commit.

Output to stdout.
To redirect the output to a file, you may use:
    $ releaseros generate > release-note.md`,
		Action: generate,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "config",
				Required: false,
				Usage:    "The path of the configuration file to use.",
			},
			&commandline.DebugFlag,
		},
		Before: func(clictx *cli.Context) error {
			commandline.SetLogLevel(clictx)
			return nil
		},
	}
}

func loadConfigFromCliContext(c *cli.Context) (config.Config, error) {
	configFilePath := c.Path("config")
	if configFilePath == "" {
		return config.LoadDefaultConfig()
	}

	return config.LoadFromFilePath(configFilePath)
}

func generate(clictx *cli.Context) error {
	config, err := loadConfigFromCliContext(clictx)
	if err != nil {
		return err
	}

	releaseNote, err := releasenote.NewGenerator().Generate(context.TODO(), config)
	if err != nil {
		return err
	}

	fmt.Fprintf(clictx.App.Writer, "%s", releaseNote)

	return nil
}
