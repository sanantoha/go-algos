package main

import "fmt"

// O(w * h) time | O(w * h) space
func longestIncreasingPath(matrix [][]int) int {
	if matrix == nil {
		return 0
	}

	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[i]))
	}

	longestIncPath := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			longestIncPath = max(dfs(matrix, row, col, cache), longestIncPath)
		}
	}
	return longestIncPath
}

func dfs(matrix [][]int, startRow int, startCol int, cache [][]int) int {
	if cache[startRow][startCol] > 0 {
		return cache[startRow][startCol]
	}

	val := matrix[startRow][startCol]

	for _, p := range getNeighbors(matrix, startRow, startCol) {
		row := p[0]
		col := p[1]
		if val < matrix[row][col] {
			cache[startRow][startCol] = max(dfs(matrix, row, col, cache), cache[startRow][startCol])
		}
	}
	cache[startRow][startCol]++
	return cache[startRow][startCol]
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

/**
 * https://leetcode.com/explore/interview/card/apple/350/trees-and-graphs/2049/
 * @param args
 */
func main() {

	input := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	fmt.Println(longestIncreasingPath(input)) // 4

	input1 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}

	fmt.Println(longestIncreasingPath(input1)) // 4
}
