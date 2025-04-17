package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"testing"
)

func TestIsCycleHappyPath(t *testing.T) {

	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3}}}
	root.Next.Next.Next = &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 7, Next: root}}}}

	if !isCycle(root) {
		t.Error("input list has cyrcle and test should return true but it returns false")
	}
}

func TestIsCycle(t *testing.T) {
	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2}}}

	if isCycle(root) {
		t.Error("input list has not cyrcle and test should return false but it returns true")
	}
}

func TestIsCycleWithoutSpaceHappyPath(t *testing.T) {

	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3}}}
	root.Next.Next.Next = &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 7, Next: root}}}}

	if !isCycleWithoutSpace(root) {
		t.Error("input list has cyrcle and test should return true but it returns false")
	}
}

func TestIsCycleWithoutSpace(t *testing.T) {
	root := &list.ListNode{Val: 0, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2}}}

	if isCycleWithoutSpace(root) {
		t.Error("input list has not cyrcle and test should return false but it returns true")
	}
}
