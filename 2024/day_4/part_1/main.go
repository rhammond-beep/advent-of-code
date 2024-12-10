package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Day 4 - Ceres Search

	Lets do a word search, words are allowed to be:
	1) Horizontal
	2) Vertical
	3) diagonal
	4) wirtten backwards
	5) Overlapping other words

	Find all the valid instances of "XMAS" within the search space.
*/
func main() {
	wordSearch := &WordSearch{SearchSpace: ReadChallengeInput("../day_4_input.txt"), SearchTerm: "XMAS"}

}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}

func ReadChallengeInput(filepath string) (searchSpace []string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		searchSpace = append(searchSpace, scanner.Text())
	}
	return
}

/*
interface for specifying a strategy for finding a term within the search space
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
