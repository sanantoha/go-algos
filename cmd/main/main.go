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
		"symmetric_tree.go":                           true,
		"prim_min_spanning_tree.go":                   true,
		"min_swaps_transform_string_to_palindrome.go": true,
		"lru_cache.go":                                true,
		"heap_sort.go":                                true,
		"reverse_list.go":                             true,
		"levenshtein_distance.go":                     true,
		"bst_successor_search.go":                     true,
		"all_paths_from_source_target.go":             true,
		"level_order_binary_tree_traverse.go":         true,
		"all_elements_in_two_binary_search_trees.go":  true,
		"subarray_sort.go":                            true,
		"search_in_rotated_sorted_array.go":           true,
		"branch_sums.go":                              true,
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

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 8,
				},
				Right: &tree.TreeNode{
					Val: 9,
				},
			},
			Right: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 10,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(branchSums(root))     // [15, 16, 18, 10, 11]
	fmt.Println(branchSumsIter(root)) // [11, 10, 18, 16, 15]
}

func branchSums(root *tree.TreeNode) []int {
	return nil
}

func branchSumsIter(root *tree.TreeNode) []int {
	return nil
}
