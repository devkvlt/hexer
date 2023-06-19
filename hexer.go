// This source file contains the core functions behind hexer. All input and
// output colors are in the hex format.

package hexer

import (
	"errors"
	"math"
)

// Luminance returns the relative Luminance of a color.
// https://en.wikipedia.org/wiki/Relative_luminance
func Luminance(c string) (float64, error) {
	rgb, err := hex2rgb(c)
	if err != nil {
		return 0, err
	}
	vr := rgb.r / 255
	vg := rgb.g / 255
	vb := rgb.b / 255
	f := func(ch float64) float64 {
		if ch <= 0.04045 {
			return ch / 12.92
		}
		return math.Pow((ch+0.055)/1.055, 2.4)
	}
	return 0.2126*f(vr) + 0.7152*f(vg) + 0.0722*f(vb), nil
}

// ContrastRatio returns the CR of 2 colors.
func ContrastRatio(c1, c2 string) (float64, error) {
	l1, err := Luminance(c1)
	if err != nil {
		return 0, err
	}
	l2, err := Luminance(c2)
	if err != nil {
		return 0, err
	}
	if l1 > l2 {
		return (l1 + 0.05) / (l2 + 0.05), nil
	}
	return (l2 + 0.05) / (l1 + 0.05), nil
}

// Invert returns the inverse of a color and an error.
func Invert(hex string) (string, error) {
	c, err := hex2rgb(hex)
	if err != nil {
		return "", err
	}
	c.r = 255 - c.r
	c.g = 255 - c.g
	c.b = 255 - c.b
	return rgb2hex(c), nil
}

// Red returns the red channel of a color and an error.
func Red(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.r, err
}

// Green returns the green channel of a color and an error.
func Green(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.g, err
}

// Blue returns the blue channel of a color and an error.
func Blue(hex string) (float64, error) {
	rgb, err := hex2rgb(hex)
	return rgb.b, err
}

// Hue returns the hue channel of a color and an error.
func Hue(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.h, err
}

// Saturation returns the saturation channel of a color and an error.
func Saturation(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.s, err
}

// Lightness returns the lightness channel of a color and an error.
func Lightness(hex string) (float64, error) {
	hsl, err := hex2hsl(hex)
	return hsl.l, err
}

// SetRed sets the red channel of a color to the provided value and returns the
// result and an error.
func SetRed(hex string, r float64) (string, error) {
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

// SetGreen sets the green channel of a color to the provided value and returns
// the result and an error.
func SetGreen(hex string, g float64) (string, error) {
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

// SetBlue sets the blue channel of a color to the provided value and returns
// the result and an error.
func SetBlue(hex string, b float64) (string, error) {
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

// SetHue sets the hue channel of a color to the provided value and returns the
// result and an error.
func SetHue(hex string, h float64) (string, error) {
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

// SetSaturation sets the saturation channel of a color to the provided value
// and returns the result and an error.
func SetSaturation(hex string, s float64) (string, error) {
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

// SetLightness sets the lightness channel of a color to the provided value and
// returns the result and an error.
func SetLightness(hex string, l float64) (string, error) {
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

// Lighten lightens a color by adding the provided value to its lightness
// channel and returns the result and an error.
func Lighten(hex string, a float64) (string, error) {
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

// Darken darkens a color by subtracting the provided value from its lightness
// channel and returns the result and an error.
func Darken(hex string, a float64) (string, error) {
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

// Saturate saturates a color by adding the provided value to its saturation
// channel and returns the result and an error.
func Saturate(hex string, a float64) (string, error) {
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

// Desaturate unsaturates a color by subtracting the provided value from its
// saturation channel and returns the result and an error.
func Desaturate(hex string, a float64) (string, error) {
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

// Mix mixes two colors using the given weight (of the first color) and returns
// the result and an error.
func Mix(c1, c2 string, w float64) (string, error) {
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

// MixPalette makes a slice of colors representing a palette of n equally
// distanced colors between c1 and c2 (c1 and c2 included) and returns it with
// an error.
func MixPalette(c1, c2 string, n int) ([]string, error) {
	if n < 2 {
		return []string{}, errors.New("n must be >= 2")
	}
	palette := make([]string, n)
	for i := 0; i < n; i++ {
		w := float64(i) / float64(n-1)
		mix, err := Mix(c1, c2, w)
		if err != nil {
			return []string{}, err
		}
		palette[n-1-i] = mix
	}
	return palette, nil
}

// DarkPalette
func DarkPalette(c string, n int) ([]string, error) {
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

// LightPalette
func LightPalette(c string, n int) ([]string, error) {
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

// AdjustToContrastRatio adjusts the lightness of the color c1 so that the
// contrast ratio between c1 and c2 is as close as possible to the given
// contrast ratio cr0.
func AdjustToContrastRatio(c1, c2 string, cr0 float64) (string, error) {
	const tolerance = 0.1
	const adjustStep = 1

	var adjust func(string, float64) (string, error)

	if cr0 < 1 || cr0 > 21 {
		return "", errors.New("invalid contrast ratio")
	}

	l1, err := Lightness(c1)
	if err != nil {
		return "", err
	}

	l2, err := Lightness(c2)
	if err != nil {
		return "", err
	}

	cr, _ := ContrastRatio(c1, c2)

	if (l1 > l2 && cr > cr0) || (l1 <= l2 && cr <= cr0) {
		adjust = Darken
	} else {
		adjust = Lighten
	}

	for math.Abs(cr0-cr) > tolerance {
		// fmt.Println(c1, c2, cr) // DEBUG:
		c1, _ = adjust(c1, adjustStep)
		prevCR := cr
		cr, _ = ContrastRatio(c1, c2)
		if prevCR == cr {
			break
		}
	}

	return c1, nil
}
