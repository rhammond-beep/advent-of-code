package main

import (
	"fmt"
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
	graph := day6.CreateGraph(input)
	graph.PopulateEdges(&lab)

	fmt.Println(graph)
	fmt.Println("-----------Verticies------------")

	for _, node := range graph.Nodes {
		fmt.Println("-----------Node------------")
		fmt.Printf("reference: %v\n point: %v\n edges: %v\n", &node, node.Val, node.Edges)
		fmt.Println("---------------------------")
	}
}
