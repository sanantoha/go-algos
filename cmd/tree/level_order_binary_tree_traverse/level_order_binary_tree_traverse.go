package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(n) space
func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*tree.TreeNode, 1)
	queue[0] = root

	res := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)

		subRes := make([]int, 0)

		for size > 0 {
			size--

			curr := queue[0]
			queue = queue[1:]

			if curr == nil {
				continue
			}

			subRes = append(subRes, curr.Val)

			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}

		if len(subRes) > 0 {
			res = append(res, subRes)
		}
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
			Val: 10,
			Left: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 6,
				},
				Right: &tree.TreeNode{
					Val: 8,
				},
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

	// [[5], [2, 10], [1, 3, 7, 15], [6, 8, 14, 17]]
	fmt.Println(levelOrder(root))
}
