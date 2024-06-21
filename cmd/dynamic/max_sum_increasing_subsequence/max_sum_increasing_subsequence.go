package main

import (
	"fmt"
	"reflect"
	"slices"
)

// O(n ^ 2) time | O(n) space
func maxSumIncreasingSubsequence(arr []int) [][]int {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	msi := make([]int, len(arr))
	prev := make([]int, len(arr))

	for i, v := range arr {
		msi[i] = v
		prev[i] = -1
	}

	maxIdx := 0

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && msi[i] < msi[j]+arr[i] {
				msi[i] = msi[j] + arr[i]
				prev[i] = j
			}
		}
		if msi[maxIdx] < msi[i] {
			maxIdx = i
		}
	}
	return [][]int{
		{msi[maxIdx]},
		buildSeq(arr, prev, maxIdx),
	}
}

func buildSeq(arr []int, prev []int, maxIdx int) []int {
	idx := maxIdx
	res := make([]int, 0)

	for idx >= 0 {
		res = append(res, arr[idx])
		idx = prev[idx]
	}

	slices.Reverse(res)
	return res
}

func main() {

	arr := []int{10, 70, 20, 30, 50, 11, 30}
	expected := [][]int{
		{110},
		{10, 20, 30, 50},
	}

	actual := maxSumIncreasingSubsequence(arr)
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}
