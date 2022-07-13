# Hexer

![Go version](https://img.shields.io/github/go-mod/go-version/devkvlt/hexer)
[![Go Report Card](https://goreportcard.com/badge/github.com/devkvlt/hexer)](https://goreportcard.com/report/github.com/devkvlt/hexer)
![Build](https://img.shields.io/github/workflow/status/devkvlt/hexer/build?label=build)
![Tests](https://img.shields.io/github/workflow/status/devkvlt/hexer/test?label=tests)
[![Coverage](https://coveralls.io/repos/github/devkvlt/hexer/badge.svg)](https://coveralls.io/github/devkvlt/hexer)
![License](https://img.shields.io/github/license/devkvlt/hexer)

A command line tool to manipulate hex colors 🎨.

## Installation

```bash
go install github.com/devkvlt/hexer@latest
```

## Usage

```
USAGE:
   hexer <command> <arguments...>
COMMANDS:
   mix      Mixes two colors
   lighten  Lightens a color
   setr     Sets the red component of a color
   seth     Sets the hue of a color
```

Examples:

```bash
# the last arg is the weight of the first color (number between 0 and 1)
hexer mix "#ff0000" "#ffff00" 0.2

hexer mix "#ff0000" "#ffff00" # weight defaults to 0.5

hexer setr "#ff0000" 0
```
