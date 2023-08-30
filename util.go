package hexer

import (
	"math"
)

func rDist(c1, c2 rgb) float64 { return math.Abs(c1.r - c2.r) }
func gDist(c1, c2 rgb) float64 { return math.Abs(c1.g - c2.g) }
func bDist(c1, c2 rgb) float64 { return math.Abs(c1.b - c2.b) }

func rgbDist(c1, c2 rgb) float64 {
	// normalize so that 100 becomes the max dist instead of 255
	rdn := rDist(c1, c2) * 100 / 255
	gdn := gDist(c1, c2) * 100 / 255
	bdn := bDist(c1, c2) * 100 / 255
	return math.Sqrt(rdn*rdn + gdn*gdn + bdn*bdn)
}

func sameRGB(c1, c2 rgb, eps float64) bool {
	return rgbDist(c1, c2) < eps
}

func hDist(c1, c2 hsl) float64 { return math.Abs(math.Mod(c1.h-c2.h, 360)) }
func sDist(c1, c2 hsl) float64 { return math.Abs(c1.s - c2.s) }
func lDist(c1, c2 hsl) float64 { return math.Abs(c1.l - c2.l) }

func hslDist(c1, c2 hsl) float64 {
	// normalize so that 100 becomes the max dist instead of 360
	hdn := hDist(c1, c2) * 100 / 360
	sdn := sDist(c1, c2)
	ldn := lDist(c1, c2)
	return math.Sqrt(hdn*hdn + sdn*sdn + ldn*ldn)
}

func sameHSL(c1, c2 hsl, eps float64) bool {
	return hslDist(c1, c2) < eps
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
	return sameRGB(rgb1, rgb2, 1)
}

func samePalette(p1, p2 []string) bool {
	l1, l2 := len(p1), len(p2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if !sameHEX(p1[i], p2[i]) {
			return false
		}
	}
	return true
}
