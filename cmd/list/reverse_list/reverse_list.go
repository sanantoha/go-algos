package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func reverseList(head *list.ListNode) *list.ListNode {
	return nil
}

func main() {
	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}

	fmt.Println(head) // 4 -> 3 -> 2 -> 1
}
