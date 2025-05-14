package main

import "testing"

func TestNumberOfIsland(t *testing.T) {

	grid := [][]int{
		{0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	exp := 3

	res := numberOfIsland(grid)

	if exp != res {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
