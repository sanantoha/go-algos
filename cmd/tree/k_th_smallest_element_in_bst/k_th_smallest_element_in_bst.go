package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func kthSmallestElement(root *tree.TreeNode, k int) int {
	return -1
}

func main() {

	root := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val:   2,
			Left:  &tree.TreeNode{Val: 1},
			Right: &tree.TreeNode{Val: 3},
		},
		Right: &tree.TreeNode{
			Val: 10,
			Left: &tree.TreeNode{
				Val:   7,
				Left:  &tree.TreeNode{Val: 6},
				Right: &tree.TreeNode{Val: 8},
			},
			Right: &tree.TreeNode{
				Val:   15,
				Left:  &tree.TreeNode{Val: 14},
				Right: &tree.TreeNode{Val: 17},
			},
		},
	}

	fmt.Println(kthSmallestElement(root, 4))
}
