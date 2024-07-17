package main

import "fmt"

// O(w * h) time | O(w * h) space
func riverSizes(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}

	res := make([]int, 0)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 1 {
				res = append(res, dfsCount(matrix, row, col))
			}
		}
	}

	return res
}

func dfsCount(matrix [][]int, startRow int, startCol int) int {
	size := 0

	stack := make([][]int, 1)
	stack[0] = []int{startRow, startCol}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row := p[0]
		col := p[1]

		if matrix[row][col] != 1 {
			continue
		}
		matrix[row][col] = -1
		size++

		neighbors := getNeighbors(matrix, row, col)
		stack = append(stack, neighbors...)
	}

	return size
}

func getNeighbors(matrix [][]int, row int, col int) [][]int {
	res := make([][]int, 0)

	if row > 0 {
		res = append(res, []int{row - 1, col})
	}
	if col > 0 {
		res = append(res, []int{row, col - 1})
	}
	if row+1 < len(matrix) {
		res = append(res, []int{row + 1, col})
	}
	if col+1 < len(matrix[row]) {
		res = append(res, []int{row, col + 1})
	}

	return res
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
