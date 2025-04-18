package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node0 := &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	node1 := &list.ListNode{Val: 2, Next: node0}
	lst := &list.ListNode{Val: 1, Next: node1}

	testTable := []struct {
		name string
		node *list.ListNode
		res  *list.ListNode
		exp  *list.ListNode
	}{
		{
			name: "happy path",
			node: node0,
			res:  lst,
			exp:  &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}},
		},
		{
			name: "nil node",
			node: nil,
			res:  nil,
			exp:  nil,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			deleteNode(subtest.node)

			if !reflect.DeepEqual(subtest.exp, subtest.res) {
				t.Errorf("expected %v, but got %v", subtest.exp, subtest.res)
			}
		})
	}
}
