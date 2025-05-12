package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	testTable := []struct {
		name  string
		graph map[string][]string
		exp   []string
		err   error
	}{
		{
			name:  "happy path",
			graph: createGraph(),
			exp:   []string{"A", "B", "C", "D"},
			err:   nil,
		},
		{
			name:  "cycle in the graph",
			graph: createCyclicGraph(),
			exp:   nil,
			err:   errors.New("circle in the graph C"),
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res, err := sort(subtest.graph)

			if subtest.err != nil && (err == nil || !strings.Contains(err.Error(), "circle in the graph")) {
				t.Errorf("expected %v, but got %v", subtest.err, err)
			}

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func TestSortIter(t *testing.T) {
	testTable := []struct {
		name  string
		graph map[string][]string
		exp   []string
		err   error
	}{
		{
			name:  "happy path",
			graph: createGraph(),
			exp:   []string{"A", "B", "C", "D"},
			err:   nil,
		},
		{
			name:  "cycle in the graph",
			graph: createCyclicGraph(),
			exp:   nil,
			err:   errors.New("circle in the graph C"),
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res, err := sortIter(subtest.graph)

			if subtest.err != nil && (err == nil || !strings.Contains(err.Error(), "circle in the graph")) {
				t.Errorf("expected %v, but got %v", subtest.err, err)
			}

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func createGraph() map[string][]string {
	graph := make(map[string][]string, 0)
	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"C"}
	graph["C"] = []string{"D"}
	graph["D"] = []string{}
	return graph
}

func createCyclicGraph() map[string][]string {
	graph := make(map[string][]string, 0)
	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"C"}
	graph["C"] = []string{"D"}
	graph["D"] = []string{"A", "B"}
	return graph
}
