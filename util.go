package main

import (
	"math"
)

func min(nums ...float64) float64 {
	m := nums[0]
	for _, v := range nums {
		if m > v {
			m = v
		}
	}
	return float64(m)
}

func max(nums ...float64) float64 {
	m := nums[0]
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return float64(m)
}

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

func sameHEX(c1, c2 string) bool {
	if c1 == c2 {
		return true
	}
	rgb1, err1 := hex2rgb(c1)
	rgb2, err2 := hex2rgb(c2)
	if err1 != nil || err2 != nil {
		return false
	}
	dr := math.Abs(rgb1.r - rgb2.r)
	dg := math.Abs(rgb1.g - rgb2.g)
	db := math.Abs(rgb1.b - rgb2.b)
	d := math.Sqrt(dr*dr + dg*dg + db*db)
	return d < 2.5
	// 2.5 is actually unnoticeable
	// TODO: add minD param to sameRGB func and reuse here

}
