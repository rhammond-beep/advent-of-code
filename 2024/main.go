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

	graph := day6.BuildGraph(input)

	fmt.Println(graph)
}
