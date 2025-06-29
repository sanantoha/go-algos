package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestBranchSum(t *testing.T) {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 8,
				},
				Right: &tree.TreeNode{
					Val: 9,
				},
			},
			Right: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 10,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	exp := []int{10, 11, 15, 16, 18}

	funcs := []func(node *tree.TreeNode) []int{
		branchSums,
		branchSumsIter,
	}

	for i, fn := range funcs {
		t.Run("branchSums_"+strconv.Itoa(i), func(t *testing.T) {
			res := fn(root)
			slices.Sort(res)

			if !reflect.DeepEqual(exp, res) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
