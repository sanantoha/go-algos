package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
)

func reconstructBst(arr []int) *tree.TreeNode {
	return nil
}

func reconstructBst1(arr []int) *tree.TreeNode {
	return nil
}

func main() {

	preOrderTraversalValues := []int{10, 4, 2, 1, 3, 17, 19, 18}

	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 17,
			Right: &tree.TreeNode{
				Val: 19,
				Left: &tree.TreeNode{
					Val: 18,
				},
			},
		},
	}

	expected := getDfsOrder(root, []int{})

	actualTree := reconstructBst(preOrderTraversalValues)
	actual := getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actualTree = reconstructBst1(preOrderTraversalValues)
	actual = getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}

func getDfsOrder(root *tree.TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	arr = getDfsOrder(root.Left, arr)
	return getDfsOrder(root.Right, arr)
}
