package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

var EXP_ROOT *tree.TreeNode = &tree.TreeNode{
	Val: 10,
	Left: &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	},
	Right: &tree.TreeNode{
		Val: 17,
		Right: &tree.TreeNode{
			Val: 19,
			Left: &tree.TreeNode{
				Val: 18,
			},
		},
	},
}

func TestReconstructBST(t *testing.T) {
	testTable := []struct {
		name     string
		preorder []int
		expTree  *tree.TreeNode
	}{
		{
			name:     "happy path",
			preorder: []int{10, 4, 2, 1, 3, 17, 19, 18},
			expTree:  EXP_ROOT,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := reconstructBst(subtest.preorder)
			actual := getDfsOrder(res, []int{})
			expected := getDfsOrder(subtest.expTree, []int{})
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("expected (%v), but got (%v)", expected, actual)
			}
		})
	}
}

func TestReconstructBST1(t *testing.T) {
	testTable := []struct {
		name     string
		preorder []int
		expTree  *tree.TreeNode
	}{
		{
			name:     "happy path",
			preorder: []int{10, 4, 2, 1, 3, 17, 19, 18},
			expTree:  EXP_ROOT,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := reconstructBst1(subtest.preorder)
			actual := getDfsOrder(res, []int{})
			expected := getDfsOrder(subtest.expTree, []int{})
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("expected (%v), but got (%v)", expected, actual)
			}
		})
	}
}
