package main

import "fmt"

func minimumPassesOfMatrix(matrix [][]int) int {
	return -1
}

func main() {

	matrix := [][]int{
		{0, -1, -3, 2, 0},
		{1, -2, -5, -1, -3},
		{3, 0, 0, -4, -1},
	}

	expected := 3

	actual := minimumPassesOfMatrix(matrix)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}
