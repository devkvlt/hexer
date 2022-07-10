package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type rgb struct {
	r, g, b float64
}

type hsl struct {
	h, s, l float64
}

func rgb2hex(c rgb) string {
	r := int(math.Round(c.r))
	g := int(math.Round(c.g))
	b := int(math.Round(c.b))
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

func hex2rgb(hex string) (rgb, error) {
	if hex[:1] == "#" {
		hex = hex[1:]
	}
	err := errors.New("invalid hex color")
	if len(hex) != 6 {
		return rgb{}, err
	}
	r, errR := strconv.ParseUint(hex[:2], 16, 64)
	g, errG := strconv.ParseUint(hex[2:4], 16, 64)
	b, errB := strconv.ParseUint(hex[4:6], 16, 64)
	if len(hex) != 6 || errR != nil || errG != nil || errB != nil {
		return rgb{}, err
	}
	return rgb{float64(r), float64(g), float64(b)}, nil
}

// https://en.wikipedia.org/wiki/HSL_and_HSV#General_approach
func rgb2hsl(c rgb) hsl {
	r := c.r / 255
	g := c.g / 255
	b := c.b / 255
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	chroma := max - min
	var h, s, l float64
	if chroma == 0 {
		h = 0 // by convention
	} else if max == r {
		h = float64(((int((g - b) / chroma)) % 6) * 60)
	} else if max == g {
		h = ((b-r)/chroma + 2) * 60
	} else {
		h = ((r-g)/chroma + 4) * 60
	}
	l = (max + min) / 2
	if l == 1 || l == 0 {
		s = 0
	} else {
		s = chroma / (1 - math.Abs(2*l-1))
	}
	return hsl{h, s, l}
}

// https://en.wikipedia.org/wiki/HSL_and_HSV#HSL_to_RGB_alternative
func hsl2rgb(c hsl) rgb {
	h := c.h
	s := c.s
	l := c.l
	f := func(n float64) float64 {
		k := int(n+h/30) % 12
		a := s * math.Min(l, 1-l)
		return l - a*math.Max(math.Min(float64(k-3), math.Min(float64(9-k), 1)), -1)
	}
	r := math.Round(255 * f(0))
	g := math.Round(255 * f(8))
	b := math.Round(255 * f(4))
	return rgb{r, g, b}
}

func mix(c1, c2 string, w float64) (string, error) {
	rgb1, err1 := hex2rgb(c1)
	rgb2, err2 := hex2rgb(c2)
	if err1 != nil || err2 != nil {
		return "", errors.New("invalid hex color")
	}
	if w < 0 || w > 1 {
		return "", errors.New("invalid weight")
	}
	rgbx := rgb{
		r: w*rgb1.r + (1-w)*rgb2.r,
		g: w*rgb1.g + (1-w)*rgb2.g,
		b: w*rgb1.b + (1-w)*rgb2.b,
	}
	return rgb2hex(rgbx), nil
}

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
