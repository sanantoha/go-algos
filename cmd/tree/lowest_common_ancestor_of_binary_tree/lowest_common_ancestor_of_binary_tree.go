package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func lowestCommonAncestor(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	return nil
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

	fmt.Println(lowestCommonAncestor(root, node7, node5))
}
