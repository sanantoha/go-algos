package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}

	graph, err := grph.NewEdgeWeightedDigraph(25)
	if err != nil {
		t.Error(err)
	}
	addEdge(t, graph, 0, 1)
	addEdge(t, graph, 0, 1)
	addEdge(t, graph, 0, 2)
	addEdge(t, graph, 1, 3)
	addEdge(t, graph, 1, 4)
	addEdge(t, graph, 1, 5)
	addEdge(t, graph, 2, 6)
	addEdge(t, graph, 2, 7)
	addEdge(t, graph, 2, 8)
	addEdge(t, graph, 3, 9)
	addEdge(t, graph, 3, 10)
	addEdge(t, graph, 4, 11)
	addEdge(t, graph, 4, 12)
	addEdge(t, graph, 5, 13)
	addEdge(t, graph, 5, 14)
	addEdge(t, graph, 6, 15)
	addEdge(t, graph, 7, 16)
	addEdge(t, graph, 8, 17)
	addEdge(t, graph, 9, 18)
	addEdge(t, graph, 9, 19)
	addEdge(t, graph, 10, 20)
	addEdge(t, graph, 10, 21)
	addEdge(t, graph, 10, 22)
	addEdge(t, graph, 10, 23)
	addEdge(t, graph, 10, 24)

	res := bfs(graph, 0)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func addEdge(t *testing.T, graph *grph.EdgeWeightedDigraph, v, w int) {
	edge, err := grph.NewDirectedEdge(v, w, 0)
	if err != nil {
		t.Error(err)
	}
	graph.AddEdge(edge)
}
