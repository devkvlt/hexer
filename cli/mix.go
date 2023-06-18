package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var mix = cli.Command{
	Name:        "mix",
	Usage:       "Mix two colors",
	Description: "Mix two colors using the provided weight of the first color. If no weight is given, it is set to 0.5.",
	UsageText:   "hexer mix <color1> <color2> [weight]",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects two hex colors and an optional weight (number between 0 and 1)", name)
		// TODO: should add "run hexer help $name to get help"...
		if argc != 2 && argc != 3 {
			return cli.Exit(errStr, 1)
		}
		c1 := args.Get(0)
		c2 := args.Get(1)
		w := 0.5
		if argc == 3 {
			w_, err := strconv.ParseFloat(args.Get(2), 64)
			if err != nil {
				return cli.Exit(errStr, 1)
			}
			w = w_
		}
		m, err := hexer.Mix(c1, c2, w)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(m)
		return nil
	},
}
