package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

var root *tree.TreeNode = &tree.TreeNode{
	Val: 10,
	Left: &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
		},
	},
	Right: &tree.TreeNode{
		Val: 15,
		Left: &tree.TreeNode{
			Val: 13,
			Right: &tree.TreeNode{
				Val: 14,
			},
		},
		Right: &tree.TreeNode{
			Val: 22,
		},
	},
}
var target int = 12

var exp = 13

func TestFindClosestValueInBst(t *testing.T) {

	res := findClosestValueInBst(root, target)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestFindClosestValueInBstRec(t *testing.T) {

	res := findClosestValueInBstRec(root, target)

	if res != exp {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
