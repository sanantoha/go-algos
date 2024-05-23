package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func lowestCommonAncestor(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func main() {

	root := &tree.TreeNode{
		Val: 6,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 0,
			},
			Right: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 3,
				},
				Right: &tree.TreeNode{
					Val: 5,
				},
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

	fmt.Println(lowestCommonAncestor(root, &tree.TreeNode{Val: 0}, &tree.TreeNode{Val: 5}))

	root1 := &tree.TreeNode{
		Val:  2,
		Left: &tree.TreeNode{Val: 1},
	}

	fmt.Println(lowestCommonAncestor(root1, &tree.TreeNode{Val: 2}, &tree.TreeNode{Val: 1}))
}
