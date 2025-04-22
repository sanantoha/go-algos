package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"sort"
	"testing"
)

func TestMstGraph(t *testing.T) {
	graph := createGraph()
	expGraph := expGraph()

	resGraph := mst(graph)

	if !equalMaps(expGraph, resGraph) {
		t.Errorf("expected:\n%v, but got:\n%v", grph.PrintGraphAsAdjList(expGraph), grph.PrintGraphAsAdjList(resGraph))
	}
}

func TestMstGraph1(t *testing.T) {
	graph := createGraph1()
	expGraph := expGraph1()

	resGraph := mst(graph)

	if !equalMaps(expGraph, resGraph) {
		t.Errorf("expected:\n%v, but got:\n%v", grph.PrintGraphAsAdjList(expGraph), grph.PrintGraphAsAdjList(resGraph))
	}
}

func expGraph() map[string][]*grph.EdgeT[string] {
	/**
	  6 5
	  0: 0-1 7.00
	  1: 1-2 3.00  0-1 7.00
	  2: 1-2 3.00  2-4 3.00
	  3: 3-4 2.00
	  4: 3-4 2.00  4-5 2.00  2-4 3.00
	  5: 4-5 2.00
	*/
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 7.0)
	edge12 := grph.NewEdgeT("1", "2", 3.0)
	edge24 := grph.NewEdgeT("2", "4", 3.0)
	edge34 := grph.NewEdgeT("3", "4", 2.0)
	edge45 := grph.NewEdgeT("4", "5", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge12}
	graph["2"] = []*grph.EdgeT[string]{edge12, edge24}
	graph["3"] = []*grph.EdgeT[string]{edge34}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34, edge45}
	graph["5"] = []*grph.EdgeT[string]{edge45}
	return graph
}

func expGraph1() map[string][]*grph.EdgeT[string] {
	/*
	   7 6
	   0: 0-1 2.00  0-2 3.00
	   1: 0-1 2.00  1-6 3.00
	   2: 2-4 1.00  0-2 3.00
	   3: 3-4 5.00
	   4: 2-4 1.00  3-4 5.00
	   5: 5-6 2.00
	   6: 5-6 2.00  1-6 3.00
	*/

	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 2.0)
	edge02 := grph.NewEdgeT("0", "2", 3.0)
	edge16 := grph.NewEdgeT("1", "6", 3.0)
	edge24 := grph.NewEdgeT("2", "4", 1.0)
	edge34 := grph.NewEdgeT("3", "4", 5.0)
	edge56 := grph.NewEdgeT("5", "6", 2.0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge02}
	graph["1"] = []*grph.EdgeT[string]{edge01, edge16}
	graph["2"] = []*grph.EdgeT[string]{edge02, edge24}
	graph["3"] = []*grph.EdgeT[string]{edge34}
	graph["4"] = []*grph.EdgeT[string]{edge24, edge34}
	graph["5"] = []*grph.EdgeT[string]{edge56}
	graph["6"] = []*grph.EdgeT[string]{edge16, edge56}

	return graph
}

func equalMaps(a, b map[string][]*grph.EdgeT[string]) bool {
	if len(a) != len(b) {
		return false
	}

	for key, aSlice := range a {
		bSlice, ok := b[key]
		if !ok || len(aSlice) != len(bSlice) {
			return false
		}

		sort.Slice(aSlice, func(i, j int) bool {
			return aSlice[i].Weight < aSlice[j].Weight
		})

		sort.Slice(bSlice, func(i, j int) bool {
			return bSlice[i].Weight < bSlice[j].Weight
		})

		for i := range aSlice {
			if aSlice[i] == nil || bSlice[i] == nil {
				if aSlice[i] != bSlice[i] {
					return false
				}
				continue
			}

			if !reflect.DeepEqual(*aSlice[i], *bSlice[i]) {
				return false
			}
		}
	}

	return true
}
