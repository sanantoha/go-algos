package main

import (
	"testing"
)

func TestUniquePathIII(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}

	exp := 2

	res := uniquePath3(grid)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
