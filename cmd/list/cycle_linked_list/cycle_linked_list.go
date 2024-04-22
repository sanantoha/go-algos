package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func isCycle(head *list.ListNode) bool {
	return false
}

func isCycleWithoutSpace(head *list.ListNode) bool {
	return false
}

func main() {
	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3}}}
	root.Next.Next.Next = &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 7, Next: root}}}}

	root1 := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2}}}

	fmt.Println(isCycle(root))
	fmt.Println(isCycle(root1))

	fmt.Println(isCycleWithoutSpace(root))
	fmt.Println(isCycleWithoutSpace(root1))
}
