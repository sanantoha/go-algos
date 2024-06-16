package main

import "fmt"

func waterArea(heights []int) int {
	return 0
}

func waterArea1(heights []int) int {
	return 0
}

func main() {

	input := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}

	actual := waterArea(input)
	fmt.Println(actual == 48)

	actual = waterArea1(input)
	fmt.Println(actual == 48)
}
