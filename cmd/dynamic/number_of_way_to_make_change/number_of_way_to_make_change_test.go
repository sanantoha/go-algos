package main

import "testing"

func TestNumberOfWaysToMakeChange(t *testing.T) {
	n := 6
	input := []int{1, 5}
	exp := 2

	res := numberOfWaysToMakeChange(n, input)

	if res != exp {
		t.Errorf("expected %d, got %d", exp, res)
	}
}
