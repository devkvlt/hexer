//go:build ignore

package main

import (
	"strings"

	"github.com/devkvlt/hexer/cli/gen"
)

var cmds = []gen.Command{
	{"Red", "Set the (0-255)-based red component of the input color in the RGB model."},
	{"Green", "Set the (0-255)-based green component of the input color in the RGB model."},
	{"Blue", "Set the (0-255)-based blue component of the input color in the RGB model."},

	{"Hue", "Set the (0-360)-based hue of the input color in the HSL model."},
	{"Saturation", "Set the (0-100)-based saturation of the input color in the HSL model."},
	{"Lightness", "Set the (0-100)-based lightness of the input color in the HSL model."},
}

var templateFile = `package cli

import (
    "fmt"
	"strconv"

    "github.com/devkvlt/hexer"
    "github.com/urfave/cli/v2"
)

var set{{slice .Name 0 1 | ToLower}} = cli.Command{
    Name:        "set{{slice .Name 0 1 | ToLower}}",
	Usage:       "Set the {{ToLower .Name}}{{if IsColor .Name}} component{{end}}",
	Description: "{{.Description}}",
    UsageText:   "hexer set{{slice .Name 0 1 | ToLower}} <color> <value>",
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
		c_, err := hexer.Set{{.Name}}(c, a)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(c_)
		return nil
    },
}`

func main() {
	for _, cmd := range cmds {
		filename := "zset" + strings.ToLower(cmd.Name[:1]) + ".go"
		gen.Generate(cmd, templateFile, filename)
	}
}
