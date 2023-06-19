package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var adjust = cli.Command{
	Name:        "adjust",
	Usage:       "Adjust for contrast ratio",
	Description: "Adjust the lightness of the first color so that its contrast ratio with the second color is as close as possible of given contrast ratio of reference.",
	UsageText:   "hexer best <color1> <color2> <contrast ratio>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects two hex colors and a contrast ratio (number between 1 and 21)", name)
		// TODO: should add "run hexer help $name to get help"...
		if argc != 3 {
			return cli.Exit(errStr, 1)
		}
		c1 := args.Get(0)
		c2 := args.Get(1)
		cr, err := strconv.ParseFloat(args.Get(2), 64)
		if err != nil {
			return cli.Exit(errStr, 1)
		}
		c1, err = hexer.AdjustToContrastRatio(c1, c2, cr)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(c1)
		return nil
	},
}
