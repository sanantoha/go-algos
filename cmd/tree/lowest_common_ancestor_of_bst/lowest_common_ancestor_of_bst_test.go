package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {

	tests := []struct {
		name  string
		input *tree.TreeNode
		p     *tree.TreeNode
		q     *tree.TreeNode
		exp   *tree.TreeNode
	}{
		{
			name: "happy path 1",
			input: &tree.TreeNode{
				Val: 6,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 0,
					},
					Right: &tree.TreeNode{
						Val: 4,
						Left: &tree.TreeNode{
							Val: 3,
						},
						Right: &tree.TreeNode{
							Val: 5,
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 8,
					Left: &tree.TreeNode{
						Val: 7,
					},
					Right: &tree.TreeNode{
						Val: 9,
					},
				},
			},
			p: &tree.TreeNode{Val: 0},
			q: &tree.TreeNode{Val: 5},
			exp: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 0,
				},
				Right: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "happy path 2",
			input: &tree.TreeNode{
				Val:  2,
				Left: &tree.TreeNode{Val: 1},
			},
			p: &tree.TreeNode{Val: 2},
			q: &tree.TreeNode{Val: 1},
			exp: &tree.TreeNode{
				Val:  2,
				Left: &tree.TreeNode{Val: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lowestCommonAncestor(tt.input, tt.p, tt.q)

			if !reflect.DeepEqual(tt.exp, res) {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}

}
