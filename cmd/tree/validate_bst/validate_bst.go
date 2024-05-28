package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func validate(root *tree.TreeNode) bool {
	return false
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
			Val: 10,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 15,
			},
		},
	}

	fmt.Println(validate(root))
}
