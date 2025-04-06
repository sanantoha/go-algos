package main

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	genericFourSum(t, [][4]int{{7, 6, -1, 4}, {7, 6, 2, 1}}, fourSum)
}

func TestFourSum1(t *testing.T) {
	genericFourSum(t, [][4]int{{4, -1, 6, 7}, {1, 2, 6, 7}}, fourSum1)
}

func genericFourSum(t *testing.T, expRes [][4]int, fn func([]int, int) [][4]int) {
	testTable := []struct {
		name   string
		input  []int
		target int
		expRes [][4]int
	}{
		{
			name:   "happy path",
			input:  []int{7, 6, 4, -1, 1, 2},
			target: 16,
			expRes: expRes,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := fn(subtest.input, subtest.target)

			if !reflect.DeepEqual(subtest.expRes, res) {
				t.Errorf("expected (%v), got (%v)", subtest.expRes, res)
			}
		})
	}
}
