package hexer

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type rgb struct{ r, g, b float64 }

type hsl struct{ h, s, l float64 }

func rgb2hex(c rgb) string {
	r := int(math.Round(c.r))
	g := int(math.Round(c.g))
	b := int(math.Round(c.b))
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

func hex2rgb(hex string) (rgb, error) {
	err := errors.New("invalid hex color")
	webc, ok := WebColors[strings.ToLower(hex)]
	if ok {
		hex = webc
	}
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return rgb{}, err
	}
	r, errR := strconv.ParseUint(hex[:2], 16, 64)
	g, errG := strconv.ParseUint(hex[2:4], 16, 64)
	b, errB := strconv.ParseUint(hex[4:], 16, 64)
	if errR != nil || errG != nil || errB != nil {
		return rgb{}, err
	}
	return rgb{float64(r), float64(g), float64(b)}, nil
}

func hsl2hex(c hsl) string {
	return rgb2hex(hsl2rgb(c))
}

func hex2hsl(hex string) (hsl, error) {
	rgb, err := hex2rgb(hex)
	return rgb2hsl(rgb), err
}

// https://en.wikipedia.org/wiki/HSL_and_HSV#General_approach
func rgb2hsl(c rgb) hsl {
	r := c.r / 255
	g := c.g / 255
	b := c.b / 255
	max := max(r, g, b)
	min := min(r, g, b)
	chroma := max - min
	var h, s, l float64
	if chroma == 0 {
		h = 0 // by convention
	} else if max == r {
		// +6 to make sure the angle comes out positive
		h = math.Mod((g-b)/chroma+6, 6)
	} else if max == g {
		h = (b-r)/chroma + 2
	} else {
		h = (r-g)/chroma + 4
	}
	l = (max + min) / 2
	if l == 1 || l == 0 {
		s = 0
	} else {
		s = chroma / (1 - math.Abs(2*l-1))
	}
	h = 60 * h
	s = 100 * s
	l = 100 * l
	return hsl{h, s, l}
}

// https://en.wikipedia.org/wiki/HSL_and_HSV#HSL_to_RGB_alternative
func hsl2rgb(c hsl) rgb {
	h := c.h
	s := c.s / 100
	l := c.l / 100
	f := func(n float64) float64 {
		k := math.Mod(n+h/30, 12)
		a := s * min(l, 1-l)
		return l - a*max(-1, min(k-3, 9-k, 1))
	}
	r := 255 * f(0)
	g := 255 * f(8)
	b := 255 * f(4)
	return rgb{r, g, b}
}
