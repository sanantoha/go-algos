package main

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 5, 6, 7},
		{19, 20, 21, 22, 23, 8},
		{18, 29, 30, 31, 24, 9},
		{17, 28, 27, 26, 25, 10},
		{16, 15, 14, 13, 12, 11},
	}

	exp := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	res := spiral(matrix)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
