package command

import (
	"fmt"
	"os"

	"releaseros/internal/config"

	commandline "releaseros/internal/cli"

	"github.com/urfave/cli/v2"

	logger "github.com/rs/zerolog/log"
)

var InitCommand cli.Command

func init() {
	InitCommand = cli.Command{
		Name:  "init",
		Usage: "Initialize a new configuration file.",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "config",
				Required: false,
				Usage:    "The path of the configuration file to generate.",
				Value:    config.DefaultFilename,
			},
			&commandline.DebugFlag,
		},
		Before: func(clictx *cli.Context) error {
			commandline.SetLogLevel(clictx)
			return nil
		},
		Action: func(clictx *cli.Context) error {
			var filepath string
			if filepath = clictx.Path("config"); filepath == "" {
				return fmt.Errorf("no config file path provided")
			}
			logger.Debug().Str("filepath config", filepath).Msg("")

			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0o644)
			if err != nil {
				return err
			}
			defer file.Close()

			logger.Debug().Msg("writing default config")
			if _, err := file.Write(config.DefaultConfig); err != nil {
				return err
			}

			return nil
		},
	}
}
