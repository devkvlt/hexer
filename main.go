package main

import (
	"os"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		os.Exit(1) // TODO: is this not redundant?
	}
}
