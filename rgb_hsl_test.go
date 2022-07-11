package main

import (
	"math"
	"testing"
)

func deepCloseHSL(c1, c2 hsl) bool {
	dh := math.Abs(c1.h - c2.h)
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

func Test_rgb2hsl(t *testing.T) {
	tests := []struct {
		name string
		c    rgb
		want hsl
	}{
		{"test 1", rgb{187, 119, 177}, hsl{309, 34, 60}},
		{"test 2", rgb{38, 178, 121}, hsl{156, 65, 42}},
		{"test 3", rgb{171, 34, 23}, hsl{4, 76, 38}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb2hsl(tt.c); !deepCloseHSL(got, tt.want) {
				t.Errorf("rgb2hsl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hsl2rgb(t *testing.T) {
	tests := []struct {
		name string
		c    hsl
		want rgb
	}{
		{"test 1", hsl{309, 34, 60}, rgb{187, 119, 177}},
		{"test 2", hsl{156, 65, 42}, rgb{38, 178, 121}},
		{"test 3", hsl{4, 76, 38}, rgb{171, 34, 23}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hsl2rgb(tt.c); !deepCloseRGB(got, tt.want) {
				t.Errorf("hsl2rgb() = %v, want %v", got, tt.want)
			}
		})
	}
}
