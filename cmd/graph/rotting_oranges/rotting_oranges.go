package main

import "fmt"

func rottingOranges(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	queue := findAllRotten(grid)

	if len(queue) == 0 {
		fmt.Println("here")
		if isAtLeastOneFresh(grid) {
			return -1
		} else {
			return 0
		}
	}

	count := 0

	for len(queue) > 0 {
		size := len(queue)

		count++

		for size > 0 {
			size--
			p := queue[0]
			queue = queue[1:]
			row := p[0]
			col := p[1]

			for _, np := range getNeighbors(grid, row, col) {
				nrow := np[0]
				ncol := np[1]
				if grid[nrow][ncol] == 1 {
					grid[nrow][ncol] = 2
					queue = append(queue, []int{nrow, ncol})
				}
			}

		}
	}

	if isAtLeastOneFresh(grid) {
		return -1
	} else {
		return count - 1
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

func isAtLeastOneFresh(grid [][]int) bool {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				return true
			}
		}
	}
	return false
}

func findAllRotten(grid [][]int) [][]int {
	queue := make([][]int, 0)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 2 {
				queue = append(queue, []int{row, col})
			}
		}
	}

	return queue
}

func main() {

	input := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	res := rottingOranges(input)
	fmt.Println(res)

	input = [][]int{
		{0},
	}

	res = rottingOranges(input)
	fmt.Println(res)
}
