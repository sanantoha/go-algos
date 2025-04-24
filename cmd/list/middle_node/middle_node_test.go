package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	head := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}}}
	exp := &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}

	res := middleNode(head)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestMiddleNode1(t *testing.T) {
	head := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}
	exp := &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}

	res := middleNode(head)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
