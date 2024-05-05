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
// O(m + n) time | O(m + n) space
func getAllElements(root1 *tree.TreeNode, root2 *tree.TreeNode) []int {

	stack1 := make([]*tree.TreeNode, 0)
	curr1 := root1
	stack2 := make([]*tree.TreeNode, 0)
	curr2 := root2

	res := make([]int, 0)

	for len(stack1) > 0 || curr1 != nil || len(stack2) > 0 || curr2 != nil {
		for curr1 != nil {
			stack1 = append(stack1, curr1)
			curr1 = curr1.Left
		}

		for curr2 != nil {
			stack2 = append(stack2, curr2)
			curr2 = curr2.Left
		}

		if len(stack2) == 0 || (len(stack1) > 0 && stack1[len(stack1)-1].Val < stack2[len(stack2)-1].Val) {
			curr1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]

			res = append(res, curr1.Val)

			curr1 = curr1.Right
		} else {
			curr2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]

			res = append(res, curr2.Val)

			curr2 = curr2.Right
		}
	}

	return res
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
