package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func isSymmetricRec(root *tree.TreeNode) bool {
	return false
}

func isSymmetric(root *tree.TreeNode) bool {
	return false
}

func main() {

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

	fmt.Println(isSymmetricRec(root))
	fmt.Println(isSymmetric(root))

	fmt.Println("==================================")

	root1 := &tree.TreeNode{
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

	fmt.Println(!isSymmetricRec(root1))
	fmt.Println(!isSymmetric(root1))

	fmt.Println("==================================")

	root2 := &tree.TreeNode{
		Val: 2,
		Left: &tree.TreeNode{
			Val: 1,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(!isSymmetricRec(root2))
	fmt.Println(!isSymmetric(root2))
}
