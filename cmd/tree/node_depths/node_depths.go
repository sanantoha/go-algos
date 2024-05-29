package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func nodeDepthsRec(root *tree.TreeNode) int {
	return helper(root, 0)
}

func helper(root *tree.TreeNode, depth int) int {
	if root == nil {
		return 0
	}

	return depth + helper(root.Left, depth+1) + helper(root.Right, depth+1)
}

// O(n) time | O(h) space
func nodeDepths(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]Info, 1)
	stack[0] = Info{node: root, depth: 0}

	res := 0

	for len(stack) > 0 {
		info := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr := info.node
		depth := info.depth

		if curr == nil {
			continue
		}

		res += depth

		stack = append(stack, Info{node: curr.Left, depth: depth + 1})
		stack = append(stack, Info{node: curr.Right, depth: depth + 1})
	}
	return res
}

type Info struct {
	node  *tree.TreeNode
	depth int
}

/**
 *  Write function that takes in a Binary Tree and return the sum of its node's depths.
 *  https://www.algoexpert.io/questions/Node%20Depths
 *            1
 *          /   \
 *         2     3
 *       /  \  /  \
 *     4    5 6    7
 *   / \
 *  8   9
 */
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

	actual := nodeDepthsRec(root)
	fmt.Println(actual)
	fmt.Println(actual == 16)

	actual = nodeDepths(root)
	fmt.Println(actual)
	fmt.Println(actual == 16)
}
