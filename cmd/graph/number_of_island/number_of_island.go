package main

import "fmt"

func numberOfIsland(grid [][]int) int {
	return -1
}

func main() {

	grid := [][]int{
		{0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}

	expected := 3

	actual := numberOfIsland(grid)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}
