package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testTable := []struct {
		name     string
		preorder []int
		inorder  []int
		expTree  *tree.TreeNode
	}{
		{
			name:     "happy path",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expTree: &tree.TreeNode{Val: 3,
				Left: &tree.TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &tree.TreeNode{Val: 20,
					Left:  &tree.TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &tree.TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
		},
		{
			name:     "handle nil",
			preorder: nil,
			inorder:  nil,
			expTree:  nil,
		},
		{
			name:     "different length",
			preorder: []int{1},
			inorder:  []int{1, 2, 3},
			expTree:  nil,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := buildTree(subtest.preorder, subtest.inorder)

			if !sameTree(subtest.expTree, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expTree, res)
			}
		})
	}
}

func sameTree(t1, t2 *tree.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return sameTree(t1.Left, t2.Left) && sameTree(t1.Right, t2.Right)
}
