package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestValidate(t *testing.T) {

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

	res := validate(root)

	if !res {
		t.Errorf("expected true, but got false")
	}
}
