package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
)

// O(n) time | O(log(n)) space
func sortedArrayToBst(arr []int) *tree.TreeNode {
	return helper(arr, 0, len(arr)-1)
}

func helper(arr []int, l int, r int) *tree.TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	root := &tree.TreeNode{Val: arr[mid]}
	root.Left = helper(arr, l, mid-1)
	root.Right = helper(arr, mid+1, r)
	return root
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sortedArrayToBst(arr))
}
