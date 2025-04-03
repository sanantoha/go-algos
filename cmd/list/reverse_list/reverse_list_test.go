package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"testing"
)

func TestReverseList(t *testing.T) {
	testTable := []struct {
		name     string
		input    *list.ListNode
		expected *list.ListNode
	}{
		{
			name:     "happy path",
			input:    &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}},
			expected: &list.ListNode{Val: 5, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 1}}}}},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := reverseList(subtest.input)

			if !isEqual(subtest.expected, res) { //reflect.DeepEqual(subtest.expected, res) { // also works but not handle infinity recursion properly
				t.Errorf("expected (%v), but got (%v)", subtest.expected, res)
			}
		})
	}
}

func isEqual(a, b *list.ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a, b = a.Next, b.Next
	}
	return a == nil && b == nil
}
