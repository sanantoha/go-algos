package main

import (
	"testing"
)

func TestSortKSortedArray(t *testing.T) {

	arr := []int{3, 2, 1, 5, 4, 7, 6, 5}

	res := sortKSortedArray(arr, 3)

	for i := 1; i < len(res); i++ {
		if res[i-1] > res[i] {
			t.Errorf("array is not sorted %v > %v", res[i-1], res[i])
		}
	}
}
