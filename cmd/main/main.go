package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
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
		"symmetric_tree.go":                             true,
		"prim_min_spanning_tree.go":                     true,
		"min_swaps_transform_string_to_palindrome.go":   true,
		"lru_cache.go":                                  true,
		"heap_sort.go":                                  true,
		"reverse_list.go":                               true,
		"levenshtein_distance.go":                       true,
		"bst_successor_search.go":                       true,
		"all_paths_from_source_target.go":               true,
		"level_order_binary_tree_traverse.go":           true,
		"all_elements_in_two_binary_search_trees.go":    true,
		"subarray_sort.go":                              true,
		"search_in_rotated_sorted_array.go":             true,
		"branch_sums.go":                                true,
		"sqrt.go":                                       true,
		"populating_next_right_pointer_in_each_node.go": true,
		"disk_stacking.go":                              true,
		"max_sum_increasing_subsequence.go":             true,
		"max_depth_of_bst.go":                           true,
		"knapsack_problem.go":                           true,
		"three_sum.go":                                  true,
		"find_nodes_distance_k.go":                      true,
		"dijkstra_shortest_path_as_map.go":              true,
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

	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/dijkstra_shortest_path/dijkstraShortestPath.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	sp, err := findShortestPath(graph, "0")
	if err != nil {
		log.Fatalln(err)
	}
	//   &{map[0:0 1:5 2:8 3:4 4:7] map[0: 1:3 2:1 3:0 4:3]}
	log.Println(sp)
}

func findShortestPath(graph map[string][]*grph.EdgeT[string], start string) (*grph.Pair[map[string]float64, map[string]string], error) {
	return nil, nil
}
