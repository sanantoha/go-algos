package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
)

func mst(graph map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string] {
	return nil
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
	fmt.Println(mst(graph))
	fmt.Println("=========================================\n\n")

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
	fmt.Println(mst(graph1))
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
