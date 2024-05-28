package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space - where n and h for smallest tree
func mergeBinaryTrees(tree1, tree2 *tree.TreeNode) *tree.TreeNode {
	if tree1 == nil {
		return tree2
	}
	if tree2 == nil {
		return tree1
	}
	tree1.Val += tree2.Val
	tree1.Left = mergeBinaryTrees(tree1.Left, tree2.Left)
	tree1.Right = mergeBinaryTrees(tree1.Right, tree2.Right)
	return tree1
}

// O(n) time | O(h) space - where n and h for smallest tree
func mergeBinaryTreesIter(tree1, tree2 *tree.TreeNode) *tree.TreeNode {
	if tree1 == nil {
		return tree2
	}
	if tree2 == nil {
		return tree1
	}

	stack1 := make([]*tree.TreeNode, 1)
	stack1[0] = tree1
	stack2 := make([]*tree.TreeNode, 1)
	stack2[0] = tree2

	for len(stack1) > 0 {
		curr1 := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		curr2 := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]

		if curr2 == nil {
			continue
		}
		curr1.Val += curr2.Val

		if curr1.Left == nil {
			curr1.Left = curr2.Left
		} else {
			stack1 = append(stack1, curr1.Left)
			stack2 = append(stack2, curr2.Left)
		}

		if curr1.Right == nil {
			curr1.Right = curr2.Right
		} else {
			stack1 = append(stack1, curr1.Right)
			stack2 = append(stack2, curr2.Right)
		}
	}

	return tree1
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
