package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	tasks := make([]string, 0)

	err := filepath.WalkDir("./cmd", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil // Continue walking
		}
		if !d.IsDir() { // Print only files
			//fmt.Println(path)
			fileName := filepath.Base(path)
			tasks = append(tasks, fileName)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

	processedTasks := map[string]bool{
		"symmetric_tree.go":                           true,
		"prim_min_spanning_tree.go":                   true,
		"min_swaps_transform_string_to_palindrome.go": true,
		"lru_cache.go":                                true,
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if strings.HasSuffix(task, ".txt") || strings.HasSuffix(task, "test.go") || task == "main.go" {
			continue
		}
		if _, exists := processedTasks[task]; exists {
			continue
		}
		fmt.Println(task)
	}

	runTask()
}

func runTask() {

	lruCache := NewLRUCache(2)

	lruCache.put(1, 1)           // cache is {1=1}
	lruCache.put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lruCache.get(1)) // return 1

	lruCache.put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lruCache.get(2)) // returns -1 (not found)
	lruCache.put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lruCache.get(1)) // return -1 (not found)
	fmt.Println(lruCache.get(3)) // return 3
	fmt.Println(lruCache.get(4)) // return 4
}

type LRUCache struct {
}

func NewLRUCache(capacity int) *LRUCache {
	return nil
}

func (lc *LRUCache) put(key int, value int) {

}

func (lc *LRUCache) get(key int) int {
	return -1
}
