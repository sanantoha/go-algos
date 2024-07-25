package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(n) time | O(1) space
func findMode(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*tree.TreeNode, 0)
	curr := root

	res := make([]int, 0)
	maxCnt := 0
	count := 0
	prev := math.MinInt

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev == curr.Val {
			count++
		} else {
			count = 1
		}

		if count == maxCnt {
			res = append(res, curr.Val)
		} else if count > maxCnt {
			maxCnt = max(maxCnt, count)

			res = make([]int, 0)
			res = append(res, curr.Val)
		}

		prev = curr.Val
		curr = curr.Right
	}

	return res
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(findMode(root)) // [2]

	root1 := &tree.TreeNode{
		Val: 0,
	}

	fmt.Println(findMode(root1)) // [0]

	root2 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 5,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(findMode(root2)) // [3, 5, 7]
}
