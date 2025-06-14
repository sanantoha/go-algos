package main

import (
	"strconv"
	"testing"
)

func TestWaterArea(t *testing.T) {

	input := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}
	exp := 48

	funcs := []func([]int) int{waterArea, waterArea1}

	for i, fn := range funcs {
		t.Run("waterArea"+strconv.Itoa(i), func(t *testing.T) {
			actual := fn(input)
			if actual != exp {
				t.Errorf("expected %v, but got %v", exp, actual)
			}
		})
	}
}
