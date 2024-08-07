package main

import "fmt"

func sort(graph map[string][]string) []string {
	return nil
}

func sortIter(graph map[string][]string) []string {
	return nil
}

func main() {
	graph := make(map[string][]string, 0)

	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"C"}
	graph["C"] = []string{"D"}
	//graph["D"] = []string{"A", "B"}
	graph["D"] = []string{}

	fmt.Println(sort(graph))

	fmt.Println(sortIter(graph))
}
