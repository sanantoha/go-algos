package main

import "fmt"

// O(n * l ^ 2) time | O(n * l ^ 2) space
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" {
		return 0
	}
	endWordExists := false
	for _, v := range wordList {
		if v == endWord {
			endWordExists = true
			break
		}
	}

	if !endWordExists {
		return 0
	}

	adjList := createAdjList(wordList)

	queue := make([]*Node, 1)
	queue[0] = &Node{beginWord, 1}

	visited := make(map[string]bool, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if endWord == node.word {
			return node.level
		}

		if visited[node.word] {
			continue
		}
		visited[node.word] = true

		for i := 0; i < len(node.word); i++ {
			key := node.word[:i] + "*" + node.word[i+1:]
			for _, word := range adjList[key] {

				queue = append(queue, &Node{word, node.level + 1})
			}
		}

	}

	return 0
}

type Node struct {
	word  string
	level int
}

func createAdjList(wordList []string) map[string][]string {
	adjList := make(map[string][]string, 0)

	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			key := word[:i] + "*" + word[i+1:]
			lst, ok := adjList[key]
			if !ok {
				lst = make([]string, 0)
			}
			lst = append(lst, word)
			adjList[key] = lst
		}
	}
	return adjList
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
