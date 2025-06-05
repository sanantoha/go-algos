package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestGetAllElements(t *testing.T) {

	root1 := &tree.TreeNode{
		Val: 9,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 12,
			Left: &tree.TreeNode{
				Val: 11,
			},
			Right: &tree.TreeNode{
				Val: 15,
			},
		},
	}

	root2 := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
		Right: &tree.TreeNode{
			Val: 16,
		},
	}

	expected := []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 15, 16}

	res := getAllElements(root1, root2)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, but got %v", expected, res)
	}
}
