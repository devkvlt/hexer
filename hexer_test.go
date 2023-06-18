package hexer

import (
	"math"
	"strings"
	"testing"
)

var getTestCases = []struct {
	name    string
	hex     string
	r       float64
	g       float64
	b       float64
	h       float64
	s       float64
	l       float64
	wantErr bool
}{
	{"test 1", "#C966B2", 201, 102, 178, 314, 48, 59, false},
	{"test 2", "#ACE1A8", 172, 225, 168, 116, 49, 77, false},
	{"test 3: bad input", "#ACE8", 0, 0, 0, 0, 0, 0, true},
}

func Test_Red(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Red(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Red() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.r) > 0.5 {
				t.Errorf("Red() = %v, want %v", got, tt.r)
			}
		})
	}
}

func Test_Green(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Green(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Green() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.g) > 0.5 {
				t.Errorf("Green() = %v, want %v", got, tt.g)
			}
		})
	}
}

func Test_Blue(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Blue(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Blue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.b) > 0.5 {
				t.Errorf("Blue() = %v, want %v", got, tt.b)
			}
		})
	}
}

func Test_getHue(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hue(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.h) > 1 {
				t.Errorf("getHue() = %v, want %v", got, tt.h)
			}
		})
	}
}

func Test_Saturation(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Saturation(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Saturation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.s) > 1 {
				t.Errorf("Saturation() = %v, want %v", got, tt.s)
			}
		})
	}
}

func Test_Lightness(t *testing.T) {
	for _, tt := range getTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lightness(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lightness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.l) > 1 {
				t.Errorf("Lightness() = %v, want %v", got, tt.l)
			}
		})
	}
}

func Test_setRed(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		r       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 45, "#2D5521", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: red out of range", "#BC5521", 300, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetRed(tt.hex, tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("setRed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setRed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setGreen(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		r       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 45, "#BC2D21", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: green out of range", "#BC5521", 300, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetGreen(tt.hex, tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("setGreen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setGreen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setBlue(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		r       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 45, "#BC552D", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: blue out of range", "#BC5521", 300, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetBlue(tt.hex, tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("setBlue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setHue(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		h       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 67, "#AABC21", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: hue out of range", "#BC5521", 400, "", true},
		{"test 4: hue out of range", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetHue(tt.hex, tt.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("setHue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setHue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setSaturation(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		h       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 45, "#9F5D3C", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: hue out of range", "#BC5521", 400, "", true},
		{"test 4: hue out of range", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetSaturation(tt.hex, tt.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("setSaturation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setSaturation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setLightness(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		h       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 23, "#642D12", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: hue out of range", "#BC5521", 400, "", true},
		{"test 4: hue out of range", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetLightness(tt.hex, tt.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("setLightness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setLightness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lighten(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		dl      float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 20, "#E38B5E", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: dl too big", "#BC5521", 100, "#FFFFFF", false},
		{"test 4: dl negative", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lighten(tt.hex, tt.dl)
			if (err != nil) != tt.wantErr {
				t.Errorf("lighten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("lighten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_darken(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		dl      float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 20, "#642D12", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: amount too big", "#BC5521", 100, "#000000", false},
		{"test 4: negative amount", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Darken(tt.hex, tt.dl)
			if (err != nil) != tt.wantErr {
				t.Errorf("darken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("darken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saturate(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		dl      float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 20, "#D04D0B", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: amount too big", "#BC5521", 100, "#DB4900", false},
		{"test 4: negative amount", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Saturate(tt.hex, tt.dl)
			if (err != nil) != tt.wantErr {
				t.Errorf("saturate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("saturate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsaturate(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		dl      float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 20, "#A45B37", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: amount too big", "#BC5521", 100, "#6E6E6E", false},
		{"test 4: negative amount", "#BC5521", -10, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Desaturate(tt.hex, tt.dl)
			if (err != nil) != tt.wantErr {
				t.Errorf("unsaturate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("unsaturate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mix(t *testing.T) {
	tests := []struct {
		name    string
		c1      string
		c2      string
		w       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#FF0000", "#FFFF00", 0.5, "#FF8000", false},
		{"test 2", "#FF000", "#FFFF00", 0.5, "", true},
		{"test 3", "#FF0000", "#FFFF00", 1.2, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mix(tt.c1, tt.c2, tt.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("mix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.EqualFold(got, tt.want) {
				t.Errorf("mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mixPalette(t *testing.T) {
	tests := []struct {
		name    string
		c1, c2  string
		n       int
		want    []string
		wantErr bool
	}{
		{"test 1", "#ACE1A8", "#E16BC3", 6, []string{"#ACE1A8", "#b7c9ad", "#c1b2b3", "#cc9ab8", "#d683be", "#E16BC3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MixPalette(tt.c1, tt.c2, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("mixPalette() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !samePalette(got, tt.want) {
				t.Errorf("mixPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_darkPalette(t *testing.T) {
	tests := []struct {
		name    string
		c       string
		n       int
		want    []string
		wantErr bool
	}{
		{"test 1", "#A92388", 5, []string{"#A92388", "#7F1A66", "#551144", "#2A0922", "#000000"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DarkPalette(tt.c, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("darkPalette() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !samePalette(got, tt.want) {
				t.Errorf("darkPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lightPalette(t *testing.T) {
	tests := []struct {
		name    string
		c       string
		n       int
		want    []string
		wantErr bool
	}{
		{"test 1", "#A92388", 4, []string{"#A92388", "#DC56BB", "#EEAADD", "#FFFFFF"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LightPalette(tt.c, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("lightPalette() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !samePalette(got, tt.want) {
				t.Errorf("lightPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invert(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    string
		wantErr bool
	}{
		{"test 1", "#0066cc", "#ff9933", false},
		{"test 2", "#b8cbde", "#473421", false},
		{"test 3", "#b71a92", "#48e56d", false},
		{"test 4", "#b71a", "", true},
		{"test 5", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Invert(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("invert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contrastRatio(t *testing.T) {
	tests := []struct {
		name    string
		c1, c2  string
		want    float64
		wantErr bool
	}{
		{"test 1", "#7B3131", "#ACC8E5", 5.22, false},
		{"test 2", "#B67272", "#ACC8E5", 2.15, false},
		{"test 3", "#C1A6A6", "#000000", 9.27, false},
		{"test 4", "#C1A6", "#000000", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ContrastRatio(tt.c1, tt.c2)
			if (err != nil) != tt.wantErr {
				t.Errorf("contrastRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.want) > 0.1 {
				t.Errorf("contrastRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
