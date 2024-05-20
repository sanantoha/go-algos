package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func maxDepth(root *tree.TreeNode) int {
	return -1
}

func maxDepthIter(root *tree.TreeNode) int {
	return -1
}

func maxDepthBfs(root *tree.TreeNode) int {
	return -1
}

func main() {

	root := &tree.TreeNode{
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
	}

	// 4
	fmt.Println(maxDepth(root))
	fmt.Println(maxDepthIter(root))
	fmt.Println(maxDepthBfs(root))

	fmt.Println("==================================")

	root1 := &tree.TreeNode{
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
	}

	// 3
	fmt.Println(maxDepth(root1))
	fmt.Println(maxDepthIter(root1))
	fmt.Println(maxDepthBfs(root1))

	fmt.Println("==================================")

	root2 := &tree.TreeNode{
		Val: 1,
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	// 2
	fmt.Println(maxDepth(root2))
	fmt.Println(maxDepthIter(root2))
	fmt.Println(maxDepthBfs(root2))
}
