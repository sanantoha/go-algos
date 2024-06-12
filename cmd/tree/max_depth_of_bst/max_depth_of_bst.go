package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func maxDepth(root *tree.TreeNode) int {
	return maxDepthRec(root, 0)
}

func maxDepthRec(root *tree.TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	return max(maxDepthRec(root.Left, depth), maxDepthRec(root.Right, depth))
}

type Info struct {
	node  *tree.TreeNode
	depth int
}

// O(n) time | O(h) space
func maxDepthIter(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]Info, 1)
	stack[0] = Info{
		node:  root,
		depth: 0,
	}

	maxDepthVal := 0

	for len(stack) > 0 {
		info := stack[0]
		stack = stack[1:]

		depth := info.depth
		curr := info.node

		maxDepthVal = max(depth, maxDepthVal)

		if curr == nil {
			continue
		}

		stack = append(stack, Info{node: curr.Left, depth: depth + 1})
		stack = append(stack, Info{node: curr.Right, depth: depth + 1})
	}

	return maxDepthVal
}

// O(n) time | O(n) space
func maxDepthBfs(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*tree.TreeNode, 1)
	queue[0] = root

	depth := 0

	for len(queue) > 0 {
		size := len(queue)

		depth++

		for size > 0 {
			size--

			curr := queue[0]
			queue = queue[1:]

			if curr == nil {
				continue
			}

			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}
	}

	return depth - 1
}

func main() {

	root := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 10,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 20,
				Right: &tree.TreeNode{
					Val: 25,
				},
			},
		},
	}

	// 4
	fmt.Println(maxDepth(root))
	fmt.Println(maxDepthIter(root))
	fmt.Println(maxDepthBfs(root))

	fmt.Println("==================================")

	root1 := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 9,
		},
		Right: &tree.TreeNode{
			Val: 20,
			Left: &tree.TreeNode{
				Val: 15,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	// 3
	fmt.Println(maxDepth(root1))
	fmt.Println(maxDepthIter(root1))
	fmt.Println(maxDepthBfs(root1))

	fmt.Println("==================================")

	root2 := &tree.TreeNode{
		Val: 1,
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	// 2
	fmt.Println(maxDepth(root2))
	fmt.Println(maxDepthIter(root2))
	fmt.Println(maxDepthBfs(root2))
}
