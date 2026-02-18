package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	// TreeNode{Val=3, Left=TreeNode{Val=9, Left=nil, Right=nil}, Right=TreeNode{Val=20, Left=TreeNode{Val=15, Left=nil, Right=nil}, Right=TreeNode{Val=7, Left=nil, Right=nil}}}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	return nil
}
