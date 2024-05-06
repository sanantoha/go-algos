package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func binaryTreeDiameter(root *tree.TreeNode) bool {
	return false
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 8,
					Left: &tree.TreeNode{
						Val: 9,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
				Right: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(binaryTreeDiameter(root))
}
