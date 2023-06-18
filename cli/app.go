package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

//go:generate go run gen/set.go
//go:generate go run gen/get.go
//go:generate go run gen/modify.go

var App = &cli.App{
	Usage:       "hex color manipulator",
	Description: "Hexer is a command line tool that inspects and manipulates hex colors. All input colors must be in long hex format (6 hex digits prefixed with #.)",
	Authors:     []*cli.Author{{Name: "https://github.com/devkvlt"}},
	HideHelp:    true,
	CommandNotFound: func(ctx *cli.Context, s string) {
		badCmd := ctx.Args().First()
		fmt.Printf("Command %q not found. Run `hexer` or `hexer help` to view all the available commands.\n", badCmd)
		os.Exit(1)
	},
	Commands: []*cli.Command{
		&help,

		&red,
		&green,
		&blue,
		&hue,
		&saturation,
		&lightness,
		&luminance,

		&contratio,

		&setr,
		&setg,
		&setb,
		&seth,
		&sets,
		&setl,

		&lighten,
		&darken,
		&saturate,
		&desaturate,

		&invert,

		&mix,
		&grad,
		&lightgrad,
		&darkgrad,
	},
}
