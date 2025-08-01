package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestIsSubtree(t *testing.T) {
	root1 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 2,
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
		},
	}

	subTree := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(root1, subTree)) // true

	root2 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 0,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
		},
	}

	fmt.Println(isSubtree(root2, subTree)) // false

	root3 := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 4,
						Left: &tree.TreeNode{
							Val: 1,
						},
						Right: &tree.TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}

	fmt.Println(isSubtree(root3, subTree)) // true

	tests := []struct {
		name    string
		input   *tree.TreeNode
		subTree *tree.TreeNode
		exp     bool
	}{
		{
			name:    "happy path",
			input:   root1,
			subTree: subTree,
			exp:     true,
		},
		{
			name:    "not sub tree",
			input:   root2,
			subTree: subTree,
			exp:     false,
		},
		{
			name:    "happy path2",
			input:   root3,
			subTree: subTree,
			exp:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubtree(tt.input, tt.subTree)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
