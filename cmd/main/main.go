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

	root := &tree.TreeNode{Val: -1}
	root.Left = &tree.TreeNode{Val: -2}
	root.Left.Left = &tree.TreeNode{Val: -4}
	root.Left.Right = &tree.TreeNode{Val: 2}
	root.Left.Left.Left = &tree.TreeNode{Val: 3}
	root.Left.Left.Right = &tree.TreeNode{Val: 2}

	root.Right = &tree.TreeNode{Val: -3}
	root.Right.Left = &tree.TreeNode{Val: 8}
	root.Right.Right = &tree.TreeNode{Val: 3}

	expected := 6

	actual := evaluateExpressionTree(root)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}

func evaluateExpressionTree(root *tree.TreeNode) int {
	return 0
}
