package day6

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
	Node *Node
}

/*
Walk through the input, defining the unconnected verticies
*/
func CreateGraph(input []string) Graph {
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
Accept a map of the lab as input to populate the edges belonging to the graph
*/
func (g *Graph) PopulateEdges(lab *Map) {
	var previousVertex *Node

	for {
		foundBarrier, exited := lab.WalkUntilBarrierFound()

		currentVertex := g.Nodes[foundBarrier]

		if previousVertex != nil {
			previousVertex.addEdge(currentVertex)
		}

		previousVertex = currentVertex

		lab.SetGuardDirection(foundBarrier)

		if exited {
			break
		}
	}
}

/*
Create an edge between two nodes (should this be a directed one??)
*/
func (n *Node) addEdge(cn *Node) {
	n.Edges = append(n.Edges, &Edge{Node: cn})
}
