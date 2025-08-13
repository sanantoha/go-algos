package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/tree"
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
		"symmetric_tree.go":                           true,
		"prim_min_spanning_tree.go":                   true,
		"min_swaps_transform_string_to_palindrome.go": true,
		"lru_cache.go":                                true,
		"heap_sort.go":                                true,
		"reverse_list.go":                             true,
		"levenshtein_distance.go":                     true,
		"bst_successor_search.go":                     true,
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

	node5 := &tree.Node{Key: 5}
	node9 := &tree.Node{Key: 9}
	node11 := &tree.Node{Key: 11}
	node12 := &tree.Node{Key: 12}
	node14 := &tree.Node{Key: 14}
	node20 := &tree.Node{Key: 20}
	node25 := &tree.Node{Key: 25}

	node5.Parent = node9

	node9.Left = node5
	node9.Right = node12
	node9.Parent = node20

	node11.Parent = node12

	node12.Left = node11
	node12.Right = node14
	node12.Parent = node9

	node14.Parent = node12

	node20.Left = node9
	node20.Right = node25

	node25.Parent = node20

	//         20
	//       /    \
	//      9      25
	//    /   \
	//   5     12
	//        /   \
	//       11   14

	fmt.Println(node20)
	fmt.Println(findInOrderSuccessor(node5).Key == 9)
	fmt.Println(findInOrderSuccessor(node9).Key == 11)
	fmt.Println(findInOrderSuccessor(node11).Key == 12)
	fmt.Println(findInOrderSuccessor(node12).Key == 14)
	fmt.Println(findInOrderSuccessor(node14).Key == 20)
	fmt.Println(findInOrderSuccessor(node20).Key == 25)
	fmt.Println(findInOrderSuccessor(node25) == nil)
	fmt.Println(findInOrderSuccessor(node20).Key == 25)
}

func findInOrderSuccessor(root *tree.Node) *tree.Node {
	return nil
}
