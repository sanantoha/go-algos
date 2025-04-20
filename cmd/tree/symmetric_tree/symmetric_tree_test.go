package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestSymmetricTreeRec(t *testing.T) {
	root := createSymmetricTreeNode()

	res := isSymmetricRec(root)
	if !res {
		t.Errorf("expected true, but return %v", res)
	}
}

func TestSymmetricTreeRecForNonSymmetricTree(t *testing.T) {
	root := createNonSymmetricTreeNode()

	res := isSymmetricRec(root)
	if res {
		t.Errorf("expected false, but return %v", res)
	}
}

func TestSymmetricTreeRecForSimpleNonSymmetricTree(t *testing.T) {
	root := createNonSymmetricSimpleTreeNode()

	res := isSymmetricRec(root)
	if res {
		t.Errorf("expected false, but return %v", res)
	}
}

func TestSymmetricTree(t *testing.T) {
	root := createSymmetricTreeNode()

	res := isSymmetric(root)
	if !res {
		t.Errorf("expected true, but return %v", res)
	}
}

func TestSymmetricTreeForNonSymmetricTree(t *testing.T) {
	root := createNonSymmetricTreeNode()

	res := isSymmetric(root)
	if res {
		t.Errorf("expected false, but return %v", res)
	}
}

func TestSymmetricTreeForSimpleNonSymmetricTree(t *testing.T) {
	root := createNonSymmetricSimpleTreeNode()

	res := isSymmetric(root)
	if res {
		t.Errorf("expected false, but return %v", res)
	}
}

func createSymmetricTreeNode() *tree.TreeNode {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	return root
}

func createNonSymmetricTreeNode() *tree.TreeNode {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}

	return root
}

func createNonSymmetricSimpleTreeNode() *tree.TreeNode {
	root := &tree.TreeNode{
		Val: 2,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	return root
}
