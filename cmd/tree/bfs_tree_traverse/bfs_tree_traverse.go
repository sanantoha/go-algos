package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func bfs(root *tree.TreeNode) []int {
	return nil
}

func main() {

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
			Val: 10,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 15,
				Left: &tree.TreeNode{
					Val: 13,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	// 5, 2, 10, 1, 3, 7, 15, 13, 17
	fmt.Println(bfs(root))
}
