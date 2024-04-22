package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func addTwoNumbers(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	return nil
}

func main() {

	l1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 0, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 9}}}}
	l2 := &list.ListNode{Val: 7, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2}}}

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2)) // 8 -> 3 -> 1 -> 0 -> 1
}
