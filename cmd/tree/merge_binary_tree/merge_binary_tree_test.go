package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestMergeBinaryTree(t *testing.T) {
	exp := expTree()
	tree1 := createTree1()
	tree2 := createTree2()

	actual := mergeBinaryTrees(tree1, tree2)

	if !reflect.DeepEqual(exp, actual) {
		t.Errorf("expected %v, but got %v", exp, actual)
	}
}

func TestMergeBinaryTreeIter(t *testing.T) {
	exp := expTree()
	tree1 := createTree1()
	tree2 := createTree2()

	actual := mergeBinaryTreesIter(tree1, tree2)

	if !reflect.DeepEqual(exp, actual) {
		t.Errorf("expected %v, but got %v", exp, actual)
	}
}

func expTree() *tree.TreeNode {
	t := &tree.TreeNode{Val: 2}
	t.Left = &tree.TreeNode{Val: 8}
	t.Right = &tree.TreeNode{Val: 11}
	t.Left.Left = &tree.TreeNode{Val: 9}
	t.Left.Right = &tree.TreeNode{Val: 4}
	t.Right.Left = &tree.TreeNode{Val: 7}
	t.Right.Right = &tree.TreeNode{Val: 6}
	return t
}

func createTree2() *tree.TreeNode {
	return &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 2,
			},
		},
		Right: &tree.TreeNode{
			Val: 9,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
	}
}

func createTree1() *tree.TreeNode {
	return &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}
}
