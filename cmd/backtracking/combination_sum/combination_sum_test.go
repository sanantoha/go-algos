package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	arr := []int{2, 3, 5, 7}
	exp := [][]int{
		{2, 2, 3},
		{2, 5},
		{7},
	}

	res := combinationSum(arr, 7)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but to %v", exp, res)
	}
}
