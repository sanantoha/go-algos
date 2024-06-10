package main

import "fmt"

func minNumberOfCoinsForChange(n int, denoms []int) int {
	return -1
}

func main() {

	input := []int{1, 5, 10}
	actual := minNumberOfCoinsForChange(7, input)
	fmt.Println(actual)
	fmt.Println(actual == 3)
}
