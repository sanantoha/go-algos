package main

import (
	"container/heap"
	"fmt"
	"math"
	"reflect"
	"slices"
)

// O(w * h * log(w * h)) time | O(w * h) space
func aStarAlgorithm(startRow int, startCol int, endRow int, endCol int, graph [][]int) [][]int {
	nodes := initNodes(graph)
	endNode := nodes[endRow][endCol]
	startNode := nodes[startRow][startCol]
	startNode.distanceFromStart = 0
	startNode.distanceToEnd = manhattanDistance(startNode, endNode)

	h := &NodeHeap{}
	heap.Push(h, startNode)

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*Node)

		if minNode == endNode {
			break
		}

		tentativeDistance := minNode.distanceFromStart + 1

		for _, neighbor := range getNeighbors(nodes, minNode) {
			if neighbor.val == 1 {
				continue
			}

			if tentativeDistance >= neighbor.distanceFromStart {
				continue
			}
			neighbor.distanceFromStart = tentativeDistance
			neighbor.distanceToEnd = tentativeDistance + manhattanDistance(neighbor, endNode)
			neighbor.cameFrom = minNode

			if !h.Update(neighbor) {
				heap.Push(h, neighbor)
			}
		}
	}
	return reconstruct(endNode)
}

func reconstruct(endNode *Node) [][]int {
	res := make([][]int, 0)

	node := endNode

	for node != nil {
		res = append(res, []int{node.x, node.y})
		node = node.cameFrom
	}

	slices.Reverse(res)
	return res
}

func getNeighbors(nodes [][]*Node, node *Node) []*Node {
	res := make([]*Node, 0)
	row := node.x
	col := node.y

	if row > 0 {
		res = append(res, nodes[row-1][col])
	}
	if col > 0 {
		res = append(res, nodes[row][col-1])
	}
	if row+1 < len(nodes) {
		res = append(res, nodes[row+1][col])
	}
	if col+1 < len(nodes[row]) {
		res = append(res, nodes[row][col+1])
	}
	return res
}

func manhattanDistance(start, end *Node) int {
	return abs(start.x-end.x) + abs(start.y-end.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func initNodes(graph [][]int) [][]*Node {
	nodes := make([][]*Node, 0)

	for i := 0; i < len(graph); i++ {
		nodes = append(nodes, make([]*Node, 0))
		for j := 0; j < len(graph[i]); j++ {
			nodes[i] = append(nodes[i], NewNode(i, j, graph[i][j]))
		}
	}
	return nodes
}

type Node struct {
	x                 int
	y                 int
	val               int
	distanceFromStart int
	distanceToEnd     int
	cameFrom          *Node
}

func NewNode(x, y, val int) *Node {
	return &Node{
		x:                 x,
		y:                 y,
		val:               val,
		distanceFromStart: math.MaxInt,
		distanceToEnd:     math.MaxInt,
	}
}

type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h *NodeHeap) Update(node *Node) bool {
	for idx, v := range *h {
		if v == node {
			heap.Fix(h, idx)
			return true
		}
	}
	return false
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].distanceToEnd < h[j].distanceToEnd
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func main() {

	startRow := 0
	startCol := 1
	endRow := 4
	endCol := 3

	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}

	expected := [][]int{
		{0, 1},
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{3, 1},
		{4, 1},
		{4, 2},
		{4, 3},
	}

	actual := aStarAlgorithm(startRow, startCol, endRow, endCol, graph)
	fmt.Println(actual)

	fmt.Println(reflect.DeepEqual(expected, actual))
}
