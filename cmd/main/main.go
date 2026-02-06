package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/sanantoha/go-algos/internals/list"
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

	common := &list.ListNode{Val: 8, Next: &list.ListNode{Val: 4, Next: &list.ListNode{Val: 5}}}
	l1 := &list.ListNode{Val: 4, Next: &list.ListNode{Val: 1, Next: common}}
	l2 := &list.ListNode{Val: 5, Next: &list.ListNode{Val: 6, Next: &list.ListNode{Val: 1, Next: common}}}

	fmt.Println(getIntersectionNode(l1, l2))
	fmt.Println(getIntersectionNode1(l1, l2))
}

func getIntersectionNode(l *list.ListNode, r *list.ListNode) *list.ListNode {
	return nil
}

func getIntersectionNode1(l *list.ListNode, r *list.ListNode) *list.ListNode {
	return nil
}
