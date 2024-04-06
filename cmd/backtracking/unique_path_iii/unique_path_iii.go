package main

import "fmt"

/**
 * https://leetcode.com/problems/unique-paths-iii/description/
 *
 * You are given an m x n integer array grid where grid[i][j] could be:
 *
 * 1 representing the starting square. There is exactly one starting square.
 * 2 representing the ending square. There is exactly one ending square.
 * 0 representing empty squares we can walk over.
 * -1 representing obstacles that we cannot walk over.
 * Return the number of 4-directional walks from the starting square to the ending square,
 * that walk over every non-obstacle square exactly once.
 *
 * Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
 * Output: 2
 * Explanation: We have the following two paths:
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
 * 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
 */

// O(3 ^ n) time | O(n) space
func uniquePath3(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	zeros := 0
	startRow := 0
	startCol := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				zeros++
			} else if grid[row][col] == 1 {
				startRow = row
				startCol = col
			}
		}
	}

	return paths(grid, startRow, startCol, zeros)
}

func paths(grid [][]int, row, col, zeros int) int {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] == -1 {
		return 0
	}
	if grid[row][col] == 2 {
		if zeros == -1 {
			return 1
		} else {
			return 0
		}
	}

	zeros--
	tmp := grid[row][col]
	grid[row][col] = -1

	count := paths(grid, row-1, col, zeros) + paths(grid, row, col-1, zeros) + paths(grid, row+1, col, zeros) +
		paths(grid, row, col+1, zeros)

	grid[row][col] = tmp

	return count
}

func main() {

	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}

	fmt.Println(uniquePath3(grid)) // 2
}
