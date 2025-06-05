package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
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
		"level_order_binary_tree_traverse.go":                true,
		"find_mode_in_bst.go":                                true,
		"min_number_of_jumps.go":                             true,
		"left_view_binary_tree.go":                           true,
		"max_sum_increasing_subsequence.go":                  true,
		"longest_incr_subseq.go":                             true,
		"minimal_haviest_set_a.go":                           true,
		"topological_sort_as_map.go":                         true,
		"insert_greatest_common_divisor.go":                  true,
		"breadth_first_search.go":                            true,
		"number_of_island.go":                                true,
		"heap_sort.go":                                       true,
		"deep_copy_arbitrary_pointer.go":                     true,
		"merge_binary_tree.go":                               true,
		"reverse_string.go":                                  true,
		"minimum_window_substring.go":                        true,
		"balanced_brackets.go":                               true,
		"search_in_rotated_sorted_array.go":                  true,
		"all_paths_from_source_target.go":                    true,
		"min_number_of_coins_for_change.go":                  true,
		"three_sum.go":                                       true,
		"populating_next_right_pointer_in_each_node.go":      true,
		"generate_matrix.go":                                 true,
		"binary_tree_diameter.go":                            true,
		"lowest_common_ancestor_of_binary_tree.go":           true,
		"staircase_traversal.go":                             true,
		"word_search.go":                                     true,
		"unique_paths.go":                                    true,
		"balanced_binary_tree.go":                            true,
		"find_closest_value_in_bst.go":                       true,
		"counting_sort.go":                                   true,
		"all_elements_in_two_binary_search_trees.go":         true,
		"knapsack_problem.go":                                true,
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

	input := [][]int{
		{1, 2}, {4, 3}, {5, 6}, {6, 7},
	}

	expected := [][]int{{10}, {1, 3}}

	res := knapsackProblem(input, 10)
	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(res, expected))
}

func knapsackProblem(items [][]int, capacity int) [][]int {
	return nil
}
