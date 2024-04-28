package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

// O(l + r) time | O(l) space
func getIntersectionNode(l *list.ListNode, r *list.ListNode) *list.ListNode {

	set := make(map[*list.ListNode]bool)

	currL := l
	for currL != nil {
		set[currL] = true
		currL = currL.Next
	}

	currR := r
	for currR != nil {
		_, ok := set[currR]
		if ok {
			return currR
		}
		currR = currR.Next
	}
	return nil
}

// O(l + r) time | O(1) space
func getIntersectionNode1(l *list.ListNode, r *list.ListNode) *list.ListNode {

	currL := l
	currR := r

	for currL != currR {
		if currL != nil {
			currL = currL.Next
		} else {
			currL = r
		}
		if currR != nil {
			currR = currR.Next
		} else {
			currR = l
		}
	}
	return currL
}

func main() {
	common := &list.ListNode{Val: 8, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	l1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 1, Next: common}}
	l2 := &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 1, Next: common}}}

	fmt.Println(getIntersectionNode(l1, l2))
	fmt.Println(getIntersectionNode1(l1, l2))
}
