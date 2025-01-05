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

/*
After placing in the new obstacle, do we visit the same edge twice?
*/
func (g *Graph) DetectCycle(lab *Map) (bool, []Node) {
	edgesTraversed := make(map[Point]*Edge, 0) // for a given position, have we gone through a given edge (essentially a set implementation)
	nodesTraversed := make([]Node, 0)

	for {
		foundBarrier, exited := lab.WalkUntilBarrierFound()
		currentVertex := g.Nodes[foundBarrier]
		if currentVertex != nil {

			nodesTraversed = append(nodesTraversed, *currentVertex)
		}

		edge := &Edge{Direction: lab.Direction, Node: currentVertex}

		_, exists := edgesTraversed[foundBarrier]
		if exists {
			return true, nodesTraversed
		}

		edgesTraversed[foundBarrier] = edge

		lab.SetGuardDirection(foundBarrier)

		if exited {
			break
		}
	}

	return false, nodesTraversed
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
