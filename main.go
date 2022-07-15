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
			Action: handleMix,
		},
		{
			Name:   "lighten",
			Usage:  "Lightens a color",
			Action: handleLighten,
		},
		{
			Name:   "darken",
			Usage:  "Darkens a color",
			Action: handleDarken,
		},
		{
			Name:   "setr",
			Usage:  "Sets the red component of a color",
			Action: handleSetr,
		},
		{
			Name:   "seth",
			Usage:  "Sets the hue of a color",
			Action: handleSeth,
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
