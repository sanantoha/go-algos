package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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
		"spiral_matrix_traverse.go":                  true,
		"binary_tree_zigzag_level_order_traverse.go": true,
		"lru_cache.go":                               true,
		"climbing_stars.go":                          true,
		"quick_sort.go":                              true,
		"maximum_subarray.go":                        true,
		"insert_sort.go":                             true,
		"first_duplicate_value.go":                   true,
		"permutations.go":                            true,
		"merge_intervals.go":                         true,
		"river_sizes.go":                             true,
		"bellman_ford_as_map.go":                     true,
		"add_two_numbers.go":                         true,
		"top_k_frequent_words.go":                    true,
		"find_nodes_distance_k.go":                   true,
		"kth_smallest_element.go":                    true,
		"clone_graph.go":                             true,
		"same_tree.go":                               true,
		"palindrome_linked_list.go":                  true,
		"string_permutations.go":                     true,
		"binary_search.go":                           true,
		"string_without_aaa_or_bbb.go":               true,
		"zig_zag_traverse.go":                        true,
		"depth_first_search.go":                      true,
		"levenshtein_distance.go":                    true,
		"best_time_to_buy_and_sell_stocks.go":        true,
		"subarray_sort.go":                           true,
		"dfs_tree_traverse.go":                       true,
		"reverse_words_in_string.go":                 true,
		"stable_internships.go":                      true,
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if _, exists := processedTasks[task]; exists {
			continue
		}
		fmt.Println(task)
	}

	runTask()
}

func runTask() {

	interns := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 2, 0},
	}

	teams := [][]int{
		{2, 1, 0},
		{0, 1, 2},
		{0, 1, 2},
	}

	expected := [][]int{
		{0, 1},
		{1, 0},
		{2, 2},
	}

	result := stableInternships(interns, teams)
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	fmt.Println(result)
	fmt.Println(reflect.DeepEqual(expected, result))
}

func stableInternships(interns [][]int, teams [][]int) [][]int {
	return nil
}
