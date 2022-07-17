package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func handleMix(ctx *cli.Context) error {
	argc := ctx.NArg()
	if argc < 2 {
		return cli.Exit("command mix requires at least 3 arguments", 1)
	} else if argc > 3 {
		return cli.Exit("too many arguments", 1)
	}
	c1 := ctx.Args().Get(0)
	c2 := ctx.Args().Get(1)
	w := 0.5
	if argc == 3 {
		w_, err := strconv.ParseFloat(ctx.Args().Get(2), 64)
		if err != nil {
			return cli.Exit("invalid weight, weight has to be a number between 0 and 1", 1)
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

func handleLighten(ctx *cli.Context) error {
	argc := ctx.NArg()
	if argc < 2 {
		return cli.Exit("command lighten requires 2 arguments", 1)
	} else if argc > 2 {
		return cli.Exit("too many arguments", 1)
	}
	c := ctx.Args().Get(0)
	a, err := strconv.ParseFloat(ctx.Args().Get(1), 64)
	if err != nil {
		return cli.Exit("amount must be a number >= 0", 1)
	}
	c_, err := lighten(c, a)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(c_)
	return nil
}

func handleDarken(ctx *cli.Context) error {
	argc := ctx.NArg()
	if argc < 2 {
		return cli.Exit("command darken requires 2 arguments", 1)
	} else if argc > 2 {
		return cli.Exit("too many arguments", 1)
	}
	c := ctx.Args().Get(0)
	a, err := strconv.ParseFloat(ctx.Args().Get(1), 64)
	if err != nil {
		return cli.Exit("amount must be a number >= 0", 1)
	}
	c_, err := darken(c, a)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(c_)
	return nil
}

func handleSetr(ctx *cli.Context) error {
	argc := ctx.NArg()
	if argc < 2 {
		return cli.Exit("command setr requires 2 arguments", 1)
	} else if argc > 2 {
		return cli.Exit("too many arguments", 1)
	}
	c := ctx.Args().Get(0)
	r, err := strconv.ParseFloat(ctx.Args().Get(1), 64)
	if err != nil {
		return cli.Exit("invalid r, r must be a number between 0 and 255", 1)
	}
	newC, err := setRed(c, r)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(newC)
	return nil
}

func handleSeth(ctx *cli.Context) error {
	argc := ctx.NArg()
	if argc < 2 {
		return cli.Exit("command seth requires 2 arguments", 1)
	} else if argc > 2 {
		return cli.Exit("too many arguments", 1)
	}
	c := ctx.Args().Get(0)
	h, err := strconv.ParseFloat(ctx.Args().Get(1), 64)
	if err != nil {
		return cli.Exit("invalid h, h must be a number between 0 and 360", 1)
	}
	newC, err := setHue(c, h)
	if err != nil {
		return cli.Exit(err, 1)
	}
	fmt.Println(newC)
	return nil
}
