package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(n) time | O(1) space
func getMinimumDifference(root *tree.TreeNode) int {
	if root == nil {
		return -1
	}

	minDiff := math.MaxInt
	prev := math.MaxInt

	stack := make([]*tree.TreeNode, 0)
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		minDiff = min(minDiff, abs(curr.Val-prev))

		prev = curr.Val
		curr = curr.Right
	}
	return minDiff
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {

	root := &tree.TreeNode{
		Val: 4,
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
			Val: 6,
		},
	}

	fmt.Println(getMinimumDifference(root)) // 1

	root1 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 0,
		},
		Right: &tree.TreeNode{
			Val: 48,
			Left: &tree.TreeNode{
				Val: 12,
			},
			Right: &tree.TreeNode{
				Val: 50,
			},
		},
	}

	fmt.Println(getMinimumDifference(root1)) // 2
}
