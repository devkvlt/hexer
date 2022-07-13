package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:   "mix",
			Usage:  "Mixes two colors",
			Action: actionMix,
		},
		{
			Name:   "lighten",
			Usage:  "Lightens a color",
			Action: actionLighten,
		},
		{
			Name:   "setr",
			Usage:  "Sets the red component of a color",
			Action: actionSetRed,
		},
		{
			Name:   "seth",
			Usage:  "Sets the hue of a color",
			Action: actionSetHue,
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
