package main

import (
	"container/heap"
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
)

// O(E * log(V)) time | O(V) space
func mst(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	ngraph := grph.NewEdgeWeightedGraph(graph.V)

	h := &NodeHeap{}

	start := 0
	for _, edge := range graph.Adj(start) {
		heap.Push(h, &Node{start, edge})
	}

	inTree := 1
	visited := make([]bool, graph.V)
	visited[start] = true

	for h.Len() > 0 && inTree < graph.V {
		node := heap.Pop(h).(*Node)
		from := node.from
		minEdge := node.edge
		to := minEdge.Other(from)

		if visited[to] {
			continue
		}
		visited[to] = true
		inTree++

		nedge := grph.NewEdge(minEdge.Either(), minEdge.Other(minEdge.Either()), minEdge.Weight)
		ngraph.AddEdge(nedge)
		for _, edge := range graph.Adj(to) {
			heap.Push(h, &Node{to, edge})
		}
	}

	if inTree < graph.V {
		return nil
	}

	return ngraph
}

type Node struct {
	from int
	edge *grph.Edge
}

type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].edge.Weight < h[j].edge.Weight
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
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
}
