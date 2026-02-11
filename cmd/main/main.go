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

	preOrderTraversalValues := []int{10, 4, 2, 1, 3, 17, 19, 18}

	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 17,
			Right: &tree.TreeNode{
				Val: 19,
				Left: &tree.TreeNode{
					Val: 18,
				},
			},
		},
	}

	expected := getDfsOrder(root, []int{})

	actualTree := reconstructBst(preOrderTraversalValues)
	actual := getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actualTree = reconstructBst1(preOrderTraversalValues)
	actual = getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}

func reconstructBst(preorder []int) *tree.TreeNode {
	return nil
}

func reconstructBst1(preorder []int) *tree.TreeNode {
	return nil
}

func getDfsOrder(root *tree.TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	arr = getDfsOrder(root.Left, arr)
	return getDfsOrder(root.Right, arr)
}
