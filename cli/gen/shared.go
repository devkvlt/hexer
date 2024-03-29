package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Command struct {
	Name        string
	Description string
}

func Generate(cmd Command, templateFile, filename string) {
	header := "// Code generated by cli/gen/" + filepath.Base(os.Args[0]) + ".go. DO NOT EDIT.\n\n"

	tmpl, err := template.New("template").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
		"IsColor": func(s string) bool { return s == "Red" || s == "Green" || s == "Blue" },
	}).Parse(header + templateFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	outfile, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer outfile.Close()

	err = tmpl.Execute(outfile, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
