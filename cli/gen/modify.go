//go:build ignore

package main

import (
	"strings"

	"github.com/devkvlt/hexer/cli/gen"
)

var cmds = []gen.Command{
	{"Lighten", "Lighten a color by adding the given amount to its (0-100)-based lightness component in the HSL model."},
	{"Darken", "Darken a color by subtracting the given amount from its (0-100)-based lightness component in the HSL model."},
	{"Saturate", "Saturate a color by adding the given amount to its (0-100)-based saturation component in the HSL model."},
	{"Desaturate", "Desaturate a color by subtracting the given amount from its (0-100)-based saturation component in the HSL model."},
}

var templateFile = `package cli

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/hexer"
	"github.com/urfave/cli/v2"
)

var {{ToLower .Name}} = cli.Command{
	Name:        "{{ToLower .Name}}",
	Usage:       "{{.Name}}",
	Description: "{{.Description}}",
	UsageText:   "hexer {{ToLower .Name}} <color> <amount>",
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
		c_, err := hexer.{{.Name}}(c, a)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(c_)
		return nil
	},
}`

func main() {
	for _, cmd := range cmds {
		filename := "z" + strings.ToLower(cmd.Name) + ".go"
		gen.Generate(cmd, templateFile, filename)
	}
}
