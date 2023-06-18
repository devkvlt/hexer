package hexer

import (
	"testing"
)

func Test_min(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"test 1", []float64{1, 2, 3}, float64(1)},
		{"test 2", []float64{4, -2, 0}, float64(-2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"test 1", []float64{1, 2, 3}, float64(3)},
		{"test 2", []float64{4, -2, 0}, float64(4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
