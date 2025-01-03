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
	Direction string // Do we need the direction strictly?
	Node      *Node
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
Function to recursively walk the graph in the same order as the guard until a null node is found
*/
func (g *Graph) WalkGraph() {

}

/*
Create an edge between two adjecent nodes (should this be a directed one??)
as in, should I be storing the relationship on both sides here? Probably not.

Should be able to use this to Place in an extra obstacle to create a cycle.
*/
func (n *Node) AddEdge(cn *Node, direction string) {
	n.Edges = append(n.Edges, &Edge{Node: cn, Direction: direction})
}
