package main

import (
	"math"
	"strings"
	"testing"
)

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
			got, err := mix(tt.c1, tt.c2, tt.w)
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

func sameHEX(c1, c2 string) bool {
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
		{"test 4: dl too small", "#BC5521", -100, "#000000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lighten(tt.hex, tt.dl)
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

func Test_setR(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		r       float64
		want    string
		wantErr bool
	}{
		{"test 1", "#BC5521", 45, "#2D5521", false},
		{"test 2: invalid color input", "#BC552", 20, "", true},
		{"test 3: red component out of range", "#BC5521", 300, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setR(tt.hex, tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("setR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameHEX(got, tt.want) {
				t.Errorf("setR() = %v, want %v", got, tt.want)
			}
		})
	}
}
