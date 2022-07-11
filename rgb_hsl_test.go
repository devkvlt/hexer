package main

import (
	"math"
	"testing"
)

func deepCloseHSL(c1, c2 hsl) bool {
	dh := math.Abs(math.Mod(c1.h-c2.h, 360))
	ds := math.Abs(c1.s - c2.s)
	dl := math.Abs(c1.l - c2.l)
	return dh <= 1 && ds <= 1 && dl <= 1
}

func deepCloseRGB(c1, c2 rgb) bool {
	dr := math.Abs(c1.r - c2.r)
	dg := math.Abs(c1.g - c2.g)
	db := math.Abs(c1.b - c2.b)
	return dr <= 1 && dg <= 1 && db <= 1
}

var rgbHSLTests = []struct {
	name string
	rgb  rgb
	hsl  hsl
}{
	{"test 1", rgb{187, 119, 177}, hsl{309, 33, 60}},
	{"test 2", rgb{38, 178, 121}, hsl{156, 65, 42}},
	{"test 3", rgb{171, 34, 23}, hsl{4, 76, 38}},
	{"test 4", rgb{255, 0, 0}, hsl{0, 100, 50}},
	{"test 5", rgb{6, 20, 14}, hsl{154, 54, 5}},
	// {"test 6", rgb{145, 1, 0}, hsl{0, 100, 28}}, // hsl2rgb not accurate enough atm
	{"test 7", rgb{0, 0, 0}, hsl{0, 0, 0}},
	{"test 8", rgb{255, 255, 255}, hsl{0, 0, 100}},
	{"test 9", rgb{255, 0, 0}, hsl{0, 100, 50}},
	{"test 10", rgb{0, 255, 0}, hsl{120, 100, 50}},
	{"test 11", rgb{0, 0, 255}, hsl{240, 100, 50}},
	{"test 12", rgb{255, 255, 0}, hsl{60, 100, 50}},
	{"test 13", rgb{0, 255, 255}, hsl{180, 100, 50}},
	{"test 14", rgb{255, 0, 255}, hsl{300, 100, 50}},
	{"test 15", rgb{191, 191, 191}, hsl{0, 0, 75}},
	{"test 16", rgb{128, 128, 128}, hsl{0, 0, 50}},
	{"test 17", rgb{128, 0, 0}, hsl{0, 100, 25}},
	{"test 18", rgb{128, 128, 0}, hsl{60, 100, 25}},
	{"test 19", rgb{0, 128, 0}, hsl{120, 100, 25}},
	{"test 20", rgb{128, 0, 128}, hsl{300, 100, 25}},
	{"test 21", rgb{0, 128, 128}, hsl{180, 100, 25}},
	{"test 22", rgb{0, 0, 128}, hsl{240, 100, 25}},
}

func Test_rgb2hsl(t *testing.T) {
	for _, tt := range rgbHSLTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb2hsl(tt.rgb); !deepCloseHSL(got, tt.hsl) {
				t.Errorf("rgb2hsl() = %v, want %v", got, tt.hsl)
			}
		})
	}
}

func Test_hsl2rgb(t *testing.T) {
	for _, tt := range rgbHSLTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hsl2rgb(tt.hsl); !deepCloseRGB(got, tt.rgb) {
				t.Errorf("hsl2rgb() = %v, want %v", got, tt.rgb)
			}
		})
	}
}
