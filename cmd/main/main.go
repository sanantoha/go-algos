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
		"symmetric_tree.go":                             true,
		"prim_min_spanning_tree.go":                     true,
		"min_swaps_transform_string_to_palindrome.go":   true,
		"lru_cache.go":                                  true,
		"heap_sort.go":                                  true,
		"reverse_list.go":                               true,
		"levenshtein_distance.go":                       true,
		"bst_successor_search.go":                       true,
		"all_paths_from_source_target.go":               true,
		"level_order_binary_tree_traverse.go":           true,
		"all_elements_in_two_binary_search_trees.go":    true,
		"subarray_sort.go":                              true,
		"search_in_rotated_sorted_array.go":             true,
		"branch_sums.go":                                true,
		"sqrt.go":                                       true,
		"populating_next_right_pointer_in_each_node.go": true,
		"disk_stacking.go":                              true,
		"max_sum_increasing_subsequence.go":             true,
		"max_depth_of_bst.go":                           true,
		"knapsack_problem.go":                           true,
		"three_sum.go":                                  true,
		"find_nodes_distance_k.go":                      true,
		"dijkstra_shortest_path_as_map.go":              true,
		"intersection_linked_list.go":                   true,
		"depth_first_search.go":                         true,
		"bfs_tree_traverse.go":                          true,
		"binary_search.go":                              true,
		"find_pivot_index.go":                           true,
		"longest_decr_subseq.go":                        true,
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

	arr := []int{5, 6, 7, 6, 5, 4, 3, 10, 14, 12} // 5

	fmt.Println(lds(arr))
	fmt.Println(lds1(arr))
	fmt.Println(ldsList(arr))
	fmt.Println(ldsList1(arr))

	arr1 := []int{100, 10, 9, 8, 7, 6, 5, 90, 80, 70, 60, 50, 40, 30, 20} // 9

	fmt.Println(lds(arr1))
	fmt.Println(lds1(arr1))
	fmt.Println(ldsList(arr1))
	fmt.Println(ldsList1(arr1))
}

func lds(arr []int) int {
	return -1
}

func lds1(arr []int) int {
	return -1
}

func ldsList(arr []int) []int {
	return nil
}

func ldsList1(arr []int) []int {
	return nil
}
