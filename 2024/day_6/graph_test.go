package day6

import (
	"fmt"
	"testing"
)

func TestShouldConstructGraphCorrectly(t *testing.T) {
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

	graph := BuildGraph(input)

	fmt.Println(graph)
}
