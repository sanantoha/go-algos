package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"runtime"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	funcs := []func(*list.ListNode, *list.ListNode) *list.ListNode{
		getIntersectionNode,
		getIntersectionNode1,
	}

	common := &list.ListNode{Val: 8, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	l1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 1, Next: common}}
	l2 := &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 1, Next: common}}}

	exp := common

	for _, fn := range funcs {
		fnPtr := reflect.ValueOf(fn).Pointer()
		funcName := runtime.FuncForPC(fnPtr).Name()
		t.Run(funcName, func(t *testing.T) {
			res := fn(l1, l2)

			if !reflect.DeepEqual(exp, res) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
