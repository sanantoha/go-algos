package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

/**
 * Given the head of a linked list head, in which each node contains an integer value.
 *
 * Between every pair of adjacent nodes, insert a new node with a value equal to the greatest common
 * divisor of them.
 *
 * Return the linked list after insertion.
 *
 * The greatest common divisor of two numbers is the largest positive integer that evenly divides
 * both numbers.
 */
// O(n) time | O(1) space
func insertGreatestCommonDivisors(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}

	curr := head

	for curr.Next != nil {
		next := curr.Next
		node := &list.ListNode{Val: gcd(curr.Val, curr.Next.Val)}
		curr.Next = node
		node.Next = next
		curr = next
	}
	return head
}

func gcd(x, y int) int {
	for y != 0 {
		tmp := y
		y = x % y
		x = tmp
	}
	return x
}

func main() {
	lst := &list.ListNode{Val: 18, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 3}}}}

	// 18 -> 6 -> 10 -> 3 -> null
	fmt.Println(lst)

	res := insertGreatestCommonDivisors(lst)

	// 18 -> 6 -> 6 -> 2 -> 10 -> 1 -> 3 -> null
	fmt.Println(res)
}
