package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func isBalanced(root *tree.TreeNode) bool {
	return false
}

func main() {

	root1 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 9,
		},
		Right: &tree.TreeNode{
			Val: 20,
			Left: &tree.TreeNode{
				Val: 15,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	root2 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 4,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isBalanced(root1))
	fmt.Println(!isBalanced(root2))
}
