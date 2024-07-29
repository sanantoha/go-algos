package main

import "fmt"

// O(w ^ 2 * h ^ 2) time | O(w * h) space
func largestIsland(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	largest := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 1 {
				largest = max(largest, dfsSize(matrix, row, col))
			}
		}
	}
	return largest
}

func dfsSize(matrix [][]int, startRow int, startCol int) int {
	size := 1
	stack := make([][]int, 0)
	stack = append(stack, getNeighbors(matrix, startRow, startCol)...)

	visited := make([][]bool, len(matrix))
	for row := 0; row < len(matrix); row++ {
		visited[row] = make([]bool, len(matrix[row]))
	}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row := p[0]
		col := p[1]

		if visited[row][col] || matrix[row][col] != 0 {
			continue
		}
		visited[row][col] = true
		size++

		stack = append(stack, getNeighbors(matrix, row, col)...)
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
