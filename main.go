package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 4 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	flag := args[0]
	if flag != "-m" {
		fmt.Println("bad flag")
		os.Exit(1)
	}
	c1 := args[1]
	c2 := args[2]
	w, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Println("bad weight")
		os.Exit(1)
	}
	m, err := mix(c1, c2, w)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(m)
}
