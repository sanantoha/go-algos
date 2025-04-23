package main

import "testing"

func TestHouseRobber(t *testing.T) {
	nums := []int{4, 1, 2, 7, 5, 3, 1}
	exp := 14

	res := rob(nums)

	if exp != res {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
