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
	graph.PrintGraphVisulisation()

	// _, nodes, error := graph.WalkGraphFromNode(day6.Point{X: 6, Y: 4}, day6.Point{X: 3, Y: 2}, "north")

	// if error != nil {
	// 	fmt.Printf("Error Encountered while walking graph: %v", error.Error())
	// }

	//for i, node := range nodes {
	//	fmt.Printf("Order Visited: %v Node Value: %v\n", i+1, node)
	//}

}
