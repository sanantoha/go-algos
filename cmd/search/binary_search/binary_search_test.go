package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		exp    int
	}{
		{
			name:   "happy path",
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 80,
			exp:    7,
		},
		{
			name:   "happy path max value",
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 100,
			exp:    9,
		},
		{
			name:   "happy path min value",
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 10,
			exp:    0,
		},
		{
			name:   "target not found 32",
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 32,
			exp:    -4,
		},
		{
			name:   "target not found 101",
			input:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 101,
			exp:    -11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := binarySearch(tt.input, tt.target)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
