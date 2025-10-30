package main

import (
	"reflect"
	"testing"
)

func TestZigZagTraverse(t *testing.T) {
	matrix := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}

	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	res := zigZagTraverse(matrix)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
