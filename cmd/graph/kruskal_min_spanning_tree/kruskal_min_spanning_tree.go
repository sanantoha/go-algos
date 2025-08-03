package main

import (
	"container/heap"
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"sort"
)

// O(E * log(E)) time | O(E + V) space
func mst(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	ngraph := grph.NewEdgeWeightedGraph(graph.V)

	edges := graph.Edges()

	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Weight == edges[j].Weight { // not required for tests only
			return edges[i].V < edges[j].V
		}
		return edges[i].Weight < edges[j].Weight
	})

	parents := makeSet(graph.V)
	ranks := make([]int, graph.V)
	for v := 0; v < graph.V; v++ {
		ranks[v] = 0
	}

	for _, minEdge := range edges {
		v := minEdge.Either()
		u := minEdge.Other(v)

		pv := find(parents, v)
		pu := find(parents, u)

		if pv != pu {
			union(parents, ranks, pv, pu)
			edge := grph.NewEdge(v, u, minEdge.Weight)
			err := ngraph.AddEdge(edge)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	return ngraph
}

// O(E * log(E)) time | O(E + V) space
func mst1(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	ngraph := grph.NewEdgeWeightedGraph(graph.V)

	h := &EdgeHeap{}

	for _, edge := range graph.Edges() {
		heap.Push(h, edge)
	}

	parents := makeSet(graph.V)
	ranks := make([]int, graph.V)
	for v := 0; v < graph.V; v++ {
		ranks[v] = 0
	}

	for h.Len() > 0 {
		minEdge := heap.Pop(h).(*grph.Edge)

		v := minEdge.Either()
		u := minEdge.Other(v)

		pv := find(parents, v)
		pu := find(parents, u)

		if pv != pu {
			union(parents, ranks, pv, pu)
			edge := grph.NewEdge(v, u, minEdge.Weight)
			err := ngraph.AddEdge(edge)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	return ngraph
}

func makeSet(size int) []int {
	parents := make([]int, size)
	for v := 0; v < size; v++ {
		parents[v] = v
	}
	return parents
}

func find(parents []int, v int) int {
	if parents[v] == v {
		return v
	}
	parents[v] = find(parents, parents[v])
	return parents[v]
}

func union(parents []int, ranks []int, pv int, pu int) {
	if ranks[pv] < ranks[pu] {
		parents[pv] = pu
	} else if ranks[pv] > ranks[pu] {
		parents[pu] = pv
	} else {
		ranks[pu]++
		parents[pv] = pu
	}
}

type EdgeHeap []*grph.Edge

func (h EdgeHeap) Len() int {
	return len(h)
}

func (h EdgeHeap) Less(i, j int) bool {
	if h[i].Weight == h[j].Weight {
		return h[i].V < h[j].V
	}
	return h[i].Weight < h[j].Weight
}

func (h EdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(*grph.Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
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
	   0: 0-1 7.00
	   1: 1-2 3.00  0-1 7.00
	   2: 1-2 3.00  2-4 3.00
	   3: 3-4 2.00
	   4: 3-4 2.00  4-5 2.00  2-4 3.00
	   5: 4-5 2.00
	*/
	fmt.Println(graph)
	fmt.Println("=========================================")
	fmt.Println(mst(graph))
	fmt.Println("=========================================")
	fmt.Println(mst1(graph))
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
	fmt.Println(mst(graph1))
	fmt.Println("=========================================")
	fmt.Println(mst1(graph1))
}
