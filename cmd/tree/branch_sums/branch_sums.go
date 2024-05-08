package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func branchSums(root *tree.TreeNode) []int {
	return nil
}

func branchSumsIter(root *tree.TreeNode) []int {
	return nil
}

func main() {

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

	fmt.Println(branchSums(root))     // [15, 16, 18, 10, 11]
	fmt.Println(branchSumsIter(root)) // [11, 10, 18, 16, 15]
}
