package main

import (
	"fmt"
	"testing"
)

func TestMstGraph(t *testing.T) {

	tests := []struct {
		name  string
		input [][]int
		exp   int
	}{
		{
			name: "case 1",
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			exp: 4,
		},
		{
			name: "case 2",
			input: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			exp: -1,
		},
		{
			name: "case 3",
			input: [][]int{
				{0, 2},
			},
			exp: 0,
		},
		{
			name: "case 4",
			input: [][]int{
				{0},
			},
			exp: 0,
		},
		{
			name: "case 5",
			input: [][]int{
				{1},
			},
			exp: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := rottingOranges(test.input)

			if res != test.exp {
				t.Errorf("expected %v, but got %v", test.exp, res)
			}
		})
	}

	fmt.Println("test rotting")
}
