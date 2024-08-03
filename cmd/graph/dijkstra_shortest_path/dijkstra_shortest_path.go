package main

import (
	"container/heap"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"math"
)

// O((E + V) * log(V)) time | O(V) space
func findShortestPath(graph *grph.EdgeWeightedDigraph, start int) (*grph.ShortestPath, error) {
	shortest := make([]float64, graph.V)
	prev := make([]int, graph.V)
	for v := 0; v < graph.V; v++ {
		shortest[v] = math.MaxFloat64
		prev[v] = -1
	}
	shortest[start] = 0

	h := &DirectedEdgeHeap{}
	edge, err := grph.NewDirectedEdge(start, start, 0)
	if err != nil {
		return nil, err
	}
	heap.Push(h, edge)

	for h.Len() > 0 {
		minEdge := heap.Pop(h).(*grph.DirectedEdge)

		adj, err := graph.Adj(minEdge.To())
		if err != nil {
			return nil, err
		}
		for _, edge := range adj {
			relax(shortest, prev, edge, h)
		}
	}

	return &grph.ShortestPath{
		Shortest: shortest,
		Prev:     prev,
	}, nil
}

func relax(shortest []float64, prev []int, edge *grph.DirectedEdge, h *DirectedEdgeHeap) {
	newWeight := shortest[edge.From()] + edge.Weight
	if newWeight < shortest[edge.To()] {
		shortest[edge.To()] = newWeight
		prev[edge.To()] = edge.From()
		heap.Push(h, edge)
	}
}

type DirectedEdgeHeap []*grph.DirectedEdge

func (h DirectedEdgeHeap) Len() int {
	return len(h)
}

func (h DirectedEdgeHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h DirectedEdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DirectedEdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(*grph.DirectedEdge))
}

func (h *DirectedEdgeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/dijkstra_shortest_path/dijkstraShortestPath.txt")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(graph)
	log.Println("=====================================")
	sp, err := findShortestPath(graph, 0)
	if err != nil {
		log.Fatalln(err)
	}
	//  &{[0 5 8 4 7] [-1 3 1 0 3]}
	log.Println(sp)
}
