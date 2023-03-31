package cli

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	urfavecli "github.com/urfave/cli/v2"
)

var DebugFlag = urfavecli.BoolFlag{
	Name:    "debug",
	Aliases: []string{"verbose", "v", "vv", "vvv"},
	Usage:   "Enable debug mode.",
}

func SetLogLevel(clictx *urfavecli.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.Disabled)
	if clictx.Bool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
