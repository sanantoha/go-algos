package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(1) space
func reverseList(head *list.ListNode) *list.ListNode {
	curr := head
	var prev *list.ListNode = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head)
	fmt.Println(reverseList(head)) // 5 -> 4 -> 3 -> 2 -> 1
}
