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
			Val: 10,
			Left: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 6,
				},
				Right: &tree.TreeNode{
					Val: 8,
				},
			},
			Right: &tree.TreeNode{
				Val: 15,
				Left: &tree.TreeNode{
					Val: 14,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	// [[5], [2, 10], [1, 3, 7, 15], [6, 8, 14, 17]]
	fmt.Println(levelOrder(root))
}

func levelOrder(root *tree.TreeNode) [][]int {
	return nil
}
