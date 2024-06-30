package main

import (
	"fmt"
)

// O(m * n) time | O(m * n) space
func uniquePaths(m, n int) int {

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/*
 *  https://leetcode.com/problems/unique-paths/
 */
func main() {

	fmt.Println(uniquePaths(3, 2)) // 3

	fmt.Println(uniquePaths(3, 7)) // 28
}
