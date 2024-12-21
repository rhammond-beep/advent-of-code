package day4

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

/*
 */
func SolveDay4Part1Fail() {
	wordSearch := &WordSearch{SearchSpace: helper.ReadChallengeInput("day_4_input.txt"), SearchTerm: "XMAS"}
	reversedSearch := &WordSearch{SearchSpace: helper.ReadChallengeInput("day_4_input.txt"), SearchTerm: "SAMX"}
	occurrences := 0

	hs := HorizontalSearch{}
	vs := VerticalSearch{}
	ds := DiagonalSearch{}

	occurrences += hs.FindTermOccurrences(wordSearch)
	occurrences += vs.FindTermOccurrences(wordSearch)
	occurrences += ds.FindTermOccurrences(wordSearch)
	occurrences += hs.FindTermOccurrences(reversedSearch)
	occurrences += vs.FindTermOccurrences(reversedSearch)
	occurrences += ds.FindTermOccurrences(reversedSearch)

	fmt.Println(occurrences)
}

/*
interface for specifying a strategy for finding a term within the search space
for reverse address can just change the searchTerm and re-run
*/
type SearchStrategy interface {
	FindTermOccurrences(ws *WordSearch) int
}

/*
helper type for representing a WordSearch to be undertaken
*/
type WordSearch struct {
	SearchSpace []string
	SearchTerm  string
}
