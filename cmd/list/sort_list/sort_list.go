package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func sort(head *list.ListNode) *list.ListNode {
	return nil
}

func main() {

	head := &list.ListNode{Val: 9, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 5,
		Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8}}}}}}

	res := sort(head)

	fmt.Println(res)
}
