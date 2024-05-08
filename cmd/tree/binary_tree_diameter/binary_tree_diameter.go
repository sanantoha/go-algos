package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
)

// O(n) time | O(h) space
func binaryTreeDiameter(root *tree.TreeNode) int {
	diameter, _ := helper(root)
	return diameter
}

func helper(root *tree.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	ld, lh := helper(root.Left)
	rd, rh := helper(root.Right)

	height := int(math.Max(float64(lh), float64(rh))) + 1

	maxDiameter := math.Max(float64(ld), float64(rd))
	diameter := int(math.Max(maxDiameter, float64(lh+rh)))

	return diameter, height
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 8,
					Left: &tree.TreeNode{
						Val: 9,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
				Right: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}

	fmt.Println(binaryTreeDiameter(root))
	fmt.Println(binaryTreeDiameter(root) == 6)
}
