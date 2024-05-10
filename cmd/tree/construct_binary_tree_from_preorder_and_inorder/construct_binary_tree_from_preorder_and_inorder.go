package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	return nil
}

func main() {

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	fmt.Println(root)
}
