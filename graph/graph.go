package graph

import (
	"fmt"
)

type graphMap map[string]*node

type Graph struct {
	vertices graphMap
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(graphMap),
	}
}

// Adds a new node to graph if it doesn't exist
func (g *Graph) AddNode(name string) *node {
	// If node doesn't exists GetNodeByName returns error
	// If error create new node
	if _, err := g.GetNodeByName(name); err != nil {
		g.vertices[name] = NewNode(name)
	}

	return g.vertices[name]
}

func (g *Graph) GetNodeByName(name string) (*node, error) {
	if _, ok := g.vertices[name]; !ok {
		return nil, fmt.Errorf("node with name %s doesn't exist", name)
	}

	return g.vertices[name], nil
}

func (g *Graph) AddTarget(source, target string) {
	// AddNode will only create a new Node if doesn't exist
	sourceNode := g.AddNode(source)
	targetNode := g.AddNode(target)

	// AddEdge handles bidirectional paths
	sourceNode.AddEdge(targetNode)
}

func (g *Graph) GetTargets(name string) ([]*node, error) {
	node, err := g.GetNodeByName(name)
	if err != nil {
		return nil, err
	}

	return node.GetEdges(), nil
}

// Breadth First Search
// Returns whether you can be teleported to destination from source
func (g *Graph) TeleportToBFS(source, destination string) bool {
	queue := NewQueue()

	// Keep track of visited nodes
	visited := make(map[string]bool)

	// Start by checking the root node
	root, err := g.GetNodeByName(source)
	// Node doesn't exist
	if err != nil {
		return false
	}

	_, err = g.GetNodeByName(destination)
	// Destination node doesn't exist
	if err != nil {
		return false
	}

	queue.Enqueue(root)

	// While the queue isn't empty
	for queue.Length() != 0 {

		subtree, _ := queue.Dequeue()

		if subtree.GetName() == destination {
			// Don't teleport to self
			if subtree.GetName() == source {
				return false
			}
			return true
		}

		// Root is not the destination
		// Search edges on root and see if they are the destination
		for _, node := range subtree.GetEdges() {

			// Already checked this node
			if visited[node.GetName()] {
				continue
			}

			// Don't search same node more than once
			if !queue.Contains(node) {
				queue.Enqueue(node)
			}

			// Mark this node as visited
			visited[node.GetName()] = true
		}
	}

	return false
}

// Does NOT return a unique list of results..
func (g *Graph) DepthDFS(source string, depth int) []string {
	result := []string{}

	// Base case depth reached
	if depth < 0 {
		return result
	}

	// Start by checking the root node
	root, err := g.GetNodeByName(source)
	// Node doesn't exist
	if err != nil {
		return result
	}

	// Add root name
	result = append(result, root.GetName())

	// Decrement depth
	depth--

	// Recursively call DFS with decremented depth for each connected edge
	for _, v := range root.GetEdges() {
		elements := g.DepthDFS(v.GetName(), depth)
		result = append(result, elements...)
	}

	return result
}

// Filter out source node and duplicates from the DFSDepth algo
func (g *Graph) GetEdgesAtDepth(source string, depth int) []string {
	result := []string{}

	// get all edges including duplicates..
	elements := g.DepthDFS(source, depth)

	for _, el := range elements {
		// Exclude source node
		if el == source {
			continue
		}

		found := false

		// remove dups
		for _, v := range result {
			if v == el {
				found = true
				break
			}
		}

		if !found {
			result = append(result, el)
		}
	}

	return result
}

// Won't honor unique cycles
func (g *Graph) cycleDFS(start string, edge *node, visited map[string]bool, found *bool) {
	node := edge.GetName()

	// To have a cycle, you must have a backedge
	if visited[node] {
		if node == start {
			*found = true
		}
		return
	}

	visited[node] = true

	for _, v := range edge.GetEdges() {
		g.cycleDFS(start, v, visited, found)
	}
}

func (g *Graph) CheckUniqueCycle(source string) bool {
	node, err := g.GetNodeByName(source)
	if err != nil {
		return false
	}

	edges := node.GetEdges()
	numLeaves := 0

	for _, v := range edges {
		if v.IsLeaf() {
			numLeaves++
		}
	}

	// ie If a node has 2 edges but one is a leaf, it can't have a unique cycle
	if len(edges)-numLeaves <= 1 {
		return false
	}

	return true
}

// Won't honor unique cycles
func (g *Graph) DetectCycles(source string) bool {
	visited := make(map[string]bool)

	root, err := g.GetNodeByName(source)
	if err != nil {
		return false
	}

	var (
		found bool
	)

	g.cycleDFS(source, root, visited, &found)
	return found
}
