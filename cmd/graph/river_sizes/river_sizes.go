package main

import "fmt"

func riverSizes(matrix [][]int) []int {
	return nil
}

func main() {

	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	fmt.Println(riverSizes(matrix)) // [2, 1, 5, 2, 2]
}
