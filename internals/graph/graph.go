package graph

import (
	"errors"
	"fmt"
	"math"
)

type DirectedEdge struct {
	V      int
	W      int
	Weight float64
}

func NewDirectedEdge(v, w int, weight float64) (*DirectedEdge, error) {
	if v < 0 || w < 0 {
		return nil, errors.New("vertex names must be nonnegative integers")
	}
	if math.IsNaN(weight) {
		return nil, errors.New("Weight is NaN")
	}

	instance := DirectedEdge{
		V:      v,
		W:      w,
		Weight: weight,
	}
	return &instance, nil
}

func (e *DirectedEdge) from() int {
	return e.V
}

func (e *DirectedEdge) to() int {
	return e.W
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %5.2f", e.V, e.W, e.Weight)
}

type EdgeWeightedDigraph struct {
	V   int
	E   int
	Adj []DirectedEdge
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: make([]*Node, 0),
	}
}
