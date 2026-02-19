package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/sanantoha/go-algos/internals/tree"
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
		"min_number_of_jumps.go":                             true,
		"intersection_linked_list.go":                        true,
		"find_mode_in_bst.go":                                true,
		"heap_sort.go":                                       true,
		"odd_even_linked_list.go":                            true,
		"subarray_sort.go":                                   true,
		"rotting_oranges.go":                                 true,
		"construct_binary_tree_from_preorder_and_inorder.go": true,
		"all_elements_in_two_binary_search_trees.go":         true,
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

	root1 := &tree.TreeNode{
		Val: 9,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 12,
			Left: &tree.TreeNode{
				Val: 11,
			},
			Right: &tree.TreeNode{
				Val: 15,
			},
		},
	}

	root2 := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
		Right: &tree.TreeNode{
			Val: 16,
		},
	}

	expected := []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 15, 16}

	res := getAllElements(root1, root2)

	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(expected, res))
}

func getAllElements(root1, root2 *tree.TreeNode) []int {
	return nil
}
