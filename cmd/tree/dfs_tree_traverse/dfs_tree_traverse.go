package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func preOrder(root *tree.TreeNode) []int {

	stack := make([]*tree.TreeNode, 1)
	stack = append(stack, root)

	res := make([]int, 0)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == nil {
			continue
		}

		res = append(res, curr.Val)

		stack = append(stack, curr.Right)
		stack = append(stack, curr.Left)
	}

	return res
}

// O(n) time | O(h) space
func inOrder(root *tree.TreeNode) []int {

	res := make([]int, 0)

	stack := make([]*tree.TreeNode, 0)
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, curr.Val)

		curr = curr.Right
	}

	return res
}

// O(n) time | O(h) space
func postOrder(root *tree.TreeNode) []int {

	res := make([]int, 0)

	fst := make([]*tree.TreeNode, 1)
	fst = append(fst, root)
	snd := make([]*tree.TreeNode, 0)

	for len(fst) > 0 {
		curr := fst[len(fst)-1]
		fst = fst[:len(fst)-1]
		if curr == nil {
			continue
		}

		snd = append(snd, curr)

		fst = append(fst, curr.Left)
		fst = append(fst, curr.Right)
	}

	for len(snd) > 0 {
		curr := snd[len(snd)-1]
		snd = snd[:len(snd)-1]
		res = append(res, curr.Val)
	}

	return res
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
			Val: 8,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(preOrder(root)) // 5 2 1 3 8 7 9

	fmt.Println(inOrder(root)) // 1 2 3 5 7 8 9

	fmt.Println(postOrder(root)) // 1 3 2 7 9 8 5
}
