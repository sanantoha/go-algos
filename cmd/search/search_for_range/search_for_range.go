package main

import "fmt"

// O(log(n)) time | O(1) space
func searchRange(arr []int, target int) []int {
	l := leftBinarySearch(arr, target)

	if l < 0 {
		return []int{}
	}
	r := rightBinarySearch(arr, l, target)
	return []int{l, r}
}

func rightBinarySearch(arr []int, l int, target int) int {
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2
		if target >= arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func leftBinarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2
		if target <= arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if len(arr) > 0 && arr[l] == target {
		return l
	} else {
		return -(l + 1)
	}
}

func main() {

	arr := []int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 10}
	fmt.Println(searchRange(arr, 8)) // [3, 9]

	fmt.Println(searchRange(arr, 6)) // []

	arr1 := []int{1}
	fmt.Println(searchRange(arr1, 1)) // [0, 0]

	arr2 := []int{}
	fmt.Println(searchRange(arr2, 0)) // []
}
