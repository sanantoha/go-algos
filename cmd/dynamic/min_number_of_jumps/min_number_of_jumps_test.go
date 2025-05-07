package main

import "testing"

var input = []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}

var exp = 4

func TestMinNumberOfJumps(t *testing.T) {

	res := minNumberOfJumps(input)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestMinNumberOfJumps1(t *testing.T) {

	res := minNumberOfJumps1(input)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
