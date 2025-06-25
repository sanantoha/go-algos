package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	expShortest := map[string]float64{
		"0": 0,
		"1": 5,
		"2": 8,
		"3": 4,
		"4": 7,
	}
	expPrev := map[string]string{
		"0": "",
		"1": "3",
		"2": "1",
		"3": "0",
		"4": "3",
	}

	graph := createGraph()

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	sp, err := findShortestPath(graph, "0")
	if err != nil {
		t.Errorf("Failed to find shortest path: %v", err)
	}
	//   &{map[0:0 1:5 2:8 3:4 4:7] map[0: 1:3 2:1 3:0 4:3]}
	fmt.Println(sp)

	if !reflect.DeepEqual(sp.Fst, expShortest) {
		t.Errorf("expected prev %v, got %v", expShortest, sp.Fst)
	}

	if !reflect.DeepEqual(sp.Snd, expPrev) {
		t.Errorf("expected prev %v, got %v", expPrev, sp.Snd)
	}
}

func createGraph() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 6.0)
	edge03 := grph.NewEdgeT("0", "3", 4.0)

	edge12 := grph.NewEdgeT("1", "2", 3.0)
	edge13 := grph.NewEdgeT("1", "3", 2.0)

	edge24 := grph.NewEdgeT("2", "4", 4.0)

	edge31 := grph.NewEdgeT("3", "1", 1.0)
	edge32 := grph.NewEdgeT("3", "2", 9.0)
	edge34 := grph.NewEdgeT("3", "4", 3.0)

	edge42 := grph.NewEdgeT("4", "2", 5.0)
	edge40 := grph.NewEdgeT("4", "0", 7.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge03}
	graph["1"] = []*grph.EdgeT[string]{edge13, edge12}
	graph["2"] = []*grph.EdgeT[string]{edge24}
	graph["3"] = []*grph.EdgeT[string]{edge31, edge32, edge34}
	graph["4"] = []*grph.EdgeT[string]{edge42, edge40}
	return graph
}
