package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"strconv"
	"testing"
)

func TestDfs(t *testing.T) {
	graph := createGraph()
	fmt.Println(grph.PrintGraphAsAdjList(graph))

	start := "0"

	exp := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}

	funcs := []func(map[string][]*grph.EdgeT[string], string) []string{
		dfsRec,
		dfsIter,
	}

	for i, fn := range funcs {
		t.Run("dfs"+strconv.Itoa(i), func(t *testing.T) {
			res := fn(graph, start)
			fmt.Println(res)

			if !reflect.DeepEqual(exp, res) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}

}

func createGraph() map[string][]*grph.EdgeT[string] {
	graph := make(map[string][]*grph.EdgeT[string], 0)

	edge01 := grph.NewEdgeT("0", "1", 0)
	edge12 := grph.NewEdgeT("1", "2", 0)
	edge23 := grph.NewEdgeT("2", "3", 0)
	edge34 := grph.NewEdgeT("3", "4", 0)
	edge25 := grph.NewEdgeT("3", "5", 0)
	edge16 := grph.NewEdgeT("1", "6", 0)
	edge07 := grph.NewEdgeT("0", "7", 0)
	edge78 := grph.NewEdgeT("7", "8", 0)

	edge89 := grph.NewEdgeT("8", "9", 0)
	edge910 := grph.NewEdgeT("9", "10", 0)
	edge811 := grph.NewEdgeT("8", "11", 0)
	edge1112 := grph.NewEdgeT("11", "12", 0)
	edge1213 := grph.NewEdgeT("12", "13", 0)
	edge1314 := grph.NewEdgeT("13", "14", 0)

	graph["0"] = []*grph.EdgeT[string]{edge01, edge07}
	graph["1"] = []*grph.EdgeT[string]{edge12, edge16}
	graph["2"] = []*grph.EdgeT[string]{edge23, edge25}
	graph["3"] = []*grph.EdgeT[string]{edge34}
	graph["7"] = []*grph.EdgeT[string]{edge78}
	graph["8"] = []*grph.EdgeT[string]{edge89, edge811}
	graph["9"] = []*grph.EdgeT[string]{edge910}
	graph["11"] = []*grph.EdgeT[string]{edge1112}
	graph["12"] = []*grph.EdgeT[string]{edge1213}
	graph["13"] = []*grph.EdgeT[string]{edge1314}
	return graph
}
