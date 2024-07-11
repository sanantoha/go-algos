package main

import (
	"fmt"
	"slices"
)

// O(n * log(n)) time | O(n) space
func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		} else {
			return 0
		}
	})

	merged := make([][]int, 0)

	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func main() {

	intervals := [][]int{
		{1, 5},
		{3, 7},
		{4, 6},
		{6, 8},
	}
	fmt.Println(merge(intervals))

	intervals1 := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals1))

	intervals2 := [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(intervals2))
}
