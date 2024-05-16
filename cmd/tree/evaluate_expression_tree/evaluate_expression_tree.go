package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func evaluateExpressionTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Val > 0 {
		return root.Val
	}

	lv := evaluateExpressionTree(root.Left)
	rv := evaluateExpressionTree(root.Right)

	if root.Val == -1 {
		return lv + rv
	} else if root.Val == -2 {
		return lv - rv
	} else if root.Val == -3 {
		return lv / rv
	} else if root.Val == -4 {
		return lv * rv
	}
	return 0
}

func main() {

	root := &tree.TreeNode{Val: -1}
	root.Left = &tree.TreeNode{Val: -2}
	root.Left.Left = &tree.TreeNode{Val: -4}
	root.Left.Right = &tree.TreeNode{Val: 2}
	root.Left.Left.Left = &tree.TreeNode{Val: 3}
	root.Left.Left.Right = &tree.TreeNode{Val: 2}

	root.Right = &tree.TreeNode{Val: -3}
	root.Right.Left = &tree.TreeNode{Val: 8}
	root.Right.Right = &tree.TreeNode{Val: 3}

	expected := 6

	actual := evaluateExpressionTree(root)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}
