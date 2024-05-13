package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func findClosestValueInBst(root *tree.TreeNode, target int) int {
	return -1
}

func findClosestValueInBstRec(root *tree.TreeNode, target int) int {
	return -1
}

func main() {

	root := &tree.TreeNode{
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

	fmt.Println(findClosestValueInBst(root, 12))

	fmt.Println(findClosestValueInBstRec(root, 12))
}
