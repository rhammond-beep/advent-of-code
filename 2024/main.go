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
	graph := day6.CreateEmptyGraph(input)
	graph.CreateEdges(&lab)

	_, nodes, error := graph.WalkGraphFromNode(day6.Point{X: 0, Y: 4}, day6.Point{X: 7, Y: 8}, "east")

	if error != nil {
		fmt.Printf("Error Encountered while walking tree: %v", error.Error())
	}

	for i, node := range nodes {
		fmt.Printf("Order Visited: %v Node Value: %v\n", i, node)
	}

	// fmt.Println("***********Graph************")

	// for _, node := range graph.Nodes {
	// 	fmt.Println("-----------Node------------")
	// 	fmt.Printf("reference: %v\n point: %v\n edges: %v\n", &node, node.Val, node.Edges)
	// 	fmt.Println("---------------------------")
	// 	fmt.Println("-----------Edges-----------")
	// 	for _, edge := range node.Edges {
	// 		fmt.Printf("edge ref: %v\n direction: %v\n", &edge, edge.Direction)
	// 	}
	// 	fmt.Println("---------------------------")
	// }
	// fmt.Println("********************************")

}
