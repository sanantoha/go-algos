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
		"spiral_matrix_traverse.go":                          true,
		"binary_tree_zigzag_level_order_traverse.go":         true,
		"lru_cache.go":                                       true,
		"climbing_stars.go":                                  true,
		"quick_sort.go":                                      true,
		"maximum_subarray.go":                                true,
		"insert_sort.go":                                     true,
		"first_duplicate_value.go":                           true,
		"permutations.go":                                    true,
		"merge_intervals.go":                                 true,
		"river_sizes.go":                                     true,
		"bellman_ford_as_map.go":                             true,
		"add_two_numbers.go":                                 true,
		"top_k_frequent_words.go":                            true,
		"find_nodes_distance_k.go":                           true,
		"kth_smallest_element.go":                            true,
		"clone_graph.go":                                     true,
		"same_tree.go":                                       true,
		"palindrome_linked_list.go":                          true,
		"string_permutations.go":                             true,
		"binary_search.go":                                   true,
		"string_without_aaa_or_bbb.go":                       true,
		"zig_zag_traverse.go":                                true,
		"depth_first_search.go":                              true,
		"levenshtein_distance.go":                            true,
		"best_time_to_buy_and_sell_stocks.go":                true,
		"subarray_sort.go":                                   true,
		"dfs_tree_traverse.go":                               true,
		"reverse_words_in_string.go":                         true,
		"stable_internships.go":                              true,
		"reconstruct_bst.go":                                 true,
		"merge_two_sorted_list.go":                           true,
		"dijkstra_shortest_path.go":                          true,
		"construct_binary_tree_from_preorder_and_inorder.go": true,
		"reverse_list.go":                                    true,
		"merge_sort.go":                                      true,
		"palindromic_substrings.go":                          true,
		"product_of_array.go":                                true,
		"next_greater_element.go":                            true,
		"four_sum.go":                                        true,
		"first_unique_character_in_string.go":                true,
		"reverse_binary_tree.go":                             true,
		"sqrt.go":                                            true,
		"word_break.go":                                      true,
		"longest_nondecr_subseq.go":                          true,
		"evaluate_expression_tree.go":                        true,
		"sort_list.go":                                       true,
		"search_for_range.go":                                true,
		"first_permutation_is_substr_second_str.go":          true,
		"dfs_tree_traverse_rec.go":                           true,
		"min_heap.go":                                        true,
		"cycle_linked_list.go":                               true,
		"delete_node_in_linked_list.go":                      true,
		"a_star_algo.go":                                     true,
		"binary_tree_tilt.go":                                true,
		"symmetric_tree.go":                                  true,
		"prim_min_spanning_tree_as_map.go":                   true,
		"greater_common_divisor.go":                          true,
		"house_robber.go":                                    true,
		"middle_node.go":                                     true,
		"surrounded_regions.go":                              true,
		"breadth_first_search_as_map.go":                     true,
		"topological_sort.go":                                true,
		"validate_bst.go":                                    true,
		"node_depths.go":                                     true,
		"powerset.go":                                        true,
		"rotate_array.go":                                    true,
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

	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr1, 3)
	fmt.Println(arr1) // 5, 6, 7, 1, 2, 3, 4

	arr11 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(arr11, 3)
	fmt.Println(arr11) // 5, 6, 7, 1, 2, 3, 4

	arr2 := []int{-1, -100, 3, 99}
	rotate(arr2, 2)
	fmt.Println(arr2) // 3, 99, -1, -100

	arr22 := []int{-1, -100, 3, 99}
	rotate1(arr22, 2)
	fmt.Println(arr22) // 3, 99, -1, -100

	arr3 := []int{1, 2, 3}
	rotate(arr3, 2)
	fmt.Println(arr3) // 2, 3, 1

	arr33 := []int{1, 2, 3}
	rotate1(arr33, 2)
	fmt.Println(arr33) // 2, 3, 1
}

func rotate(arr []int, k int) []int {
	return nil
}

func rotate1(arr []int, k int) []int {
	return nil
}
