package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
	log "github.com/sirupsen/logrus"
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
		"symmetric_tree.go":                             true,
		"prim_min_spanning_tree.go":                     true,
		"min_swaps_transform_string_to_palindrome.go":   true,
		"lru_cache.go":                                  true,
		"heap_sort.go":                                  true,
		"reverse_list.go":                               true,
		"levenshtein_distance.go":                       true,
		"bst_successor_search.go":                       true,
		"all_paths_from_source_target.go":               true,
		"level_order_binary_tree_traverse.go":           true,
		"all_elements_in_two_binary_search_trees.go":    true,
		"subarray_sort.go":                              true,
		"search_in_rotated_sorted_array.go":             true,
		"branch_sums.go":                                true,
		"sqrt.go":                                       true,
		"populating_next_right_pointer_in_each_node.go": true,
		"disk_stacking.go":                              true,
		"max_sum_increasing_subsequence.go":             true,
		"max_depth_of_bst.go":                           true,
		"knapsack_problem.go":                           true,
		"three_sum.go":                                  true,
		"find_nodes_distance_k.go":                      true,
		"dijkstra_shortest_path_as_map.go":              true,
		"intersection_linked_list.go":                   true,
		"depth_first_search.go":                         true,
		"bfs_tree_traverse.go":                          true,
		"binary_search.go":                              true,
		"find_pivot_index.go":                           true,
		"longest_decr_subseq.go":                        true,
		"median_of_two_sorted_arrays.go":                true,
		"dfs_tree_traverse.go":                          true,
		"binary_tree_zigzag_level_order_traverse.go":    true,
		"regular_expressions.go":                        true,
		"dfs_tree_traverse_rec.go":                      true,
		"spiral_matrix_traverse.go":                     true,
		"stable_internships.go":                         true,
		"merge_two_sorted_list.go":                      true,
		"breadth_first_search.go":                       true,
		"validate_starting_city.go":                     true,
		"string_permutations.go":                        true,
		"balanced_binary_tree.go":                       true,
		"generate_parentheses.go":                       true,
		"deep_copy_arbitrary_pointer.go":                true,
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

	root := &list.ArbitraryListNode{Val: 1}
	second := &list.ArbitraryListNode{Val: 2}
	third := &list.ArbitraryListNode{Val: 3}
	four := &list.ArbitraryListNode{Val: 4}
	five := &list.ArbitraryListNode{Val: 5}

	root.Next = second
	second.Next = third
	third.Next = four
	four.Next = five

	second.Arbitrary = five
	third.Arbitrary = root
	five.Arbitrary = second

	copyNode := deepCopy(root)

	assertArbitraryListNode(root, copyNode)
	fmt.Println(true)
}

func deepCopy(root *list.ArbitraryListNode) *list.ArbitraryListNode {
	return nil
}

func assertArbitraryListNode(root *list.ArbitraryListNode, copy *list.ArbitraryListNode) {
	c1 := root
	c2 := copy

	for c1 != nil && c2 != nil {
		if c1.Val != c2.Val {
			log.Fatalf("%d != %d", c1.Val, c2.Val)
		}
		if c1 == c2 {
			log.Fatalf("%v == %v", c1, c2)
		}
		if (c1.Arbitrary == nil && c2.Arbitrary != nil) || (c1.Arbitrary != nil && c2.Arbitrary == nil) {
			log.Fatalf("%v != %v", c1.Arbitrary, c2.Arbitrary)
		}
		if c1.Arbitrary != nil && c1.Arbitrary == c2.Arbitrary {
			log.Fatalf("%v == %v", c1.Arbitrary, c2.Arbitrary)
		}

		c1 = c1.Next
		c2 = c2.Next
	}

	if c1 != nil || c2 != nil {
		log.Fatalf("%v != %v", c1, c2)
	}
}
