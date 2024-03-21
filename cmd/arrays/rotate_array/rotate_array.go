package main

import "fmt"

// O(n) time | O(n) space
func rotate(arr []int, k int) {
	clone := make([]int, len(arr))
	for i, v := range arr {
		clone[i] = v
	}

	for i, _ := range arr {
		idx := (i + k) % len(arr)
		arr[idx] = clone[i]
	}
}

// O(n) time | O(1) space
func rotate1(arr []int, k int) {
	n := k % len(arr)
	reverse(arr, 0, len(arr)-1)
	reverse(arr, 0, n-1)
	reverse(arr, n, len(arr)-1)
}

func reverse(arr []int, l, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr1, 3)
	fmt.Println(arr1) // 5, 6, 7, 1, 2, 3, 4

	arr11 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(arr11, 3)
	fmt.Println(arr11) // 5, 6, 7, 1, 2, 3, 4

	arr2 := []int{-1, -100, 3, 99}
	rotate(arr2, 2)
	fmt.Println(arr2) // 3, 99, -1, -100

	arr22 := []int{-1, -100, 3, 99}
	rotate1(arr22, 2)
	fmt.Println(arr22) // 3, 99, -1, -100

	arr3 := []int{1, 2, 3}
	rotate(arr3, 2)
	fmt.Println(arr3) // 2, 3, 1

	arr33 := []int{1, 2, 3}
	rotate1(arr33, 2)
	fmt.Println(arr33) // 2, 3, 1
}
