package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

type Index struct {
	idx int
}

// O(n) time | O(h) space
func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}

	inorderMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderMap[v] = i
	}
	index := &Index{idx: 0}
	return helper(preorder, inorderMap, index, 0, len(preorder)-1)
}

func helper(preorder []int, inorderMap map[int]int, index *Index, l, r int) *tree.TreeNode {
	if l > r {
		return nil
	}

	val := preorder[index.idx]
	index.idx++

	root := &tree.TreeNode{Val: val}
	root.Left = helper(preorder, inorderMap, index, l, inorderMap[val]-1)
	root.Right = helper(preorder, inorderMap, index, inorderMap[val]+1, r)

	return root
}

func main() {

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	// TreeNode{Val=3, Left=TreeNode{Val=9, Left=nil, Right=nil}, Right=TreeNode{Val=20, Left=TreeNode{Val=15, Left=nil, Right=nil}, Right=TreeNode{Val=7, Left=nil, Right=nil}}}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}
