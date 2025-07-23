package main

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		exp  float64
	}{
		{
			name: "pow(4, 2)",
			x:    4,
			n:    2,
			exp:  16,
		},
		{
			name: "pow(2, 4)",
			x:    2,
			n:    4,
			exp:  16,
		},
		{
			name: "pow(2, -2)",
			x:    2,
			n:    -2,
			exp:  0.25,
		},
		{
			name: "pow(2.1, 3)",
			x:    2.1,
			n:    3,
			exp:  9.261000000000001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pow(tt.x, tt.n)

			if !safeCompare(tt.exp, res) {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}

func safeCompare(a, b float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) {
		return false // NaN is not comparable
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return a == b // Compare infinities directly
	}
	return math.Abs(a-b) < 1e-9
}
