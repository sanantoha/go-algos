package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestSortedArrayToBst(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 8,
			Left: &tree.TreeNode{
				Val: 6,
				Right: &tree.TreeNode{
					Val: 7,
				},
			},
			Right: &tree.TreeNode{
				Val: 9,
				Right: &tree.TreeNode{
					Val: 10,
				},
			},
		},
	}

	res := sortedArrayToBst(arr)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
