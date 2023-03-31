package command

import (
	"fmt"

	"releaseros/internal/config"
	"releaseros/internal/context"
	"releaseros/internal/releasenote"

	commandline "releaseros/internal/cli"

	"github.com/urfave/cli/v2"
)

var GenerateCommand cli.Command

func init() {
	GenerateCommand = cli.Command{
		Name:   "generate",
		Usage:  "Do the generation.",
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

	ctx := context.New(config)

	releaseNote, err := releasenote.NewGenerator().Generate(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(clictx.App.Writer, "%s", releaseNote)

	return nil
}
