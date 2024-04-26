package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func getIntersectionNode(l *list.ListNode, r *list.ListNode) *list.ListNode {
	return nil
}

func getIntersectionNode1(l *list.ListNode, r *list.ListNode) *list.ListNode {
	return nil
}

func main() {
	common := &list.ListNode{Val: 8, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	l1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 1, Next: common}}
	l2 := &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 1, Next: common}}}

	fmt.Println(getIntersectionNode(l1, l2))
	fmt.Println(getIntersectionNode1(l1, l2))
}
