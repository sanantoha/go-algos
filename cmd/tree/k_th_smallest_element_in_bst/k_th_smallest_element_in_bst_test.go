package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestKthSmallestElement(t *testing.T) {

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

	exp := 5

	res := kthSmallestElement(root, 4)

	if res != exp {
		t.Errorf("kthSmallestElement returned %d, expected %d", res, exp)
	}
}
