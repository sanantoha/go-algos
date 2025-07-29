package main

import (
	"testing"
)

func TestMinimumPassesOfMatrix(t *testing.T) {
	matrix := [][]int{
		{0, -1, -3, 2, 0},
		{1, -2, -5, -1, -3},
		{3, 0, 0, -4, -1},
	}
	expected := 3

	actual := minimumPassesOfMatrix(matrix)

	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
