package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(1) space
func isPalindrome(head *list.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fst := head
	snd := reverse(getMiddle(head))

	for snd != nil {
		if fst.Val != snd.Val {
			return false
		}
		fst = fst.Next
		snd = snd.Next
	}
	return true
}

func getMiddle(head *list.ListNode) *list.ListNode {

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *list.ListNode) *list.ListNode {
	curr := head
	var prev *list.ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {

	lst := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 1}}}}}}
	lst1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 1}}}}}

	fmt.Println(isPalindrome(lst))
	fmt.Println(isPalindrome(lst1))
}
