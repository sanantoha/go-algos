package main

import "fmt"

func kthSmallestElement(arr []int, k int) int {
	return -1
}

func main() {
	arr := []int{8, 2, 1, 6, 9, 3, 45, 6, 7, 13}

	for i := 0; i < len(arr); i++ {
		fmt.Println(kthSmallestElement(arr, i+1))
	}

	fmt.Println(kthSmallestElement(arr, -1))
	fmt.Println(kthSmallestElement(arr, 0))
	fmt.Println(kthSmallestElement(arr, 11))
}
