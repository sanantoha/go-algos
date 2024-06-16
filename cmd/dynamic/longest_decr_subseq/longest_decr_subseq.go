package main

import (
	"fmt"
	"math"
	"slices"
)

// O(n ^ 2) time | O(n) space
func lds(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	lds := make([]int, len(arr))
	for i, _ := range arr {
		lds[i] = 1
	}

	maxVal := math.MinInt

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] && lds[i] < lds[j]+1 {
				lds[i] = lds[j] + 1
			}
		}

		if maxVal < lds[i] {
			maxVal = lds[i]
		}
	}
	return maxVal
}

func lds1(arr []int) int {
	return -1
}

// O(n ^ 2) time | O(n) space
func ldsList(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	lds := make([]int, len(arr))
	prev := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		lds[i] = 1
		prev[i] = -1
	}

	maxIdx := 0

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] && lds[i] < lds[j]+1 {
				lds[i] = lds[j] + 1
				prev[i] = j
			}
		}
		if lds[maxIdx] < lds[i] {
			maxIdx = i
		}
	}
	return buildSeq(arr, prev, maxIdx)
}

func buildSeq(arr []int, prev []int, maxIdx int) []int {
	res := make([]int, 0)

	idx := maxIdx

	for idx >= 0 {
		res = append(res, arr[idx])
		idx = prev[idx]
	}

	slices.Reverse(res)
	return res
}

func ldsList1(arr []int) []int {
	return nil
}

func main() {

	arr := []int{5, 6, 7, 6, 5, 4, 3, 10, 14, 12} // 5

	fmt.Println(lds(arr))
	fmt.Println(lds1(arr))
	fmt.Println(ldsList(arr))
	fmt.Println(ldsList1(arr))

	arr1 := []int{100, 10, 9, 8, 7, 6, 5, 90, 80, 70, 60, 50, 40, 30, 20} // 9

	fmt.Println(lds(arr1))
	fmt.Println(lds1(arr1))
	fmt.Println(ldsList(arr1))
	fmt.Println(ldsList1(arr1))
}
