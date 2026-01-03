package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 || len(prerequisites) == 0 {
		return true
	}

	adjList := createAdjList(numCourses, prerequisites)
	fmt.Println(adjList)
	return false
}

func createAdjList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)

	for i, _ := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, p := range prerequisites {
		adjList[p[0]] = append(adjList[p[0]], p[1])
	}

	return adjList
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	return false
}

func main() {

	fmt.Println(
		canFinish(2, [][]int{}),
	)

	fmt.Println(
		canFinish(2, [][]int{
			{1, 0},
		}),
	)

	fmt.Println(
		!canFinish(2, [][]int{
			{1, 0},
			{0, 1},
		}),
	)

	fmt.Println(
		!canFinish(4, [][]int{
			{1, 0},
			{2, 1},
			{3, 2},
			{0, 3},
		}),
	)

	fmt.Println(
		canFinish(4, [][]int{
			{1, 0},
			{2, 1},
			{3, 2},
		}),
	)

	fmt.Println("================================================")

	fmt.Println(
		canFinish1(2, [][]int{}),
	)

	fmt.Println(
		canFinish1(2, [][]int{
			{1, 0},
		}),
	)

	fmt.Println(
		!canFinish1(2, [][]int{
			{1, 0},
			{0, 1},
		}),
	)

	fmt.Println(
		!canFinish1(4, [][]int{
			{1, 0},
			{2, 1},
			{3, 2},
			{0, 3},
		}),
	)

	fmt.Println(
		canFinish1(4, [][]int{
			{1, 0},
			{2, 1},
			{3, 2},
		}),
	)
}
