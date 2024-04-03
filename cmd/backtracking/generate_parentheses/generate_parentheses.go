package main

import "fmt"

// O((4 ^ n) / sqrt(n)) time | O((4 ^ n) / sqrt(n)) space // found in leetcode https://leetcode.com/problems/generate-parentheses/solution/
func generateParentheses(cnt int) []string {
	res := make([]string, 0)
	backtrack(cnt, 0, 0, "", &res)
	return res
}

func backtrack(cnt, open, close int, ans string, res *[]string) {
	if cnt*2 == (open + close) {
		*res = append(*res, ans)
		return
	}

	if open < cnt {
		backtrack(cnt, open+1, close, ans+string('('), res)
	}
	if close < open {
		backtrack(cnt, open, close+1, ans+string(')'), res)
	}
}

func main() {

	fmt.Println(generateParentheses(3))

	fmt.Println(generateParentheses(2))
}
