package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	node7 := &tree.TreeNode{Val: 7}
	node5 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 6,
		},
		Right: &tree.TreeNode{
			Val:  2,
			Left: node7,
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
	}

	node1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 0,
		},
		Right: &tree.TreeNode{
			Val: 8,
		},
	}

	root := &tree.TreeNode{
		Val:   3,
		Left:  node5,
		Right: node1,
	}

	res := lowestCommonAncestor(root, node7, node5)

	if !reflect.DeepEqual(node5, res) {
		t.Errorf("expected %v, but got %v", node5, res)
	}
}
