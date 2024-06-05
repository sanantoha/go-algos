package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(n) time | O(h) space
func validate(root *tree.TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *tree.TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val < minVal || root.Val >= maxVal {
		return false
	}

	return helper(root.Left, minVal, root.Val) && helper(root.Right, root.Val, maxVal)
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
