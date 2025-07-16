package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestMst(t *testing.T) {
	tests := []struct {
		name     string
		graph    *grph.EdgeWeightedGraph
		expGraph *grph.EdgeWeightedGraph
	}{
		{
			name:     "happy path 6*9",
			graph:    createGraph(),
			expGraph: createExpectedGraph(),
		},
		{
			name:     "happy path 7*10",
			graph:    createGraph1(),
			expGraph: createExpectedGraph1(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mst(tt.graph)

			if !reflect.DeepEqual(tt.expGraph, res) {
				t.Errorf("expected %v, but got %v", tt.expGraph, res)
			}
		})
	}
}

func createGraph() *grph.EdgeWeightedGraph {
	graph := grph.NewEdgeWeightedGraph(6)

	graph.AddEdge(grph.NewEdge(0, 1, 7.0))
	graph.AddEdge(grph.NewEdge(0, 2, 8.0))
	graph.AddEdge(grph.NewEdge(1, 2, 3.0))
	graph.AddEdge(grph.NewEdge(1, 3, 6.0))
	graph.AddEdge(grph.NewEdge(2, 3, 4.0))
	graph.AddEdge(grph.NewEdge(2, 4, 3.0))
	graph.AddEdge(grph.NewEdge(3, 4, 2.0))
	graph.AddEdge(grph.NewEdge(3, 5, 5.0))
	graph.AddEdge(grph.NewEdge(4, 5, 2.0))

	return graph
}

func createExpectedGraph() *grph.EdgeWeightedGraph {
	/*
	   6 5
	   0: 0-1 7.00
	   1: 1-2 3.00  0-1 7.00
	   2: 1-2 3.00  2-4 3.00
	   3: 3-4 2.00
	   4: 3-4 2.00  4-5 2.00  2-4 3.00
	   5: 4-5 2.00
	*/

	graph := grph.NewEdgeWeightedGraph(6)

	graph.AddEdge(grph.NewEdge(0, 1, 7.0))
	graph.AddEdge(grph.NewEdge(1, 2, 3.0))
	graph.AddEdge(grph.NewEdge(2, 4, 3.0))
	graph.AddEdge(grph.NewEdge(3, 4, 2.0))
	graph.AddEdge(grph.NewEdge(4, 5, 2.0))

	return graph
}

func createGraph1() *grph.EdgeWeightedGraph {
	graph := grph.NewEdgeWeightedGraph(7)

	graph.AddEdge(grph.NewEdge(0, 1, 2.0))
	graph.AddEdge(grph.NewEdge(0, 2, 3.0))
	graph.AddEdge(grph.NewEdge(0, 3, 7.0))
	graph.AddEdge(grph.NewEdge(1, 2, 6.0))
	graph.AddEdge(grph.NewEdge(1, 6, 3.0))
	graph.AddEdge(grph.NewEdge(2, 4, 1.0))
	graph.AddEdge(grph.NewEdge(2, 5, 8.0))
	graph.AddEdge(grph.NewEdge(3, 4, 5.0))
	graph.AddEdge(grph.NewEdge(4, 5, 4.0))
	graph.AddEdge(grph.NewEdge(5, 6, 2.0))

	return graph
}

func createExpectedGraph1() *grph.EdgeWeightedGraph {
	/*
	   7 6
	   0: 0-1 2.00000  0-2 3.00000
	   1: 0-1 2.00000  1-6 3.00000
	   2: 2-4 1.00000  0-2 3.00000
	   3: 3-4 5.00000
	   4: 2-4 1.00000  3-4 5.00000
	   5: 5-6 2.00000
	   6: 5-6 2.00000  1-6 3.00000
	*/
	graph := grph.NewEdgeWeightedGraph(7)

	graph.AddEdge(grph.NewEdge(0, 1, 2.0))
	graph.AddEdge(grph.NewEdge(0, 2, 3.0))
	graph.AddEdge(grph.NewEdge(1, 6, 3.0))
	graph.AddEdge(grph.NewEdge(2, 4, 1.0))
	graph.AddEdge(grph.NewEdge(3, 4, 5.0))
	graph.AddEdge(grph.NewEdge(5, 6, 2.0))

	return graph
}
