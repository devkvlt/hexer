package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var grad = cli.Command{
	Name:        "grad",
	Usage:       "Make gradient colors",
	Description: "Make a pelette of n colors representing the gradient colors between the two input colors. n must be greater than or equal 2. The first and last colors of the palette are the input colors.",
	UsageText:   "hexer grad <color1> <color2> <n>",
	Action: func(ctx *cli.Context) error {
		argc := ctx.NArg()
		args := ctx.Args()
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects two hex colors and an integer >= 2", name)
		if argc != 3 {
			return cli.Exit(errStr, 1)
		}
		c1 := args.Get(0)
		c2 := args.Get(1)
		n, _ := strconv.Atoi(args.Get(2))
		palette, err := hexer.MixPalette(c1, c2, n)
		if err != nil {
			return cli.Exit(err, 1)
		}
		for _, v := range palette {
			fmt.Println(v)
		}
		return nil
	},
}
