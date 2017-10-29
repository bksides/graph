package graph

type Node interface {}

type Property interface {}

type NodePair struct {
	n1, n2 Node
}

type Edge struct {
	nodes NodePair,
	properties map[string]Property
}

type Graph interface {
	Nodes() []Node
	Edges() map[NodePair]map[string]Property
	Neighbors(n Node) ([]Node, error)
	EdgeInfo(n NodePair) map[string]Property
	AddNode(n Node) error
	AddNodes(n ...Node) map[Node]error
	RemoveNode(n Node) error
	RemoveNodes(n ...Node) map[Node]error
	AddEdge(e Edge) error
	AddEdges(e ...Edge) map[Edge]error
}

type graph struct{
	nodes []Node
	edges map[NodePair]map[string]Property
}

func (this graph) Nodes() []Node {
	nodes = make([]Node, len(this.nodes))
	copy(this.nodes, nodes)
	return nodes
}

func (this graph) Edges() map[NodePair]map[string]Property {
	
}