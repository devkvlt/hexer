//go:build ignore

package main

import (
	"strings"

	"github.com/devkvlt/hexer/cli/gen"
)

var cmds = []gen.Command{
	{"Red", "Get the (0-255)-based red component of the input color in the RGB model."},
	{"Green", "Get the (0-255)-based green component of the input color in the RGB model."},
	{"Blue", "Get the (0-255)-based blue component of the input color in the RGB model."},

	{"Hue", "Get the (0-360)-based hue of the input color in the HSL model."},
	{"Saturation", "Get the (0-100)-based saturation of the input color in the HSL model."},
	{"Lightness", "Get the (0-100)-based lightness of the input color in the HSL model."},

	{"Luminance", "Get the relative luminance of the input color, i.e. the Y component in the XYZ model."},
}

var templateFile = `package cli

import (
    "fmt"

    "github.com/devkvlt/hexer"
    "github.com/urfave/cli/v2"
)

var {{ToLower .Name}} = cli.Command{
	Name:        "{{ToLower .Name}}",
	Usage:       "Get the {{ToLower .Name}}{{if IsColor .Name}} component{{end}}",
	Description: "{{.Description}}",
	UsageText:   "hexer {{ToLower .Name}} <color>",
    Action: func(ctx *cli.Context) error {
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color", name)
		if ctx.NArg() != 1 {
			return cli.Exit(errStr, 1)
		}
		c := ctx.Args().Get(0)
		x, err := hexer.{{.Name}}(c)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(x)
		return nil
    },
}`

func main() {
	for _, cmd := range cmds {
		filename := "z" + strings.ToLower(cmd.Name) + ".go"
		gen.Generate(cmd, templateFile, filename)
	}
}
