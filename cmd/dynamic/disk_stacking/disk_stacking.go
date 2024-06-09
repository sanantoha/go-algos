package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

// O(n ^ 2) time | O(n) space
func diskStacking(disks [][]int) [][]int {
	if disks == nil || len(disks) == 0 {
		return nil
	}

	sort.Slice(disks, func(i, j int) bool {
		return disks[i][2] < disks[j][2]
	})

	heights := make([]int, len(disks))
	prev := make([]int, len(disks))
	for i, d := range disks {
		heights[i] = d[2]
		prev[i] = -1
	}

	maxIdx := 0

	for i := 1; i < len(disks); i++ {
		for j := 0; j < i; j++ {
			if less(disks[j], disks[i]) && heights[i] <= heights[j]+disks[i][2] {
				heights[i] = heights[j] + disks[i][2]
				prev[i] = j
			}
		}
		if heights[maxIdx] < heights[i] {
			maxIdx = i
		}
	}

	return buildSeq(disks, prev, maxIdx)
}

func buildSeq(disks [][]int, prev []int, maxIdx int) [][]int {

	idx := maxIdx
	res := make([][]int, 0)

	for idx >= 0 {
		res = append(res, disks[idx])
		idx = prev[idx]
	}

	slices.Reverse(res)
	return res
}

func less(d1, d2 []int) bool {
	return d1[0] < d2[0] && d1[1] < d2[1] && d1[2] < d2[2]
}

func main() {

	input := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{2, 2, 8},
		{2, 3, 4},
		{2, 2, 1},
		{4, 4, 5},
	}

	expected := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{4, 4, 5},
	}

	actual := diskStacking(input)
	fmt.Println(actual)

	fmt.Println(reflect.DeepEqual(actual, expected))
}
