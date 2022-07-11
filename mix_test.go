package main

import (
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
