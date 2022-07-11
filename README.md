# Hexer

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/devkvlt/hexer)
[![Go Report Card](https://goreportcard.com/badge/github.com/devkvlt/hexer)](https://goreportcard.com/report/github.com/devkvlt/hexer)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/devkvlt/hexer/build?label=build)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/devkvlt/hexer/test?label=tests)
[![Coverage Status](https://coveralls.io/repos/github/devkvlt/hexer/badge.svg)](https://coveralls.io/github/devkvlt/hexer)
![GitHub](https://img.shields.io/github/license/devkvlt/hexer)

A command line tool to manipulate hex colors 🎨.

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
