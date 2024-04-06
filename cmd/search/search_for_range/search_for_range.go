package main

import "fmt"

func searchRange(arr []int, target int) [2]int {
	return [2]int{}
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
