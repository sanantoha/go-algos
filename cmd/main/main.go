package main

import (
	"fmt"
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
		"min_number_of_jumps.go":      true,
		"intersection_linked_list.go": true,
		"find_mode_in_bst.go":         true,
		"heap_sort.go":                true,
		"odd_even_linked_list.go":     true,
		"subarray_sort.go":            true,
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

	output := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	fmt.Println(subarraySort(output))

	output1 := []int{1, 2, 4, 7, 10, 11, 7, 12, 7, 7, 16, 18, 19}
	fmt.Println(subarraySort(output1))

	output2 := []int{1, 2}
	fmt.Println(subarraySort(output2))
}

func subarraySort(arr []int) []int {
	return nil
}
