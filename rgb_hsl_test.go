package main

import (
	"math"
	"testing"
)

// sameHSL checks if two hsl colors are close enough to be considered the same
// in tests. The metric used is a simple distance.
func sameHSL(c1, c2 hsl) bool {
	dh := math.Abs(math.Mod(c1.h-c2.h, 360))
	ds := math.Abs(c1.s - c2.s)
	dl := math.Abs(c1.l - c2.l)
	d := math.Sqrt(dh*dh + ds*ds + dl*dl)
	return d < 1
}

// sameRGB does the same as sameHSL but for rgb colors. Take a shot every time
// you read "same" 🤡.
func sameRGB(c1, c2 rgb) bool {
	dr := math.Abs(c1.r - c2.r)
	dg := math.Abs(c1.g - c2.g)
	db := math.Abs(c1.b - c2.b)
	d := math.Sqrt(dr*dr + dg*dg + db*db)
	return d < 1
}

func Test_rgb2hsl(t *testing.T) {
	tests := []struct {
		name string
		in   rgb
		out  hsl
	}{
		{"test 1", rgb{187, 119, 177}, hsl{309, 33, 60}},
		{"test 2", rgb{38, 178, 121}, hsl{156, 65, 42}},
		{"test 3", rgb{171, 34, 23}, hsl{4, 76, 38}},
		{"test 4", rgb{255, 0, 0}, hsl{0, 100, 50}},
		{"test 5", rgb{6, 20, 14}, hsl{154, 54, 5}},
		{"test 6", rgb{145, 1, 0}, hsl{0, 100, 28}},
		{"test 7", rgb{0, 0, 0}, hsl{0, 0, 0}},
		{"test 8", rgb{255, 255, 255}, hsl{0, 0, 100}},
		{"test 9", rgb{255, 0, 0}, hsl{0, 100, 50}},
		{"test 10", rgb{0, 255, 0}, hsl{120, 100, 50}},
		{"test 11", rgb{0, 0, 255}, hsl{240, 100, 50}},
		{"test 12", rgb{255, 255, 0}, hsl{60, 100, 50}},
		{"test 13", rgb{0, 255, 255}, hsl{180, 100, 50}},
		{"test 14", rgb{255, 0, 255}, hsl{300, 100, 50}},
		{"test 15", rgb{191, 191, 191}, hsl{0, 0, 75}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb2hsl(tt.in); !sameHSL(got, tt.out) {
				t.Errorf("rgb2hsl() = %v, want %v", got, tt.out)
			}
		})
	}
}

func Test_hsl2rgb(t *testing.T) {
	tests := []struct {
		name string
		in   hsl
		out  rgb
	}{
		{"test 1", hsl{309, 33, 60}, rgb{187, 119, 177}},
		{"test 2", hsl{156, 65, 42}, rgb{37, 177, 121}},
		{"test 3", hsl{4, 76, 38}, rgb{171, 33, 23}},
		{"test 4", hsl{0, 100, 50}, rgb{255, 0, 0}},
		{"test 5", hsl{154, 54, 5}, rgb{6, 20, 14}},
		{"test 6", hsl{0, 100, 28}, rgb{143, 0, 0}},
		{"test 7", hsl{0, 0, 0}, rgb{0, 0, 0}},
		{"test 8", hsl{0, 0, 100}, rgb{255, 255, 255}},
		{"test 9", hsl{0, 100, 50}, rgb{255, 0, 0}},
		{"test 10", hsl{120, 100, 50}, rgb{0, 255, 0}},
		{"test 11", hsl{240, 100, 50}, rgb{0, 0, 255}},
		{"test 12", hsl{60, 100, 50}, rgb{255, 255, 0}},
		{"test 13", hsl{180, 100, 50}, rgb{0, 255, 255}},
		{"test 14", hsl{300, 100, 50}, rgb{255, 0, 255}},
		{"test 15", hsl{0, 0, 75}, rgb{191, 191, 191}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hsl2rgb(tt.in); !sameRGB(got, tt.out) {
				t.Errorf("hsl2rgb() = %v, want %v", got, tt.out)
			}
		})
	}
}
