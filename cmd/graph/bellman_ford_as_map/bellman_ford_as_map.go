package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func findShortestPath(graph map[string][]*grph.EdgeT[string], start string) *grph.Pair[map[string]float64, map[string]string] {
	return nil
}

func findNegativeWeightCycle(graph map[string][]*grph.EdgeT[string], path *grph.Pair[map[string]float64, map[string]string]) []string {
	return nil
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/bellman_ford/bellmanFord.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	sp := findShortestPath(graph, "0")
	fmt.Println(sp)
	cycle := findNegativeWeightCycle(graph, sp)
	fmt.Println(cycle)
}
