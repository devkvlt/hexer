// Code generated by cli/gen/get.go. DO NOT EDIT.

package cli

import (
    "fmt"

    "github.com/devkvlt/hexer"
    "github.com/urfave/cli/v2"
)

var saturation = cli.Command{
	Name:        "saturation",
	Usage:       "Get the saturation",
	Description: "Get the (0-100)-based saturation of the input color in the HSL model.",
	UsageText:   "hexer saturation <color>",
    Action: func(ctx *cli.Context) error {
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color", name)
		if ctx.NArg() != 1 {
			return cli.Exit(errStr, 1)
		}
		c := ctx.Args().Get(0)
		x, err := hexer.Saturation(c)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(x)
		return nil
    },
}