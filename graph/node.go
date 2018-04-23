package graph

import "fmt"

type node struct {
	name  string
	edges []*node
}

func NewNode(name string) *node {
	return &node{
		name:  name,
		edges: []*node{},
	}
}

func (n node) String() string {
	return fmt.Sprintf("%s", n.GetName())
}

func (n *node) AddEdge(node *node) {
	// edges should be a set
	for _, v := range n.edges {
		if node.GetName() == v.GetName() {
			// Already exists in list
			return
		}
	}

	n.edges = append(n.edges, node)

	// Add bidirectional Edge
	node.AddEdge(n)
}

func (n *node) AddEdges(nodes []*node) {
	for _, v := range nodes {
		n.AddEdge(v)
	}
}

func (n *node) GetName() string {
	return n.name
}

func (n *node) GetEdges() []*node {
	return n.edges
}

func (n *node) IsLeaf() bool {
	return len(n.edges) == 1
}
