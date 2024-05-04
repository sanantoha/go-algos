package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
)

/**
 * https://leetcode.com/problems/all-elements-in-two-binary-search-trees/
 *
 * Given two binary search trees root1 and root2,
 * return a list containing all the integers from both trees sorted in ascending order.
 */
func getAllElements(root1 *tree.TreeNode, root2 *tree.TreeNode) []int {
	return nil
}

func main() {

	root1 := &tree.TreeNode{
		Val: 9,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 12,
			Left: &tree.TreeNode{
				Val: 11,
			},
			Right: &tree.TreeNode{
				Val: 15,
			},
		},
	}

	root2 := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
		Right: &tree.TreeNode{
			Val: 16,
		},
	}

	expected := []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 15, 16}

	res := getAllElements(root1, root2)

	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(expected, res))
}
