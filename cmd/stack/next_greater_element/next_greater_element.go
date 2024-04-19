package main

import "fmt"

func nextGreaterElement(arr []int) []int {
	return nil
}

func nextGreaterElement1(arr []int) []int {
	return nil
}

func main() {

	input := []int{2, 5, -3, -4, 6, 7, 2}
	actual := nextGreaterElement(input)
	fmt.Println(actual) // [5, 6, 6, 6, 7, -1, 5]

	actual = nextGreaterElement1(input)
	fmt.Println(actual) // [5, 6, 6, 6, 7, -1, 5]
}
