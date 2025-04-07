package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestReverse(t *testing.T) {
	testReverse(t, reverse)
}

func TestReverseIter(t *testing.T) {
	testReverse(t, reverseIter)
}

func testReverse(t *testing.T, fn func(node *tree.TreeNode)) {
	testTable := []struct {
		name    string
		input   *tree.TreeNode
		expTree *tree.TreeNode
	}{
		{
			name: "happy path",
			input: &tree.TreeNode{
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
						Val: 15,
						Left: &tree.TreeNode{
							Val: 14,
						},
						Right: &tree.TreeNode{
							Val: 17,
						},
					},
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			expTree: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 10,
					Left: &tree.TreeNode{
						Val: 7,
					},
					Right: &tree.TreeNode{
						Val: 15,
						Left: &tree.TreeNode{
							Val: 17,
						},
						Right: &tree.TreeNode{
							Val: 14,
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			fn(subtest.input)

			if !sameTree(subtest.expTree, subtest.input) {
				t.Errorf("expected (%v), got (%v)", subtest.expTree, subtest.input)
			}
		})
	}
}

func sameTree(exp *tree.TreeNode, act *tree.TreeNode) bool {
	if exp == nil && act == nil {
		return true
	}
	if exp == nil || act == nil {
		return false
	}
	if exp.Val != act.Val {
		return false
	}
	return sameTree(exp.Left, act.Left) && sameTree(exp.Right, act.Right)
}
