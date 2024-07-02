package main

import (
	"fmt"
)

// O(n ^ 3) time | O(n) space
func wordBreak(s string, wordDict []string) bool {
	if s == "" || wordDict == nil || len(wordDict) == 0 {
		return false
	}

	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			w := s[j:i]
			if dp[j] {
				if _, exist := wordSet[w]; exist {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)]
}

func main() {

	fmt.Println(wordBreak("leetcode", []string{"leet", "code"})) // true

	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"})) // true

	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}
