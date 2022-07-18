package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func mixActionFunc(ctx *cli.Context) error {
	argc := ctx.NArg()
	args := ctx.Args()
	name := ctx.Command.FullName()
	errStr := fmt.Sprintf("command %q expects two hex colors and an optional number between 0 and 1", name)
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
	m, err := mix(c1, c2, w)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(m)
	return nil
}

// makeActionFunc turns a core function of two arguments into an action function
// to handle cli context.
func makeActionFunc(f func(string, float64) (string, error)) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
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
		c_, err := f(c, a)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(c_)
		return nil
	}
}

// makeActionFunc2 does same as makeActionFunc but for get functions.
// TODO: find good names for these funcs.
func makeActionFunc2(f func(string) (float64, error)) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		name := ctx.Command.FullName()
		errStr := fmt.Sprintf("command %q expects a hex color", name)
		if ctx.NArg() != 1 {
			return cli.Exit(errStr, 1)
		}
		c := ctx.Args().Get(0)
		x, err := f(c)
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(x)
		return nil
	}
}

func mixPaletteActionFunc(ctx *cli.Context) error {
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
	palette, err := mixPalette(c1, c2, n)
	if err != nil {
		return cli.Exit(err, 1)
	}
	for _, v := range palette {
		fmt.Println(v)
	}
	return nil
}

func darkPaletteActionFunc(ctx *cli.Context) error {
	argc := ctx.NArg()
	args := ctx.Args()
	name := ctx.Command.FullName()
	errStr := fmt.Sprintf("command %q expects a hex color and an integer >= 2", name)
	if argc != 2 {
		return cli.Exit(errStr, 1)
	}
	c := args.Get(0)
	n, _ := strconv.Atoi(args.Get(1))
	palette, err := darkPalette(c, n)
	if err != nil {
		return cli.Exit(err, 1)
	}
	for _, v := range palette {
		fmt.Println(v)
	}
	return nil
}

func lightPaletteActionFunc(ctx *cli.Context) error {
	argc := ctx.NArg()
	args := ctx.Args()
	name := ctx.Command.FullName()
	errStr := fmt.Sprintf("command %q expects a hex color and an integer >= 2", name)
	if argc != 2 {
		return cli.Exit(errStr, 1)
	}
	c := args.Get(0)
	n, _ := strconv.Atoi(args.Get(1))
	palette, err := lightPalette(c, n)
	if err != nil {
		return cli.Exit(err, 1)
	}
	for _, v := range palette {
		fmt.Println(v)
	}
	return nil
}

func invertActionFunc(ctx *cli.Context) error {
	argc := ctx.NArg()
	args := ctx.Args()
	name := ctx.Command.FullName()
	errStr := fmt.Sprintf("command %q expects a hex color", name)
	if argc != 1 {
		return cli.Exit(errStr, 1)
	}
	c := args.Get(0)
	i, err := invert(c)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(i)
	return nil
}

func contrastRatioActionFunc(ctx *cli.Context) error {
	argc := ctx.NArg()
	args := ctx.Args()
	name := ctx.Command.FullName()
	errStr := fmt.Sprintf("command %q expects 2 hex colors", name)
	if argc != 2 {
		return cli.Exit(errStr, 1)
	}
	c1 := args.Get(0)
	c2 := args.Get(1)
	cr, err := contrastRatio(c1, c2)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(cr)
	return nil
}
