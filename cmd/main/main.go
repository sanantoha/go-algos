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
		"heap_sort.go":                                true,
		"reverse_list.go":                             true,
		"levenshtein_distance.go":                     true,
		"bst_successor_search.go":                     true,
		"all_paths_from_source_target.go":             true,
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

	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	actualRec := allPathsSourceTargetRec(graph)
	actual := allPathsSourceTarget(graph)

	fmt.Println(actualRec) // [[0 1 3] [0 2 3]]
	fmt.Println(actual)    // [[0 2 3] [0 1 3]]

	graph1 := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}

	actualRec = allPathsSourceTargetRec(graph1)
	actual = allPathsSourceTarget(graph1)

	fmt.Println(actualRec) // [[0 4] [0 3 4] [0 1 3 4] [0 1 2 3 4] [0 1 4]]
	fmt.Println(actual)    // [[0 1 4] [0 1 2 3 4] [0 1 3 4] [0 3 4] [0 4]]
}

func allPathsSourceTargetRec(graph [][]int) [][]int {
	return nil
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return nil
}
