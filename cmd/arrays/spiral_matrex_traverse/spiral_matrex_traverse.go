package main

import "fmt"

func spiral(matrix [][]int) []int {
	return nil
}

func main() {

	matrix := [][]int{
		{1, 2, 3, 5, 6, 7},
		{19, 20, 21, 22, 23, 8},
		{18, 29, 30, 31, 24, 9},
		{17, 28, 27, 26, 25, 10},
		{16, 15, 14, 13, 12, 11},
	}

	fmt.Println(spiral(matrix))
}
