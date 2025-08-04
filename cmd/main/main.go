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
		"symmetric_tree.go": true,
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
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(isSymmetricRec(root))
	fmt.Println(isSymmetric(root))

	fmt.Println("==================================")

	root1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(!isSymmetricRec(root1))
	fmt.Println(!isSymmetric(root1))

	fmt.Println("==================================")

	root2 := &tree.TreeNode{
		Val: 2,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(!isSymmetricRec(root2))
	fmt.Println(!isSymmetric(root2))
}

func isSymmetricRec(root *tree.TreeNode) bool {
	return false
}

func isSymmetric(root *tree.TreeNode) bool {
	return false
}
