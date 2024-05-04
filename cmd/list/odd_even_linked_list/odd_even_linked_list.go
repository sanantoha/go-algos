package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(n) time | O(1) space
func oddEvenList(head *list.ListNode) *list.ListNode {

	curr := head
	dummyEven := &list.ListNode{}
	dummyOdd := &list.ListNode{}
	even := dummyEven
	odd := dummyOdd
	idx := 1

	for curr != nil {
		if idx%2 == 0 {
			even.Next = curr
			even = even.Next
		} else {
			odd.Next = curr
			odd = odd.Next
		}
		idx++
		curr = curr.Next
	}
	even.Next = nil
	odd.Next = dummyEven.Next
	return dummyOdd.Next
}

// O(n) time | O(1) space
func oddEvenList1(head *list.ListNode) *list.ListNode {
	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
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
	fmt.Println(oddEvenList1(head1))
}
