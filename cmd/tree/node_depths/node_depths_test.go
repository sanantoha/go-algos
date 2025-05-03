package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

var exp int = 16

func TestNodeDepthsRec(t *testing.T) {
	root := createTreeNode()

	res := nodeDepthsRec(root)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestNodeDepths(t *testing.T) {
	root := createTreeNode()

	res := nodeDepths(root)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func createTreeNode() *tree.TreeNode {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 8,
				},
				Right: &tree.TreeNode{
					Val: 9,
				},
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

	return root
}
