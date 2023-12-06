package day12

import "fmt"

type Node struct {
	value string
	visit int
}

type Graph struct {
	adjList   map[Node][]Node
	PathCount int
}

func NewGraph() *Graph {
	g := &Graph{
		adjList:   make(map[Node][]Node),
		PathCount: 0,
	}
	return g
}

func (g *Graph) AddEdge(source, dest Node) {
	g.adjList[source] = append(g.adjList[source], dest)
}

func (g *Graph) FindAllPaths(source, dest Node) {
	visited := make(map[Node]int)

	var localPaths []Node
	localPaths = append(localPaths, source)

	g.findAllPathsInternal(source, dest, visited, localPaths)
}

func (g *Graph) findAllPathsInternal(source, dest Node, visited map[Node]int, paths []Node) {
	if source == dest {
		g.PathCount = g.PathCount + 1
		fmt.Println("------------------------------------")
		fmt.Println(paths)
		//fmt.Println(g.pathCount)
		return
	}
	// Add one visit to current node
	if val, ok := visited[source]; ok {
		visited[source] = val + 1
	} else {
		visited[source] = 1
	}

	for _, node := range g.adjList[source] {
		// Not visited the number of times
		if visited[node] < node.visit || node.visit < 0 {
			paths = append(paths, node)
			g.findAllPathsInternal(node, dest, visited, paths)
			paths = paths[:len(paths)-1]
		}
	}
	delete(visited, source)
}
