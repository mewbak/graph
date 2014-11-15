// package graph implements generic graph
package graph

type Vertex interface{}

type Edge interface{}

type Graph map[Vertex]map[Vertex]Edge

func New() Graph {
	return make(Graph)
}

func (g Graph) AddDirected(from, to Vertex, e Edge) {
	if _, ok := g[from]; !ok {
		g[from] = make(map[Vertex]Edge)
	}
	g[from][to] = e
}

func (g Graph) AddUndirected(from, to Vertex, e Edge) {
	g.AddDirected(from, to, e)
	g.AddDirected(to, from, e)
}

func (g Graph) RemoveDirected(from, to Vertex) {
	if es, ok := g[from]; ok {
		delete(es, to)
	}
}

func (g Graph) RemoveUndirected(from, to Vertex) {
	g.RemoveDirected(from, to)
	g.RemoveDirected(to, from)
}

func (g Graph) GetEdges(from Vertex) map[Vertex]Edge {
	return g[from]
}

func (g Graph) GetEdge(from, to Vertex) Edge {
	es := g.GetEdges(from)
	if es == nil {
		return nil
	}
	return es[to]
}
