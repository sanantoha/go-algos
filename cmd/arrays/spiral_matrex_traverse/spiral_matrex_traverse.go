package main

import "fmt"

// O(w * h) time | O(w * h) space
func spiral(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	startRow, startCol := 0, 0
	endRow := len(matrix) - 1
	endCol := len(matrix[endRow]) - 1

	for startRow <= endRow && startCol <= endCol {
		for i := startCol; i <= endCol; i++ {
			res = append(res, matrix[startRow][i])
		}
		startRow++

		for i := startRow; i <= endRow; i++ {
			res = append(res, matrix[i][endCol])
		}
		endCol--

		if startRow <= endRow && startCol <= endCol {
			for i := endCol; i >= startCol; i-- {
				res = append(res, matrix[endRow][i])
			}
			endRow--

			for i := endRow; i >= startRow; i-- {
				res = append(res, matrix[i][startCol])
			}
			startCol++
		}
	}
	return res
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
