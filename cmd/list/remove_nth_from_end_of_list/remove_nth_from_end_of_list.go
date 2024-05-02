package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(1) space
func removeNthFromEnd(head *list.ListNode, n int) *list.ListNode {

	k := n

	curr := head

	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}

	if curr == nil {
		if k == 0 {
			return head.Next
		} else {
			return head
		}
	}

	fst := curr
	snd := head

	for fst.Next != nil {
		snd = snd.Next
		fst = fst.Next
	}

	next := snd.Next
	if next != nil {
		snd.Next = next.Next
	}
	return head
}

func main() {

	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head)

	res := removeNthFromEnd(head, 10)

	fmt.Println(res)
}
