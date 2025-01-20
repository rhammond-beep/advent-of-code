package day_10

type Point struct {
	X int
	Y int
}

type Graph struct {
	Nodes map[Point]*Node
}

/*
Encode the relationship between given nodes
*/
type Edge struct {
	Node *Node
}

/*
An Obstacle in this given instance, the subject of our interest.
*/
type Node struct {
	Value int
	Type  string
	Edges []*Edge // A List of connected Edges
}

/*
From a given node, we need to see how many 9's we can walk to
*/
func (n *Node) Walk(validTrails int) int {
	if n.Value == 9 {
		validTrails += 1
	}

	for _, edge := range n.Edges {
		if edge.Node.Value == (n.Value + 1) {
			edge.Node.Walk(validTrails)
		}
	}

	return validTrails
}

/*
Create an edge between two adjecent nodes (should this be a directed one??)
as in, should I be storing the relationship on both sides here? Probably not.

Should be able to use this to Place in an extra obstacle to create a cycle.
*/
func (n *Node) AddEdge(cn *Node, direction string) {
	n.Edges = append(n.Edges, &Edge{Node: cn})
}
