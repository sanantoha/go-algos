package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func findShortestPath(graph *grph.EdgeWeightedDigraph, start int) *grph.ShortestPath {
	return nil
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/dijkstra_shortest_path/dijkstraShortestPath.txt")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(graph)
	log.Println("=====================================")
	sp := findShortestPath(graph, 0)
	log.Println(sp)
}
