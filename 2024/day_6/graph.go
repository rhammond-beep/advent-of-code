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
	Edges map[Point]*Edge
}

/*
Encode the relationship between given nodes
*/
type Edge struct {
	Node *Node
}

/*
Walk through the input, building up the graph as we go along
*/
func BuildGraph(input []string) Graph {
	nodes := make(map[Point]*Node)

	var previousNode *Node

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				edges := make(map[Point]*Edge)
				edge := &Edge{Node: previousNode}
				edges[point] = edge

				node := &Node{Val: point, Edges: edges}
				nodes[point] = node
				previousNode = node
			}
		}
	}

	return Graph{Nodes: nodes}
}
