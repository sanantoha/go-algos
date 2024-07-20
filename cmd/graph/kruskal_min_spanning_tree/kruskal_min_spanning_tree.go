package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
)

func mst(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	return nil
}

func mst1(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	return nil
}

func main() {

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
	/*
	   6 5
	   0: 0-1 7.00000
	   1: 1-2 3.00000  0-1 7.00000
	   2: 1-2 3.00000  2-4 3.00000
	   3: 3-4 2.00000
	   4: 3-4 2.00000  4-5 2.00000  2-4 3.00000
	   5: 4-5 2.00000
	*/
	fmt.Println(graph)
	fmt.Println("=========================================")
	fmt.Println(mst(graph))
	fmt.Println("=========================================")
	fmt.Println("\n\n")
	fmt.Println("=========================================")

	graph1 := grph.NewEdgeWeightedGraph(7)
	graph1.AddEdge(grph.NewEdge(0, 1, 2.0))
	graph1.AddEdge(grph.NewEdge(0, 2, 3.0))
	graph1.AddEdge(grph.NewEdge(0, 3, 7.0))
	graph1.AddEdge(grph.NewEdge(1, 2, 6.0))
	graph1.AddEdge(grph.NewEdge(1, 6, 3.0))
	graph1.AddEdge(grph.NewEdge(2, 4, 1.0))
	graph1.AddEdge(grph.NewEdge(2, 5, 8.0))
	graph1.AddEdge(grph.NewEdge(3, 4, 5.0))
	graph1.AddEdge(grph.NewEdge(4, 5, 4.0))
	graph1.AddEdge(grph.NewEdge(5, 6, 2.0))

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
	fmt.Println(graph1)
	fmt.Println("=========================================")
	fmt.Println(mst1(graph1))
}
