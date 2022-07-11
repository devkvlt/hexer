# Hexer

[![Go Report Card](https://goreportcard.com/badge/github.com/devkvlt/hexer)](https://goreportcard.com/report/github.com/devkvlt/hexer) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/devkvlt/hexer)

A command line tool to manipulate hex colors.

## Installation

```bash
go install github.com/devkvlt/hexer@latest
```

## Usage

Only flag `-m` (mix colors) is available at the moment.

```bash
# the last arg is the weight of the first color (number between 0 and 1)
hexer -m "#ff0000" "#ffff00" 0.5
```
