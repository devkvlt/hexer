package hexer

import (
	"reflect"
	"strings"
	"testing"
)

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
			if got := rgb2hsl(tt.in); !sameHSL(got, tt.out, 0.5) {
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
			if got := hsl2rgb(tt.in); !sameRGB(got, tt.out, 0.5) {
				t.Errorf("hsl2rgb() = %v, want %v", got, tt.out)
			}
		})
	}
}

func Test_hex2rgb(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    rgb
		wantErr bool
	}{
		{"test 1", "#FF6478", rgb{255, 100, 120}, false},
		{"test 2", "#ff6478", rgb{255, 100, 120}, false},
		{"test 3", "#FF68", rgb{0, 0, 0}, true},
		{"test 4", "#FF684305", rgb{0, 0, 0}, true},
		{"test 5", "FF6478", rgb{255, 100, 120}, false},
		{"test 6", "#ZWF647", rgb{0, 0, 0}, true},
		{"test 7", "", rgb{0, 0, 0}, true},
		{"test 8", "Red", rgb{255, 0, 0}, false},
		{"test 9", "red", rgb{255, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hex2rgb(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hex2rgb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hex2rgb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rgb2hex(t *testing.T) {
	tests := []struct {
		name string
		c    rgb
		want string
	}{
		{"test 1", rgb{255, 100, 120}, "#FF6478"},
		{"test 2", rgb{255, 255, 255}, "#FFFFFF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb2hex(tt.c); !strings.EqualFold(got, tt.want) {
				t.Errorf("rgb2hex() = %v, want %v", got, tt.want)
			}
		})
	}
}
