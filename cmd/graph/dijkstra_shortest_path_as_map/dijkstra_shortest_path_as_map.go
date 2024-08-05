package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func findShortestPath(graph map[string][]*grph.EdgeT[string], start string) (*grph.Pair[map[string]float64, map[string]string], error) {
	return nil, nil
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/dijkstra_shortest_path/dijkstraShortestPath.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	sp, err := findShortestPath(graph, "0")
	if err != nil {
		log.Fatalln(err)
	}
	//  &{[0 5 8 4 7] [-1 3 1 0 3]}
	log.Println(sp)
}
