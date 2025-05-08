package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestListView(t *testing.T) {
	/**
	 * left view of binary tree
	 *          1
	 *        /   \
	 *       2     3
	 *           /  \
	 *          4    8
	 *        /  \
	 *       5    6
	 *             \
	 *              7
	 *  output: [1, 2, 4, 5, 7]
	 */
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 5,
				},
				Right: &tree.TreeNode{
					Val: 6,
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 8,
			},
		},
	}

	exp := []int{1, 2, 4, 5, 7}

	res := leftView(root)
	fmt.Println(res)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
