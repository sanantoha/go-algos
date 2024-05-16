package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func leftView(root *tree.TreeNode) []int {
	return nil
}

func main() {

	/**
	 * left view of binary tree
	 *          1
	 *        /   \
	 *       2     3
	 *           /  \
	 *          4    8
	 *        /  \
	 *       5    6
	 *             \
	 *              7
	 *  output: [1, 2, 4, 5, 7]
	 */
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 5,
				},
				Right: &tree.TreeNode{
					Val: 6,
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 8,
			},
		},
	}

	// 1 2 4 5 7
	fmt.Println(leftView(root))
}
