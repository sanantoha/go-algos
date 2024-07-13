package graph

import (
	"errors"
	"fmt"
	"math"
	"strings"
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
	return fmt.Sprintf("%d->%d %4.2f", e.V, e.W, e.Weight)
}

type EdgeWeightedDigraph struct {
	V   int
	E   int
	adj [][]*DirectedEdge
}

func NewEdgeWeightedDigraph(V int) (*EdgeWeightedDigraph, error) {
	if V < 0 {
		return nil, errors.New("Number of vertices in a Digraph must be nonnegative")
	}
	adj := make([][]*DirectedEdge, V)
	for v := 0; v < V; v++ {
		adj[v] = make([]*DirectedEdge, 0)
	}

	graph := EdgeWeightedDigraph{
		V:   V,
		E:   0,
		adj: adj,
	}
	return &graph, nil
}

func (e *EdgeWeightedDigraph) AddEdge(edge *DirectedEdge) {
	v := edge.from()
	e.adj[v] = append(e.adj[v], edge)
	e.E++
}

func (e *EdgeWeightedDigraph) Adj(v int) ([]*DirectedEdge, error) {
	if v < 0 || v >= e.V {
		return nil, errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, e.V-1))
	}
	return e.adj[v], nil
}

func (e *EdgeWeightedDigraph) Edges() []*DirectedEdge {
	list := make([]*DirectedEdge, 0)

	for v := 0; v < e.V; v++ {
		for _, edge := range e.adj[v] {
			list = append(list, edge)
		}
	}

	return list
}

func (e *EdgeWeightedDigraph) String() string {
	var builder = strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d %d\n", e.V, e.E))

	for v := 0; v < e.V; v++ {
		builder.WriteString(fmt.Sprintf("%d: ", v))
		for _, edge := range e.adj[v] {
			builder.WriteString(fmt.Sprintf("%v  ", edge))
		}
		builder.WriteString("\n")
	}
	return builder.String()
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
