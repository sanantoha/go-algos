package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	"strconv"
	"testing"
)

func TestMst(t *testing.T) {

	funcs := []func(map[string][]*grph.EdgeT[string]) map[string][]*grph.EdgeT[string]{
		mst,
		mst1,
	}

	tests := []struct {
		name     string
		graph    map[string][]*grph.EdgeT[string]
		expGraph map[string][]*grph.EdgeT[string]
	}{
		{
			name:     "graph 1",
			graph:    createGraph(),
			expGraph: expGraph(),
		},
		{
			name:     "graph 2",
			graph:    createGraph1(),
			expGraph: expGraph1(),
		},
	}

	for i, fn := range funcs {
		for _, tt := range tests {
			t.Run(tt.name+" func "+strconv.Itoa(i), func(t *testing.T) {
				if got := fn(tt.graph); !grph.EqualMaps(tt.expGraph, got) {
					t.Errorf("expected:\n%v, but got:\n%v", grph.PrintGraphAsAdjList(tt.expGraph), grph.PrintGraphAsAdjList(got))
				}
			})
		}
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
