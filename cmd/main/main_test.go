package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

var ROOT *tree.TreeNode = &tree.TreeNode{
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
		Val: 8,
		Left: &tree.TreeNode{
			Val: 7,
		},
		Right: &tree.TreeNode{
			Val: 9,
		},
	},
}

func TestPreOrder(t *testing.T) {
	testTable := []struct {
		name   string
		tree   *tree.TreeNode
		extRes []int
	}{
		{
			name:   "happy path",
			tree:   ROOT,
			extRes: []int{5, 2, 1, 3, 8, 7, 9},
		},
		{
			name:   "nil value",
			tree:   nil,
			extRes: []int{},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := preOrder(subtest.tree)
			if !reflect.DeepEqual(subtest.extRes, res) {
				t.Errorf("expected res (%v), got (%v)", subtest.extRes, res)
			}
		})
	}
}

func TestInOrder(t *testing.T) {
	testTable := []struct {
		name   string
		tree   *tree.TreeNode
		expRes []int
	}{
		{
			name:   "happy path",
			tree:   ROOT,
			expRes: []int{1, 2, 3, 5, 7, 8, 9},
		},
		{
			name:   "nil value",
			tree:   nil,
			expRes: []int{},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := inOrder(subtest.tree)

			if !reflect.DeepEqual(subtest.expRes, res) {
				t.Errorf("expected res (%v), got (%v)", subtest.expRes, res)
			}
		})
	}
}

func TestPostOrder(t *testing.T) {
	testTable := []struct {
		name   string
		tree   *tree.TreeNode
		expRes []int
	}{
		{
			name:   "happy path",
			tree:   ROOT,
			expRes: []int{1, 3, 2, 7, 9, 8, 5},
		},
		{
			name:   "nil value",
			tree:   nil,
			expRes: []int{},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := postOrder(subtest.tree)

			if !reflect.DeepEqual(subtest.expRes, res) {
				t.Errorf("expected %v, but got %v", subtest.expRes, res)
			}
		})
	}
}
