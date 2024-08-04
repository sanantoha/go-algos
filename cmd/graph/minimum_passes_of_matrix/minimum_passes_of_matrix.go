package main

import "fmt"

// O(w * h) time | O(w * h) space
func minimumPassesOfMatrix(matrix [][]int) int {

	count := 0

	queue := getAllPositives(matrix)

	for len(queue) > 0 {
		size := len(queue)

		count++

		for size > 0 {
			size--

			curr := queue[0]
			queue = queue[1:]
			row := curr[0]
			col := curr[1]

			for _, p := range getNeighbors(matrix, row, col) {
				nrow := p[0]
				ncol := p[1]
				if matrix[nrow][ncol] < 0 {
					queue = append(queue, []int{nrow, ncol})
					matrix[nrow][ncol] *= -1
				}
			}
		}
	}

	if isNegative(matrix) {
		return -1
	}
	return count - 1
}

func isNegative(matrix [][]int) bool {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] < 0 {
				return true
			}
		}
	}
	return false
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

func getAllPositives(matrix [][]int) [][]int {
	q := make([][]int, 0)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] > 0 {
				q = append(q, []int{row, col})
			}
		}
	}
	return q
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
