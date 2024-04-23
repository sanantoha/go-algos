package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(max(l1, l2)) time | O(max(l1, l2)) space
func addTwoNumbers(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {

	c1 := l1
	c2 := l2
	carry := 0

	dummy := &list.ListNode{}
	curr := dummy

	for c1 != nil || c2 != nil {
		v1 := 0
		if c1 != nil {
			v1 = c1.Val
		}
		v2 := 0
		if c2 != nil {
			v2 = c2.Val
		}
		sum := v1 + v2 + carry

		curr.Next = &list.ListNode{Val: sum % 10}
		carry = sum / 10

		curr = curr.Next
		if c1 != nil {
			c1 = c1.Next
		}
		if c2 != nil {
			c2 = c2.Next
		}
	}

	if carry > 0 {
		curr.Next = &list.ListNode{Val: carry}
	}

	return dummy.Next
}

func main() {

	l1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 0, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 9}}}}
	l2 := &list.ListNode{Val: 7, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2}}}

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2)) // 8 -> 3 -> 1 -> 0 -> 1
}
