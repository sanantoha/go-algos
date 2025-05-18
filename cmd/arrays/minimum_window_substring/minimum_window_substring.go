package main

import (
	"fmt"
	"math"
)

// O(s + t) time | O(s + t) space
func minimumWindowSubstring(s string, t string) string {

	dictT := make(map[rune]int)
	for _, c := range t {
		dictT[c] += 1
	}

	formed := 0
	required := len(dictT)

	l, r := 0, 0
	ansL, ansR := 0, 0
	ans := math.MaxInt

	wordCount := make(map[rune]int)

	for r < len(s) {
		c := rune(s[r])

		wordCount[c] += 1

		dcnt, ok := dictT[c]
		if ok && dcnt == wordCount[c] {
			formed++
		}

		for l <= r && formed == required {
			c = rune(s[l])

			if ans > (r - l + 1) {
				ans = r - l + 1
				ansL = l
				ansR = r
			}

			wordCount[c] -= 1

			dcnt, ok = dictT[c]
			if ok && wordCount[c] < dcnt {
				formed--
			}

			l++
		}
		r++
	}

	if ans == math.MaxInt {
		return ""
	}
	return s[ansL : ansR+1]
}

func main() {

	fmt.Println(minimumWindowSubstring("ADOBECODEBANC", "ABC")) // "BANC"

	fmt.Println(minimumWindowSubstring("a", "a")) // a

	fmt.Println(minimumWindowSubstring("a", "aa")) // ""
}
