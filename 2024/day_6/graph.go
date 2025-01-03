package day6

import "errors"

/*
The space representing the patrol path of the guard and the obstructions he encounters
*/
type Graph struct {
	Nodes map[Point]*Node
}

/*
An Obstacle in this given instance, the subject of our interest.
*/
type Node struct {
	Val   Point
	Edges []*Edge // A List of connected Edges
}

/*
Encode the relationship between given nodes
*/
type Edge struct {
	Direction string // Do we need the direction strictly?
	Node      *Node
}

var directionMap = map[string]string{
	"north": "east",
	"east":  "south",
	"south": "west",
	"west":  "north",
}

/*
Walk through the input, defining the unconnected verticies
*/
func CreateEmptyGraph(input []string) Graph {
	nodes := make(map[Point]*Node)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				edges := make([]*Edge, 0)
				node := &Node{Val: point, Edges: edges}
				nodes[point] = node
			}
		}
	}

	return Graph{Nodes: nodes}
}

/*
Accept a map of the lab as input to populate the edges based on the walk of the guard
As we create the edges within the graph, return the nodes we visit in order.
*/
func (g *Graph) CreateEdges(lab *Map) []*Node {
	var previousVertex *Node
	nodes_walked := make([]*Node, 0)

	for {
		foundBarrier, exited := lab.WalkUntilBarrierFound()

		currentVertex := g.Nodes[foundBarrier]
		nodes_walked = append(nodes_walked, currentVertex)

		if previousVertex != nil {
			previousVertex.AddEdge(currentVertex, lab.Direction)
		}

		previousVertex = currentVertex

		lab.SetGuardDirection(foundBarrier)

		if exited {
			break
		}
	}

	return nodes_walked
}

/*
given that two point correspond to a node in the graph, Try to walk from the start point to the endpoint,
returning a boolean to indicate if walk was successful along with the nodes traversed, otherwise returning false and nil
*/
func (g *Graph) WalkGraphFromNode(startPoint, endPoint Point, direction string) (bool, []*Node, error) {
	currentNode, ok := g.Nodes[startPoint]
	if !ok {
		return false, nil, errors.New("No node at the specified starting position")
	}

	endNode, ok := g.Nodes[endPoint]
	if !ok {
		return false, nil, errors.New("No node at the specified end position")
	}
	path := make([]*Node, 0)

	for {
		nextNode, error := currentNode.Walk(direction)
		path = append(path, nextNode)

		if error != nil {
			return false, path, error
		}

		if nextNode == endNode {
			return true, path, nil
		}

		direction = directionMap[direction]
		currentNode = nextNode
	}
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

	return nil, errors.New("No edge with that direction found")
}

/*
Create an edge between two adjecent nodes (should this be a directed one??)
as in, should I be storing the relationship on both sides here? Probably not.

Should be able to use this to Place in an extra obstacle to create a cycle.
*/
func (n *Node) AddEdge(cn *Node, direction string) {
	n.Edges = append(n.Edges, &Edge{Node: cn, Direction: direction})
}
