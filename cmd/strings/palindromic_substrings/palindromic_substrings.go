package main

import "fmt"

// O(n ^ 2) time | O(1) space
func countSubstrings(str string) int {
	if str == "" {
		return 0
	}

	count := 0

	for i := 0; i < len(str); i++ {
		count += countPalindrome(str, i, i)
		count += countPalindrome(str, i, i+1)
	}
	return count
}

func countPalindrome(str string, l, r int) int {
	count := 0
	for l >= 0 && r < len(str) {
		if str[l] != str[r] {
			break
		}
		count++
		l--
		r++
	}

	return count
}

func main() {

	fmt.Println(countSubstrings("abc")) // 3

	fmt.Println(countSubstrings("aaa")) // 6

	fmt.Println(countSubstrings("aabbbaa")) // 14

	fmt.Println(countSubstrings("aaab")) // 7
}
