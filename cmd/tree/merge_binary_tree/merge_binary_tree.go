package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func mergeBinaryTrees(tree1, tree2 *tree.TreeNode) *tree.TreeNode {
	return nil
}

func mergeBinaryTreesIter(tree1, tree2 *tree.TreeNode) *tree.TreeNode {
	return nil
}

func main() {

	tree1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	tree2 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 2,
			},
		},
		Right: &tree.TreeNode{
			Val: 9,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
	}

	actual := mergeBinaryTrees(tree1, tree2)

	fmt.Println(actual.Val == 2)
	fmt.Println(actual.Left.Val == 8)
	fmt.Println(actual.Left.Left.Val == 9)
	fmt.Println(actual.Left.Right.Val == 4)
	fmt.Println(actual.Right.Val == 11)
	fmt.Println(actual.Right.Left.Val == 7)
	fmt.Println(actual.Right.Right.Val == 6)

	fmt.Println("==========================================")

	tree3 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	tree4 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 2,
			},
		},
		Right: &tree.TreeNode{
			Val: 9,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
	}

	actual = mergeBinaryTreesIter(tree3, tree4)

	fmt.Println(actual.Val == 2)
	fmt.Println(actual.Left.Val == 8)
	fmt.Println(actual.Left.Left.Val == 9)
	fmt.Println(actual.Left.Right.Val == 4)
	fmt.Println(actual.Right.Val == 11)
	fmt.Println(actual.Right.Left.Val == 7)
	fmt.Println(actual.Right.Right.Val == 6)
}
