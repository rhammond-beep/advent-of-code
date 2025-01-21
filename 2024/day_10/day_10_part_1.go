package day_10

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"strings"
)

type Point struct {
	Row int
	Col int
}

type TopologicalMap struct {
	PeaksVisited map[Point]bool
	TrailHeads   []*Point
}

func SolveDay10Part1() int {
	puzzleInput := helper.ReadChallengeInput("./day_10/day_10_input.txt")

	lava_map := make([][]int, len(puzzleInput))
	peak_visited := make(map[Point]bool)
	trail_heads := make([]*Point, 0)

	for i := 0; i < len(puzzleInput); i++ {
		lava_line := make([]int, len(puzzleInput[0]))
		for j := 0; j < len(puzzleInput[0]); j++ {
			var builder strings.Builder
			builder.WriteRune(rune(puzzleInput[i][j])) // single byte can be just a rune
			lava_line[j] = helper.ExtractInt(builder.String())

			point := &Point{Row: i, Col: j}

			if lava_line[j] == 0 {
				trail_heads = append(trail_heads, point)
			}

			if lava_line[j] == 9 {
				peak_visited[*point] = false
			}

		}
		lava_map[i] = lava_line
	}

	topologicalMap := &TopologicalMap{PeaksVisited: peak_visited, TrailHeads: trail_heads}
	return topologicalMap.NavigateFromTrailHeadToPeaks(lava_map)
}

/*
A trail head is a 0 position Node within the map, starting from these trail heads

how many 9's (Mountain peaks) can you navigate to? For each trail head this represents a score
For each trail head, add up the score, this example is 36

My initial thoughts here is to build a graph, but that seems costly, surely I can just>
Find the trail head "0" in the map. Recursively backtrack through each each neighbouring node if
we have a valid jump
*/
func (tm *TopologicalMap) NavigateFromTrailHeadToPeaks(lavaMap [][]int) int {
	score := 0

	for _, th := range tm.TrailHeads {
		tm.navigatePeaks(th.Row, th.Col, lavaMap)
		score += tm.countVisited()
		tm.clearVisited()
	}

	return score
}

func (tm *TopologicalMap) countVisited() int {
	peaks_visited := 0
	for _, pv := range tm.PeaksVisited {
		if pv {
			peaks_visited += 1
		}
	}
	return peaks_visited
}

func (tm *TopologicalMap) clearVisited() {
	for point := range tm.PeaksVisited {
		tm.PeaksVisited[point] = false
	}
}

/*
Given the position within the map (row, col), try to navigate to the sorrounding column

update score through the recursive calls if a valid path has been found
*/
func (tm *TopologicalMap) navigatePeaks(row, col int, lava_map [][]int) {
	current_value := lava_map[row][col]

	if current_value == 9 {
		tm.PeaksVisited[Point{Row: row, Col: col}] = true
	}

	if row > 0 && lava_map[row-1][col] == current_value+1 {
		tm.navigatePeaks(row-1, col, lava_map)
	}

	if col > 0 && lava_map[row][col-1] == current_value+1 {
		tm.navigatePeaks(row, col-1, lava_map)
	}

	if row < len(lava_map)-1 && lava_map[row+1][col] == current_value+1 {
		tm.navigatePeaks(row+1, col, lava_map)
	}

	if col < len(lava_map)-1 && lava_map[row][col+1] == current_value+1 {
		tm.navigatePeaks(row, col+1, lava_map)
	}
}
