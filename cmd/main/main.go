package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
	"github.com/sanantoha/go-algos/internals/tree"
)

func main() {
	head := &list.ListNode{Val: 1}

	root := &tree.TreeNode{1, &tree.TreeNode{2, nil, nil}, &tree.TreeNode{3, &tree.TreeNode{4, nil, nil}, &tree.TreeNode{5, nil, nil}}}

	fmt.Println(head)
	fmt.Println(root)
}
