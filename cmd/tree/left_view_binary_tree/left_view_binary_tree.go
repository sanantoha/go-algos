package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func leftView(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	queue := make([]*tree.TreeNode, 1) // use slice for queue not optimal solution
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)

		isFirst := true

		for size > 0 {
			size--

			curr := queue[0]
			queue = queue[1:]

			if curr == nil {
				continue
			}

			if isFirst {
				res = append(res, curr.Val)
				isFirst = false
			}

			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}
	}

	return res
}

func main() {

	/**
	 * left view of binary tree
	 *          1
	 *        /   \
	 *       2     3
	 *           /  \
	 *          4    8
	 *        /  \
	 *       5    6
	 *             \
	 *              7
	 *  output: [1, 2, 4, 5, 7]
	 */
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 5,
				},
				Right: &tree.TreeNode{
					Val: 6,
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 8,
			},
		},
	}

	// 1 2 4 5 7
	fmt.Println(leftView(root))
}
