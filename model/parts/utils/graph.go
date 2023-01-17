package utils

import (
	"fmt"
)

// Graph structure, node Ids in array and edges in map
type Graph struct {
	nodes []*Node
	edges map[int][]*Edge
}

// Edge that connects to nodes with attributes about the connection
type Edge struct {
	fromNodeId int
	toNodeId   int
	attrs      EdgeAttrs
}

// Edge attributes structure,
// "a2b" show how much this node asked from other node,
// "last" is for the last forgiveness time
type EdgeAttrs struct {
	a2b  int
	last int
}

// func (g *Graph) Init() {
// 	g.nodes = make([]*Node, 1000)
// 	g.edges = make(map[int][]*Edge)
// }

// Returns all nodes
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// AddNode will add a Node to a graph
func (g *Graph) AddNode(nodeId int) error {
	if contains(g.nodes, nodeId) {
		err := fmt.Errorf("node %d already exists", nodeId)
		return err
	} else {
		v := &Node{
			id: nodeId,
		}
		g.nodes = append(g.nodes, v)
	}
	return nil
}

// AddEdge will add an edge from a node to a node
func (g *Graph) AddEdge(edge *Edge) error {
	toNode := g.getNode(edge.toNodeId)
	fromNode := g.getNode(edge.fromNodeId)
	if toNode == nil || fromNode == nil {
		return fmt.Errorf("not a valid edge from %d ---> %d", edge.fromNodeId, edge.toNodeId)
	} else {
		newEdges := append(g.edges[edge.fromNodeId], edge)
		g.edges[edge.fromNodeId] = newEdges
		return nil
	}
}

// getNode will return a node point if exists or return nil
func (g *Graph) getNode(nodeId int) *Node {
	for i, v := range g.nodes {
		if v.id == nodeId {
			return g.nodes[i]
		}
	}
	return nil
}

func contains(v []*Node, id int) bool {
	for _, v := range v {
		if v.id == id {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.nodes {
		fmt.Printf("%d : ", v.id)
		for _, i := range v.adj {
			for _, v := range i {
				fmt.Printf("%d ", v.id)
			}
		}
		fmt.Println()
	}
}

// func PrintEgDirectedGraph() {
// 	g := &Graph{}
// 	g.AddNode(1)
// 	g.AddNode(2)
// 	g.AddNode(3)
// 	g.AddEdge(1, 2)
// 	g.AddEdge(2, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 1)
// 	g.Print()
// }

/*
// Call In Main:: datastructures.PrintEgDirectedGraph()
// Output:
// 1 : 3
// 2 : 1
// 3 : 2 1
*/
