package cli

import (
	"github.com/urfave/cli/v2"
)

var help = cli.Command{
	Name:        "help",
	Usage:       "Show help",
	Description: "Show help for a command.",
	UsageText:   "hexer help [command]",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Present() {
			return cli.ShowCommandHelp(ctx, args.First())
		}
		_ = cli.ShowAppHelp(ctx)
		return nil
	},
}
