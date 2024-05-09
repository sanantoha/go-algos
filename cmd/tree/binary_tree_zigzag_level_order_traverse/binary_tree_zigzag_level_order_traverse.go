package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func zigZag(root *tree.TreeNode) [][]int {

	queue := make([]*tree.TreeNode, 1)
	queue = append(queue, root)

	res := make([][]int, 0)
	idx := 0

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

		if idx%2 == 1 {
			subRes = reverse(subRes)
		}
		idx++
		if len(subRes) > 0 {
			res = append(res, subRes)
		}

	}
	return res
}

func reverse(arr []int) []int {
	l := 0
	r := len(arr) - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	return arr
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

	// [[5], [10, 2], [1, 3, 7, 15], [17, 14]]
	fmt.Println(zigZag(root))
}
