package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(n) space
func isCycle(head *list.ListNode) bool {
	if head == nil {
		return false
	}

	set := make(map[*list.ListNode]bool)

	curr := head

	for curr != nil {
		if _, exists := set[curr]; exists {
			return true
		}
		set[curr] = true
		curr = curr.Next
	}
	return false
}

// O(n) time | O(1) space
func isCycleWithoutSpace(head *list.ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head

	for fast != slow {

		if fast == nil || fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}

func main() {
	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3}}}
	root.Next.Next.Next = &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 7, Next: root}}}}

	root1 := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2}}}

	fmt.Println(isCycle(root))
	fmt.Println(!isCycle(root1))

	fmt.Println(isCycleWithoutSpace(root))
	fmt.Println(!isCycleWithoutSpace(root1))
}
