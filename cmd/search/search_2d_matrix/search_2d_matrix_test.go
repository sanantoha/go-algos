package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]int
		key    int
		exp    bool
	}{
		{
			name: "happy path 1",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			key: 21,
			exp: true,
		},
		{
			name: "happy path 2",
			matrix: [][]int{
				{1, 4},
				{2, 5},
			},
			key: 5,
			exp: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := searchMatrix(tt.matrix, tt.key)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
