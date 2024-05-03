package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n * log(n)) time | O(log(n)) space - for recursion
func sort(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddle(head)
	left := sort(head)
	right := sort(mid)
	return merge(left, right)
}

func merge(left *list.ListNode, right *list.ListNode) *list.ListNode {

	dummy := &list.ListNode{}
	curr := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	if left != nil {
		curr.Next = left
	} else {
		curr.Next = right
	}

	return dummy.Next
}

func getMiddle(head *list.ListNode) *list.ListNode {
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	next := slow.Next
	slow.Next = nil
	return next
}

func main() {

	head := &list.ListNode{Val: 9, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 5,
		Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8}}}}}}

	res := sort(head)

	fmt.Println(res)
}
