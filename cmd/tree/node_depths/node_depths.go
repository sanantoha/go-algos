package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func nodeDepthsRec(root *tree.TreeNode) int {
	return -1
}

func nodeDepths(root *tree.TreeNode) int {
	return -1
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
	/*
		TreeNode root = new TreeNode(1,
		                new TreeNode(2,
		                        new TreeNode(4,
		                                new TreeNode(8),
		                                new TreeNode(9)),
		                        new TreeNode(5)),
		                new TreeNode(3,
		                        new TreeNode(6),
		                        new TreeNode(7)));

		        System.out.println(nodeDepths(root));
		        System.out.println(nodeDepths(root) == 16);
		        System.out.println(nodeDepthsRec(root));
		        System.out.println(nodeDepthsRec(root) == 16);
	*/

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
