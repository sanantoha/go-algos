package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {

	tests := []struct {
		name  string
		input *tree.TreeNode
		exp   int
	}{
		{
			name: "first tree",
			input: &tree.TreeNode{
				Val: 4,
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
					Val: 6,
				},
			},
			exp: 1,
		},
		{
			name: "second tree",
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 0,
				},
				Right: &tree.TreeNode{
					Val: 48,
					Left: &tree.TreeNode{
						Val: 12,
					},
					Right: &tree.TreeNode{
						Val: 50,
					},
				},
			},
			exp: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getMinimumDifference(tt.input)

			if tt.exp != res {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
