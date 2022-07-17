// This source file contains the core functions behind hexer. All input and
// output colors are in the hex format.

package main

import (
	"errors"
)

// getRed returns the red channel of a color and an error.
func getRed(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.r, err
}

// getGreen returns the green channel of a color and an error.
func getGreen(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.g, err
}

// getBlue returns the blue channel of a color and an error.
func getBlue(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.b, err
}

// getHue returns the hue channel of a color and an error.
func getHue(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.h, err
}

// getSaturation returns the saturation channel of a color and an error.
func getSaturation(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.s, err
}

// getLightness returns the lightness channel of a color and an error.
func getLightness(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.l, err
}

// setRed sets the red channel of a color to the provided value and returns the
// result and an error.
func setRed(hex string, r float64) (string, error) {
	rgb, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	if r < 0 || r > 255 {
		return "", errors.New("red out of range (0-255)")
	}
	rgb.r = r
	return rgb2hex(rgb), nil
}

// setGreen sets the green channel of a color to the provided value and returns
// the result and an error.
func setGreen(hex string, g float64) (string, error) {
	rgb, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	if g < 0 || g > 255 {
		return "", errors.New("green out of range (0-255)")
	}
	rgb.g = g
	return rgb2hex(rgb), nil
}

// setBlue sets the blue channel of a color to the provided value and returns
// the result and an error.
func setBlue(hex string, b float64) (string, error) {
	rgb, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	if b < 0 || b > 255 {
		return "", errors.New("blue out of range (0-255)")
	}
	rgb.b = b
	return rgb2hex(rgb), nil
}

// setHue sets the hue channel of a color to the provided value and returns the
// result and an error.
func setHue(hex string, h float64) (string, error) {
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	if h < 0 || h > 360 {
		return "", errors.New("hue out of range (0-360)")
	}
	hsl.h = h
	return hsl2hex(hsl), nil
}

// setSaturation sets the saturation channel of a color to the provided value
// and returns the result and an error.
func setSaturation(hex string, s float64) (string, error) {
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	if s < 0 || s > 100 {
		return "", errors.New("saturation out of range (0-100)")
	}
	hsl.s = s
	return hsl2hex(hsl), nil
}

// setLightness sets the lightness channel of a color to the provided value and
// returns the result and an error.
func setLightness(hex string, l float64) (string, error) {
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	if l < 0 || l > 100 {
		return "", errors.New("lightness out of range (0-100)")
	}
	hsl.l = l
	return hsl2hex(hsl), nil
}

// lighten lightens a color by adding the provided value to its lightness
// channel and returns the result and an error.
func lighten(hex string, a float64) (string, error) {
	if a < 0 {
		return "", errors.New("amount must be >= 0")
	}
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	l := hsl.l + a
	hsl.l = min(l, 100)
	return hsl2hex(hsl), nil
}

// darken darkens a color by subtracting the provided value from its lightness
// channel and returns the result and an error.
func darken(hex string, a float64) (string, error) {
	if a < 0 {
		return "", errors.New("amount must be >= 0")
	}
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	l := hsl.l - a
	hsl.l = max(l, 0)
	return hsl2hex(hsl), nil
}

// saturate saturates a color by adding the provided value to its saturation
// channel and returns the result and an error.
func saturate(hex string, a float64) (string, error) {
	if a < 0 {
		return "", errors.New("amount must be >= 0")
	}
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	s := hsl.s + a
	hsl.s = min(s, 100)
	return hsl2hex(hsl), nil
}

// unsaturate unsaturates a color by subtracting the provided value from its
// saturation channel and returns the result and an error.
func unsaturate(hex string, a float64) (string, error) {
	if a < 0 {
		return "", errors.New("amount must be >= 0")
	}
	hsl, err := hex2hsl(hex)
	if err != nil {
		return "", err
	}
	s := hsl.s - a
	hsl.s = max(s, 0)
	return hsl2hex(hsl), nil
}

// mix mixes two colors using the given weight (of the first color) and returns
// the result and an error.
func mix(c1, c2 string, w float64) (string, error) {
	rgb1, err := hex2rgb(c1)
	if err != nil {
		return "", err
	}
	rgb2, err := hex2rgb(c2)
	if err != nil {
		return "", err
	}
	if w < 0 || w > 1 {
		return "", errors.New("weight must be between 0 and 1")
	}
	rgbx := rgb{
		w*rgb1.r + (1-w)*rgb2.r,
		w*rgb1.g + (1-w)*rgb2.g,
		w*rgb1.b + (1-w)*rgb2.b,
	}
	return rgb2hex(rgbx), nil
}

// makeMixPalette makes a slice of colors representing a palette of n equally
// distanced colors between c1 and c2 (c1 and c2 included) and returns it with
// an error.
func makeMixPalette(c1, c2 string, n int) ([]string, error) {
	if n < 2 {
		return []string{}, errors.New("n must be >= 2")
	}
	palette := make([]string, n)
	for i := 0; i < n; i++ {
		w := float64(i) / float64(n-1)
		mix, err := mix(c1, c2, w)
		if err != nil {
			return []string{}, err
		}
		palette[i] = mix
	}
	return palette, nil
}

// makeDarkPalette
func makeDarkPalette(c string, n int) ([]string, error) {
	if n < 2 {
		return []string{}, errors.New("n must be >= 2")
	}
	hsl, err := hex2hsl(c)
	if err != nil {
		return []string{}, err
	}
	l0 := hsl.l
	palette := make([]string, n)
	for i := 0; i < n; i++ {
		l := l0 - float64(i)*l0/float64(n-1)
		hsl.l = l
		palette[i] = hsl2hex(hsl)
	}
	return palette, nil
}

// makeLightPalette
func makeLightPalette(c string, n int) ([]string, error) {
	if n < 2 {
		return []string{}, errors.New("n must be >= 2")
	}
	hsl, err := hex2hsl(c)
	if err != nil {
		return []string{}, err
	}
	l0 := hsl.l
	palette := make([]string, n)
	for i := 0; i < n; i++ {
		l := l0 + float64(i)*(100-l0)/float64(n-1)
		hsl.l = l
		palette[i] = hsl2hex(hsl)
	}
	return palette, nil
}
