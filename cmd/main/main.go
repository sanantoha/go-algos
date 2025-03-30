package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math/rand"
	"os"
	"path/filepath"
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
		"spiral_matrix_traverse.go":                  true,
		"binary_tree_zigzag_level_order_traverse.go": true,
		"lru_cache.go":                               true,
		"climbing_stars.go":                          true,
		"quick_sort.go":                              true,
		"maximum_subarray.go":                        true,
		"insert_sort.go":                             true,
		"first_duplicate_value.go":                   true,
		"permutations.go":                            true,
		"merge_intervals.go":                         true,
		"river_sizes.go":                             true,
		"bellman_ford_as_map.go":                     true,
		"add_two_numbers.go":                         true,
		"top_k_frequent_words.go":                    true,
		"find_nodes_distance_k.go":                   true,
		"kth_smallest_element.go":                    true,
		"clone_graph.go":                             true,
		"same_tree.go":                               true,
		"palindrome_linked_list.go":                  true,
		"string_permutations.go":                     true,
		"binary_search.go":                           true,
		"string_without_aaa_or_bbb.go":               true,
		"zig_zag_traverse.go":                        true,
		"depth_first_search.go":                      true,
		"levenshtein_distance.go":                    true,
		"best_time_to_buy_and_sell_stocks.go":        true,
		"subarray_sort.go":                           true,
		"dfs_tree_traverse.go":                       true,
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if _, exists := processedTasks[task]; exists {
			continue
		}
		fmt.Println(task)
	}

	runTask()
}

func runTask() {

	root := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 8,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(preOrder(root)) // 5 2 1 3 8 7 9

	fmt.Println(inOrder(root)) // 1 2 3 5 7 8 9

	fmt.Println(postOrder(root)) // 1 3 2 7 9 8 5
}

// O(n) time | O(h) space
func preOrder(root *tree.TreeNode) []int {

	stack := make([]*tree.TreeNode, 1)
	stack[0] = root

	res := make([]int, 0)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == nil {
			continue
		}

		res = append(res, curr.Val)

		stack = append(stack, curr.Right)
		stack = append(stack, curr.Left)
	}

	return res
}

// O(n) time | O(h) space
func inOrder(root *tree.TreeNode) []int {

	stack := make([]*tree.TreeNode, 0)
	curr := root

	res := make([]int, 0)

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)

		curr = curr.Right
	}

	return res
}

// O(n) time | O(h) space
func postOrder(root *tree.TreeNode) []int {

	res := make([]int, 0)

	fst := make([]*tree.TreeNode, 1)
	fst[0] = root
	snd := make([]*tree.TreeNode, 0)

	for len(fst) > 0 {
		curr := fst[len(fst)-1]
		fst = fst[:len(fst)-1]
		if curr == nil {
			continue
		}
		snd = append(snd, curr)

		fst = append(fst, curr.Left)
		fst = append(fst, curr.Right)
	}

	for len(snd) > 0 {
		v := snd[len(snd)-1]
		snd = snd[:len(snd)-1]
		res = append(res, v.Val)
	}

	return res
}
