package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

/**
 * https://leetcode.com/problems/binary-tree-tilt/
 *
 * Given the root of a binary tree, return the sum of every tree node's tilt.
 *
 * The tilt of a tree node is the absolute difference between the sum of all left subtree node values and all right
 * subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0.
 * The rule is similar if the node does not have a right child.
 *
 *         1
 *       /  \
 *     2     3
 *
 * Input: root = [1,2,3]
 * Output: 1
 * Explanation:
 * Tilt of node 2 : |0-0| = 0 (no children)
 * Tilt of node 3 : |0-0| = 0 (no children)
 * Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
 * Sum of every tilt : 0 + 0 + 1 = 1
 *
 *
 *                  4
 *               /    \
 *             2       9
 *           /  \       \
 *         3     5       7
 *
 *
 * Input: root = [4,2,9,3,5,null,7]
 * Output: 15
 * Explanation:
 * Tilt of node 3 : |0-0| = 0 (no children)
 * Tilt of node 5 : |0-0| = 0 (no children)
 * Tilt of node 7 : |0-0| = 0 (no children)
 * Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
 * Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
 * Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
 * Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15
 */

// O(n) time | O(h) space
func findTilt(root *tree.TreeNode) int {
	return helper(root).tilt
}

func helper(root *tree.TreeNode) *Info {
	if root == nil {
		return &Info{sum: 0, tilt: 0}
	}

	li := helper(root.Left)
	ri := helper(root.Right)

	sum := li.sum + ri.sum + root.Val
	tilt := abs(li.sum-ri.sum) + li.tilt + ri.tilt

	return &Info{sum: sum, tilt: tilt}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Info struct {
	sum  int
	tilt int
}

func main() {

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(findTilt(root) == 1)

	root1 := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 9,
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(findTilt(root1) == 15)

	root2 := &tree.TreeNode{
		Val: 21,
		Left: &tree.TreeNode{
			Val: 7,
			Left: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 3,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			Right: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 14,
			Left: &tree.TreeNode{
				Val: 2,
			},
			Right: &tree.TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(findTilt(root2) == 9)
}
