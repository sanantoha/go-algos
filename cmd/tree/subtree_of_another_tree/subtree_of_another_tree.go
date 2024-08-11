package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(m * n) time | O(n) space
func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {
	return false
}

func isSameTree(root1 *tree.TreeNode, root2 *tree.TreeNode) bool {

	stack := make([]*tree.TreeNode, 2)
	stack[0] = root1
	stack[1] = root2

	for len(stack) > 0 {

	}
	return false
}

/**
 * https://leetcode.com/problems/subtree-of-another-tree/
 */
func main() {

	root1 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 2,
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
		},
	}

	subTree := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(root1, subTree)) // true

	root2 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 0,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
		},
	}

	fmt.Println(isSubtree(root2, subTree)) // false

	root3 := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 4,
						Left: &tree.TreeNode{
							Val: 1,
						},
						Right: &tree.TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}

	fmt.Println(isSubtree(root3, subTree)) // true
}
