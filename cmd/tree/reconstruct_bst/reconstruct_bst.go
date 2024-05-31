package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"math"
	"reflect"
)

// O(n ^ 2) time | O(n ^ 2) space
func reconstructBst(arr []int) *tree.TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	val := arr[0]
	lst := arr[1:]

	root := &tree.TreeNode{Val: val}
	leftLst := make([]int, 0)
	rightLst := make([]int, 0)
	for _, v := range lst {
		if v < val {
			leftLst = append(leftLst, v)
		} else {
			rightLst = append(rightLst, v)
		}
	}
	root.Left = reconstructBst(leftLst)
	root.Right = reconstructBst(rightLst)
	return root
}

// O(n) time | O(n) space
func reconstructBst1(arr []int) *tree.TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	return helper(arr, CreateIndexer(0), math.MinInt, math.MaxInt)
}

type Indexer struct {
	idx *int
}

func CreateIndexer(i int) Indexer {
	return Indexer{idx: &i}
}

func (idx *Indexer) Get() int {
	return *idx.idx
}

func (idx *Indexer) Inc() {
	*idx.idx++
}

func helper(arr []int, indexer Indexer, minVal, maxVal int) *tree.TreeNode {
	if indexer.Get() == len(arr) {
		return nil
	}

	val := arr[indexer.Get()]
	if val < minVal || val >= maxVal {
		return nil
	}
	indexer.Inc()
	root := &tree.TreeNode{Val: val}
	root.Left = helper(arr, indexer, minVal, val)
	root.Right = helper(arr, indexer, val, maxVal)
	return root
}

func main() {

	preOrderTraversalValues := []int{10, 4, 2, 1, 3, 17, 19, 18}

	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 17,
			Right: &tree.TreeNode{
				Val: 19,
				Left: &tree.TreeNode{
					Val: 18,
				},
			},
		},
	}

	expected := getDfsOrder(root, []int{})

	actualTree := reconstructBst(preOrderTraversalValues)
	actual := getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actualTree = reconstructBst1(preOrderTraversalValues)
	actual = getDfsOrder(actualTree, []int{})
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}

func getDfsOrder(root *tree.TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	arr = getDfsOrder(root.Left, arr)
	return getDfsOrder(root.Right, arr)
}
