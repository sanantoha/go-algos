package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func branchSums(root *tree.TreeNode) []int {
	res := make([]int, 0)
	backtrack(root, 0, &res)
	return res
}

func backtrack(root *tree.TreeNode, sum int, res *[]int) {
	if root == nil {
		return
	}

	sum += root.Val
	if root.Left == nil && root.Right == nil {
		*res = append(*res, sum)
	}

	backtrack(root.Left, sum, res)
	backtrack(root.Right, sum, res)
}

func branchSumsIter(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	stack := make([]Info, 0)
	stack = append(stack, Info{node: root, sum: 0})

	for len(stack) > 0 {
		info := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if info.node == nil {
			continue
		}

		curr := info.node
		sum := info.sum
		sum += curr.Val
		if curr.Left == nil && curr.Right == nil {
			res = append(res, sum)
		}

		stack = append(stack, Info{node: curr.Left, sum: sum})
		stack = append(stack, Info{node: curr.Right, sum: sum})
	}

	return res
}

type Info struct {
	node *tree.TreeNode
	sum  int
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
