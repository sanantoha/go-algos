package main

import (
	"reflect"
	"testing"
)

func TestMaxSumIncreasingSubsequence(t *testing.T) {

	arr := []int{10, 70, 20, 30, 50, 11, 30}
	expected := [][]int{
		{110},
		{10, 20, 30, 50},
	}

	res := maxSumIncreasingSubsequence(arr)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, but got %v", expected, res)
	}
}
