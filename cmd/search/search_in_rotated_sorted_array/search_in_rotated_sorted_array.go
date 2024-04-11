package main

import "fmt"

// O(log(n)) time | O(1) space
func search(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	smallestIdx := findSmallestIdx(arr)
	l := 0
	r := len(arr) - 1

	if target >= arr[smallestIdx] && target <= arr[smallestIdx] {
		l = smallestIdx
	} else {
		r = smallestIdx - 1
	}

	for l <= r {
		mid := l + (r-l)/2
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

func findSmallestIdx(arr []int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= arr[r] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// O(log(n)) time | O(1) space
func search1(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if target == arr[mid] {
			return mid
		}

		if arr[l] <= arr[mid] {
			if target >= arr[l] && target < arr[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > arr[mid] && target <= arr[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -(l + 1)
}

func main() {

	arr := []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35}
	target := 90

	fmt.Println(findSmallestIdx(arr))
	fmt.Println(search(arr, target))
	fmt.Println(search1(arr, target))
}
