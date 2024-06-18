package main

import (
	"fmt"
	"math"
)

// O(n ^ 2) time | O(n) space
func lnds(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	lnds := make([]int, len(arr))
	for i, _ := range arr {
		lnds[i] = 1
	}

	maxVal := math.MinInt

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] <= arr[i] && lnds[i] < lnds[j]+1 {
				lnds[i] = lnds[j] + 1
			}
		}
		if maxVal < lnds[i] {
			maxVal = lnds[i]
		}
	}
	return maxVal
}

// O(n * log(n)) time | O(n) space
func lnds1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	res := make([]int, 1)
	res[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prev := res[len(res)-1]
		if prev <= arr[i] {
			res = append(res, arr[i])
		} else {
			j := binarySearch(res, arr[i])
			if j < 0 {
				j = -(j + 1)
			}
			res[j] = arr[i]
		}
	}
	return len(res)
}

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := int(uint(l+r) >> 1)
		if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if arr[l] == target {
		return l
	} else {
		return -(l + 1)
	}
}

func lndsList(arr []int) []int {
	return nil
}

func lndsList1(arr []int) []int {
	return nil
}

func main() {

	arr := []int{-2, -1, 2, 3, 4, 5, 2, 2, 2, 2, 3} // 8
	// arr := []int{-1, 3, 4, 5, 2, 2, 2, 2} // 5

	fmt.Println(lnds(arr))
	fmt.Println(lnds1(arr))
	fmt.Println(lndsList(arr))
	fmt.Println(lndsList1(arr))
}
