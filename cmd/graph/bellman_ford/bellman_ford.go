package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func findShortestPath(graph *grph.EdgeWeightedDigraph, start int) *grph.ShortestPath {
	return nil
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
