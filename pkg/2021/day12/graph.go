package day12

type Node struct {
	value string
	visit int
}

type Graph struct {
	adjList   map[Node][]Node
	pathCount int
}

func NewGraph() *Graph {
	g := &Graph{
		adjList:   make(map[Node][]Node),
		pathCount: 0,
	}
	return g
}

func (g *Graph) AddEdge(source, dest Node) {
	g.adjList[source] = append(g.adjList[source], dest)
}

func (g *Graph) PathCount() int {
	return g.pathCount
}

func (g *Graph) FindAllPaths(source, dest Node) {
	//fmt.Println(g.adjList)
	visited := make(map[Node]int)

	var localPaths []Node
	localPaths = append(localPaths, source)

	g.findAllPathsInternal(source, dest, visited, localPaths)
}

func (g *Graph) findAllPathsInternal(source, dest Node, visited map[Node]int, paths []Node) {
	if source == dest {
		g.pathCount = g.pathCount + 1
		//fmt.Println("------------------------------------")
		//fmt.Println(visited)
		//fmt.Println(paths)
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
		if node.visit < 0 || visited[node] < node.visit {
			paths = append(paths, node)
			g.findAllPathsInternal(node, dest, visited, paths)
			paths = paths[:len(paths)-1]
		}
	}
	delete(visited, source)
}
