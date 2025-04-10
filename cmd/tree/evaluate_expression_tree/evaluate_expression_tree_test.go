package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestEvaluateExpressionTree(t *testing.T) {
	root := &tree.TreeNode{Val: -1}
	root.Left = &tree.TreeNode{Val: -2}
	root.Left.Left = &tree.TreeNode{Val: -4}
	root.Left.Right = &tree.TreeNode{Val: 2}
	root.Left.Left.Left = &tree.TreeNode{Val: 3}
	root.Left.Left.Right = &tree.TreeNode{Val: 2}

	root.Right = &tree.TreeNode{Val: -3}
	root.Right.Left = &tree.TreeNode{Val: 8}
	root.Right.Right = &tree.TreeNode{Val: 3}

	exp := 6

	res := evaluateExpressionTree(root)

	if exp != res {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
