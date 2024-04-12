package main

import "fmt"

// O(w + h) time | O(1) space
func searchMatrix(matrix [][]int, key int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	row := 0
	col := len(matrix[row]) - 1

	for row < len(matrix) && col >= 0 {

		if matrix[row][col] == key {
			return true
		} else if matrix[row][col] < key {
			row++
		} else {
			col--
		}
	}
	return false
}

func main() {

	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println(searchMatrix(matrix, 21))

	matrix1 := [][]int{
		{1, 4},
		{2, 5},
	}

	fmt.Println(searchMatrix(matrix1, 5))
}
