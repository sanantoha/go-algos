package main

import (
	"fmt"
	"strings"
)

// O(s1 ^ 2 * s1! + s2 * s1!) time | O(s1 * s1!) space
func findPermutation(s1, s2 string) bool {
	for _, perm := range permute(s1) {
		if strings.Contains(s2, perm) {
			return true
		}
	}
	return false
}

func permute(str string) []string {
	res := make([]string, 0)
	backtrack(str, "", &res)
	return res
}

func backtrack(str string, ans string, res *[]string) {
	if len(str) == 0 {
		*res = append(*res, ans)
		return
	}

	for i := 0; i < len(str); i++ {
		rem := str[:i] + str[i+1:]
		backtrack(rem, ans+string(str[i]), res)
	}
}

// O(s1 + s2) time | O(1) space
func findPermutation1(s1, s2 string) bool {

	alphas := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		alphas[s1[i]-'a']++
	}

	for i := 0; i < len(s2); i++ {
		alphas[s2[i]-'a']--

		if i >= len(s1) {
			alphas[s2[i-len(s1)]-'a']++
		}
		if allZeros(alphas) {
			return true
		}
	}
	return false
}

func allZeros(alphas []int) bool {
	for _, v := range alphas {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findPermutation("abc", "hdflebacworld"))
	fmt.Println(findPermutation1("abc", "hdflebacworld"))

	fmt.Println(findPermutation("abbc", "hbbcadflebdworld"))
	fmt.Println(findPermutation1("abbc", "hbbcadflebdworld"))
}
