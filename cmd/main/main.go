package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
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
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if strings.HasSuffix(task, ".txt") || strings.HasSuffix(task, "test.go") {
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

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(findTilt(root) == 1)

	root1 := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 9,
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(findTilt(root1) == 15)

	root2 := &tree.TreeNode{
		Val: 21,
		Left: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 3,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			Right: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 14,
			Left: &tree.TreeNode{
				Val: 2,
			},
			Right: &tree.TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(findTilt(root2) == 9)
}

func findTilt(root *tree.TreeNode) int {
	return 0
}
