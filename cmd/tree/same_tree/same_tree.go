package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func isSameTree(t1, t2 *tree.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
}

// O(n) time | O(h) space
func isSameTreeIter(t1, t2 *tree.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	stack := make([]*tree.TreeNode, 2)
	stack[0] = t1
	stack[1] = t2

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

func main() {

	t1 := &tree.TreeNode{
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

	t2 := &tree.TreeNode{
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

	fmt.Println(isSameTree(t1, t2))
	fmt.Println(isSameTreeIter(t1, t2))

	fmt.Println("================================")

	t11 := &tree.TreeNode{
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

	t22 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 4,
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

	fmt.Println(!isSameTree(t11, t22))
	fmt.Println(!isSameTreeIter(t11, t22))
}
