package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"math"
)

// O(E * V) time | O(V) space
func findShortestPath(graph *grph.EdgeWeightedDigraph, start int) *grph.ShortestPath {
	shortest := make([]float64, graph.V)
	prev := make([]int, graph.V)
	for i := 0; i < graph.V; i++ {
		shortest[i] = math.MaxFloat64
		prev[i] = -1
	}
	shortest[start] = 0

	for v := 0; v < graph.V-1; v++ {
		for _, edge := range graph.Edges() {
			relax(shortest, prev, edge)
		}
	}

	return &grph.ShortestPath{
		Shortest: shortest,
		Prev:     prev,
	}
}

func relax(shortest []float64, prev []int, edge *grph.DirectedEdge) {
	newWeight := shortest[edge.From()] + edge.Weight
	if newWeight < shortest[edge.To()] {
		shortest[edge.To()] = newWeight
		prev[edge.To()] = edge.From()
	}
}

func findNegativeWeightCycle(graph *grph.EdgeWeightedDigraph, path *grph.ShortestPath) []int {
	return nil
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/bellman_ford/bellmanFord.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sp := findShortestPath(graph, 0)
	fmt.Println(sp) // ShortestPath{Shortest=[-9.0, -20.0, -18.0, -2.0, -11.0], Prev=[4, 2, 4, 0, 1]}

	circle := findNegativeWeightCycle(graph, sp)
	fmt.Println(circle) // [1, 2, 4]
}
