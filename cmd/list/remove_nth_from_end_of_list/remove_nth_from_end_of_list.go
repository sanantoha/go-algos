package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func removeNthFromEnd(head *list.ListNode, n int) *list.ListNode {
	return nil
}

func main() {

	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head)

	res := removeNthFromEnd(head, 4)

	fmt.Println(res)
}
