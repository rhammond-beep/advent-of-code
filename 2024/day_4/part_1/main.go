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
	occurrences := 0

	hs := HorizontalSearch{}
	vs := VerticalSearch{}
	ds := DiagonalSearch{}

	occurrences += hs.FindTermOccurrences(wordSearch)
	occurrences += vs.FindTermOccurrences(wordSearch)
	occurrences += ds.FindTermOccurrences(wordSearch)

	fmt.Println(occurrences)

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
for reverse address can just change the searchTerm and re-run
*/
type SearchStrategy interface {
	FindTermOccurrences(ws *WordSearch) int
}

type HorizontalSearch struct {
}

type VerticalSearch struct {
}

type DiagonalSearch struct {
}

/*
check each row for any occuring instances of the search term
*/
func (hs *HorizontalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	ub := len(ws.SearchTerm)
	for _, row := range ws.SearchSpace {
		// For each row, split up into chunks based on the length of the target word
		for i := ub; i < len(row)-ub; i += 1 {
			word := row[i-ub : i]
			if word == ws.SearchTerm {
				occurrences += 1
			}
		}
	}
	return
}

/*
	 for each character in the column
		Take the next n characters.
		Check if they match the search term
			if yes, increment occurrences
*/
func (vs *VerticalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	ub := len(ws.SearchTerm)
	columns := ws.SearchSpace[0]
	for i := 0; i < len(columns); i += 1 {
		offset := 0
		var sb strings.Builder
		for j := offset; j < (offset + ub); j += 1 {
			sb.WriteByte(ws.SearchSpace[j][i])
		}
		offset += 1
		if sb.String() == ws.SearchTerm {
			occurrences += 1
		}
	}

	return
}

/*
 */
func (ds *DiagonalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {

	return
}

/*
helper type for representing a WordSearch to be undertaken
*/
type WordSearch struct {
	SearchSpace []string
	SearchTerm  string
}
