package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func maxPathSum(root *tree.TreeNode) int {
	return -1
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 5,
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

	actual := maxPathSum(root)
	fmt.Println(actual)
	fmt.Println(actual == 18)
}
