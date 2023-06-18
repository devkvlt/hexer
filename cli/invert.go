package cli

import (
	"fmt"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var invert = cli.Command{
	Name:        "invert",
	Usage:       "Invert",
	Description: "Invert a color.",
	UsageText:   "hexer invert <color>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color", name)
		if argc != 1 {
			return cli.Exit(errStr, 1)
		}
		c := args.Get(0)
		i, err := hexer.Invert(c)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(i)
		return nil
	},
}
