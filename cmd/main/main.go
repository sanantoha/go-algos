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
		"symmetric_tree.go":                                  true,
		"prim_min_spanning_tree.go":                          true,
		"min_swaps_transform_string_to_palindrome.go":        true,
		"lru_cache.go":                                       true,
		"heap_sort.go":                                       true,
		"reverse_list.go":                                    true,
		"levenshtein_distance.go":                            true,
		"bst_successor_search.go":                            true,
		"all_paths_from_source_target.go":                    true,
		"level_order_binary_tree_traverse.go":                true,
		"all_elements_in_two_binary_search_trees.go":         true,
		"subarray_sort.go":                                   true,
		"search_in_rotated_sorted_array.go":                  true,
		"branch_sums.go":                                     true,
		"sqrt.go":                                            true,
		"populating_next_right_pointer_in_each_node.go":      true,
		"disk_stacking.go":                                   true,
		"max_sum_increasing_subsequence.go":                  true,
		"max_depth_of_bst.go":                                true,
		"knapsack_problem.go":                                true,
		"three_sum.go":                                       true,
		"find_nodes_distance_k.go":                           true,
		"dijkstra_shortest_path_as_map.go":                   true,
		"intersection_linked_list.go":                        true,
		"depth_first_search.go":                              true,
		"bfs_tree_traverse.go":                               true,
		"binary_search.go":                                   true,
		"find_pivot_index.go":                                true,
		"longest_decr_subseq.go":                             true,
		"median_of_two_sorted_arrays.go":                     true,
		"dfs_tree_traverse.go":                               true,
		"binary_tree_zigzag_level_order_traverse.go":         true,
		"regular_expressions.go":                             true,
		"dfs_tree_traverse_rec.go":                           true,
		"spiral_matrix_traverse.go":                          true,
		"stable_internships.go":                              true,
		"merge_two_sorted_list.go":                           true,
		"breadth_first_search.go":                            true,
		"validate_starting_city.go":                          true,
		"string_permutations.go":                             true,
		"balanced_binary_tree.go":                            true,
		"generate_parentheses.go":                            true,
		"deep_copy_arbitrary_pointer.go":                     true,
		"optimal_freelancing.go":                             true,
		"construct_binary_tree_from_preorder_and_inorder.go": true,
		"clone_graph.go":                                     true,
		"topological_sort_as_map.go":                         true,
		"counting_sort.go":                                   true,
		"kth_smallest_element.go":                            true,
		"first_duplicate_value.go":                           true,
		"search_2d_matrix.go":                                true,
		"unique_path_iii.go":                                 true,
		"first_permutation_is_substr_second_str.go":          true,
		"four_sum.go":                                        true,
		"merge_sort.go":                                      true,
		"cycle_linked_list.go":                               true,
		"binary_tree_diameter.go":                            true,
		"greater_common_divisor.go":                          true,
		"same_tree.go":                                       true,
		"sparse_matrix_multiplication.go":                    true,
		"top_k_frequent_words.go":                            true,
		"find_closest_value_in_bst.go":                       true,
		"number_of_island.go":                                true,
		"reverse_words_in_string.go":                         true,
		"select_sort.go":                                     true,
		"sort_k_sorted_array.go":                             true,
		"subarray_sum_equals_k.go":                           true,
		"reverse_binary_tree.go":                             true,
		"unique_paths.go":                                    true,
		"node_depths.go":                                     true,
		"insert_sort.go":                                     true,
		"insert_greatest_common_divisor.go":                  true,
		"powerset.go":                                        true,
		"find_mode_in_bst.go":                                true,
		"rotate_array.go":                                    true,
		"subtree_of_another_tree.go":                         true,
		"zig_zag_traverse.go":                                true,
		"depth_first_search_as_map.go":                       true,
		"k_th_smallest_element_in_bst.go":                    true,
		"palindromic_substrings.go":                          true,
		"pow.go":                                             true,
		"one_edit.go":                                        true,
		"merge_intervals.go":                                 true,
		"search_for_range.go":                                true,
		"max_path_sum_in_binary_tree.go":                     true,
		"add_two_numbers.go":                                 true,
		"phone_number_mnemonic.go":                           true,
		"product_of_array.go":                                true,
		"min_rewards.go":                                     true,
		"minimum_absolute_difference_in_bst.go":              true,
		"evaluate_expression_tree.go":                        true,
		"water_area.go":                                      true,
		"surrounded_regions.go":                              true,
		"best_time_to_buy_and_sell_stocks.go":                true,
		"left_view_binary_tree.go":                           true,
		"generate_matrix.go":                                 true,
		"longest_increasing_path_in_matrix.go":               true,
		"longest_incr_subseq.go":                             true,
		"minimal_haviest_set_a.go":                           true,
		"palindrome_linked_list.go":                          true,
		"number_of_way_to_make_change.go":                    true,
		"combination_sum.go":                                 true,
		"longest_nondecr_subseq.go":                          true,
		"middle_node.go":                                     true,
		"word_ladder.go":                                     true,
		"validate_bst.go":                                    true,
		"balanced_brackets.go":                               true,
		"word_break.go":                                      true,
		"rotting_oranges.go":                                 true,
		"quick_sort.go":                                      true,
		"dijkstra_shortest_path.go":                          true,
		"convert_sorted_array_to_bst.go":                     true,
		"find_peak_element.go":                               true,
		"largest_range.go":                                   true,
		"lowest_common_ancestor_of_bst.go":                   true,
		"sort_list.go":                                       true,
		"reverse_integer.go":                                 true,
		"breadth_first_search_as_map.go":                     true,
		"min_heap.go":                                        true,
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

func updateMatrix(matrix [][]int) int {
	return -1
}
