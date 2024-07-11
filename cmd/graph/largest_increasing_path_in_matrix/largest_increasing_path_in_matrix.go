package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	return -1
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
