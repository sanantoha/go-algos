package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLargestIsland(t *testing.T) {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
	}

	expected := 8

	funcs := []func([][]int) int{
		largestIsland,
		largestIsland1,
	}

	for i, fn := range funcs {
		t.Run("largestIsland"+strconv.Itoa(i), func(t *testing.T) {
			actual := fn(matrix)
			fmt.Println(actual)

			if actual != expected {
				t.Errorf("expected %v, but got %v", expected, actual)
			}
		})
	}

}
