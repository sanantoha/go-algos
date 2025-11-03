package main

import (
	"fmt"
	"math"
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

	fmt.Println(pow(4, 2))
	fmt.Println(pow(2, 4))
	fmt.Println(pow(2.0, -2))

	fmt.Println(pow(2.1, 3))
	fmt.Println(math.Pow(2.1, 3))
}

func pow(x float64, n int) float64 {
	return 0.0
}
