package main

import (
	"reflect"
	"strings"
	"testing"
)

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
		{"test 5", "FF6478", rgb{0, 0, 0}, true},
		{"test 6", "", rgb{0, 0, 0}, true},
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
