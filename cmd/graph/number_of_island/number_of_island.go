package main

import "fmt"

// O(w * h) time | O(w * h) space
func numberOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	number := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				dfs(grid, row, col)
				number++
			}
		}
	}

	return number
}

func dfs(grid [][]int, startRow int, startCol int) {
	stack := make([][]int, 1)
	stack[0] = []int{startRow, startCol}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row := p[0]
		col := p[1]

		if grid[row][col] != 1 {
			continue
		}
		grid[row][col] = -1

		stack = append(stack, getNeighbors(grid, row, col)...)
	}
}

func getNeighbors(grid [][]int, row int, col int) [][]int {
	res := make([][]int, 0)

	if row > 0 {
		res = append(res, []int{row - 1, col})
	}
	if col > 0 {
		res = append(res, []int{row, col - 1})
	}
	if row+1 < len(grid) {
		res = append(res, []int{row + 1, col})
	}
	if col+1 < len(grid[row]) {
		res = append(res, []int{row, col + 1})
	}

	return res
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
