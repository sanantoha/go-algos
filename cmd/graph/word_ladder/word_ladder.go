package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return -1
}

/**
 * https://leetcode.com/problems/word-ladder/
 */
func main() {

	actual := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(actual)
	fmt.Println(actual == 5)

	actual = ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	fmt.Println(actual)
	fmt.Println(actual == 0)

	actual = ladderLength("MAMA", "SIRI", []string{"SAMA", "SIMA", "SIRA", "SIRI", "MISA", "DISA"})
	fmt.Println(actual)
	fmt.Println(actual == 5)
}
