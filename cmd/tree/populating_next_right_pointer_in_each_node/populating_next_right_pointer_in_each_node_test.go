package main

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	exp := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
				Next: &Node{
					Val: 5,
					Next: &Node{
						Val: 6,
						Next: &Node{
							Val: 7,
						},
					},
				},
			},
			Right: &Node{
				Val: 5,
				Next: &Node{
					Val: 6,
					Next: &Node{
						Val: 7,
					},
				},
			},
			Next: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
					Next: &Node{
						Val: 7,
					},
				},
				Right: &Node{
					Val: 7,
				},
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
				Next: &Node{
					Val: 7,
				},
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	res := connect(root)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
