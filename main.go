package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:   "invert",
			Usage:  "Inverts a color",
			Action: invertActionFunc,
		},
		{
			Name:   "contrast",
			Usage:  "Prints the contrast ration of 2 colors",
			Action: contrastRatioActionFunc,
		},

		{
			Name:   "getr",
			Usage:  "Prints the red component of a color",
			Action: makeActionFunc2(getRed),
		},
		{
			Name:   "getg",
			Usage:  "Prints the green component of a color",
			Action: makeActionFunc2(getGreen),
		},
		{
			Name:   "getb",
			Usage:  "Prints the blue component of a color",
			Action: makeActionFunc2(getBlue),
		},
		{
			Name:   "geth",
			Usage:  "Prints the hue component of a color",
			Action: makeActionFunc2(getHue),
		},
		{
			Name:   "gets",
			Usage:  "Prints the saturation component of a color",
			Action: makeActionFunc2(getSaturation),
		},
		{
			Name:   "getl",
			Usage:  "Prints the lightness component of a color",
			Action: makeActionFunc2(getLightness),
		},
		{
			Name:   "luminance",
			Usage:  "Prints the luminance of a color",
			Action: makeActionFunc2(luminance),
		},
		{
			Name:   "lighten",
			Usage:  "Lightens a color",
			Action: makeActionFunc(lighten),
		},
		{
			Name:   "darken",
			Usage:  "Darkens a color",
			Action: makeActionFunc(darken),
		},
		{
			Name:   "saturate",
			Usage:  "Saturates a color",
			Action: makeActionFunc(saturate),
		},
		{
			Name:   "unsaturate",
			Usage:  "Unsaturates a color",
			Action: makeActionFunc(unsaturate),
		},
		{
			Name:   "setr",
			Usage:  "Sets the red component of a color",
			Action: makeActionFunc(setRed),
		},
		{
			Name:   "setg",
			Usage:  "Sets the green component of a color",
			Action: makeActionFunc(setGreen),
		},
		{
			Name:   "setb",
			Usage:  "Sets the blue component of a color",
			Action: makeActionFunc(setBlue),
		},
		{
			Name:   "seth",
			Usage:  "Sets the hue of a color",
			Action: makeActionFunc(setHue),
		},
		{
			Name:   "sets",
			Usage:  "Sets the saturation of a color",
			Action: makeActionFunc(setSaturation),
		},
		{
			Name:   "setl",
			Usage:  "Sets the lightness of a color",
			Action: makeActionFunc(setLightness),
		},
		{
			Name:   "mix",
			Usage:  "Mixes two colors",
			Action: mixActionFunc,
		},
		{
			Name:   "mixpal",
			Usage:  "Makes a pelette of colors between two colors",
			Action: mixPaletteActionFunc,
		},

		{
			Name:   "darkpal",
			Usage:  "Makes a dark palette for a color",
			Action: darkPaletteActionFunc,
		},

		{
			Name:   "lightpal",
			Usage:  "Makes a light palette for a color",
			Action: lightPaletteActionFunc,
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
