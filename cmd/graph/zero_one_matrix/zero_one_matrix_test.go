package main

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		exp   [][]int
	}{
		{
			name: "case1",
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "case2",
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := updateMatrix(tt.input)

			if !reflect.DeepEqual(tt.exp, res) {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
