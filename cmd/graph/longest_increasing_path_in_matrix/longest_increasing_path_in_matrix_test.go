package main

import (
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {

	tests := []struct {
		name  string
		input [][]int
		exp   int
	}{
		{
			name: "happy path 1",
			input: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			exp: 4,
		},
		{
			name: "happy path 2",
			input: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			exp: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestIncreasingPath(tt.input)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
