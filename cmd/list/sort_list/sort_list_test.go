package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	head := &list.ListNode{Val: 9, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 5,
		Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8}}}}}}

	exp := &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4,
		Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 9}}}}}}

	act := sort(head)

	if !reflect.DeepEqual(exp, act) {
		t.Errorf("expected %v, got %v", exp, act)
	}
}
