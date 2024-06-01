package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func reverse(root *tree.TreeNode) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left

	reverse(root.Left)
	reverse(root.Right)
}

// O(n) time | O(h) space
func reverseIter(root *tree.TreeNode) {
	if root == nil {
		return
	}

	stack := make([]*tree.TreeNode, 1)
	stack[0] = root

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == nil {
			continue
		}

		left := curr.Left
		right := curr.Right
		curr.Left = right
		curr.Right = left

		stack = append(stack, curr.Left)
		stack = append(stack, curr.Right)
	}
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
					Val: 14,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	fmt.Println(root)

	reverse(root)

	fmt.Println(root)
	fmt.Println(root.Val == 5)
	fmt.Println(root.Left.Val == 10)
	fmt.Println(root.Right.Val == 2)
	fmt.Println(root.Left.Left.Val == 15)
	fmt.Println(root.Left.Right.Val == 7)
	fmt.Println(root.Right.Left.Val == 3)
	fmt.Println(root.Right.Right.Val == 1)
	fmt.Println(root.Left.Left.Left.Val == 17)
	fmt.Println(root.Left.Left.Right.Val == 14)

	fmt.Println("==============================")

	root1 := &tree.TreeNode{
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
					Val: 14,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	fmt.Println(root1)

	reverseIter(root1)

	fmt.Println(root1)
	fmt.Println(root1.Val == 5)
	fmt.Println(root1.Left.Val == 10)
	fmt.Println(root1.Right.Val == 2)
	fmt.Println(root1.Left.Left.Val == 15)
	fmt.Println(root1.Left.Right.Val == 7)
	fmt.Println(root1.Right.Left.Val == 3)
	fmt.Println(root1.Right.Right.Val == 1)
	fmt.Println(root1.Left.Left.Left.Val == 17)
	fmt.Println(root1.Left.Left.Right.Val == 14)
}
