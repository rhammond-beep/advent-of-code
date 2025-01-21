package day_10

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"strings"
)

func SolveDay10Part1() int {

	puzzleInput := helper.ReadChallengeInput("./day_10/day_10_input.txt")
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

	return calculateTrailHeadsScore(lava_map)
}

/*
A trail head is a 0 position Node within the map, starting from these trail heads

how many 9's (Mountain peaks) can you navigate to? For each trail head this represents a score
For each trail head, add up the score, this example is 36

My initial thoughts here is to build a graph, but that seems costly, surely I can just>
Find the trail head "0" in the map. Recursively backtrack through each each neighbouring node if
we have a valid jump
*/
func calculateTrailHeadsScore(lava_map [][]int) (score int) {
	for i := 0; i < len(lava_map); i++ {
		for j := 0; j < len(lava_map); j++ {
			if lava_map[i][j] == 0 {
				score += navigatePeaks(i, j, lava_map)
			}
		}
	}
	return
}

/*
Given the position within the map (row, col), try to navigate to the sorrounding column
update score through the recursive calls if a valid path has been found
this will count the number of distinct valid paths to a given trail
*/
func navigatePeaks(row, col int, lava_map [][]int) int {
	score := 0
	current_value := lava_map[row][col]

	if current_value == 9 {
		return 1
	}

	if row > 0 && lava_map[row-1][col] == current_value+1 {
		score += navigatePeaks(row-1, col, lava_map)
	}

	if col > 0 && lava_map[row][col-1] == current_value+1 {
		score += navigatePeaks(row, col-1, lava_map)
	}

	if row < len(lava_map)-1 && lava_map[row+1][col] == current_value+1 {
		score += navigatePeaks(row+1, col, lava_map)
	}

	if col < len(lava_map)-1 && lava_map[row][col+1] == current_value+1 {
		score += navigatePeaks(row, col+1, lava_map)
	}

	return score
}
