package main

import (
	"fmt"
	"slices"
)

// O(n * l ^ 2) time | O(n * l ^ 2) space
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if beginWord == "" || endWord == "" || !slices.Contains(wordList, endWord) {
		return nil
	}

	res := make([][]string, 0)

	queue := make([]*Node, 1)
	initPath := make([]string, 1)
	initPath[0] = beginWord
	queue[0] = &Node{
		word: beginWord,
		path: initPath,
	}

	adjList := createAdjList(wordList)

	visited := make(map[string]bool, 0)

	for len(queue) > 0 {
		size := len(queue)

		currVisited := make(map[string]bool, 0)

		for size > 0 {
			size--

			node := queue[0]
			queue = queue[1:]

			if endWord == node.word {
				res = append(res, node.path)
				continue
			}

			for i := 0; i < len(node.word); i++ {
				key := node.word[:i] + "*" + node.word[i+1:]

				for _, word := range adjList[key] {
					if visited[word] {
						continue
					}
					currVisited[word] = true
					path := make([]string, len(node.path))
					copy(path, node.path)
					path = append(path, word)

					queue = append(queue, &Node{word, path})
				}
			}
		}

		for key, value := range currVisited {
			visited[key] = value
		}
	}

	return res
}

type Node struct {
	word string
	path []string
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

func main() {

	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"

	// [[hit, hot, dot, dog, cog], [hit, hot, lot, log, cog]]
	actual := findLadders(beginWord, endWord, wordList)
	fmt.Println(actual)
}
