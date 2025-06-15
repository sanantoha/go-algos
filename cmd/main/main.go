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
	fmt.Println(grph.PrintGraphAsAdjList(mst1(graph)))
	fmt.Println("\n\n")
	fmt.Println("=========================================")

	graph1 := createGraph1()

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
	fmt.Println(grph.PrintGraphAsAdjList(graph1))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst(graph1)))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst1(graph1)))
}

func mst(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {

	return nil
}

func mst1(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {

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
