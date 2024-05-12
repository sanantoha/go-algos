package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func preOrderRec(root *tree.TreeNode) []int {
	return nil
}

func inOrderRec(root *tree.TreeNode) []int {
	return nil
}

func postOrderRec(root *tree.TreeNode) []int {
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

	fmt.Println(preOrderRec(root)) // 5 2 1 3 8 7 9

	fmt.Println(inOrderRec(root)) // 1 2 3 5 7 8 9

	fmt.Println(postOrderRec(root)) // 1 3 2 7 9 8 5
}
