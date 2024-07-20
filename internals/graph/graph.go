package graph

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	V      int
	W      int
	Weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	if v < 0 || w < 0 {
		log.Fatalln(errors.New("vertex names must be nonnegative integers"))
	}
	if math.IsNaN(weight) {
		log.Fatalln(errors.New("Weight is NaN"))
	}

	instance := Edge{
		V:      v,
		W:      w,
		Weight: weight,
	}

	return &instance
}

func (e *Edge) Either() int {
	return e.V
}

func (e *Edge) Other(v int) int {
	if v == e.V {
		return e.W
	} else if v == e.W {
		return e.V
	} else {
		panic("Illegal state")
	}
}

func (e *Edge) CompareTo(other *Edge) int {
	if e.Weight < other.Weight {
		return -1
	} else if e.Weight > other.Weight {
		return 1
	} else {
		return 0
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.V, e.W, e.Weight)
}

type DirectedEdge Edge

func NewDirectedEdge(v, w int, weight float64) (*DirectedEdge, error) {
	if v < 0 || w < 0 {
		return nil, errors.New("vertex names must be nonnegative integers")
	}
	if math.IsNaN(weight) {
		return nil, errors.New("Weight is NaN")
	}

	instance := DirectedEdge{
		V:      v,
		W:      w,
		Weight: weight,
	}
	return &instance, nil
}

func (e *DirectedEdge) From() int {
	return e.V
}

func (e *DirectedEdge) To() int {
	return e.W
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %5.2f", e.V, e.W, e.Weight)
}

type Digraph struct {
	V   int
	E   int
	adj [][]int
}

func NewDigraph(V int) (*Digraph, error) {
	if V < 0 {
		return nil, errors.New("Number of vertices in a Digraph must be nonnegative")
	}

	adj := make([][]int, V)
	for v := 0; v < V; v++ {
		adj[v] = make([]int, 0)
	}

	graph := Digraph{
		V:   V,
		E:   0,
		adj: adj,
	}
	return &graph, nil
}

func NewDigraphFromFile(filePath string) (*Digraph, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	str := string(file)

	scanner := bufio.NewScanner(strings.NewReader(str))

	scanner.Scan()
	V, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	graph, err := NewDigraph(V)
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	E, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, err
	}

	for i := 0; i < E; i++ {
		scanner.Scan()

		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, errors.New("invalid edge format, it should contains from, to vertexes")
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil || v < 0 || v >= V {
			return nil, errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[0], V-1))
		}
		w, err := strconv.Atoi(parts[1])
		if err != nil || w < 0 || w >= V {
			return nil, errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[1], V-1))
		}

		err = graph.AddEdge(v, w)
		if err != nil {
			return nil, err
		}
	}

	return graph, nil
}

func (g *Digraph) AddEdge(v int, w int) error {
	if v < 0 || v >= g.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	if w < 0 || w >= g.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", w, g.V-1))
	}
	g.adj[v] = append(g.adj[v], w)
	g.E++
	return nil
}

func (g *Digraph) Adj(v int) ([]int, error) {
	if v < 0 || v >= g.V {
		return nil, errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	return g.adj[v], nil
}

func (g *Digraph) String() string {
	var builder = strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d vertices, %d edges\n", g.V, g.E))

	for v := 0; v < g.V; v++ {
		builder.WriteString(fmt.Sprintf("%d: ", v))
		for _, w := range g.adj[v] {
			builder.WriteString(fmt.Sprintf("%d ", w))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

type EdgeWeightedGraph struct {
	V   int
	E   int
	adj [][]*Edge
}

func NewEdgeWeightedGraph(V int) *EdgeWeightedGraph {
	if V < 0 {
		log.Fatalln(errors.New("Number of vertices in a Digraph must be nonnegative"))
	}
	adj := make([][]*Edge, V)
	for v := 0; v < V; v++ {
		adj[v] = make([]*Edge, 0)
	}

	graph := EdgeWeightedGraph{
		V:   V,
		E:   0,
		adj: adj,
	}
	return &graph
}

func NewEdgeWeightedGraphFromFile(filePath string) *EdgeWeightedGraph {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	str := string(file)

	scanner := bufio.NewScanner(strings.NewReader(str))

	scanner.Scan()
	V, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	graph := NewEdgeWeightedGraph(V)
	if err != nil {
		log.Fatalln(err)
	}

	scanner.Scan()
	E, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < E; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			log.Fatalln(errors.New("invalid edge format, it should contains from, to vertexes and edge weight"))
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil || v < 0 || v >= V {
			log.Fatalln(errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[0], V-1)))
		}
		w, err := strconv.Atoi(parts[1])
		if err != nil || w < 0 || w >= V {
			log.Fatalln(errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[1], V-1)))
		}

		weight, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			log.Fatalln(errors.New(fmt.Sprintf("Invalid weight %v", parts[2])))
		}

		edge := Edge{
			V:      v,
			W:      w,
			Weight: weight,
		}
		graph.AddEdge(&edge)
	}

	return graph
}

func (g *EdgeWeightedGraph) AddEdge(edge *Edge) {
	v := edge.Either()
	w := edge.Other(v)
	if v < 0 || v >= g.V {
		log.Fatalln(errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1)))
	}
	if w < 0 || w >= g.V {
		log.Fatalln(errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", w, g.V-1)))
	}
	g.adj[v] = append(g.adj[v], edge)
	g.adj[w] = append(g.adj[w], edge)
	g.E++
}

func (g *EdgeWeightedGraph) String() string {
	var builder = strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d %d\n", g.V, g.E))

	for v := 0; v < g.V; v++ {
		builder.WriteString(fmt.Sprintf("%d: ", v))
		for _, edge := range g.adj[v] {
			builder.WriteString(fmt.Sprintf("%v   ", edge))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type EdgeWeightedDigraph struct {
	V   int
	E   int
	adj [][]*DirectedEdge
}

func NewEdgeWeightedDigraph(V int) (*EdgeWeightedDigraph, error) {
	if V < 0 {
		return nil, errors.New("Number of vertices in a Digraph must be nonnegative")
	}
	adj := make([][]*DirectedEdge, V)
	for v := 0; v < V; v++ {
		adj[v] = make([]*DirectedEdge, 0)
	}

	graph := EdgeWeightedDigraph{
		V:   V,
		E:   0,
		adj: adj,
	}
	return &graph, nil
}

func NewEdgeWeightedDigraphFromFile(filePath string) (*EdgeWeightedDigraph, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	str := string(file)

	scanner := bufio.NewScanner(strings.NewReader(str))

	scanner.Scan()
	V, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	graph, err := NewEdgeWeightedDigraph(V)
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	E, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	for i := 0; i < E; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, errors.New("invalid edge format, it should contains from, to vertexes and edge weight")
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil || v < 0 || v >= V {
			return nil, errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[0], V-1))
		}
		w, err := strconv.Atoi(parts[1])
		if err != nil || w < 0 || w >= V {
			return nil, errors.New(fmt.Sprintf("Can not parse vertex %v or it is not between 0 and %d", parts[1], V-1))
		}

		weight, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid weight %v", parts[2]))
		}

		edge, err := NewDirectedEdge(v, w, weight)
		if err != nil {
			return nil, err
		}
		graph.AddEdge(edge)
	}

	return graph, nil
}

func (g *EdgeWeightedDigraph) AddEdge(edge *DirectedEdge) {
	v := edge.From()
	g.adj[v] = append(g.adj[v], edge)
	g.E++
}

func (g *EdgeWeightedDigraph) Adj(v int) ([]*DirectedEdge, error) {
	if v < 0 || v >= g.V {
		return nil, errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	return g.adj[v], nil
}

func (g *EdgeWeightedDigraph) Edges() []*DirectedEdge {
	list := make([]*DirectedEdge, 0)

	for v := 0; v < g.V; v++ {
		for _, edge := range g.adj[v] {
			list = append(list, edge)
		}
	}

	return list
}

func (g *EdgeWeightedDigraph) Outdegree(v int) (int, error) {
	if v < 0 || v >= g.V {
		return 0, errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	return len(g.adj[v]), nil
}

func (g *EdgeWeightedDigraph) String() string {
	var builder = strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d %d\n", g.V, g.E))

	for v := 0; v < g.V; v++ {
		builder.WriteString(fmt.Sprintf("%d: ", v))
		for _, edge := range g.adj[v] {
			builder.WriteString(fmt.Sprintf("%v   ", edge))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type ShortestPath struct {
	Shortest []float64
	Prev     []int
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: make([]*Node, 0),
	}
}
