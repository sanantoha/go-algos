package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func maxPathSum(root *tree.TreeNode) int {
	return helper(root).maxPathSum
}

func helper(root *tree.TreeNode) Info {
	if root == nil {
		return Info{0, 0}
	}

	li := helper(root.Left)
	ri := helper(root.Right)

	maxPathSumChildAsBranch := max(li.maxPathSumAsBranch, ri.maxPathSumAsBranch)
	maxPathSumAsBranch := max(maxPathSumChildAsBranch+root.Val, root.Val)

	maxPathSumAsRootNode := max(maxPathSumAsBranch, li.maxPathSumAsBranch+root.Val+ri.maxPathSumAsBranch)
	maxPathSum := max(li.maxPathSum, ri.maxPathSum, maxPathSumAsRootNode)

	return Info{maxPathSum: maxPathSum, maxPathSumAsBranch: maxPathSumAsBranch}
}

type Info struct {
	maxPathSum         int
	maxPathSumAsBranch int
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	actual := maxPathSum(root)
	fmt.Println(actual)
	fmt.Println(actual == 18)
}
