package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	testTable := []struct {
		name  string
		input *tree.TreeNode
		exp   []int
	}{
		{
			name: "simple tree",
			input: &tree.TreeNode{
				Val: 1,
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 2,
					},
				},
			},
			exp: []int{2},
		},
		{
			name: "one node tree",
			input: &tree.TreeNode{
				Val: 0,
			},
			exp: []int{0},
		},
		{
			name: "happy path",
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 3,
					Left: &tree.TreeNode{
						Val: 1,
					},
					Right: &tree.TreeNode{
						Val: 3,
					},
				},
				Right: &tree.TreeNode{
					Val: 7,
					Left: &tree.TreeNode{
						Val: 5,
					},
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			exp: []int{3, 5, 7},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := findMode(subtest.input)

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
