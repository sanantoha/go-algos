package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestBinaryTreeDiameter(t *testing.T) {
	exp := 6

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 8,
					Left: &tree.TreeNode{
						Val: 9,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
				Right: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	res := binaryTreeDiameter(root)

	if exp != res {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
