package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func isSymmetricRec(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}
	return helper(root.Left, root.Right)
}

func helper(left *tree.TreeNode, right *tree.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

// O(n) time | O(h) space
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}

	stack := make([]*tree.TreeNode, 2)
	stack[0] = root.Left
	stack[1] = root.Right

	for len(stack) > 0 {
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := stack[len((stack))-1]
		stack = stack[:len(stack)-1]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left)
		stack = append(stack, right.Right)
		stack = append(stack, left.Right)
		stack = append(stack, right.Left)
	}

	return true
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(isSymmetricRec(root))
	fmt.Println(isSymmetric(root))

	fmt.Println("==================================")

	root1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(!isSymmetricRec(root1))
	fmt.Println(!isSymmetric(root1))

	fmt.Println("==================================")

	root2 := &tree.TreeNode{
		Val: 2,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(!isSymmetricRec(root2))
	fmt.Println(!isSymmetric(root2))
}
