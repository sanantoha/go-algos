package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(l1 + l2) time | O(1) space
func mergeTwoLists(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	dummy := &list.ListNode{}

	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}

func main() {

	head1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}}}
	head2 := &list.ListNode{Val: 7, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 16}}}}

	fmt.Println(mergeTwoLists(head1, head2))
}
