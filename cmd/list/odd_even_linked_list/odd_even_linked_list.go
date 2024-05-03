package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func oddEvenList(head *list.ListNode) *list.ListNode {
	return nil
}

func oddEvenList1(head *list.ListNode) *list.ListNode {
	return nil
}

func main() {

	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}
	head1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head)
	// 1 -> 3 -> 5 -> 2 -> 4
	fmt.Println(oddEvenList(head))

	fmt.Println("============================================")

	fmt.Println(head1)
	// 1 -> 3 -> 5 -> 2 -> 4
	fmt.Println(oddEvenList(head1))
}
