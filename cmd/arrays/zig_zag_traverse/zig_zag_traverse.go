package main

import "fmt"

// O(w * h) time | O(w * h) space
func zigZagTraverse(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	row, col := 0, 0
	height := len(matrix) - 1
	width := len(matrix[height]) - 1

	isGoingDown := true

	for row <= height && col <= width {
		res = append(res, matrix[row][col])

		if isGoingDown {
			if row == height || col == 0 {
				isGoingDown = false
				if row == height {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		} else {
			if row == 0 || col == width {
				isGoingDown = true
				if col == width {
					row++
				} else {
					col++
				}
			} else {
				row--
				col++
			}
		}
	}
	return res
}

func main() {

	matrix := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}

	fmt.Println(zigZagTraverse(matrix))
}
