package main

import (
	"fmt"
	"sort"
)

// O(n * log(n)) time | O(1) space
func largestRange(arr []int) [2]int {
	if arr == nil || len(arr) == 0 {
		return [2]int{-1, -1}
	}
	sort.Ints(arr)

	start := arr[0]
	currStart := arr[0]
	end := arr[0]

	length := 1
	maxLen := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1]+1 {
			length++
		} else {
			currStart = arr[i]
			length = 1
		}

		if maxLen < length {
			maxLen = length
			start = currStart
			end = arr[i]
		}
	}
	return [2]int{start, end}
}

// O(n) time | O(n) space
func largestRange1(arr []int) [2]int {

	bestRange := [2]int{-1, -1}
	maxLen := 0

	mp := make(map[int]bool)

	for _, v := range arr {
		mp[v] = true
	}

	for _, v := range arr {

		if !mp[v] {
			continue
		}

		length := 1

		l := v - 1
		for l >= 0 {
			if _, ok := mp[l]; ok {
				mp[l] = false
				l--
				length++
			} else {
				break
			}
		}

		r := v + 1
		for r < len(arr) {
			if _, ok := mp[r]; ok {
				mp[r] = false
				r++
				length++
			} else {
				break
			}
		}

		if maxLen < length {
			maxLen = length
			bestRange = [2]int{l + 1, r - 1}
		}
	}

	return bestRange
}

func main() {

	fmt.Println(largestRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}))

	fmt.Println(largestRange1([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}))
}
