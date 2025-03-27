package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"path/filepath"
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
		"spiral_matrix_traverse.go":                  true,
		"binary_tree_zigzag_level_order_traverse.go": true,
		"lru_cache.go":                               true,
		"climbing_stars.go":                          true,
		"quick_sort.go":                              true,
		"maximum_subarray.go":                        true,
		"insert_sort.go":                             true,
		"first_duplicate_value.go":                   true,
		"permutations.go":                            true,
		"merge_intervals.go":                         true,
		"river_sizes.go":                             true,
		"bellman_ford_as_map.go":                     true,
		"add_two_numbers.go":                         true,
		"top_k_frequent_words.go":                    true,
		"find_nodes_distance_k.go":                   true,
		"kth_smallest_element.go":                    true,
		"clone_graph.go":                             true,
		"same_tree.go":                               true,
		"palindrome_linked_list.go":                  true,
		"string_permutations.go":                     true,
		"binary_search.go":                           true,
		"string_without_aaa_or_bbb.go":               true,
		"zig_zag_traverse.go":                        true,
		"depth_first_search.go":                      true,
	}

	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	for _, task := range tasks {
		if _, exists := processedTasks[task]; exists {
			continue
		}
		fmt.Println(task)
	}

	runTask()
}

func runTask() {

	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/depth_first_search/dfs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(graph)

	log.Println("\n")
	log.Println("=======================================")
	log.Println(dfsRec(graph, 0))
	log.Println("=======================================")
	log.Println(dfsIter(graph, 0))
}

func dfsRec(graph *grph.EdgeWeightedDigraph, start int) []int {
	return nil
}

func dfsIter(graph *grph.EdgeWeightedDigraph, start int) []int {
	return nil
}
