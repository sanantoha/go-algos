package main

import (
	"fmt"
	"math"
	"slices"
)

// O(n ^ 2) time | O(n) space
func lis0(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	lis := make([]int, len(arr))
	for i, _ := range arr {
		lis[i] = 1
	}

	maxVal := math.MinInt

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
		if maxVal < lis[i] {
			maxVal = lis[i]
		}
	}
	return maxVal
}

// O(n * log(n)) time | O(n) space
func lis(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	res := make([]int, 1)
	res[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prev := res[len(res)-1]
		if prev < arr[i] {
			res = append(res, arr[i])
		} else {
			j, _ := slices.BinarySearch(res, arr[i])
			res[j] = arr[i]
		}
	}
	return len(res)
}

// O(n ^ 2) time | O(n) space
func lisList0(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}

	lis := make([]int, len(arr))
	prev := make([]int, len(arr))
	for i, _ := range arr {
		lis[i] = 1
		prev[i] = -1
	}

	maxIdx := 0

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
				prev[i] = j
			}
		}
		if lis[maxIdx] < lis[i] {
			maxIdx = i
		}
	}

	return buildSeq(arr, prev, maxIdx)
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

// O(n * log(n)) time | O(n) space
func lisList(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}

	indices := make([]int, len(arr)+1)
	for i, _ := range indices {
		indices[i] = -1
	}
	prev := make([]int, len(arr))
	for i, _ := range arr {
		prev[i] = -1
	}

	length := 0

	for i := 0; i < len(arr); i++ {
		maxLen := binarySearch(arr, indices, 1, length, arr[i])
		indices[maxLen] = i
		prev[i] = indices[maxLen-1]
		length = int(math.Max(float64(length), float64(maxLen)))
	}
	return buildSeq(arr, prev, indices[length])
}

func binarySearch(arr []int, indices []int, l int, r int, target int) int {
	for l <= r {
		mid := l + (r-l)/2
		if target <= arr[indices[mid]] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

	arr0 := []int{1, 2, 3, 6, -100, -90, -80, -70, -60, 7, 8, 9, 10, -50, -40} // 9
	fmt.Println(lis0(arr0))
	fmt.Println(lis(arr0))
	fmt.Println(lisList0(arr0))
	fmt.Println(lisList(arr0))
	fmt.Println("==============================")

	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80} // 6
	fmt.Println(lis0(arr))
	fmt.Println(lis(arr))
	fmt.Println(lisList0(arr))
	fmt.Println(lisList(arr))
	fmt.Println("==============================")

	arr1 := []int{4, 10, 4, 3, 8, 9} // 3
	fmt.Println(lis0(arr1))
	fmt.Println(lis(arr1))
	fmt.Println(lisList0(arr1))
	fmt.Println(lisList(arr1))
	fmt.Println("==============================")

	arr2 := []int{10, 9, 2, 5, 3, 7, 101, 18} // 4
	fmt.Println(lis0(arr2))
	fmt.Println(lis(arr2))
	fmt.Println(lisList0(arr2))
	fmt.Println(lisList(arr2))
	fmt.Println("==============================")

	arr3 := []int{1, -10, 20, 30, 2, 3, 4, 5} // 5
	fmt.Println(lis0(arr3))
	fmt.Println(lis(arr3))
	fmt.Println(lisList0(arr3))
	fmt.Println(lisList(arr3))
}
