package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	grph "github.com/sanantoha/go-algos/internals/graph"
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
		"house_robber.go":                                    true,
		"minimum_passes_of_matrix.go":                        true,
		"reverse_string.go":                                  true,
		"kruskal_min_spanning_tree_as_map.go":                true,
		"staircase_traversal.go":                             true,
		"minimum_window_substring.go":                        true,
		"merge_binary_tree.go":                               true,
		"maximum_subarray.go":                                true,
		"next_greater_element.go":                            true,
		"min_number_of_jumps.go":                             true,
		"permutations.go":                                    true,
		"first_unique_character_in_string.go":                true,
		"remove_nth_from_end_of_list.go":                     true,
		"word_search.go":                                     true,
		"string_without_aaa_or_bbb.go":                       true,
		"delete_node_in_linked_list.go":                      true,
		"lowest_common_ancestor_of_binary_tree.go":           true,
		"bellman_ford.go":                                    true,
		"word_ladderII.go":                                   true,
		"same_bst.go":                                        true,
		"prim_min_spanning_tree_as_map.go":                   true,
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

	graph := createGraph()
	/*
	   6 5
	   0: 0-1 7.00
	   1: 1-2 3.00  0-1 7.00
	   2: 1-2 3.00  2-4 3.00
	   3: 3-4 2.00
	   4: 3-4 2.00  4-5 2.00  2-4 3.00
	   5: 4-5 2.00
	*/

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst(graph)))
	fmt.Println("=========================================")

	graph1 := createGraph1()

	/*
	   7 6
	   0: 0-1 2.00  0-2 3.00
	   1: 0-1 2.00  1-6 3.00
	   2: 2-4 1.00  0-2 3.00
	   3: 3-4 5.00
	   4: 2-4 1.00  3-4 5.00
	   5: 5-6 2.00
	   6: 5-6 2.00  1-6 3.00
	*/
	fmt.Println(grph.PrintGraphAsAdjList(graph1))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst(graph1)))
}

func mst(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {
	return nil
}

func createGraph() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 7.0)
	edge02 := grph.NewEdgeT("0", "2", 8.0)
	edge12 := grph.NewEdgeT("1", "2", 3.0)
	edge13 := grph.NewEdgeT("1", "3", 6.0)
	edge23 := grph.NewEdgeT("2", "3", 4.0)
	edge24 := grph.NewEdgeT("2", "4", 3.0)
	edge34 := grph.NewEdgeT("3", "4", 2.0)
	edge35 := grph.NewEdgeT("3", "5", 5.0)
	edge45 := grph.NewEdgeT("4", "5", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge02}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge12, edge13}
	graph["2"] = []*grph.EdgeT[string]{edge02, edge12, edge23, edge24}
	graph["3"] = []*grph.EdgeT[string]{edge13, edge23, edge34, edge35}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34, edge45}
	graph["5"] = []*grph.EdgeT[string]{edge35, edge45}
	return graph
}

func createGraph1() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 2.0)
	edge02 := grph.NewEdgeT("0", "2", 3.0)
	edge03 := grph.NewEdgeT("0", "3", 7.0)
	edge12 := grph.NewEdgeT("1", "2", 6.0)
	edge16 := grph.NewEdgeT("1", "6", 3.0)
	edge24 := grph.NewEdgeT("2", "4", 1.0)
	edge25 := grph.NewEdgeT("2", "5", 8.0)
	edge34 := grph.NewEdgeT("3", "4", 5.0)
	edge45 := grph.NewEdgeT("4", "5", 4.0)
	edge56 := grph.NewEdgeT("5", "6", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge02, edge03}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge12, edge16}
	graph["2"] = []*grph.EdgeT[string]{edge02, edge12, edge24, edge25}
	graph["3"] = []*grph.EdgeT[string]{edge03, edge34}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34, edge45}
	graph["5"] = []*grph.EdgeT[string]{edge25, edge45, edge56}
	graph["6"] = []*grph.EdgeT[string]{edge16, edge56}

	return graph
}
