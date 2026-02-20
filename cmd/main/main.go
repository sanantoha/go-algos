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
		"min_number_of_jumps.go":                             true,
		"intersection_linked_list.go":                        true,
		"find_mode_in_bst.go":                                true,
		"heap_sort.go":                                       true,
		"odd_even_linked_list.go":                            true,
		"subarray_sort.go":                                   true,
		"rotting_oranges.go":                                 true,
		"construct_binary_tree_from_preorder_and_inorder.go": true,
		"all_elements_in_two_binary_search_trees.go":         true,
		"zero_one_matrix.go":                                 true,
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

	input0 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	input1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	fmt.Println(updateMatrix(input0))

	fmt.Println(updateMatrix(input1))
}

func updateMatrix(matrix [][]int) [][]int {
	return nil
}
