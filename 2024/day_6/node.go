package day6

import (
	"errors"
	"fmt"
)

/*
An Obstacle in this given instance, the subject of our interest.
*/
type Node struct {
	Position Point
	Type     string
	Edges    []*Edge // A List of connected Edges
}

/*
Walk a step in the specified direction, return an error if not possible
*/
func (n *Node) Walk(direction string) (*Node, error) {
	for _, edge := range n.Edges {
		if edge.Direction == direction {
			return edge.Node, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No edge with direction %v found from node %v", direction, n))
}

/*
Create an edge between two adjecent nodes (should this be a directed one??)
as in, should I be storing the relationship on both sides here? Probably not.

Should be able to use this to Place in an extra obstacle to create a cycle.
*/
func (n *Node) AddEdge(cn *Node, direction string) {
	n.Edges = append(n.Edges, &Edge{Node: cn, Direction: direction})
}
