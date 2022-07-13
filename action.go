package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func actionMix(ctx *cli.Context) error {
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
