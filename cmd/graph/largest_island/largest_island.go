package main

import "fmt"

func largestIsland(matrix [][]int) int {
	return -1
}

func largestIsland1(matrix [][]int) int {
	return -1
}

func main() {

	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
	}

	expected := 8

	actual := largestIsland(matrix)
	fmt.Println(actual)
	fmt.Println(actual == expected)

	actual = largestIsland1(matrix)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}
