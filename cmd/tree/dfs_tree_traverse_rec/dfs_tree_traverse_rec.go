package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func preOrderRec(root *tree.TreeNode) []int {
	res := make([]int, 0)
	preOrderRecHelper(root, &res)
	return res
}

func preOrderRecHelper(root *tree.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrderRecHelper(root.Left, res)
	preOrderRecHelper(root.Right, res)
}

// O(n) time | O(h) space
func inOrderRec(root *tree.TreeNode) []int {
	res := make([]int, 0)
	inOrderRecHelper(root, &res)
	return res
}

func inOrderRecHelper(root *tree.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderRecHelper(root.Left, res)
	*res = append(*res, root.Val)
	inOrderRecHelper(root.Right, res)
}

// O(n) time | O(h) space
func postOrderRec(root *tree.TreeNode) []int {
	res := make([]int, 0)
	postOrderRecHelper(root, &res)
	return res
}

func postOrderRecHelper(root *tree.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postOrderRecHelper(root.Left, res)
	postOrderRecHelper(root.Right, res)
	*res = append(*res, root.Val)
}

func main() {

	root := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 8,
			Left: &tree.TreeNode{
				Val: 7,
			},
			Right: &tree.TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(preOrderRec(root)) // 5 2 1 3 8 7 9

	fmt.Println(inOrderRec(root)) // 1 2 3 5 7 8 9

	fmt.Println(postOrderRec(root)) // 1 3 2 7 9 8 5
}
