package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	exp := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
		"15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}

	graph := map[string][]*grph.EdgeT[string]{
		"2": {
			grph.NewEdgeT("2", "6", 0.00),
			grph.NewEdgeT("2", "7", 0.00),
			grph.NewEdgeT("2", "8", 0.00),
		},
		"9": {
			grph.NewEdgeT("9", "18", 0.00),
			grph.NewEdgeT("9", "19", 0.00),
		},
		"3": {
			grph.NewEdgeT("3", "9", 0.00),
			grph.NewEdgeT("3", "10", 0.00),
		},
		"4": {
			grph.NewEdgeT("4", "11", 0.00),
			grph.NewEdgeT("4", "12", 0.00),
		},
		"5": {
			grph.NewEdgeT("5", "13", 0.00),
			grph.NewEdgeT("5", "14", 0.00),
		},
		"6": {
			grph.NewEdgeT("6", "15", 0.00),
		},
		"7": {
			grph.NewEdgeT("7", "16", 0.00),
		},
		"8": {
			grph.NewEdgeT("8", "17", 0.00),
		},
		"0": {
			grph.NewEdgeT("0", "1", 0.00),
			grph.NewEdgeT("0", "2", 0.00),
		},
		"1": {
			grph.NewEdgeT("1", "3", 0.00),
			grph.NewEdgeT("1", "4", 0.00),
			grph.NewEdgeT("1", "5", 0.00),
		},
		"10": {
			grph.NewEdgeT("10", "20", 0.00),
			grph.NewEdgeT("10", "21", 0.00),
			grph.NewEdgeT("10", "22", 0.00),
			grph.NewEdgeT("10", "23", 0.00),
			grph.NewEdgeT("10", "24", 0.00),
		},
	}
	fmt.Println(grph.PrintGraphAsAdjList(graph))

	res := bfs(graph, "0")

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
