package main

import "fmt"

func maximumSubarray(arr []int) int {
	return -1
}

func main() {

	arr := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4} // 19
	fmt.Println(maximumSubarray(arr))

	arr1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 6
	fmt.Println(maximumSubarray(arr1))

	arr2 := []int{3, 4, -6, 7, 8, -18, 100} // 100
	fmt.Println(maximumSubarray(arr2))
}
