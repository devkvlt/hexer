# Hexer

![Go version](https://img.shields.io/github/go-mod/go-version/devkvlt/hexer)
[![Go Report Card](https://goreportcard.com/badge/github.com/devkvlt/hexer)](https://goreportcard.com/report/github.com/devkvlt/hexer)
![Build](https://img.shields.io/github/actions/workflow/status/devkvlt/hexer/build.yml?branch=main&label=build)
![Tests](https://img.shields.io/github/actions/workflow/status/devkvlt/hexer/test.yml?branch=main&label=tests)
[![Coverage](https://github.com/devkvlt/hexer/wiki/coverage.svg)](https://raw.githack.com/wiki/devkvlt/hexer/coverage.html)
![License](https://img.shields.io/github/license/devkvlt/hexer)

A command line tool and Go library to inspect and manipulate hex colors.

## Installation

```bash
go install github.com/devkvlt/hexer/cmd/hexer@latest
```

## Usage

Run `hexer` or `hexer help` to get a list of available commands.

Run `hexer help <cmd>` to get detailed help about the command `<cmd>`.

Here are some examples of usage:

```bash
# Mix red and green
hexer mix "#ff0000" "#008000"

# Darken red by 10
hexer darken "#ff0000" 10

# Get the contrast ratio between red and green
hexer contratio "#ff0000" "#008000"

# Make 5 colors representing gradient colors between red and green
hexer grad "#ff0000" "#008000" 5
```

## Usage as a library

```bash
go get github.com/devkvlt/hexer
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/devkvlt/hexer"
)

func main() {
	red := "#ff0000"
	green := "#008000"

	mix, err := hexer.Mix(red, green, 0.3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mix)
}
```
