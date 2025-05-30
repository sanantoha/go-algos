package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	root1 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 9,
		},
		Right: &tree.TreeNode{
			Val: 20,
			Left: &tree.TreeNode{
				Val: 15,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	root2 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 4,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	tests := []struct {
		name   string
		input  *tree.TreeNode
		expRes bool
	}{
		{
			name:   "happy path",
			input:  root1,
			expRes: true,
		},
		{
			name:   "inbalanced tree",
			input:  root2,
			expRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isBalanced(tt.input)

			if tt.expRes != res {
				t.Errorf("expected %v, but got %v", tt.expRes, res)
			}
		})
	}
}
