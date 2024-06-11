package main

import "fmt"

func minNumberOfJumps(arr []int) int {
	return -1
}

func minNumberOfJumps1(arr []int) int {
	return -1
}

func main() {

	input := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}

	actual := minNumberOfJumps(input)
	fmt.Println(actual)
	fmt.Println(actual == 4)

	actual = minNumberOfJumps1(input)
	fmt.Println(actual)
	fmt.Println(actual == 4)
}
