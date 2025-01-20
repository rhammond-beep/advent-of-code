package day_10

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"strings"
)

func SolveDay10Part1() {

	puzzleInput := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
	// need to transform this into a matrix of ints

	lava_map := make([][]int, len(puzzleInput))

	for i := 0; i < len(puzzleInput); i++ {
		lava_line := make([]int, len(puzzleInput[0]))
		for j := 0; j < len(puzzleInput[0]); j++ {
			var builder strings.Builder
			builder.WriteRune(rune(puzzleInput[i][j])) // single byte can be just a rune
			lava_line[j] = helper.ExtractInt(builder.String())
		}
		lava_map[i] = lava_line
	}

	// A trail head is a 0 position Node within the map, starting from these trail heads
	// how many 9's (Mountain peaks) can you navigate to? For each trail head this represents a score
	// For each trail head, add up the score, this example is 36

	// My initial thoughts here is to build a graph, but that seems costly, surely I can just>
	// Find the trail head "0" in the map. Recursively backtrack through each each neighbouring node if
	// we have a valid jump
}

func findTrailHeads() {

}
