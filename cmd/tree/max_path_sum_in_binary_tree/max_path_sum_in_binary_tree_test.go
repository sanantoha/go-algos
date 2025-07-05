package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestMaxPathSum(t *testing.T) {

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

	exp := 18

	actual := maxPathSum(root)
	fmt.Println(actual)

	if actual != exp {
		t.Errorf("expected %v but got %v", exp, actual)
	}
}
