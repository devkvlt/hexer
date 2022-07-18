# Hexer

![Go version](https://img.shields.io/github/go-mod/go-version/devkvlt/hexer)
[![Go Report Card](https://goreportcard.com/badge/github.com/devkvlt/hexer)](https://goreportcard.com/report/github.com/devkvlt/hexer)
![Build](https://img.shields.io/github/workflow/status/devkvlt/hexer/build?label=build)
![Tests](https://img.shields.io/github/workflow/status/devkvlt/hexer/test?label=tests)
[![Coverage](https://coveralls.io/repos/github/devkvlt/hexer/badge.svg)](https://coveralls.io/github/devkvlt/hexer)
![License](https://img.shields.io/github/license/devkvlt/hexer)

A command line tool to inspect and manipulate hex colors 🎨 with a simple and straightforward API.

## Installation

```bash
go install github.com/devkvlt/hexer@latest
```

## Usage

```
USAGE:
   hexer  command [command options] [arguments...]

DESCRIPTION:
   Hexer is a command line tool that inspects and manipulates hex colors. All input colors must be in long hex format (6 hex digits prefixed with #.)

COMMANDS:
   help        show help
   red         print red
   green       print green
   blue        print blue
   hue         print hue
   saturation  print saturation
   lightness   print lightness
   luminance   print luminance
   contratio   print contrast ratio
   setr        set red
   setg        set green
   setb        set blue
   seth        set hue
   sets        set saturation
   setl        set lightness
   lighten     lighten color
   darken      darken color
   saturate    saturate color
   desaturate  desaturate color
   invert      invert color
   mix         mix colors
   grad        make gradient colors
   darkgrad    make dark gradient colors
   lightgrad   make light gradient colors
```
