package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n * m) time | O(n) space
func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {

	stack := make([]*tree.TreeNode, 1)
	stack[0] = root

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == nil {
			continue
		}

		if isSameTree(curr, subRoot) {
			return true
		}

		stack = append(stack, curr.Right)
		stack = append(stack, curr.Left)
	}

	return false
}

func isSameTree(root1 *tree.TreeNode, root2 *tree.TreeNode) bool {

	stack := make([]*tree.TreeNode, 2)
	stack[0] = root1
	stack[1] = root2

	for len(stack) > 0 {
		curr1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr1 == nil && curr2 == nil {
			continue
		}
		if curr1 == nil || curr2 == nil {
			return false
		}
		if curr1.Val != curr2.Val {
			return false
		}

		stack = append(stack, curr1.Left)
		stack = append(stack, curr2.Left)
		stack = append(stack, curr1.Right)
		stack = append(stack, curr2.Right)
	}
	return true
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
