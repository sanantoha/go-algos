package main

import (
	"fmt"
	"sort"
)

// O(n * log(n)) time | O(1) space
func minimalHaviestSetA(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	sort.Ints(arr)

	l := 0
	r := len(arr) - 1

	leftSum := arr[l]
	rightSum := arr[r]

	for l < r {
		if leftSum < rightSum {
			l++
			leftSum += arr[l]
		} else {
			r--
			rightSum += arr[r]
		}
	}
	return arr[l:]
}

func main() {
	arr := []int{6, 4, 2, 3, 1, 5}

	fmt.Println(minimalHaviestSetA(arr))
}
