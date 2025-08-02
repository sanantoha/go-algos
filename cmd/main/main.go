package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
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
		"bellman_ford.go":                                    true,
		"sparse_matrix_multiplication.go":                    true,
		"min_rewards.go":                                     true,
		"longest_decr_subseq.go":                             true,
		"water_area.go":                                      true,
		"unique_path_iii.go":                                 true,
		"kruskal_min_spanning_tree_as_map.go":                true,
		"phone_number_mnemonic.go":                           true,
		"apartment_hunting.go":                               true,
		"subarray_sum_equals_k.go":                           true,
		"k_th_smallest_element_in_bst.go":                    true,
		"number_of_way_to_make_change.go":                    true,
		"optimal_freelancing.go":                             true,
		"dijkstra_shortest_path_as_map.go":                   true,
		"sort_k_sorted_array.go":                             true,
		"valid_ip_address.go":                                true,
		"word_ladderII.go":                                   true,
		"branch_sums.go":                                     true,
		"minimum_absolute_difference_in_bst.go":              true,
		"largest_island.go":                                  true,
		"depth_first_search_as_map.go":                       true,
		"one_edit.go":                                        true,
		"find_peak_element.go":                               true,
		"max_path_sum_in_binary_tree.go":                     true,
		"regular_expressions.go":                             true,
		"combination_sum.go":                                 true,
		"longest_increasing_path_in_matrix.go":               true,
		"intersection_linked_list.go":                        true,
		"generate_parentheses.go":                            true,
		"lowest_common_ancestor_of_bst.go":                   true,
		"remove_nth_from_end_of_list.go":                     true,
		"validate_starting_city.go":                          true,
		"search_2d_matrix.go":                                true,
		"reverse_integer.go":                                 true,
		"prim_min_spanning_tree.go":                          true,
		"median_of_two_sorted_arrays.go":                     true,
		"disk_stacking.go":                                   true,
		"bst_successor_search.go":                            true,
		"largest_range.go":                                   true,
		"convert_sorted_array_to_bst.go":                     true,
		"min_swaps_transform_string_to_palindrome.go":        true,
		"pow.go":                       true,
		"select_sort.go":               true,
		"word_ladder.go":               true,
		"odd_even_linked_list.go":      true,
		"find_pivot_index.go":          true,
		"minimum_passes_of_matrix.go":  true,
		"bfs_tree_traverse.go":         true,
		"max_depth_of_bst.go":          true,
		"subtree_of_another_tree.go":   true,
		"same_bst.go":                  true,
		"kruskal_min_spanning_tree.go": true,
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

	graph := grph.NewEdgeWeightedGraph(6)

	graph.AddEdge(grph.NewEdge(0, 1, 7.0))
	graph.AddEdge(grph.NewEdge(0, 2, 8.0))
	graph.AddEdge(grph.NewEdge(1, 2, 3.0))
	graph.AddEdge(grph.NewEdge(1, 3, 6.0))
	graph.AddEdge(grph.NewEdge(2, 3, 4.0))
	graph.AddEdge(grph.NewEdge(2, 4, 3.0))
	graph.AddEdge(grph.NewEdge(3, 4, 2.0))
	graph.AddEdge(grph.NewEdge(3, 5, 5.0))
	graph.AddEdge(grph.NewEdge(4, 5, 2.0))
	/*
	   6 5
	   0: 0-1 7.00
	   1: 1-2 3.00  0-1 7.00
	   2: 1-2 3.00  2-4 3.00
	   3: 3-4 2.00
	   4: 3-4 2.00  4-5 2.00  2-4 3.00
	   5: 4-5 2.00
	*/
	fmt.Println(graph)
	fmt.Println("=========================================")
	fmt.Println(mst(graph))
	fmt.Println("=========================================")
	fmt.Println(mst1(graph))
	fmt.Println("\n\n")
	fmt.Println("=========================================")

	graph1 := grph.NewEdgeWeightedGraph(7)
	graph1.AddEdge(grph.NewEdge(0, 1, 2.0))
	graph1.AddEdge(grph.NewEdge(0, 2, 3.0))
	graph1.AddEdge(grph.NewEdge(0, 3, 7.0))
	graph1.AddEdge(grph.NewEdge(1, 2, 6.0))
	graph1.AddEdge(grph.NewEdge(1, 6, 3.0))
	graph1.AddEdge(grph.NewEdge(2, 4, 1.0))
	graph1.AddEdge(grph.NewEdge(2, 5, 8.0))
	graph1.AddEdge(grph.NewEdge(3, 4, 5.0))
	graph1.AddEdge(grph.NewEdge(4, 5, 4.0))
	graph1.AddEdge(grph.NewEdge(5, 6, 2.0))

	/*
	   7 6
	   0: 0-1 2.00000  0-2 3.00000
	   1: 0-1 2.00000  1-6 3.00000
	   2: 2-4 1.00000  0-2 3.00000
	   3: 3-4 5.00000
	   4: 2-4 1.00000  3-4 5.00000
	   5: 5-6 2.00000
	   6: 5-6 2.00000  1-6 3.00000
	*/
	fmt.Println(graph1)
	fmt.Println("=========================================")
	fmt.Println(mst(graph1))
	fmt.Println("=========================================")
	fmt.Println(mst1(graph1))
}

func mst(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	return nil
}

func mst1(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	return nil
}
