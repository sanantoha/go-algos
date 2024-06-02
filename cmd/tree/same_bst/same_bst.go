package main

import (
	"fmt"
	"math"
)

// O(n ^ 2) time | O(n ^ 2) space
func sameBsts(arr1, arr2 []int) bool {
	if (arr1 == nil || len(arr1) == 0) && (arr2 == nil || len(arr2) == 0) {
		return true
	}

	if arr1 == nil || len(arr1) == 0 || arr2 == nil || len(arr2) == 0 {
		return false
	}

	if arr1[0] != arr2[0] {
		return false
	}

	li := partition(arr1)
	ri := partition(arr2)

	return sameBsts(li.less, ri.less) && sameBsts(li.great, ri.great)
}

type Info struct {
	less  []int
	great []int
}

func partition(arr []int) Info {
	less := make([]int, 0)
	great := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v < arr[0] {
			less = append(less, v)
		} else {
			great = append(great, v)
		}
	}

	return Info{less: less, great: great}
}

// O(n ^ 2) time | O(d) space
func sameBsts1(arr1, arr2 []int) bool {
	return helper(arr1, arr2, 0, 0, math.MinInt, math.MaxInt)
}

func helper(arr1 []int, arr2 []int, idx1 int, idx2 int, minVal int, maxVal int) bool {
	if idx1 == -1 || idx2 == -1 {
		return idx1 == idx2
	}
	if arr1[idx1] != arr2[idx2] {
		return false
	}

	lessIdx1 := lessIdx(arr1, idx1, minVal)
	lessIdx2 := lessIdx(arr2, idx2, minVal)
	greatIdx1 := greaterOrEqualIdx(arr1, idx1, maxVal)
	greatIdx2 := greaterOrEqualIdx(arr2, idx2, maxVal)
	v := arr1[idx1]
	return helper(arr1, arr2, lessIdx1, lessIdx2, minVal, v) && helper(arr1, arr2, greatIdx1, greatIdx2, v, maxVal)
}

func lessIdx(arr []int, idx int, minVal int) int {
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] < arr[idx] && arr[i] >= minVal {
			return i
		}
	}
	return -1
}

func greaterOrEqualIdx(arr []int, idx int, maxVal int) int {
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] >= arr[idx] && arr[i] < maxVal {
			return i
		}
	}
	return -1
}

func main() {

	arr1 := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	arr2 := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}

	fmt.Println(sameBsts(arr1, arr2))
	fmt.Println(sameBsts1(arr1, arr2))
}
