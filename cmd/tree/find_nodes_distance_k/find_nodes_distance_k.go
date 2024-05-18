package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
	"reflect"
	"sort"
)

// O(n) time | O(n) space
func findNodesDistanceK(root *tree.TreeNode, target int, k int) []int {
	if root == nil || k < 0 {
		return nil
	}

	parentsMap := make(map[int]*tree.TreeNode)
	enrichParents(parentsMap, root, nil)

	targetNode := findTargetNode(parentsMap, root, target)
	if targetNode == nil {
		return nil
	}

	visited := make(map[*tree.TreeNode]bool)
	visited[targetNode] = true

	res := make([]int, 0)

	queue := make([]Info, 1) // use slice as queue (not optimal time complexity)
	queue[0] = Info{node: targetNode, distance: 0}

	for len(queue) > 0 {
		info := queue[0]
		queue = queue[1:]

		curr := info.node
		distance := info.distance

		if distance == k {
			res = append(res, curr.Val)
			for len(queue) > 0 {
				info = queue[0]
				queue = queue[1:]
				res = append(res, info.node.Val)
			}
			break
		}

		nodes := make([]*tree.TreeNode, 3)
		nodes = append(nodes, curr.Left)
		nodes = append(nodes, curr.Right)
		nodes = append(nodes, parentsMap[curr.Val])

		for _, node := range nodes {
			if node == nil || visited[node] {
				continue
			}

			visited[node] = true

			queue = append(queue, Info{node: node, distance: distance + 1})
		}
	}

	return res
}

type Info struct {
	node     *tree.TreeNode
	distance int
}

func findNodesDistanceKRec(root *tree.TreeNode, target int, k int) []int {
	if root == nil || k < 0 {
		return nil
	}

	parentsMap := make(map[int]*tree.TreeNode)
	enrichParents(parentsMap, root, nil)

	targetNode := findTargetNode(parentsMap, root, target)
	if targetNode == nil {
		return nil
	}

	res := make([]int, 0)
	visited := make(map[*tree.TreeNode]bool)
	findNodes(targetNode, 0, k, parentsMap, visited, &res)
	return res
}

func findNodes(node *tree.TreeNode, distance int, k int, parentsMap map[int]*tree.TreeNode, visited map[*tree.TreeNode]bool, res *[]int) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true

	if distance == k {
		*res = append(*res, node.Val)
	} else {
		findNodes(node.Left, distance+1, k, parentsMap, visited, res)
		findNodes(node.Right, distance+1, k, parentsMap, visited, res)
		findNodes(parentsMap[node.Val], distance+1, k, parentsMap, visited, res)
	}
}

func findTargetNode(parentsMap map[int]*tree.TreeNode, root *tree.TreeNode, target int) *tree.TreeNode {
	if root.Val == target {
		return root
	}

	parent := parentsMap[target]
	if parent == nil {
		return nil
	}
	if parent.Left.Val == target {
		return parent.Left
	} else {
		return parent.Right
	}
}

func enrichParents(parentsMap map[int]*tree.TreeNode, node *tree.TreeNode, parent *tree.TreeNode) {
	if node == nil {
		return
	}

	parentsMap[node.Val] = parent
	enrichParents(parentsMap, node.Left, node)
	enrichParents(parentsMap, node.Right, node)
}

func main() {

	root := &tree.TreeNode{Val: 1}
	root.Left = &tree.TreeNode{Val: 2}
	root.Right = &tree.TreeNode{Val: 3}
	root.Left.Left = &tree.TreeNode{Val: 4}
	root.Left.Right = &tree.TreeNode{Val: 5}
	root.Right.Right = &tree.TreeNode{Val: 6}
	root.Right.Right.Left = &tree.TreeNode{Val: 7}
	root.Right.Right.Right = &tree.TreeNode{Val: 8}

	target := 3
	k := 2

	expected := []int{2, 7, 8}

	actual := findNodesDistanceK(root, target, k)
	sort.Ints(actual)
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))

	actual = findNodesDistanceKRec(root, target, k)
	sort.Ints(actual)
	fmt.Println(actual)
	fmt.Println(reflect.DeepEqual(actual, expected))
}
