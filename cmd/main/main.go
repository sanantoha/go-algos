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
		"spiral_matrix_traverse.go":                          true,
		"sort_list.go":                                       true,
		"word_ladder.go":                                     true,
		"number_of_island.go":                                true,
		"min_heap.go":                                        true,
		"river_sizes.go":                                     true,
		"house_robber.go":                                    true,
		"merge_sort.go":                                      true,
		"binary_tree_diameter.go":                            true,
		"minimum_passes_of_matrix.go":                        true,
		"lru_cache.go":                                       true,
		"validate_starting_city.go":                          true,
		"largest_island.go":                                  true,
		"maximum_subarray.go":                                true,
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

	arr := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4} // 19
	fmt.Println(maximumSubarray(arr))

	arr1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 6
	fmt.Println(maximumSubarray(arr1))

	arr2 := []int{3, 4, -6, 7, 8, -18, 100} // 100
	fmt.Println(maximumSubarray(arr2))
}

func maximumSubarray(arr []int) int {
	return 0
}
