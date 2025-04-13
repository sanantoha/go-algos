package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestPreOrderRec(t *testing.T) {
	input := createTree()
	expected := []int{5, 2, 1, 3, 8, 7, 9}

	res := preOrderRec(input)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestInOrderRec(t *testing.T) {
	input := createTree()
	expected := []int{1, 2, 3, 5, 7, 8, 9}

	res := inOrderRec(input)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestPostOrderRec(t *testing.T) {
	input := createTree()
	expected := []int{1, 3, 2, 7, 9, 8, 5}

	res := postOrderRec(input)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func createTree() *tree.TreeNode {
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
			Val: 8,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 9,
			},
		},
	}

	return root
}
