package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func evaluateExpressionTree(root *tree.TreeNode) int {
	return -1
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
