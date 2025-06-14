package graph

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	V      int
	W      int
	Weight float64
}

func NewEdgeWithValidation(v, w int, weight float64) (*Edge, error) {
	if v < 0 || w < 0 {
		return nil, errors.New("vertex names must be nonnegative integers")
	}
	if math.IsNaN(weight) {
		return nil, errors.New("Weight is NaN")
	}

	instance := Edge{
		V:      v,
		W:      w,
		Weight: weight,
	}

	return &instance, nil
}

func NewEdge(v, w int, weight float64) *Edge {
	edge, err := NewEdgeWithValidation(v, w, weight)
	if err != nil {
		log.Fatalln(err)
	}
	return edge
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

func NewEdgeWeightedGraphWithValidation(V int) (*EdgeWeightedGraph, error) {
	if V < 0 {
		return nil, errors.New("Number of vertices in a Digraph must be nonnegative")
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
	return &graph, nil
}

func NewEdgeWeightedGraph(V int) *EdgeWeightedGraph {
	graph, err := NewEdgeWeightedGraphWithValidation(V)
	if err != nil {
		log.Fatalln(err)
	}
	return graph
}

func NewEdgeWeightedGraphFromFile(filePath string) (*EdgeWeightedGraph, error) {
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

	graph := NewEdgeWeightedGraph(V)

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

		edge := Edge{
			V:      v,
			W:      w,
			Weight: weight,
		}
		err = graph.AddEdge(&edge)
		if err != nil {
			return nil, err
		}
	}

	return graph, nil
}

func (g *EdgeWeightedGraph) AddEdge(edge *Edge) error {
	v := edge.Either()
	w := edge.Other(v)
	if v < 0 || v >= g.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	if w < 0 || w >= g.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", w, g.V-1))
	}
	g.adj[v] = append(g.adj[v], edge)
	g.adj[w] = append(g.adj[w], edge)
	g.E++
	return nil
}

func (g *EdgeWeightedGraph) AdjSafe(v int) ([]*Edge, error) {
	if v < 0 || v >= g.V {
		return nil, errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, g.V-1))
	}
	return g.adj[v], nil
}

func (g *EdgeWeightedGraph) Adj(v int) []*Edge {
	edges, err := g.AdjSafe(v)
	if err != nil {
		log.Fatalln(err)
	}
	return edges
}

func (g *EdgeWeightedGraph) Edges() []*Edge {
	mp := make(map[*Edge]struct{}, 0)

	for v := 0; v < g.V; v++ {
		for _, edge := range g.adj[v] {
			mp[edge] = struct{}{}
		}
	}

	edges := make([]*Edge, 0)

	for k, _ := range mp {
		edges = append(edges, k)
	}
	return edges
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

type EdgeT[T comparable] struct {
	V      T
	U      T
	Weight float64
}

func NewEdgeTSafe[T comparable](v T, u T, weight float64) (*EdgeT[T], error) {
	if math.IsNaN(weight) {
		return nil, errors.New("Weight is NaN")
	}
	return &EdgeT[T]{
		V:      v,
		U:      u,
		Weight: weight,
	}, nil
}

func NewEdgeT[T comparable](v T, u T, weight float64) *EdgeT[T] {
	edge, err := NewEdgeTSafe[T](v, u, weight)
	if err != nil {
		log.Fatalln(err)
	}
	return edge
}

func (e *EdgeT[T]) From() T {
	return e.V
}

func (e *EdgeT[T]) To() T {
	return e.U
}

func (e *EdgeT[T]) Either() T {
	return e.V
}

func (e *EdgeT[T]) Other(t T) T {
	if e.V == t {
		return e.U
	} else if e.U == t {
		return e.V
	}
	log.Fatalf("undefined vertex %s\n", t)
	return e.V
}

func NewGraphAsAdjListFromFile(filePath string) (map[string][]*EdgeT[string], error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	str := string(file)

	scanner := bufio.NewScanner(strings.NewReader(str))

	graph := make(map[string][]*EdgeT[string], 0)

	scanner.Scan() // skip V
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

		v := strings.TrimSpace(parts[0])
		w := strings.TrimSpace(parts[1])

		weight, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid weight %v", parts[2]))
		}

		edge, err := NewEdgeTSafe[string](v, w, weight)
		if err != nil {
			return nil, err
		}
		lst, ok := graph[v]
		if !ok {
			lst = make([]*EdgeT[string], 0)
		}
		lst = append(lst, edge)
		graph[v] = lst
	}

	return graph, nil
}

func PrintGraphAsAdjList(graph map[string][]*EdgeT[string]) string {
	var builder = strings.Builder{}

	edges := make(map[*EdgeT[string]]struct{}, 0)

	var subRes = strings.Builder{}
	for key, lst := range graph {
		subRes.WriteString(fmt.Sprintf("%s: ", key))
		for _, edge := range lst {
			subRes.WriteString(fmt.Sprintf("%s->%s %.2f  ", edge.V, edge.U, edge.Weight))
			edges[edge] = struct{}{}
		}
		subRes.WriteString("\n")

	}

	builder.WriteString(fmt.Sprintf("%d %d\n", len(graph), len(edges)))
	builder.WriteString(subRes.String())

	return builder.String()
}

type ShortestPath struct {
	Shortest []float64
	Prev     []int
}

type Pair[T any, E any] struct {
	Fst T
	Snd E
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

func EqualMaps(a, b map[string][]*EdgeT[string]) bool {
	if len(a) != len(b) {
		return false
	}

	for key, aSlice := range a {
		bSlice, ok := b[key]
		if !ok || len(aSlice) != len(bSlice) {
			return false
		}

		sortEdgeT(aSlice)
		sortEdgeT(bSlice)

		for i := range aSlice {
			if aSlice[i] == nil || bSlice[i] == nil {
				if aSlice[i] != bSlice[i] {
					return false
				}
				continue
			}

			if !reflect.DeepEqual(*aSlice[i], *bSlice[i]) {
				return false
			}
		}
	}

	return true
}

func sortEdgeT(aSlice []*EdgeT[string]) {
	sort.Slice(aSlice, func(i, j int) bool {
		if aSlice[i].Weight == aSlice[j].Weight {
			sa := fmt.Sprintf("%s-%s", aSlice[i].V, aSlice[i].U)
			sb := fmt.Sprintf("%s-%s", aSlice[j].V, aSlice[j].U)
			return sa <= sb
		}
		return aSlice[i].Weight < aSlice[j].Weight
	})
}
