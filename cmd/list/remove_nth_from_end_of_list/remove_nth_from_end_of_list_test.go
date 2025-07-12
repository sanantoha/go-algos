package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name  string
		input *list.ListNode
		n     int
		exp   *list.ListNode
	}{
		{
			name:  "remove 1st",
			input: createList(),
			n:     1,
			exp:   &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4}}}},
		},
		{
			name:  "remove 2nd",
			input: createList(),
			n:     2,
			exp:   &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 5}}}},
		},
		{
			name:  "remove 3rd",
			input: createList(),
			n:     3,
			exp:   &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}},
		},
		{
			name:  "remove 4th",
			input: createList(),
			n:     4,
			exp:   &list.ListNode{Val: 1, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}},
		},
		{
			name:  "remove 5th",
			input: createList(),
			n:     5,
			exp:   &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}},
		},
		{
			name:  "remove 6th",
			input: createList(),
			n:     6,
			exp:   createList(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeNthFromEnd(tt.input, tt.n)

			if !reflect.DeepEqual(tt.exp, res) {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}

func createList() *list.ListNode {
	return &list.ListNode{Val: 1, Next: &list.ListNode{Val: 2, Next: &list.ListNode{Val: 3, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}}}
}
