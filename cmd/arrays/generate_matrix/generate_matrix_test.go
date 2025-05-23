package main

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expResult [][]int
	}{
		{
			name:  "generate matrix 3x3",
			input: 3,
			expResult: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name:  "generate matrix 1x1",
			input: 1,
			expResult: [][]int{
				{1},
			},
		},
		{
			name:  "generate matrix 4x4",
			input: 4,
			expResult: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matrix := generateMatrix(test.input)

			if !reflect.DeepEqual(test.expResult, matrix) {
				t.Errorf("expected %v, but got %v", test.expResult, matrix)
			}
		})
	}
}
