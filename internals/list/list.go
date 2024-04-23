package list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	builder := strings.Builder{}

	curr := head

	for curr != nil {
		builder.WriteString(fmt.Sprintf("%d", curr.Val))
		if curr.Next != nil {
			builder.WriteString(" -> ")
		}

		curr = curr.Next
	}

	return builder.String()
}

type ArbitraryListNode struct {
	Val       int
	Next      *ArbitraryListNode
	Arbitrary *ArbitraryListNode
}

// String provides a string representation of the list starting from this node
func (n *ArbitraryListNode) String() string {
	visited := make(map[*ArbitraryListNode]bool)
	return stringify(n, visited)
}

// Helper function to create a string representation of the list
func stringify(node *ArbitraryListNode, visited map[*ArbitraryListNode]bool) string {
	if node == nil {
		return "nil"
	}
	if visited[node] {
		return fmt.Sprintf("Node(val=%d) [Already Visited]", node.Val)
	}
	visited[node] = true

	nextStr := stringify(node.Next, visited)
	arbStr := "nil"
	if node.Arbitrary != nil {
		arbStr = stringify(node.Arbitrary, visited)
	}

	return fmt.Sprintf("Node(val=%d, Next=[%s], Arbitrary=[%s])", node.Val, nextStr, arbStr)
}
