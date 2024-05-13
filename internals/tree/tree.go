package tree

import (
	"fmt"
	"strconv"
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

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func (tree *Node) String() string {
	if tree == nil {
		return "nil"
	}
	pv := ""
	if tree.Parent != nil {
		pv = strconv.Itoa(tree.Parent.Key)
	}
	return fmt.Sprintf("Node{Key=%d, Left=%s, Right=%s, Parent=%s}", tree.Key, tree.Left.String(), tree.Right.String(), pv)
}
