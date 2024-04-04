package main

import "fmt"

// O(n * 4 ^ n) time | O(n * 4 ^ n) space
func phoneNumberMnemonics(str string) []string {
	res := make([]string, 0)
	letters := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	backtrack(str, letters, 0, "", &res)

	return res
}

func backtrack(str string, letters []string, idx int, ans string, res *[]string) {
	if idx == len(str) {
		*res = append(*res, ans)
		return
	}

	c := str[idx]
	d := int(c - '0')
	letter := letters[d]
	for i := 0; i < len(letter); i++ {
		backtrack(str, letters, idx+1, ans+string(letter[i]), res)
	}
}

func main() {
	str := "1905"

	// [1w0j 1w0k 1w0l 1x0j 1x0k 1x0l 1y0j 1y0k 1y0l 1z0j 1z0k 1z0l]
	fmt.Println(phoneNumberMnemonics(str))
}
