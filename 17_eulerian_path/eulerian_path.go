package main

import (
	"fmt"
	"github.com/C6H6/algoimpl/go/graph"
)

func main() {
	g := graph.New(graph.Undirected)

	nodes := make(map[rune]graph.Node, 0)
	nodes[0] = g.MakeNode()
	nodes[1] = g.MakeNode()
	nodes[2] = g.MakeNode()
	nodes[3] = g.MakeNode()
	nodes[4] = g.MakeNode()

	for key, node := range nodes {
		*node.Value = key
	}
	g.MakeEdge(nodes[0], nodes[1])
	g.MakeEdge(nodes[1], nodes[2])
	g.MakeEdge(nodes[2], nodes[0])
	g.MakeEdge(nodes[0], nodes[3])
	g.MakeEdge(nodes[3], nodes[4])

	path := g.EulerianPath()

	for _, edge := range path{
		fmt.Println(*edge.Start.Value, "->", *edge.End.Value)
	}
}
