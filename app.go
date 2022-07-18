package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func makeGetCmd(ch string, f func(string) (float64, error)) *cli.Command {
	return &cli.Command{
		Name:        ch,
		Usage:       "print " + ch,
		Description: "Print the " + ch + " component of the input color.",
		UsageText:   "hexer " + ch + " <color>",
		Action:      makeActionFunc2(f),
	}
}

func makeSetCmd(ch string, f func(string, float64) (string, error)) *cli.Command {
	return &cli.Command{
		Name:        "set" + ch[:1],
		Usage:       "set " + ch,
		Description: "Set the " + ch + " component of the input color to the provided value.",
		UsageText:   "hexer set" + ch[:1] + " <color> <value>",
		Action:      makeActionFunc(f),
	}
}

// I redefine this help cmd action just to get rid of parasitic -h flag
func showHelp(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Present() {
		return cli.ShowCommandHelp(ctx, args.First())
	}
	_ = cli.ShowAppHelp(ctx)
	return nil
}

var app = &cli.App{
	Usage: "hex color manipulator",

	Description: "Hexer is a command line tool that inspects and manipulates hex colors. All input colors must be in long hex format (6 hex digits prefixed with #.)",

	Authors: []*cli.Author{{Name: "https://github.com/devkvlt"}},

	HideHelp: true,

	CommandNotFound: func(ctx *cli.Context, s string) {
		badCmd := ctx.Args().First()
		fmt.Printf("Command %q not found. Run `hexer` or `hexer help` to view all the available commands.\n", badCmd)
		os.Exit(1)
	},

	Commands: []*cli.Command{
		{
			Name:        "help",
			Usage:       "show help",
			Description: "Show help for a command.",
			UsageText:   "hexer help [command]",
			Action:      showHelp,
		},
		makeGetCmd("red", getRed),
		makeGetCmd("green", getGreen),
		makeGetCmd("blue", getBlue),
		makeGetCmd("hue", getHue),
		makeGetCmd("saturation", getSaturation),
		makeGetCmd("lightness", getLightness),
		{
			Name:        "luminance",
			Usage:       "print luminance",
			Description: "Print the relative luminance of the input color.",
			UsageText:   "hexer luminance <color>",
			Action:      makeActionFunc2(luminance),
		},
		{
			Name:        "contratio",
			Usage:       "print contrast ratio",
			Description: "Print the contrast ratio between two input colors.",
			UsageText:   "hexer contratio <color1> <color2>",
			Action:      contrastRatioActionFunc,
		},
		makeSetCmd("red", setRed),
		makeSetCmd("green", setGreen),
		makeSetCmd("blue", setBlue),
		makeSetCmd("hue", setHue),
		makeSetCmd("saturation", setSaturation),
		makeSetCmd("lightness", setLightness),
		{
			Name:        "lighten",
			Usage:       "lighten color",
			Description: "Lighten a color by adding the given amount to its lightness component.",
			UsageText:   "hexer lighten <color> <amount>",
			Action:      makeActionFunc(lighten),
		},
		{
			Name:        "darken",
			Usage:       "darken color",
			Description: "Darken a color by subtracting the given amount from its lightness component.",
			UsageText:   "hexer darken <color> <amount>",
			Action:      makeActionFunc(darken),
		},
		{
			Name:        "saturate",
			Usage:       "saturate color",
			Description: "Saturate a color by adding the given amount to its saturation component.",
			UsageText:   "hexer saturate <color> <amount>",
			Action:      makeActionFunc(saturate),
		},
		{
			Name:        "desaturate",
			Usage:       "desaturate color",
			Description: "Desaturate a color by subtracting the given amount from its saturation component.",
			UsageText:   "hexer desaturate <color> <amount>",
			Action:      makeActionFunc(unsaturate),
		},
		{
			Name:        "invert",
			Usage:       "invert color",
			Description: "Invert a color.",
			UsageText:   "hexer invert <color>",
			Action:      invertActionFunc,
		},
		{
			Name:        "mix",
			Usage:       "mix colors",
			Description: "Mix two colors using the provided weight . The weight is that of the first color and it is optional. If no weight is given, it is set to 0.5.",
			UsageText:   "hexer mix <color1> <color2> [weight]",
			Action:      mixActionFunc,
		},
		{
			Name:        "grad",
			Usage:       "make gradient colors",
			Description: "Make a pelette of n colors representing the gradient colors between the two input colors. n must be greater than or equal 2. The first and last colors of the palette are the input colors.",
			UsageText:   "hexer grad <color1> <color2> <n>",
			Action:      mixPaletteActionFunc,
		},
		{
			Name:        "darkgrad",
			Usage:       "make dark gradient colors",
			Description: "Make a pelette of n colors representing a gradient toward black. The first color is the input color and the last one is black (#000000).",
			UsageText:   "hexer darkgrad <color> <n>",
			Action:      darkPaletteActionFunc,
		},
		{
			Name:        "lightgrad",
			Usage:       "make light gradient colors",
			Description: "Make a pelette of n colors representing a gradient toward white. The first color is the input color and the last one is white (#ffffff).",
			UsageText:   "hexer lightgrad <color> <n>",
			Action:      lightPaletteActionFunc,
		},
	},
}
