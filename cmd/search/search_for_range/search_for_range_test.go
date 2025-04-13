package main

import (
	"reflect"
	"testing"
)

func TestSearchForRange(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		target int
		exp    []int
	}{
		{
			name:   "happy path",
			input:  []int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 10},
			target: 8,
			exp:    []int{3, 9},
		},
		{
			name:   "target not in input",
			input:  []int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 10},
			target: 6,
			exp:    []int{},
		},
		{
			name:   "one element only",
			input:  []int{1},
			target: 1,
			exp:    []int{0, 0},
		},
		{
			name:   "empty input slice",
			input:  []int{},
			target: 1,
			exp:    []int{},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := searchRange(subtest.input, subtest.target)

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, got %v", subtest.exp, res)
			}
		})
	}
}
