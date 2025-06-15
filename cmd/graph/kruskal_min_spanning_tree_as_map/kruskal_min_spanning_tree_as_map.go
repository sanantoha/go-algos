package main

import (
	"container/heap"
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"sort"
)

// O(E * log(E)) time | O(E + V) space
func mst(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {

	ngraph := make(map[string][]*grph.EdgeT[string], 0)

	parents := makeSet(graph)
	rank := makeRank(graph)

	edges := getEdges(graph)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	for _, minEdge := range edges {
		v := minEdge.From()
		u := minEdge.To()

		pv := find(parents, v)
		pu := find(parents, u)

		if pv != pu {
			nedge := *grph.NewEdgeT[string](v, u, minEdge.Weight)
			fromLst := ngraph[v]
			fromLst = append(fromLst, &nedge)
			ngraph[v] = fromLst
			toLst := ngraph[u]
			toLst = append(toLst, &nedge)
			ngraph[u] = toLst
			union(parents, rank, pv, pu)
		}
	}

	return ngraph
}

// O(E * log(E)) time | O(E + V) space
func mst1(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {
	ngraph := make(map[string][]*grph.EdgeT[string], 0)
	//
	parents := makeSet(graph)
	rank := makeRank(graph)

	h := &EdgeHeap{}

	for _, edges := range graph {
		for _, edge := range edges {
			heap.Push(h, edge)
		}
	}

	for h.Len() > 0 {
		minEdge := heap.Pop(h).(*grph.EdgeT[string])
		v := minEdge.From()
		u := minEdge.To()

		pv := find(parents, v)
		pu := find(parents, u)

		if pv != pu {
			nedge := *grph.NewEdgeT[string](v, u, minEdge.Weight)
			fromLst := ngraph[v]
			fromLst = append(fromLst, &nedge)
			ngraph[v] = fromLst
			toLst := ngraph[u]
			toLst = append(toLst, &nedge)
			ngraph[u] = toLst
			union(parents, rank, pv, pu)
		}
	}

	return ngraph
}

func union(parents map[string]string, ranks map[string]int, v string, u string) {
	if ranks[v] > ranks[u] {
		parents[u] = v
	} else if ranks[v] < ranks[u] {
		parents[v] = u
	} else {
		parents[v] = u
		ranks[u]++
	}
}

func find(parents map[string]string, v string) string {
	if parents[v] == v {
		return v
	} else {
		parents[v] = find(parents, parents[v])
		return parents[v]
	}
}

func getEdges(graph map[string][]*grph.EdgeT[string]) []*grph.EdgeT[string] {

	edges := make([]*grph.EdgeT[string], 0)
	seen := make(map[*grph.EdgeT[string]]bool, 0)

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

func makeRank(graph map[string][]*grph.EdgeT[string]) map[string]int {
	ranks := make(map[string]int, len(graph))

	for v, _ := range graph {
		ranks[v] = 0
	}

	return ranks
}

func makeSet(graph map[string][]*grph.EdgeT[string]) map[string]string {
	parents := make(map[string]string, len(graph))

	for v, _ := range graph {
		parents[v] = v
	}

	return parents
}

type EdgeHeap []*grph.EdgeT[string]

func (h EdgeHeap) Len() int {
	return len(h)
}

func (h EdgeHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h EdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(*grph.EdgeT[string]))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func main() {

	graph := createGraph()
	/*
	   6 5
	   0: 0-1 7.00
	   1: 1-2 3.00  0-1 7.00
	   2: 1-2 3.00  2-4 3.00
	   3: 3-4 2.00
	   4: 3-4 2.00  4-5 2.00  2-4 3.00
	   5: 4-5 2.00
	*/

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst(graph)))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst1(graph)))
	fmt.Println()
	fmt.Println("=========================================")

	graph1 := createGraph1()

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
	fmt.Println(grph.PrintGraphAsAdjList(graph1))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst(graph1)))
	fmt.Println("=========================================")
	fmt.Println(grph.PrintGraphAsAdjList(mst1(graph1)))
}

func createGraph() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 7.0)
	edge02 := grph.NewEdgeT("0", "2", 8.0)
	edge12 := grph.NewEdgeT("1", "2", 3.0)
	edge13 := grph.NewEdgeT("1", "3", 6.0)
	edge23 := grph.NewEdgeT("2", "3", 4.0)
	edge24 := grph.NewEdgeT("2", "4", 3.0)
	edge34 := grph.NewEdgeT("3", "4", 2.0)
	edge35 := grph.NewEdgeT("3", "5", 5.0)
	edge45 := grph.NewEdgeT("4", "5", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge02}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge12, edge13}
	graph["2"] = []*grph.EdgeT[string]{edge02, edge12, edge23, edge24}
	graph["3"] = []*grph.EdgeT[string]{edge13, edge23, edge34, edge35}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34, edge45}
	graph["5"] = []*grph.EdgeT[string]{edge35, edge45}
	return graph
}

func createGraph1() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 2.0)
	edge02 := grph.NewEdgeT("0", "2", 3.0)
	edge03 := grph.NewEdgeT("0", "3", 7.0)
	edge12 := grph.NewEdgeT("1", "2", 6.0)
	edge16 := grph.NewEdgeT("1", "6", 3.0)
	edge24 := grph.NewEdgeT("2", "4", 1.0)
	edge25 := grph.NewEdgeT("2", "5", 8.0)
	edge34 := grph.NewEdgeT("3", "4", 5.0)
	edge45 := grph.NewEdgeT("4", "5", 4.0)
	edge56 := grph.NewEdgeT("5", "6", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge02, edge03}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge12, edge16}
	graph["2"] = []*grph.EdgeT[string]{edge02, edge12, edge24, edge25}
	graph["3"] = []*grph.EdgeT[string]{edge03, edge34}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34, edge45}
	graph["5"] = []*grph.EdgeT[string]{edge25, edge45, edge56}
	graph["6"] = []*grph.EdgeT[string]{edge16, edge56}

	return graph
}
