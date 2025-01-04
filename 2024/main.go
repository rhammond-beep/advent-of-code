package main

import (
	day6 "rupert-hammond-aoc/day_6"
)

func main() {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	lab := day6.BuildMap(input)
	graph := day6.CreateEmptyGraphFromMap(&lab)
	graph.CreateEdges(&lab)

	lab.ObstacleLocations[*lab.GuardLocation] = true // ignore the guard's starting position

	for point, obstaclePresent := range lab.ObstacleLocations { // For every viable location
		if obstaclePresent { // Then a node already exists, so move on.
			continue
		}

		newObstacle := &day6.Node{Position: point, Type: "obstacle", Edges: make([]*day6.Edge, 0)} // Create the new obstacle

		// try to Knit into graph based on surronding obstacles and guard's trail through the lab
		for node_point, node := range graph.Nodes {
			edge := &day6.Edge{Node: newObstacle}

			if node_point.X == (point.X + 1) {
				_, err := node.Walk("north")
				if node_point.Y > point.Y && err != nil {
					edge.Direction = "east" // In the case where the new obstacle lies above the existing obstacle, the guard must be going east and is to the right of it, the guard must be going east
					node.Edges = append(node.Edges, edge)
				} else if node_point.Y < point.Y { //
					edge.Direction = "west"
					node.Edges = append(node.Edges, edge)
				}

			}

		}

	}

	// _, nodes, error := graph.WalkGraphFromNode(day6.Point{X: 6, Y: 4}, day6.Point{X: 3, Y: 2}, "north")

}
