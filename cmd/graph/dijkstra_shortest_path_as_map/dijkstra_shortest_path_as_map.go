package main

import (
	"container/heap"
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"math"
)

// O((E + V) * log(V)) time | O(V) space
func findShortestPath(graph map[string][]*grph.EdgeT[string], start string) (*grph.Pair[map[string]float64, map[string]string], error) {
	shortest := make(map[string]float64, len(graph))
	prev := make(map[string]string, len(graph))
	for v, _ := range graph {
		shortest[v] = math.MaxFloat64
		prev[v] = ""
	}
	shortest[start] = 0

	h := &EdgeTHeap{}
	heap.Push(h, grph.NewEdgeT[string]("0", "0", 0))

	for h.Len() > 0 {
		minEdge := heap.Pop(h).(*grph.EdgeT[string])

		for _, edge := range graph[minEdge.To()] {
			relax(shortest, prev, edge, h)
		}
	}

	return &grph.Pair[map[string]float64, map[string]string]{
		Fst: shortest,
		Snd: prev,
	}, nil
}

func relax(shortest map[string]float64, prev map[string]string, edge *grph.EdgeT[string], h *EdgeTHeap) {
	newWeight := shortest[edge.From()] + edge.Weight
	if newWeight < shortest[edge.To()] {
		shortest[edge.To()] = newWeight
		prev[edge.To()] = edge.From()
		heap.Push(h, edge)
	}
}

type EdgeTHeap []*grph.EdgeT[string]

func (h EdgeTHeap) Len() int {
	return len(h)
}

func (h EdgeTHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h EdgeTHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeTHeap) Push(x interface{}) {
	*h = append(*h, x.(*grph.EdgeT[string]))
}

func (h *EdgeTHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
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
	//   &{map[0:0 1:5 2:8 3:4 4:7] map[0: 1:3 2:1 3:0 4:3]}
	log.Println(sp)
}
