package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
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
		for i := ub; i < len(row)+1; i += 1 {
			word := row[i-ub : i]
			if word == ws.SearchTerm {
				occurrences += 1
			}
		}
	}
	return
}

/*
 */
func (vs *VerticalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	mappedSearchSpace := make([]string, len(ws.SearchSpace))

	// Map to a horizontal space
	for i := 0; i < len(ws.SearchSpace); i++ {
		var sb strings.Builder
		for j := 0; j < len(ws.SearchSpace); j++ {
			sb.WriteByte(ws.SearchSpace[j][i])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	hs := &HorizontalSearch{}
	mappedSearch := &WordSearch{SearchSpace: mappedSearchSpace, SearchTerm: ws.SearchTerm}

	return hs.FindTermOccurrences(mappedSearch)
}

/*
Diagonal Search is a little more tricky that the others,
We want to move across the diagonal lines, however we have to exclude all
whose length (number of dicrete points n) < len(ws.SearchTerm)

This seems like we've maybe got to do a little pre-processing. One thing we can
do is build up a slice containing the indexes of the diagonals [][]int

[

	[3,0], (i, j)
	[2,1],
	[1,2],
	[0,3]

]

We take the size of the diagonals len(diagonals) and remove anything whose size is
< len(ws.SearchTerm)

Now Iterate over the remaining lines with the sliding window, use this to extract out
a string compare it to the desired searchTerm.

Feels like I shouldn't have to do this preprocessing stage though... Feels pretty annoying,
I'd prefer to just do the work as part of the original loop.
*/
func (ds *DiagonalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	//	var diagonals [][]int
	//
	//	for k := len(ws.SearchTerm) - 1; k < len(ws.SearchSpace); k++ {
	//		offset := len(ws.SearchTerm)
	//		for j := 0; j <= k; j++ {
	//			i := k - j
	//			// ws.SearchSpace[i][j])
	//		}
	//	}
	//
	//	for k := len(ws.SearchSpace); k > len(ws.SearchTerm)-1; k-- {
	//		for j := 0; j <= k; j++ {
	//			i := k - j
	//			// ws.SearchSpace[i][j])
	//		}
	//	}
	//
	return
}

/*
helper type for representing a WordSearch to be undertaken
*/
type WordSearch struct {
	SearchSpace []string
	SearchTerm  string
}
