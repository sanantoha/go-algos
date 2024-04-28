package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func middleNode(head *list.ListNode) *list.ListNode {
	return nil
}

func main() {
	head := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}}}

	fmt.Println(middleNode(head))

	head1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head1)
}
