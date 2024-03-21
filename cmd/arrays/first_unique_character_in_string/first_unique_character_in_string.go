package main

import "fmt"

// O(n) time | O(d) space where d unique chars
func firstUniqChar(str string) int {

	mp := make(map[rune]int)

	for _, c := range str {
		mp[c]++
	}

	for i, c := range str {
		if mp[c] == 1 {
			return i
		}
	}
	return -1
}

func main() {

	fmt.Println(firstUniqChar("leetcode"))     // 0
	fmt.Println(firstUniqChar("loveleetcode")) // 2
	fmt.Println(firstUniqChar("aabb"))         // -1
}
