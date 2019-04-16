package graph

// Graph 无向图
type Graph struct {
	adj []LinkList
	num int
}

// NewGraph new
func NewGraph(v int) Graph {
	adj := make([]LinkList, v)
	return Graph{adj, v}
}

// AddEdge 添加一条边
func (g *Graph) AddEdge(s, t int) {
	g.adj[s].AddNode(t)
	g.adj[t].AddNode(s)
}
