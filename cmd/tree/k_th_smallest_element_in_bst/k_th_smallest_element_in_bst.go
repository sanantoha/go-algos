package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(n) space
func kthSmallestElement(root *tree.TreeNode, k int) int {
	if root == nil {
		return -1
	}

	stack := make([]*tree.TreeNode, 0)
	curr := root

	idx := 1

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if idx == k {
			return curr.Val
		}
		idx++

		curr = curr.Right
	}

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

	fmt.Println(kthSmallestElement(root, 4)) // 5
}
