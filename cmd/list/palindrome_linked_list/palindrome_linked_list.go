package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/list"
)

func isPalindrome(head *list.ListNode) bool {
	return false
}

func main() {

	lst := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 1}}}}}}
	lst1 := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 1}}}}}

	fmt.Println(isPalindrome(lst))
	fmt.Println(isPalindrome(lst1))
}
