package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func deleteNode(node *list.ListNode) {

}

func main() {
	node0 := &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	node1 := &list.ListNode{Val: 2, Next: node0}
	lst := &list.ListNode{Val: 1, Next: node1}

	fmt.Println(lst)

	deleteNode(lst)

	fmt.Println(lst)
}
