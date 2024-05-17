package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(h) time | O(1) space
func findClosestValueInBst(root *tree.TreeNode, target int) int {
	if root == nil {
		return -1
	}

	curr := root
	closest := curr.Val

	for curr != nil {

		if math.Abs(float64(target-closest)) > math.Abs(float64(target-curr.Val)) {
			closest = curr.Val
		}

		if curr.Val > target {
			curr = curr.Left
		} else if curr.Val < target {
			curr = curr.Right
		} else {
			return curr.Val
		}
	}

	return closest
}

// O(n) time | O(h) space
func findClosestValueInBstRec(root *tree.TreeNode, target int) int {
	return helper(root, root.Val, target)
}

func helper(root *tree.TreeNode, closest, target int) int {
	if root == nil {
		return closest
	}

	if math.Abs(float64(target-root.Val)) < math.Abs(float64(target-closest)) {
		closest = root.Val
	}

	if root.Val > target {
		return helper(root.Left, closest, target)
	} else if root.Val < target {
		return helper(root.Right, closest, target)
	} else {
		return root.Val
	}
}

func main() {

	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 15,
			Left: &tree.TreeNode{
				Val: 13,
				Right: &tree.TreeNode{
					Val: 14,
				},
			},
			Right: &tree.TreeNode{
				Val: 22,
			},
		},
	}

	fmt.Println(findClosestValueInBst(root, 12)) // 13

	fmt.Println(findClosestValueInBstRec(root, 12)) // 13
}
