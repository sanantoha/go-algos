package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(n) time | O(h) space
func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}
	isBalanced, _ := helper(root)
	return isBalanced
}

func helper(root *tree.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	lb, lh := helper(root.Left)
	rb, rh := helper(root.Right)

	height := int(math.Max(float64(lh), float64(rh))) + 1
	isBalanced := lb && rb && math.Abs(float64(lh-rh)) <= 1

	return isBalanced, height
}

func main() {

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

	root2 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 4,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isBalanced(root1))
	fmt.Println(!isBalanced(root2))
}
