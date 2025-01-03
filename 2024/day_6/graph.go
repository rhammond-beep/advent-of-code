package day6

import (
	"errors"
	"fmt"
)

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
	Position Point
	Type     string
	Edges    []*Edge // A List of connected Edges
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
Walk through the obstacles, defining the unconnected verticies
(We also define "Guard" as a special type of node which we start at)
*/
func CreateEmptyGraphFromMap(m *Map) Graph {
	nodes := make(map[Point]*Node)

	nodes[*m.GuardLocation] = &Node{Type: "guard", Position: *m.GuardLocation, Edges: make([]*Edge, 0)}

	for obstacle, present := range m.ObstacleLocations {
		if present {
			nodes[obstacle] = &Node{Type: "obstacle", Position: obstacle, Edges: make([]*Edge, 0)}
		}
	}

	return Graph{Nodes: nodes}
}

/*
Accept a map of the lab as input to populate the edges based on the walk of the guard
As we create the edges within the graph, return the nodes we visit in order.
*/
func (g *Graph) CreateEdges(lab *Map) []*Node {
	previousVertex := g.Nodes[*lab.GuardLocation]
	nodes_walked := make([]*Node, 0)

	for {
		foundBarrier, exited := lab.WalkUntilBarrierFound()

		currentVertex := g.Nodes[foundBarrier]
		nodes_walked = append(nodes_walked, currentVertex)
		previousVertex.AddEdge(currentVertex, lab.Direction)

		previousVertex = currentVertex

		lab.SetGuardDirection(foundBarrier)

		if exited {
			break
		}
	}

	return nodes_walked
}

func (g *Graph) PrintGraphVisulisation() {
	fmt.Println("***********Graph************")

	for _, node := range g.Nodes {
		fmt.Println("-----------Node------------")
		fmt.Printf("reference: %v\n Value: %v\n", &node, node)
		fmt.Println("---------------------------")
		fmt.Println("-----------Edges-----------")
		for _, edge := range node.Edges {
			fmt.Printf("edge ref: %v\n direction: %v\n", &edge, edge.Direction)
		}
		fmt.Println("---------------------------")
	}
	fmt.Println("********************************")
}

/*
Counts the number of edges across all Nodes
*/
func (g *Graph) CalculateGraphSize() int {
	number_edges := 0
	for _, node := range g.Nodes {
		for range node.Edges {
			number_edges += 1
		}
	}
	return number_edges
}

/*
given that two points correspond to a node in the graph, Try to walk between the two endpoints,
returning a boolean to indicate if walk was successful along with the nodes traversed, (Excluing the start and end Nodes)
otherwise, if the path in unreachable false is returned with an accompaning error message
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

		if error != nil {
			return false, path, error
		}

		if nextNode == endNode {
			return true, path, nil
		}

		path = append(path, nextNode)

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
