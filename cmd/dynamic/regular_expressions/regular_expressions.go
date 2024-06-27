package main

import "fmt"

// O(s * p) time | O(s * p) space
func isMatch(s, p string) bool {
	memo := make([][]bool, len(s)+1)
	for i, _ := range memo {
		memo[i] = make([]bool, len(p)+1)
	}
	return dp(0, 0, s, p, memo)
}

func dp(i int, j int, s string, p string, memo [][]bool) bool {
	if !memo[i][j] {
		var ans bool
		if len(p) == j {
			ans = len(s) == i
		} else {
			firstMatch := i < len(s) && (s[i] == p[j] || '.' == p[j])

			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2, s, p, memo) || firstMatch && dp(i+1, j, s, p, memo)
			} else {
				ans = firstMatch && dp(i+1, j+1, s, p, memo)
			}
		}

		memo[i][j] = ans
	}
	return memo[i][j]
}

// O(s * p) time | O(s * p) space
func isMatchIter(s, p string) bool {

	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (p[j] == '.' || s[i] == p[j])

			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func main() {

	fmt.Println(!isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("abcde", ".*"))
	fmt.Println(isMatch("abcde", ".*de"))
	fmt.Println(!isMatch("abcde", ".*dk"))

	fmt.Println("======================================")

	fmt.Println(!isMatchIter("aa", "a"))
	fmt.Println(isMatchIter("aa", "a*"))
	fmt.Println(isMatchIter("abcde", ".*"))
	fmt.Println(isMatchIter("abcde", ".*de"))
	fmt.Println(!isMatchIter("abcde", ".*dk"))

}
