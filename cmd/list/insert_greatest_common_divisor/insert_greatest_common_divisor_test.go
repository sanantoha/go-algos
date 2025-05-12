package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"testing"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
	// 18 -> 6 -> 6 -> 2 -> 10 -> 1 -> 3
	exp := &list.ListNode{Val: 18, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3}}}}}}}
	lst := &list.ListNode{Val: 18, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 3}}}}

	res := insertGreatestCommonDivisors(lst)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
