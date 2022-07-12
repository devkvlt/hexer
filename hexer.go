package main

import (
	"errors"
)

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

func setR(hex string, r float64) (string, error) {
	rgb, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	if r < 0 || r > 255 {
		return "", errors.New("red component out of range (0-255)")
	}
	rgb.r = float64(r)
	return rgb2hex(rgb), nil
}

func lighten(hex string, dl float64) (string, error) {
	rgb, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	hsl := rgb2hsl(rgb)
	l := hsl.l + dl
	if l > 100 {
		hsl.l = 100
	} else if l < 0 {
		hsl.l = 0
	} else {
		hsl.l = l
	}
	return rgb2hex(hsl2rgb(hsl)), nil
}
