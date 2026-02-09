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
		"min_number_of_jumps.go":      true,
		"intersection_linked_list.go": true,
		"find_mode_in_bst.go":         true,
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
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(findMode(root)) // [2]

	root1 := &tree.TreeNode{
		Val: 0,
	}

	fmt.Println(findMode(root1)) // [0]

	root2 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 5,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(findMode(root2)) // [3, 5, 7]
}

func findMode(root *tree.TreeNode) []int {
	return nil
}
