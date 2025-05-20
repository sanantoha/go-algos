package main

import (
	"reflect"
	"testing"
)

var input = []int{12, 3, 1, 2, -6, 5, -8, 6}
var target = 0

func TestThreeSum(t *testing.T) {
	exp := [][3]int{
		{3, 5, -8},
		{1, -6, 5},
		{2, -8, 6},
	}
	res := threeNumberSum(input, target)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestThreeSum1(t *testing.T) {
	exp := [][3]int{
		{-8, 2, 6},
		{-8, 3, 5},
		{-6, 1, 5},
	}
	res := threeNumberSum1(input, target)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
