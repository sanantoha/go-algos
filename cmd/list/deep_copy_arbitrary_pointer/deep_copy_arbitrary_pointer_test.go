package main

import (
	"github.com/sanantoha/go-algos/internals/list"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	root := &list.ArbitraryListNode{Val: 1}
	second := &list.ArbitraryListNode{Val: 2}
	third := &list.ArbitraryListNode{Val: 3}
	four := &list.ArbitraryListNode{Val: 4}
	five := &list.ArbitraryListNode{Val: 5}

	root.Next = second
	second.Next = third
	third.Next = four
	four.Next = five

	second.Arbitrary = five
	third.Arbitrary = root
	five.Arbitrary = second

	copyNode := deepCopy(root)

	assert(t, root, copyNode)
}

func assert(t *testing.T, root *list.ArbitraryListNode, copy *list.ArbitraryListNode) {
	c1 := root
	c2 := copy

	for c1 != nil && c2 != nil {
		if c1.Val != c2.Val {
			t.Errorf("%d != %d", c1.Val, c2.Val)
		}
		if c1 == c2 {
			t.Errorf("%v == %v", c1, c2)
		}
		if (c1.Arbitrary == nil && c2.Arbitrary != nil) || (c1.Arbitrary != nil && c2.Arbitrary == nil) {
			t.Errorf("%v != %v", c1.Arbitrary, c2.Arbitrary)
		}
		if c1.Arbitrary != nil && c1.Arbitrary == c2.Arbitrary {
			t.Errorf("%v == %v", c1.Arbitrary, c2.Arbitrary)
		}

		c1 = c1.Next
		c2 = c2.Next
	}

	if c1 != nil || c2 != nil {
		t.Errorf("%v != %v", c1, c2)
	}
}
