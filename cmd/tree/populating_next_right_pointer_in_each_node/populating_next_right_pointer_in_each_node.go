package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (tree *Node) String() string {
	if tree == nil {
		return "nil"
	}
	pv := ""
	if tree.Next != nil {
		pv = strconv.Itoa(tree.Next.Val)
	}
	return fmt.Sprintf("Node{Val=%d, Left=%s, Right=%s, Next=%s}", tree.Val, tree.Left.String(), tree.Right.String(), pv)
}

// O(n) time | O(n) space
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 1)
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)

		var prev *Node

		for size > 0 {
			size--

			curr := queue[0]
			queue = queue[1:]

			if curr == nil {
				continue
			}

			if prev != nil {
				prev.Next = curr
			}

			prev = curr

			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}
	}

	return root
}

func main() {

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	fmt.Println(root)

	res := connect(root)

	fmt.Println(res)
}
