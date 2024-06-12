package main

import (
	"fmt"
	"math"
)

// O(n) time | O(1) space
func subarraySort(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{-1, -1}
	}

	minVal := math.MaxInt
	maxVal := math.MinInt

	for i, v := range arr {
		if outOfOrder(arr, i, v) {
			minVal = min(minVal, v)
			maxVal = max(maxVal, v)
		}
	}

	if minVal == math.MaxInt {
		return []int{-1, -1}
	}

	l := 0
	for l < len(arr) && arr[l] <= minVal {
		l++
	}

	r := len(arr) - 1
	for r >= 0 && arr[r] >= maxVal {
		r--
	}
	return []int{l, r}
}

func outOfOrder(arr []int, idx int, val int) bool {
	if len(arr) == 1 {
		return false
	}

	if idx == 0 {
		return val > arr[idx+1]
	} else if idx == len(arr)-1 {
		return arr[idx-1] > val
	} else {
		return arr[idx-1] > val || val > arr[idx+1]
	}
}

func main() {
	output := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	fmt.Println(subarraySort(output))

	output1 := []int{1, 2, 4, 7, 10, 11, 7, 12, 7, 7, 16, 18, 19}
	fmt.Println(subarraySort(output1))

	output2 := []int{1, 2}
	fmt.Println(subarraySort(output2))
}
