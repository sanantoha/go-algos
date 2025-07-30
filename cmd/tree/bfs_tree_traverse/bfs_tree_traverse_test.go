package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {

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
					Val: 13,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	exp := []int{5, 2, 10, 1, 3, 7, 15, 13, 17}

	res := bfs(root)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
