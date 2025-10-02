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

	fmt.Println(findPermutation("abc", "hdflebacworld"))
	fmt.Println(findPermutation1("abc", "hdflebacworld"))

	fmt.Println(findPermutation("abbc", "hbbcadflebdworld"))
	fmt.Println(findPermutation1("abbc", "hbbcadflebdworld"))
}

func findPermutation(s1, s2 string) bool {
	return false
}

func findPermutation1(s1, s2 string) bool {
	return false
}
