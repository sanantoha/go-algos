package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(1) space
func deleteNode(node *list.ListNode) {

	curr := node

	for curr != nil {
		next := curr.Next
		if next != nil {
			curr.Val = next.Val
			if next.Next == nil {
				curr.Next = nil
				break
			}
		}
		curr = next
	}
}

func main() {
	node0 := &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	node1 := &list.ListNode{Val: 2, Next: node0}
	lst := &list.ListNode{Val: 1, Next: node1}

	fmt.Println(lst)

	deleteNode(node0)

	fmt.Println(lst)
}
