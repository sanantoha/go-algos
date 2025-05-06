package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testRotateFn(t, rotate)
}

func TestRotate1(t *testing.T) {
	testRotateFn(t, rotate1)
}

func testRotateFn(t *testing.T, fn func([]int, int)) {
	testTable := []struct {
		name  string
		input []int
		k     int
		exp   []int
	}{
		{
			name:  "hp1",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
			exp:   []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:  "hp2",
			input: []int{-1, -100, 3, 99},
			k:     2,
			exp:   []int{3, 99, -1, -100},
		},
		{
			name:  "hp3",
			input: []int{1, 2, 3},
			k:     2,
			exp:   []int{2, 3, 1},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			fn(subtest.input, subtest.k)

			if !reflect.DeepEqual(subtest.exp, subtest.input) {
				t.Errorf("expected %v, but got %v", subtest.exp, subtest.input)
			}
		})
	}
}
