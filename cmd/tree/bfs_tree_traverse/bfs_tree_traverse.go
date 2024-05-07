package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(n) space - consider if deque is O(1), for simplicity use slice for queue impl.
func bfs(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)

	res := make([]int, 0)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // O(n) - complexity

		if curr == nil {
			continue
		}
		res = append(res, curr.Val)

		queue = append(queue, curr.Left)
		queue = append(queue, curr.Right)
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
			},
			Right: &tree.TreeNode{
				Val: 15,
				Left: &tree.TreeNode{
					Val: 13,
				},
				Right: &tree.TreeNode{
					Val: 17,
				},
			},
		},
	}

	// 5, 2, 10, 1, 3, 7, 15, 13, 17
	fmt.Println(bfs(root))
}
