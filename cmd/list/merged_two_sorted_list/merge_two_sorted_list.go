package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func mergeTwoLists(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	return nil
}

func main() {

	head1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}}}
	head2 := &list.ListNode{Val: 7, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 16}}}}

	fmt.Println(mergeTwoLists(head1, head2))
}
