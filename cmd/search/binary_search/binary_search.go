package main

import (
	"fmt"
)

// O(log(n)) time | O(1) space
func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := int(uint(l+r) >> 1) // avoid overflow
		if target < arr[mid] {
			r = mid - 1
		} else if target > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -(l + 1)
}

func main() {

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(binarySearch(arr, 80))
}
