package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	tasks := make([]string, 0)

	err := filepath.WalkDir("./cmd", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil // Continue walking
		}
		if !d.IsDir() { // Print only files
			//fmt.Println(path)
			fileName := filepath.Base(path)
			tasks = append(tasks, fileName)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

	processedTasks := map[string]bool{
		"symmetric_tree.go":         true,
		"prim_min_spanning_tree.go": true,
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if strings.HasSuffix(task, ".txt") || strings.HasSuffix(task, "test.go") || task == "main.go" {
			continue
		}
		if _, exists := processedTasks[task]; exists {
			continue
		}
		fmt.Println(task)
	}

	runTask()
}

func runTask() {

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

func mst(graph *grph.EdgeWeightedGraph) *grph.EdgeWeightedGraph {
	return nil
}
