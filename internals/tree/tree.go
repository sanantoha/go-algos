package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return "nil"
	}
	return fmt.Sprintf("TreeNode{Val=%d, Left=%s, Right=%s}", tree.Val, tree.Left.String(), tree.Right.String())
}
