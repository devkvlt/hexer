package main

import (
	"os"

	"github.com/devkvlt/hexer/cli"
)

func main() {
	if err := cli.App.Run(os.Args); err != nil {
		os.Exit(1) // TODO: is this not redundant?
	}
}
