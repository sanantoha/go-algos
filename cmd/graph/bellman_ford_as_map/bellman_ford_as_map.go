package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"math"
)

// O(E * V) time O(V) space
func findShortestPath(graph map[string][]*grph.EdgeT[string], start string) *grph.Pair[map[string]float64, map[string]string] {
	shortest := make(map[string]float64, len(graph))
	prev := make(map[string]string, len(graph))
	for v, _ := range graph {
		shortest[v] = math.MaxFloat64
		prev[v] = ""
	}
	shortest[start] = 0
	edges := getUniqueEdges(graph)

	for v := 0; v < len(graph)-1; v++ {
		for _, edge := range edges {
			relax(shortest, prev, edge)
		}
	}
	return &grph.Pair[map[string]float64, map[string]string]{
		Fst: shortest,
		Snd: prev,
	}
}

// O(E + V) time | O(V) space
func findNegativeWeightCycle(graph map[string][]*grph.EdgeT[string], path *grph.Pair[map[string]float64, map[string]string]) []string {
	shortest := path.Fst
	prev := path.Snd

	v := ""

	edges := getUniqueEdges(graph)
	for _, edge := range edges {
		newWeight := shortest[edge.From()] + edge.Weight
		if newWeight < shortest[edge.To()] {
			v = edge.From()
		}
	}

	if v == "" {
		return nil
	}
	res := make([]string, 1)
	res[0] = v
	u := prev[v]

	for u != v {
		res = append(res, u)
		u = prev[u]
	}
	return res
}

func getUniqueEdges(graph map[string][]*grph.EdgeT[string]) []*grph.EdgeT[string] {
	seen := make(map[*grph.EdgeT[string]]bool, 0)
	edges := make([]*grph.EdgeT[string], 0)
	for _, lst := range graph {
		for _, edge := range lst {
			if !seen[edge] {
				seen[edge] = true
				edges = append(edges, edge)
			}
		}
	}
	return edges
}

func relax(shortest map[string]float64, prev map[string]string, edge *grph.EdgeT[string]) {
	newWeight := edge.Weight + shortest[edge.From()]
	if newWeight < shortest[edge.To()] {
		shortest[edge.To()] = newWeight
		prev[edge.To()] = edge.From()
	}
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/bellman_ford/bellmanFord.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	// ShortestPath{Shortest=[-9.0, -20.0, -18.0, -2.0, -11.0], Prev=[4, 2, 4, 0, 1]}
	sp := findShortestPath(graph, "0")
	fmt.Println(sp)
	cycle := findNegativeWeightCycle(graph, sp)
	fmt.Println(cycle)
}
