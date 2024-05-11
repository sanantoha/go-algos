package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func preOrder(root *tree.TreeNode) []int {
	return nil
}

func inOrder(root *tree.TreeNode) []int {
	return nil
}

func postOrder(root *tree.TreeNode) []int {
	return nil
}

func main() {

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
