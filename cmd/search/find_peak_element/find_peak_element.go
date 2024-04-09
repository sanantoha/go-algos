package main

import "fmt"

/**
 * https://leetcode.com/problems/find-peak-element/description/
 * A peak element is an element that is strictly greater than its neighbors.
 *
 * Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
 *
 * You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.
 * You must write an algorithm that runs in O(log n) time.
 */
// O(log(n)) time | O(1) space
func findPeak(arr []int) int {
	l := 0
	r := len(arr) - 1

	for l < r {
		mid := l + (r-l)/2

		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {

	arr := []int{1, 2, 1, 3, 4, 5, 7, 8, 5, 3, 2, 11, 1}
	fmt.Println(findPeak(arr)) // 7
}
