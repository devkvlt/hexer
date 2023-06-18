// Code generated by cli/genmodify.go. DO NOT EDIT.

package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var desaturate = cli.Command{
	Name:        "desaturate",
	Usage:       "Desaturate",
	Description: "Desaturate a color by subtracting the given amount from its (0-100)-based saturation component in the HSL model.",
	UsageText:   "hexer desaturate <color> <amount>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color and a number", name)
		if argc != 2 {
			return cli.Exit(errStr, 1)
		}
		a, err := strconv.ParseFloat(args.Get(1), 64)
		if err != nil {
			return cli.Exit(errStr, 1)
		}
		c := args.Get(0)
		c_, err := hexer.Desaturate(c, a)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(c_)
		return nil
	},
}