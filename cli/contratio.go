package cli

import (
	"fmt"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var contratio = cli.Command{
	Name:        "contratio",
	Usage:       "Get the contrast ratio",
	Description: "Get the contrast ratio between two input colors.",
	UsageText:   "hexer contratio <color1> <color2>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects 2 hex colors", name)
		if argc != 2 {
			return cli.Exit(errStr, 1)
		}
		c1 := args.Get(0)
		c2 := args.Get(1)
		cr, err := hexer.ContrastRatio(c1, c2)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(cr)
		return nil
	},
}
