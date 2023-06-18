package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var lightgrad = cli.Command{
	Name:        "lightgrad",
	Usage:       "Make light gradient colors",
	Description: "Make a pelette of n colors representing a gradient toward white. The first color is the input color and the last one is white (#ffffff).",
	UsageText:   "hexer lightgrad <color> <n>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color and an integer >= 2", name)
		if argc != 2 {
			return cli.Exit(errStr, 1)
		}
		c := args.Get(0)
		n, _ := strconv.Atoi(args.Get(1))
		palette, err := hexer.LightPalette(c, n)
		if err != nil {
			return cli.Exit(err, 1)
		}
		for _, v := range palette {
			fmt.Println(v)
		}
		return nil
	},
}
