package main

import "fmt"

// O(E + V) time | O(V) space
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 || len(prerequisites) == 0 {
		return true
	}

	adjList := createAdjList(numCourses, prerequisites)
	visited := make([]int, numCourses)

	for v := 0; v < numCourses; v++ {
		if visited[v] == 0 {
			if canNotFinishDfs(adjList, visited, v) {
				return false
			}
		}
	}
	return true
}

func canNotFinishDfs(adjList [][]int, visited []int, v int) bool {
	visited[v] = 1

	for _, u := range adjList[v] {
		if visited[u] == 1 {
			return true
		}
		if visited[u] == 0 {
			if canNotFinishDfs(adjList, visited, u) {
				return true
			}
		}
	}

	visited[v] = 2
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

// O(E + V) time | O(V) space
func canFinish1(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 || len(prerequisites) == 0 {
		return true
	}

	adjList := createAdjList(numCourses, prerequisites)

	cnt := make([]int, numCourses)
	for v := 0; v < numCourses; v++ {
		for _, u := range adjList[v] {
			cnt[u]++
		}
	}

	isCircle := true
	queue := make([]int, 0)

	for v := 0; v < len(cnt); v++ {
		if cnt[v] == 0 {
			queue = append(queue, v)
			isCircle = false
		}
	}

	if isCircle {
		return false
	}

	idx := 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		idx++

		for _, u := range adjList[v] {
			cnt[u]--

			if cnt[u] == 0 {
				queue = append(queue, u)
			}
		}
	}

	if idx != numCourses {
		return false
	}

	return true
}

func main() {

	fmt.Println(
		canFinish(1, [][]int{}),
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
		canFinish1(1, [][]int{}),
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
