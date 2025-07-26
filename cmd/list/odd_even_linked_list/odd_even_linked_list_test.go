package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"runtime"
	"testing"
)

func TestOddEvenList(t *testing.T) {

	exp := &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4}}}}}

	funcs := []func(node *list.ListNode) *list.ListNode{
		oddEvenList,
		oddEvenList1,
	}

	for _, fn := range funcs {
		fnPtr := reflect.ValueOf(fn).Pointer()
		funcName := runtime.FuncForPC(fnPtr).Name()
		t.Run(funcName, func(t *testing.T) {
			res := fn(createList())

			if !reflect.DeepEqual(exp, res) {

			}
		})
	}
}

func createList() *list.ListNode {
	return &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}
}
