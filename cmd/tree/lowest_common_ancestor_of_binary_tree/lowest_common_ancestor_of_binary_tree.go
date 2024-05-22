package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(h) space
func lowestCommonAncestor(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

/*
	      3
	   /     \
	  5       1
	/   \    /  \
   6     2  0    8
       / \
	  7   4
*/

func main() {

	node7 := &tree.TreeNode{Val: 7}
	node5 := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 6,
		},
		Right: &tree.TreeNode{
			Val:  2,
			Left: node7,
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
	}

	node1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 0,
		},
		Right: &tree.TreeNode{
			Val: 8,
		},
	}

	root := &tree.TreeNode{
		Val:   3,
		Left:  node5,
		Right: node1,
	}

	// TreeNode{Val=5, Left=TreeNode{Val=6, Left=nil, Right=nil}, Right=TreeNode{Val=2, Left=TreeNode{Val=7, Left=nil, Right=nil}, Right=TreeNode{Val=4, Left=nil, Right=nil}}}
	fmt.Println(lowestCommonAncestor(root, node7, node5))
}
