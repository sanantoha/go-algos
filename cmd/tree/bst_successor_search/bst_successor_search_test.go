package main

import (
	"github.com/sanantoha/go-algos/internals/tree"
	"testing"
)

func TestFindInOrderSuccessor(t *testing.T) {
	node5 := &tree.Node{Key: 5}
	node9 := &tree.Node{Key: 9}
	node11 := &tree.Node{Key: 11}
	node12 := &tree.Node{Key: 12}
	node14 := &tree.Node{Key: 14}
	node20 := &tree.Node{Key: 20}
	node25 := &tree.Node{Key: 25}

	node5.Parent = node9

	node9.Left = node5
	node9.Right = node12
	node9.Parent = node20

	node11.Parent = node12

	node12.Left = node11
	node12.Right = node14
	node12.Parent = node9

	node14.Parent = node12

	node20.Left = node9
	node20.Right = node25

	node25.Parent = node20

	//         20
	//       /    \
	//      9      25
	//    /   \
	//   5     12
	//        /   \
	//       11   14

	assert(t, node5, 9)
	assert(t, node9, 11)
	assert(t, node11, 12)
	assert(t, node12, 14)
	assert(t, node14, 20)
	assert(t, node20, 25)
	v := findInOrderSuccessor(node25)
	if v != nil {
		t.Errorf("expecte nil, but got %v", v)
	}
}

func assert(t *testing.T, node *tree.Node, exp int) {
	key := findInOrderSuccessor(node).Key
	if key != exp {
		t.Errorf("expected %v, but got %v", exp, key)
	}
}
