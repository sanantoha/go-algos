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
		"level_order_binary_tree_traverse.go":         true,
		"all_elements_in_two_binary_search_trees.go":  true,
		"subarray_sort.go":                            true,
		"search_in_rotated_sorted_array.go":           true,
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

	arr := []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35}
	target := 90

	fmt.Println(findSmallestIdx(arr))
	fmt.Println(search(arr, target))
	fmt.Println(search1(arr, target))
}

func findSmallestIdx(arr []int) int {
	return -1
}

func search(arr []int, target int) int {
	return -1
}

func search1(arr []int, target int) int {
	return -1
}
