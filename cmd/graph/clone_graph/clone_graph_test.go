package main

import (
	"testing"

	"github.com/sanantoha/go-algos/internals/graph"
)

func TestCloneGraph(t *testing.T) {
	n1 := graph.NewNode(1)
	n2 := graph.NewNode(2)
	n3 := graph.NewNode(3)
	n4 := graph.NewNode(4)

	n1.Neighbors = append(n1.Neighbors, n2)
	n1.Neighbors = append(n1.Neighbors, n4)

	n2.Neighbors = append(n2.Neighbors, n1)
	n2.Neighbors = append(n2.Neighbors, n3)

	n3.Neighbors = append(n3.Neighbors, n2)
	n3.Neighbors = append(n3.Neighbors, n4)

	n4.Neighbors = append(n4.Neighbors, n1)
	n4.Neighbors = append(n4.Neighbors, n3)

	res := cloneGraph(n1)

	if res.Val != n1.Val {
		t.Errorf("expected %v, but got %v", n1.Val, res.Val)
	}
	if len(res.Neighbors) != len(n1.Neighbors) {
		t.Errorf("node has different number of child %v != %v", len(n1.Neighbors), len(res.Neighbors))
	}
	if res == n1 {
		t.Errorf("the objects are the same %v, %v", n1, res)
	}

	res2 := res.Neighbors[0]
	if res2.Val != n2.Val {
		t.Errorf("expected %v, but got %v", n2.Val, res2.Val)
	}
	if len(res2.Neighbors) != len(n2.Neighbors) {
		t.Errorf("node has different number of child %v != %v", len(n2.Neighbors), len(res2.Neighbors))
	}
	if res2 == n2 {
		t.Errorf("the objects are the same %v, %v", n2, res2)
	}

	res4 := res.Neighbors[1]
	if res4.Val != n4.Val {
		t.Errorf("expected %v, but got %v", n4.Val, res4.Val)
	}
	if len(res4.Neighbors) != len(n4.Neighbors) {
		t.Errorf("node has different number of child %v != %v", len(n4.Neighbors), len(res4.Neighbors))
	}
	if res4 == n4 {
		t.Errorf("the objects are the same %v, %v", n4, res4)
	}

	res3 := res2.Neighbors[1]
	if res3.Val != n3.Val {
		t.Errorf("expected %v, but got %v", n3.Val, res3.Val)
	}
	if len(res3.Neighbors) != len(n3.Neighbors) {
		t.Errorf("node has different number of child %v != %v", len(n3.Neighbors), len(res3.Neighbors))
	}
	if res3 == n3 {
		t.Errorf("the objects are the same %v, %v", n3, res3)
	}
}
