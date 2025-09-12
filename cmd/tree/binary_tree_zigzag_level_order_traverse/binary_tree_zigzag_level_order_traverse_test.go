package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestZigZag(t *testing.T) {
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
				Left: &tree.TreeNode{
					Val: 14,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	exp := [][]int{{5}, {10, 2}, {1, 3, 7, 15}, {17, 14}}

	res := zigZag(root)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
