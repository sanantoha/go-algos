package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"runtime"
	"testing"
)

func TestMaxDepth(t *testing.T) {

	funcs := []func(node *tree.TreeNode) int{
		maxDepth,
		maxDepthIter,
		maxDepthBfs,
	}

	tests := []struct {
		name  string
		input *tree.TreeNode
		exp   int
	}{
		{
			name: "tree with depth = 4",
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 1,
					},
				},
				Right: &tree.TreeNode{
					Val: 10,
					Left: &tree.TreeNode{
						Val: 7,
					},
					Right: &tree.TreeNode{
						Val: 20,
						Right: &tree.TreeNode{
							Val: 25,
						},
					},
				},
			},
			exp: 4,
		},
		{
			name: "tree with depth = 3",
			input: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 9,
				},
				Right: &tree.TreeNode{
					Val: 20,
					Left: &tree.TreeNode{
						Val: 15,
					},
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			exp: 3,
		},
		{
			name: "tree with depth = 2",
			input: &tree.TreeNode{
				Val: 1,
				Right: &tree.TreeNode{
					Val: 2,
				},
			},
			exp: 2,
		},
	}

	for _, tt := range tests {
		for _, fn := range funcs {
			pc := reflect.ValueOf(fn).Pointer()
			fun := runtime.FuncForPC(pc)

			testName := fmt.Sprintf("%s %s", fun.Name(), tt.name)

			t.Run(testName, func(t *testing.T) {
				res := fn(tt.input)

				if res != tt.exp {
					t.Errorf("expected %v, but got %v", tt.exp, res)
				}
			})
		}
	}
}
