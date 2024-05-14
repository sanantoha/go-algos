package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"sort"
)

func findNodesDistanceK(root *tree.TreeNode, target int, k int) []int {
	return nil
}

func findNodesDistanceKRec(root *tree.TreeNode, target int, k int) []int {
	return nil
}

func main() {

	root := &tree.TreeNode{Val: 1}
	root.Left = &tree.TreeNode{Val: 2}
	root.Right = &tree.TreeNode{Val: 3}
	root.Left.Left = &tree.TreeNode{Val: 4}
	root.Left.Right = &tree.TreeNode{Val: 5}
	root.Right.Right = &tree.TreeNode{Val: 6}
	root.Right.Right.Left = &tree.TreeNode{Val: 7}
	root.Right.Right.Right = &tree.TreeNode{Val: 8}

	target := 3
	k := 2

	expected := []int{2, 7, 8}

	actual := findNodesDistanceK(root, target, k)
	sort.Ints(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actual = findNodesDistanceKRec(root, target, k)
	sort.Ints(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}
