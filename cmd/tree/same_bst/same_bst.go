package main

import "fmt"

func sameBsts(arr1, arr2 []int) bool {
	return false
}

func sameBsts1(arr1, arr2 []int) bool {
	return false
}

func main() {

	arr1 := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	arr2 := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}

	fmt.Println(sameBsts(arr1, arr2))
	fmt.Println(sameBsts1(arr1, arr2))
}
