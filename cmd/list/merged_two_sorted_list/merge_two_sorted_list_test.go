package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"testing"
)

func TestMergeTwoSortedList(t *testing.T) {
	testTable := []struct {
		name   string
		head1  *list.ListNode
		head2  *list.ListNode
		expRes *list.ListNode
	}{
		{
			name:   "happy path",
			head1:  &list.ListNode{Val: 4, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 19}}}},
			head2:  &list.ListNode{Val: 7, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 16}}}},
			expRes: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 7, Next: &list.ListNode{Val: 8, Next: &list.ListNode{Val: 9, Next: &list.ListNode{Val: 10, Next: &list.ListNode{Val: 15, Next: &list.ListNode{Val: 16, Next: &list.ListNode{Val: 19}}}}}}}},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := mergeTwoLists(subtest.head1, subtest.head2)

			if !compareLists(subtest.expRes, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expRes, res)
			}
		})
	}
}

func compareLists(l1, l2 *list.ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
