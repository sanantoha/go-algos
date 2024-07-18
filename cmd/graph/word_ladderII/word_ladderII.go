package main

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	return nil
}

func main() {

	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"

	// [[hit, hot, dot, dog, cog], [hit, hot, lot, log, cog]]
	actual := findLadders(beginWord, endWord, wordList)
	fmt.Println(actual)
}
